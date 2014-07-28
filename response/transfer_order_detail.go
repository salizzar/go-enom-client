package response

import (
	"encoding/xml"
)

type TODContacts struct {
	Registrant	string		`xml:"Registrant" json:"registrant"`
	AuxBilling	string		`xml:"AuxBilling" json:"auxiliar_billing"`
	Tech		string		`xml:"Tech"       json:"technical"`
	Admin		string		`xml:"Admin"      json:"administrator"`
	Billing		string		`xml:"Billing"    json:"billing"`
}

type TransferOrderDetail struct {
	Result

	Id		int		`xml:"transferorderdetail>transferorderdetailid" json:"id"`
	Sld		string		`xml:"transferorderdetail>sld"                   json:"sld"`
	Tld		string		`xml:"transferorderdetail>tld"                   json:"tld"`
	Lock		bool		`xml:"transferorderdetail>lock"                  json:"lock"`
	Renew		bool		`xml:"transferorderdetail>renew"                 json:"renew"`
	DomainPassword	string		`xml:"transferorderdetail>domainpassword"        json:"domain_password"`
	StatusId	int		`xml:"transferorderdetail>statusid"              json:"status_id"`
	StatusDesc	string		`xml:"transferorderdetail>statusdesc"            json:"status_description"`
	Price		float32		`xml:"transferorderdetail>price"                 json:"price"`
	UseContacts	int		`xml:"transferorderdetail>usecontacts"           json:"use_contacts"`
	Contacts	TODContacts	`xml:"transferorderdetail>contacts"              json:"contacts"`
}

func (response TransferOrderDetail) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

