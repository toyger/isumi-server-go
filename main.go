package main

import (
	"runtime"
	"flag"

	"github.com/toyger/isumi-server/config"
	"github.com/toyger/isumi-server/server"

)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	conf := config.Parse()

	flag.Set("bind", conf.Bind)

	server.Serve(conf)
}
