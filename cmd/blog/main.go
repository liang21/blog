package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	grpc_middlewares "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	ac_v1 "github.com/liang21/blog/api/blog/v1"
	article_v1 "github.com/liang21/blog/api/blog/v1"
	category_v1 "github.com/liang21/blog/api/blog/v1"
	tag_v1 "github.com/liang21/blog/api/blog/v1"
	blogBiz "github.com/liang21/blog/internal/blog/biz"
	blogRepo "github.com/liang21/blog/internal/blog/repository"
	blogService "github.com/liang21/blog/internal/blog/service"
	"github.com/liang21/blog/pkg/file"
	"github.com/liang21/blog/pkg/log"
	"github.com/liang21/blog/pkg/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

var (
	confPath = flag.String("conf", "D:\\terminator\\go\\blog\\configs\\comfig.yml", "config path, eg: -conf ./configs")
)

func main() {
	flag.Parse()
	// logger配置
	logOpts := &log.Options{
		Level:            "debug",
		Format:           "console",
		EnableColor:      true,
		DisableCaller:    false,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	// 初始化全局logger
	log.Init(logOpts)
	defer log.Flush()
	bs, err := file.LoadFile(*confPath)
	if err != nil {
		log.Fatalf("load config failed!", err)
	}
	engine, err := options.NewMysqlOptions(&bs.Data.Mysql)
	if err != nil {
		log.Fatalf("init mysql failed!", err)
	}

	// grpc client
	//cc, err := grpc.Dial(bs.GRPC.Addr, grpc.WithInsecure())

	lis, err := net.Listen("tcp", bs.GRPC.Addr)
	if err != nil {
		log.Fatalf("init listen grpc failed!", err)
	}
	defer lis.Close()
	s := grpc.NewServer(
		// add grpc interceptor
		grpc.UnaryInterceptor(
			grpc_middlewares.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
				grpc_validator.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(zap.NewNop()),
			)))
	// TODO: 注册服务
	// article service
	articleRepo := blogRepo.NewArticleRepo(engine)
	articleCase := blogBiz.NewArticleCase(articleRepo)
	articleService := blogService.NewArticleService(articleCase)
	article_v1.RegisterArticleServiceServer(s, articleService)
	// tag service
	tagRepo := blogRepo.NewTagRepo(engine)
	tagCase := blogBiz.NewTagCase(tagRepo)
	tagService := blogService.NewTagService(tagCase)
	tag_v1.RegisterTagServiceServer(s, tagService)
	// category service
	categoryRepo := blogRepo.NewCategoryRepo(engine)
	categoryCase := blogBiz.NewCategoryCase(categoryRepo)
	categoryService := blogService.NewCategoryService(categoryCase)
	category_v1.RegisterCategoryServiceServer(s, categoryService)
	// article category service
	articleCategoryRepo := blogRepo.NewArticleCategoryRepo(engine)
	articleCategoryCase := blogBiz.NewArticleCategoryCase(articleCategoryRepo)
	articleCategoryService := blogService.NewArticleCategoryService(articleCategoryCase)
	ac_v1.RegisterArticleCategoryServiceServer(s, articleCategoryService)
	// 启动grpc服务
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("start grpc service failed!", err)
		}
	}()
	mux := runtime.NewServeMux(
		// 返回错误结果处理
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaller runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			e := runtime.HTTPStatusError{HTTPStatus: 200, Err: err}
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaller, writer, request, &e)
		}),
		// 编码转换
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)
	rpcS := strings.Split(bs.GRPC.Addr, `:`)
	rpcPort := rpcS[1]
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// 注册服务
	if err := article_v1.RegisterArticleServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+rpcPort, opts); err != nil {
		log.Fatalf("register service failed!", err)
	}
	if err := tag_v1.RegisterTagServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+rpcPort, opts); err != nil {
		log.Fatalf("register service failed!", err)
	}
	if err := category_v1.RegisterCategoryServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+rpcPort, opts); err != nil {
		log.Fatalf("register service failed!", err)
	}
	if err := ac_v1.RegisterArticleCategoryServiceHandlerFromEndpoint(context.Background(), mux, "localhost:"+rpcPort, opts); err != nil {
		log.Fatalf("register service failed!", err)
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	pprof.Register(router)
	router.Use(GinLogger())
	router.Group("/blog").Any("/*{gateway}", gin.WrapH(mux))
	if err = router.Run(bs.HTTP.Addr); err != nil {
		panic(err)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s %v", debug.Stack(), err)
		}
	}()
}
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		log.Debugf(path,
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"ip", c.ClientIP(),
			"cost", cost,
			"user-agent", c.Request.UserAgent(),
			"error", c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
	}
}
