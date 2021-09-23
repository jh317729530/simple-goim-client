package main

import (
	"flag"
	"fmt"
	"simple-goim-client/internal/conf"
)

func main() {
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}

	fmt.Println(conf.Conf.HttpServer.Addr)
}
