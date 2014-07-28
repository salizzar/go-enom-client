package response

import (
	"encoding/xml"
)

type PasswordBit struct {
	Result

	DomainPassword	string	`xml:"DomainPassword" json:"domain_password"`
	PasswordSet	int	`xml:"password-set"   json:"password_set"`
}

func (response PasswordBit) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

