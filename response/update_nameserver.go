package response

import (
	"encoding/xml"
)

type UpdateNameserver struct {
	Result

	NsSuccess	int	`xml:"RegisterNameserver>NsSuccess" json:"ns_success"`
}

func (response UpdateNameserver) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

