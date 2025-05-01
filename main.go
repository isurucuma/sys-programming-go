package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type CliConfig struct {
	OutputWriter io.Writer
	ErrorWriter  io.Writer
}

func NewCliConfig(opts ...Option) (*CliConfig, error) {
	c := CliConfig{
		OutputWriter: os.Stdout,
		ErrorWriter:  os.Stderr,
	}

	for _, opt := range opts {
		err := opt(&c)
		if err != nil {
			return nil, err
		}
	}
	return &c, nil
}

type Option func(*CliConfig) error

func CliConfigWithOutputWriter(w io.Writer) Option {
	return func(c *CliConfig) error {
		c.OutputWriter = w
		return nil
	}
}

func CliConfigWithErrorWriter(w io.Writer) Option {
	return func(c *CliConfig) error {
		c.ErrorWriter = w
		return nil
	}
}

func (c *CliConfig) PrintArgs(args []string) error {
	if len(args) == 0 {
		return errors.New("no arguments specified")
	}

	for _, w := range args {
		if len(w)%2 == 0 {
			// this is even therefore transfer this into the standard output
			_, _ = fmt.Fprintln(c.OutputWriter, w)
		} else {
			// this has odd number of letters, then transfer it to standard error
			_, _ = fmt.Fprintln(c.ErrorWriter, w)
		}
	}
	return nil
}

func main() {
	c, err := NewCliConfig(CliConfigWithOutputWriter(os.Stdout), CliConfigWithErrorWriter(os.Stderr))
	if err != nil {
		panic(err)
	}

	if err := c.PrintArgs(os.Args[1:]); err != nil {
		panic(err)
	}
}
