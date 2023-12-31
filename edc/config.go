package edc

import (
	"log"

	edchttp "github.com/Think-iT-Labs/edc-connector-client-go/edc/transport/http"
)

type Config struct {
	AuthToken  string
	Logger     *log.Logger
	HTTPClient *edchttp.HTTPClient
	Addresses
}

type Addresses struct {
	Control    *string
	Management *string
	Protocol   *string
	Public     *string
	Default    *string
}

// NewConfig returns a new Config pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
func NewConfig() *Config {
	return &Config{}
}

// Copy will return a shallow copy of the Config object. If any additional
// configurations are provided they will be merged into the new config returned.
func (c Config) Copy() Config {
	cp := c
	return cp
}
