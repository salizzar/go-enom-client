package response

import (
	"encoding/xml"
)

type RegisterNameserver struct {
	Result

	NsSuccess	int	`xml:"RegisterNameserver>NsSuccess" json:"ns_success"`
}

func (response RegisterNameserver) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

