package ratelimit

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/Unknwon/com"
	humanize "github.com/dustin/go-humanize"
	"github.com/pkg/errors"

	"github.com/rai-project/config"
)

const (
	timeFormat = time.RFC3339Nano
)

func New(opts ...RateLimitOption) error {
	options := RateLimitOptions{
		limit: Config.RateLimit,
	}
	for _, o := range opts {
		o(&options)
	}
	tmpDir := os.TempDir()
	if !com.IsDir(tmpDir) {
		msg := "Not able to find temporary directory for " + config.App.Name + "."
		log.Debug(msg)
		return errors.New(msg)
	}
	dotKeepFilePath := filepath.Join(tmpDir, config.App.Name+".keep")
	if !com.IsFile(dotKeepFilePath) {
		com.WriteFile(dotKeepFilePath, []byte(time.Now().Format(timeFormat)))
		return nil
	}
	tbytes, err := ioutil.ReadFile(dotKeepFilePath)
	if err != nil {
		msg := "Not able to read " + dotKeepFilePath + "."
		log.WithError(err).Debug(msg)
		return errors.Wrap(err, msg)
	}
	prevTime, err := time.Parse(timeFormat, string(tbytes))
	if err != nil {
		msg := "Not able to parse time in " + dotKeepFilePath + "."
		log.WithError(err).Debug(msg)
		return errors.Wrap(err, msg)
	}
	timeDiff := time.Since(prevTime)
	if timeDiff < options.limit {
		msg := "Last submission was made " + humanize.Time(prevTime) + ". " +
			"Due to the rate limitter, submissions are not allows within a " + options.limit.String() +
			"  time window. "
		log.Debug(msg)
		return errors.New(msg)
	}
	com.WriteFile(dotKeepFilePath, []byte(time.Now().Format(timeFormat)))
	return nil
}
