package response

import (
	"encoding/xml"
)

type Registrant struct {
	OrganizationName	string	`xml:"RegistrantOrganizationName" json:"organization_name"`
	JobTitle		string	`xml:"RegistrantJobTitle"         json:"job_title"`
	FirstName		string	`xml:"RegistrantFirstName"        json:"first_name"`
	LastName		string	`xml:"RegistrantLastName"         json:"last_name"`
	Address1		string	`xml:"RegistrantAddress1"         json:"address_1"`
	Address2		string	`xml:"RegistrantAddress2"         json:"address_2"`
	City			string	`xml:"RegistrantCity"             json:"city"`
	Country			string	`xml:"RegistrantCountry"          json:"country"`
	StateProvince		string	`xml:"RegistrantStateProvince"    json:"state_provice"`
	PostalCode		string	`xml:"RegistrantPostalCode"       json:"postal_code"`
	Email			string	`xml:"RegistrantEmailAddress"     json:"email"`
	Phone			string	`xml:"RegistrantPhone"            json:"phone"`
	PhoneExtension		string	`xml:"RegistrantPhoneExt"         json:"phone_extension"`
	Fax			string	`xml:"RegistrantFax"              json:"fax"`
}

type AuxiliarBilling struct {
	OrganizationName	string	`xml:"AuxBillingOrganizationName" json:"organization_name"`
	JobTitle		string	`xml:"AuxBillingJobTitle"         json:"job_title"`
	FirstName		string	`xml:"AuxBillingFirstName"        json:"first_name"`
	LastName		string	`xml:"AuxBillingLastName"         json:"last_name"`
	Address1		string	`xml:"AuxBillingAddress1"         json:"address_1"`
	Address2		string	`xml:"AuxBillingAddress2"         json:"address_2"`
	City			string	`xml:"AuxBillingCity"             json:"city"`
	Country			string	`xml:"AuxBillingCountry"          json:"country"`
	StateProvince		string	`xml:"AuxBillingStateProvince"    json:"state_province`
	PostalCode		string	`xml:"AuxBillingPostalCode"       json:"postal_code"`
	Email			string	`xml:"AuxBillingEmailAddress"     json:"email"`
	Phone			string	`xml:"AuxBillingPhone"            json:"phone"`
	PhoneExtension		string	`xml:"AuxBillingPhoneExt"         json:"phone_extension"`
	Fax			string	`xml:"AuxBillingFax"              json:"fax"`
}

type Technical struct {
	OrganizationName	string	`xml:"TechOrganizationName"       json:"organization_name"`
	JobTitle		string	`xml:"TechJobTitle"               json:"job_title"`
	FirstName		string	`xml:"TechFirstName"              json:"first_name"`
	LastName		string	`xml:"TechLastName"               json:"last_name"`
	Address1		string	`xml:"TechAddress1"               json:"address_1"`
	Address2		string	`xml:"TechAddress2"               json:"address_2"`
	City			string	`xml:"TechCity"                   json:"city"`
	Country			string	`xml:"TechCountry"                json:"country"`
	StateProvince		string	`xml:"TechStateProvince"          json:"state_province"`
	PostalCode		string	`xml:"TechPostalCode"             json:"postal_code"`
	Email			string	`xml:"TechEmailAddress"           json:"email"`
	Phone			string	`xml:"TechPhone"                  json:"phone"`
	PhoneExtension		string	`xml:"TechPhoneExt"               json:"phone_extension"`
	Fax			string	`xml:"TechFax"                    json:"fax"`
}

type Administrator struct {
	OrganizationName	string	`xml:"AdminOrganizationName"      json:"organization_name"`
	JobTitle		string	`xml:"AdminJobTitle"              json:"job_title"`
	FirstName		string	`xml:"AdminFirstName"             json:"first_name"`
	LastName		string	`xml:"AdminLastName"              json:"last_name"`
	Address1		string	`xml:"AdminAddress1"              json:"address_1"`
	Address2		string	`xml:"AdminAddress2"              json:"address_2"`
	City			string	`xml:"AdminCity"                  json:"city"`
	Country			string	`xml:"AdminCountry"               json:"country"`
	StateProvince		string	`xml:"AdminStateProvince"         json:"state_province"`
	PostalCode		string	`xml:"AdminPostalCode"            json:"postal_code"`
	Email			string	`xml:"AdminEmailAddress"          json:"email"`
	Phone			string	`xml:"AdminPhone"                 json:"phone"`
	PhoneExtension		string	`xml:"AdminPhoneExt"              json:"phone_extension"`
	Fax			string	`xml:"AdminFax"                   json:"fax"`
}

