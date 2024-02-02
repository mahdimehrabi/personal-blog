package zerolog

import "github.com/rs/zerolog/log"

type ZeroLog struct {
}

func NewZeroLog() *ZeroLog {
	return &ZeroLog{}
}

func (l ZeroLog) Error(errMsg string) {
	log.Error().Msg(errMsg)
}
