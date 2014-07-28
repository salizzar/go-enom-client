package response

import (
	"encoding/xml"
)

type PDomainInfo struct {
	RegistryCreateDate	string		`xml:"RegistryCreateDate" json:"registry_create_date"`
	RegistryExpDate		string		`xml:"RegistryExpDate"    json:"registry_expiration_date"`
}

type Purchase struct {
	Result

	OrderID			int		`xml:"OrderID"            json:"order_id"`
	DomainInfo		PDomainInfo	`xml:"DomainInfo"         json:"domain_info"`
	TotalCharged		float32		`xml:"TotalCharged"       json:"total_charged"`
	RegistrantPartyID	string		`xml:"RegistrantPartyID"  json:"registrant_party_id"`
}

func (response Purchase) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

