package ratelimit

import (
	"time"

	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type ratelimitConfig struct {
	RateLimit time.Duration `json:"rate_limit" config:"app.rate_limit" default:"30s"`
}

var (
	Config = &ratelimitConfig{}
)

func (ratelimitConfig) ConfigName() string {
	return "RateLimit"
}

func (ratelimitConfig) SetDefaults() {
}

func (a *ratelimitConfig) Read() {
	vipertags.Fill(a)
}

func (c ratelimitConfig) String() string {
	return pp.Sprintln(c)
}

func (c ratelimitConfig) Debug() {
	log.Debug("RateLimit Config = ", c)
}

func init() {
	config.Register(Config)
}
