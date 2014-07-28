package enom

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"

	"github.com/salizzar/go-enom-client/response"
)

type Client interface {
	CommandFactory() *CommandFactory
	Execute(command *Command) (response.ResponseParser, error)
	DoRequest(comman *Command) ([]byte, error)
}

type EnomClient struct {
	Config *Config;
}

func (c EnomClient) CommandFactory() *CommandFactory {
	factory := &CommandFactory{ Config: c.Config }

	return factory
}

func (c EnomClient) Execute(command *Command) (response.ResponseParser, error) {
	body, err := c.DoRequest(command)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute Enom request: %s", err)
	}

	factory := &ResponseFactory{ ResponseType: c.Config.ResponseType }
	response, err := factory.GetResponse(body, command)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse Enom response: %s", err)
	}

	return response, nil
}

func (c EnomClient) DoRequest(command *Command) ([]byte, error) {
	data := command.ToParams()

	log.Print("URL to be called: ", c.Config.Url)
	log.Print("URL POST params: ", data)

	resp, err := http.PostForm(c.Config.Url, data)
	if err != nil {
		log.Fatal("Post Failed: ", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read: ", err)
		return nil, err
	}

	log.Print(resp.Status)
	log.Print(fmt.Printf("%s\n", body))

	return body, nil
}

