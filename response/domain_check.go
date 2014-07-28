package response

import (
	"encoding/xml"
)

type DomainCheck struct {
	Result

	ResponseCode	int	`xml:"RRPCode" json:"response_code"`
	Available	string	`xml:"RRPText" json:"available"`
}

func (response DomainCheck) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
