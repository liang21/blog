package config

import "github.com/liang21/blog/pkg/options"

type Bootstrap struct {
	Server
	Data
}
type Server struct {
	GRPC Addr
	HTTP Addr
}
type Addr struct {
	Addr string
}
type Data struct {
	Mysql options.MysqlOptions
}
