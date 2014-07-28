package response

import (
	"encoding/xml"
)

type RegLock struct {
	Result

	RRPCode		int	`xml:"RRPCode" json:"rrp_code"`
	RRPText		string	`xml:"RRPText" json:"rrp_text"`
}

func (response RegLock) Success() bool {
	return response.RRPCode == 200
}

func (response RegLock) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

