package enom

import (
	"net/url"

	"enom-client/response"
)

type Command struct {
	Name		string
	Args		url.Values
	Parser		response.ResponseParser
}

func (c *Command) ToParams() url.Values {
	return c.Args
}

func (c *Command) SetDefaultArgs(config *Config) {
	c.Args = url.Values{}

	c.Args.Set("UID", config.Username)
	c.Args.Set("PW", config.Password)
	c.Args.Set("ResponseType", config.ResponseType)
	c.Args.Set("Command", c.Name)
}

func (c *Command) AddParam(key string, value string) {
	c.Args.Set(key, value)
}

