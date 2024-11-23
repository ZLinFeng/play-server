package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

var (
	INFO  = "info"
	DEBUG = "debug"
)

type LogConfig struct {
	Dir           string `mapstructure:"log-dir"`
	Level         string `mapstructure:"level"`
	Type          string `mapstructure:"type"`
	FileSizeMb    int    `mapstructure:"file-size-mb"`
	RetentionDays int    `mapstructure:"retention-days"`
	FileLog       bool
	StdLog        bool
}

func (c *LogConfig) Check() error {
	if c.Type == "" {
		c.Type = "std"
	}
	parts := strings.Split(c.Type, ",")
	for _, part := range parts {
		part = strings.ToLower(strings.Trim(part, " "))
		if part != "file" && part != "std" {
			return errors.New(fmt.Sprintf("invalid log type: %s, only file, std are supported.", part))
		}

		if part == "file" {
			c.FileLog = true
		}
		if part == "std" {
			c.StdLog = true
		}
	}

	if c.Dir == "" {
		c.Dir = filepath.Join(getHome(), "logs")
	}
	// check dir
	if c.FileLog {
		if dir, err := os.Stat(c.Dir); err != nil {
			if os.IsNotExist(err) {
				createErr := os.MkdirAll(c.Dir, os.ModePerm)
				if createErr != nil {
					return errors.New("fatal error while create log dir.")
				}
			} else {
				return errors.New("fatal error while check log dir.")
			}
		} else {
			if !dir.IsDir() {
				return errors.New("log-dir must be a directory.")
			}
		}
	}

	c.Level = strings.ToLower(c.Level)

	switch c.Level {
	case INFO, DEBUG:

	default:
		fmt.Printf("only info and debug are supported.\n")
		fmt.Printf("%s is invalid log level, use default value: INFO", c.Level)
		c.Level = INFO
	}
	return nil
}
