package enom

import (
	"enom-client/response"
)

type ResponseFactory struct {
	ResponseType string
}

func (f ResponseFactory) GetResponse(body []byte, command *Command) (response.ResponseParser, error) {
	parser := command.Parser

	if f.ResponseType == "XML" {
		response, err := parser.UnmarshalXml(body)

		if err != nil {
			return nil, err
		}

		return response, nil
	} else {
		// text not implemented yet
		return nil, nil
	}
}

