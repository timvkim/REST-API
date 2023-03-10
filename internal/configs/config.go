package configs

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

type Config struct {
	Port   int    `long:"port" env:"HTTP_PORT" description:"HTTP port" required:"false" default:"8000"`
	DB_URL string `long:"db" env:"DATABASE" description:"db url" required:"false" default:""`
}

func New() (*Config, error) {
	defer os.Clearenv()

	c := &Config{}
	p := flags.NewParser(c, flags.Default|flags.IgnoreUnknown)

	if _, err := p.Parse(); err != nil {
		return nil, fmt.Errorf("error parsing config oprions %w", err)
	}

	return c, nil
}
