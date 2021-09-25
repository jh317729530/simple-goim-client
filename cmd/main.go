package main

import (
	"flag"
	"io"
	"math/rand"
	"os"
	"simple-goim-client/internal/conf"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	file, err := os.OpenFile("simple-goim-client.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Err(err)
		panic(err)
	}

	subLogger := log.Logger.Output(io.MultiWriter(os.Stdout, file))
	log.Logger = subLogger

	log.Info().Msg("test")

	if err := conf.Init(); err != nil {
		log.Error().Msg("Global config init failed")
		panic(err)

	}

}
