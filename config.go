package ratelimit

import (
	"time"

	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type ratelimiteConfig struct {
	RateLimit time.Duration `json:"rate_limit" config:"app.rate_limit" default:"30s"`
}

var (
	Config = &ratelimiteConfig{}
)

func (ratelimiteConfig) ConfigName() string {
	return "RateLimit"
}

func (ratelimiteConfig) SetDefaults() {
}

func (a *ratelimiteConfig) Read() {
	vipertags.Fill(a)
}

func (c ratelimiteConfig) String() string {
	return pp.Sprintln(c)
}

func (c ratelimiteConfig) Debug() {
	log.Debug("RateLimit Config = ", c)
}

func init() {
	config.Register(Config)
}
