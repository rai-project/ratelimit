package ratelimit

import (
	"time"

	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
)

type ratelimitConfig struct {
	RateLimit time.Duration `json:"rate_limit" config:"app.rate_limit" default:"120s"`
	done      chan struct{} `json:"-" config:"-"`
}

var (
	Config = &ratelimitConfig{
		done: make(chan struct{}),
	}
)

func (ratelimitConfig) ConfigName() string {
	return "RateLimit"
}

func (a *ratelimitConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

func (a *ratelimitConfig) Read() {
	vipertags.Fill(a)
}

func (c ratelimitConfig) Wait() {
	<-c.done
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
