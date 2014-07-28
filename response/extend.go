package response

import (
	"encoding/xml"
)

type EDomainInfo struct {
	RegistryExpDate	string		`xml:"RegistryExpDate" json:"registry_expiration_date"`
}

type Extend struct {
	Result

	Extension	string		`xml:"Extension"  json:"extension"`
	DomainName	string		`xml:"DomainName" json:"domain_name"`
	OrderID		int		`xml:"OrderID"    json:"order_id"`

	DomainInfo	EDomainInfo	`xml:"DomainInfo" json:"domain_info"`
}

func (response Extend) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
