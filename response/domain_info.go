package response

import (
	"encoding/xml"
)

type Registry struct {
	PartyId			string		`xml:"party-id,attr"                json:"party_id"`
	Name			string		`xml:",chardata"                    json:"name"`
}

type Status struct {
	Expiration		string		`xml:"expiration"                   json:"expiration"`
	DeleteByDate		string		`xml:"deletebydate"                 json:"delete_by_date"`
	DeleteType		string		`xml:"deletetype"                   json:"delete_type"`
	Restorable		bool		`xml:"restorable"                   json:"restorable"`
	RenewBeforeExpiration	string		`xml:"renewbeforeexpiration"        json:"renew_before_expiration"`
	Registrar		string		`xml:"registrar"                    json:"registrar"`
	RegistrationStatus	string		`xml:"registrationstatus"           json:"registration_status"`
	PurchaseStatus		string		`xml:"purchase-status"              json:"purchase_status"`
	Registry		Registry	`xml:"belongs-to"                   json:"registry"`
	EscrowHold		string		`xml:"escrowhold"                   json:"escrow_hold"`
	EscrowLiftDate		string		`xml:"escrowliftdate"               json:"escrow_lift_date"`
	AuctionHold		bool		`xml:"auctionhold"                  json:"auction_hold"`
	AuctionLiftDate		string		`xml:"auctionliftdate"              json:"auction_lift_date"`
}

type Entry struct {
	Name			string		`xml:"name,attr"                    json:"name"`
}

type DomainName struct {
	Sld			string		`xml:"sld,attr"                     json:"sld"`
	Tld			string		`xml:"tld,attr"                     json:"tld"`
	DomainNameId		int32		`xml:"domainnameid,attr"            json:"domain_name_id"`
	DomainId		string		`xml:",chardata"                    json:"domain_id"`
}

type DomainInfo struct {
	Result

	MultyLangSLD		string		`xml:"GetDomainInfo>multy-langSLD"  json:"multy_lang_sld"`
	DomainName		DomainName	`xml:"GetDomainInfo>domainname"     json:"domain_name"`
	Status			Status		`xml:"GetDomainInfo>status"         json:"status"`
	Services		[]Entry		`xml:"GetDomainInfo>services>entry" json:"services"` // TODO: partially implemented because entry tag returns a lot of different types

	ParkingEnabled		bool		`xml:"GetDomainInfo>ParkingEnabled" json:"parking_enabled"`
}

func (response DomainInfo) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

