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
		log.Debugf("Unable to find temporary directory for %s.", config.App.Name)
		return nil
	}
	dotKeepFilePath := filepath.Join(tmpDir, config.App.Name+".keep")
	if !com.IsFile(dotKeepFilePath) {
		com.WriteFile(dotKeepFilePath, []byte(time.Now().Format(timeFormat)))
		return nil
	}
	tbytes, err := ioutil.ReadFile(dotKeepFilePath)
	if err != nil {
		log.WithError(err).Debugf("Unable to read %s.", dotKeepFilePath)
		return nil
	}
	prevTime, err := time.Parse(timeFormat, string(tbytes))
	if err != nil {
		log.WithError(err).Debugf("Unable to parse time in %s.", dotKeepFilePath)
		return nil
	}
	timeDiff := time.Since(prevTime)
	if timeDiff < options.limit {
		msg := "Last submission was made " + humanize.Time(prevTime) + ". " +
			"Due to the rate limit, submissions are not allows within a " + options.limit.String() +
			"  time window. "
		log.Debug(msg)
		return errors.New(msg)
	}
	com.WriteFile(dotKeepFilePath, []byte(time.Now().Format(timeFormat)))
	return nil
}
