package response

import (
	"encoding/xml"
)

type Delete struct {
	Result

	DomainDeleted	bool	`xml:"deletedomain>domaindeleted" json:"domain_deleted"`
}

func (response Delete) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

