package response

import (
	"encoding/xml"
)

type Domain struct {
	Name			string	`xml:"DomName,attr"              json:"name"`
	NSStatus		string	`xml:"nsstatus,attr"             json:"ns_status"`
	RegistrationStatus	string	`xml:"registrationstatus,attr"   json:"registration_status"`
	ExpirationDate		string	`xml:"ExpirationDate,attr"       json:"expiration_date"`
}

type Report struct {
	Result

	FirstName	string		`xml:"GetReport>FName"           json:"first_name"`
	LastName	string		`xml:"GetReport>LName"           json:"last_name"`
	Party		string		`xml:"GetReport>Party"           json:"party"`

	Domains		[]Domain	`xml:"GetReport>ReportDetail>dn" json:"domains"`
}

func (response Report) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

