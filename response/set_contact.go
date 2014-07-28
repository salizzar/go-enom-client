package response

import (
	"encoding/xml"
)

type SetContact struct {
	Result

	PendingVerification	bool	`xml:"PendingVerification" json:"pending_verification"`
}

func (response SetContact) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

