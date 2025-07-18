package midas

import (
	"github.com/rs/zerolog"
	"time"
)

type RetryConfig[T any] struct {
	Config   T //For any config type
	Logger   *zerolog.Logger
	Attempts int
}

func RetryFunc[T any](log zerolog.Logger, cfg RetryConfig[T], start func(config T, log zerolog.Logger) error) {
	for i := 0; i < cfg.Attempts; i++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Error().Msgf("panic caught: %v", r)
				}
			}()
			time.Sleep(time.Duration(1<<i) * time.Second) //exp duration
			log.Info().Msgf("%d try of start", i)
			if err := start(cfg.Config, log); err != nil { //start thats func for starting program
				log.Error().Msgf("error starting %d try: %v", i, err)
			}
		}()
	}
}