type Billing struct {
	OrganizationName	string	`xml:"BillingOrganizationName"    json:"organization_name"`
	JobTitle		string	`xml:"BillingJobTitle"            json:"job_title"`
	FirstName		string	`xml:"BillingFirstName"           json:"first_name"`
	LastName		string	`xml:"BillingLastName"            json:"last_name"`
	Address1		string	`xml:"BillingAddress1"            json:"address_1"`
	Address2		string	`xml:"BillingAddress2"            json:"address_2"`
	City			string	`xml:"BillingCity"                json:"city"`
	Country			string	`xml:"BillingCountry"             json:"country"`
	StateProvince		string	`xml:"BillingStateProvince"       json:"state_province"`
	PostalCode		string	`xml:"BillingPostalCode"          json:"postal_code"`
	Email			string	`xml:"BillingEmailAddress"        json:"email"`
	Phone			string	`xml:"BillingPhone"               json:"phone"`
	PhoneExtension		string	`xml:"BillingPhoneExt"            json:"phone_extension"`
	Fax			string	`xml:"BillingFax"                 json:"fax"`
}

type WPPSContactData struct {
	ContactType		string	`xml:"ContactType"                json:"contact_type"`
	OrganizationName	string	`xml:"OrganizationName"           json:"organization_name"`
	JobTitle		string	`xml:"JobTitle"                   json:"job_title"`
	FirstName		string	`xml:"FirstName"                  json:"first_name"`
	LastName		string	`xml:"LastName"                   json:"last_name"`
	Address1		string	`xml:"Address1"                   json:"address_1"`
	Address2		string	`xml:"Address2"                   json:"address_2"`
	City			string	`xml:"City"                       json:"city"`
	Country			string	`xml:"Country"                    json:"country"`
	StateProvince		string	`xml:"StateProvince"              json:"state_province"`
	PostalCode		string	`xml:"PostalCode"                 json:"postal_code"`
	Email			string	`xml:"EmailAddress"               json:"email"`
	Phone			string	`xml:"Phone"                      json:"phone"`
	PhoneExtension		string	`xml:"PhoneExt"                   json:"phone_extension"`
	Fax			string	`xml:"Fax"                        json:"fax"`
	DateTimeChanged		string	`xml:"DateTimeChanged"            json:"date_time_changed"`
}

type Contacts struct {
	Result

	Registrant		Registrant	`xml:"GetContacts>Registrant"          json:"registrant"`
	AuxiliarBilling		AuxiliarBilling	`xml:"GetContacts>AuxBilling"          json:"auxiliar_billing"`
	Technical		Technical	`xml:"GetContacts>Tech"                json:"technical"`
	Administrator		Administrator	`xml:"GetContacts>Admin"               json:"administrator"`
	Billing			Billing		`xml:"GetContacts>Billing"             json:"billing"`

	PendingVerification	bool		`xml:"GetContacts>PendingVerification" json:"pending_verification"`
	WPPSAllowed		int		`xml:"GetContacts>WPPSAllowed"         json:"wpps_allowed"`
	WPPSExists		int		`xml:"GetContacts>WPPSExists"          json:"wpps_exists"`
	WPPSEnabled		int		`xml:"GetContacts>WPPSEnabled"         json:"wpps_enabled"`
	WPPSExpirationDate	string		`xml:"GetContacts>WPPSExpDate"         json:"wpps_expiration_date"`
	WPPSAutoRenew		string		`xml:"GetContacts>WPPSAutoRenew"       json:"wpps_auto_renew"`

	WPPSContactData		WPPSContactData	`xml:"GetContacts>WPPSContactData"     json:"wpps_contact_data"`
}

func (response Contacts) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)

	return response, err
}

