package cmd

import (
	"github.com/konoui/alfred-tldr/pkg/tldr"
	"github.com/konoui/go-alfred"
)

type envs struct {
	formatFunc                 func(string) string
	openURLModKey              alfred.ModKey
	isUpdateDBRecommendEnabled bool
}

type config struct {
	platform       tldr.Platform
	language       string
	update         bool
	updateWorkflow bool
	confirm        bool
	fuzzy          bool
	version        bool
	tldrClient     *tldr.Tldr
	fromEnv        envs
}

func newConfig() *config {
	cfg := new(config)
	cfg.fromEnv.formatFunc = getCommandFormatFunc()
	cfg.fromEnv.openURLModKey = getOpenURLMod()
	cfg.fromEnv.isUpdateDBRecommendEnabled = isUpdateDBRecommendEnabled()
	return cfg
}
