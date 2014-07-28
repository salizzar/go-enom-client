package response

import (
	"encoding/xml"
)

type RRPInfo struct {
	Nameservers	[]string	`xml:"nameserver>nameserver"        json:"nameservers"`
	UpdatedDate	string		`xml:"updated-date"                 json:"updated_date"`
	CreatedDate	string		`xml:"created-date"                 json:"created_date"`
	ExpirationDate	string		`xml:"registration-expiration-date" json:"expiration_date"`
	Status		[]string	`xml:"status>status"                json:"statuses"`
}

type Contact struct {
	Type		string	`xml:"ContactType,attr"                     json:"type"`
	Organization	string	`xml:"Organization"                         json:"organization"`
	FirstName	string	`xml:"FName"                                json:"first_name"`
	LastName	string	`xml:"LName"                                json:"last_name"`
	Address		string	`xml:"Address1"                             json:"address_1"`
	AddressNumber	string	`xml:"Address2"                             json:"address_2"`
	City		string	`xml:"City"                                 json:"city"`
	State		string	`xml:"StateProvince"                        json:"state_province"`
	PostalCode	string	`xml:"PostalCode"                           json:"postal_code"`
	Country		string	`xml:"Country"                              json:"country"`
	PhoneNumber	string	`xml:"Phone"                                json:"phone"`
	PhoneExtension	string	`xml:"PhoneExt"                             json:"phone_extension"`
	FaxNumber	string	`xml:"Fax"                                  json:"fax"`
	Email		string	`xml:"EmailAddress"                         json:"email"`
}

type WhoisContact struct {
	Result

	Contacts	[]Contact	`xml:"GetWhoisContacts>contacts>contact" json:"contacts"`
	RRPInfo		RRPInfo		`xml:"GetWhoisContacts>rrp-info"         json:"rrp_info"`
	DomainExpired	bool		`xml:"GetWhoisContacts>DomainExpired"    json:"domain_expired"`
}

func (response WhoisContact) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

