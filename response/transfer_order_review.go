package response

import (
	"encoding/xml"
)

type TransferOrderDetailed struct {
	DetailId			int			`xml:"transferorderdetailid" json:"detail_id"`
	Sld				string			`xml:"sld"                   json:"sld"`
	Tld				string			`xml:"tld"                   json:"tld"`
	Price				float32			`xml:"price"                 json:"price"`
	AbleToLock			bool			`xml:"abletolock"            json:"able_to_lock"`
	Lock				bool			`xml:"lock"                  json:"lock"`
	Renew				bool			`xml:"renew"                 json:"renew"`
	DomainPassword			string			`xml:"domainpassword"        json:"domain_password"`
	UseContacts			int			`xml:"usecontacts"           json:"use_contacts"`
	RequestAuthInfo			bool			`xml:"reqauthinfo"           json:"request_auth_info"`
	AuthInfo			string			`xml:"authinfo"              json:"auth_info"`
	RRProcessor			string			`xml:"RRProcessor"           json:"rr_processor"`
	RegistrantPartyID		string			`xml:"RegistrantPartyID"     json:"registrant_party_id"`
	RegistrantROID			string			`xml:"RegistrantROID"        json:"registrant_roid"`
	PremiumDomain			bool			`xml:"PremiumDomain"         json:"premium_domain"`
}

type TransferOrderReview struct {
	Result

	TransferOrderDetailed		TransferOrderDetailed	`xml:"transferorderreview>transferorderdetail" json:"transfer_order_detail"`

	TransferOrderDetailCount	int			`xml:"transferorderdetailcount"                json:"transfer_order_detail_count"`
	TransferOrderDetailEuCount	int			`xml:"transferorderdetaileucount"              json:"transfer_order_detail_eu_count"`
	TransferOrderDetailCaCount	int			`xml:"transferorderdetailcacount"              json:"transfer_order_detail_ca_count"`
	TransferOrderDetailDeCount	int			`xml:"transferorderdetaildecount"              json:"transfer_order_detail_de_count"`
	TransferOrderDetailBeCount	int			`xml:"transferorderdetailbecount"              json:"transfer_order_detail_be_count"`
	TransferTotalPrice		float32			`xml:"transfertotalprice"                      json:"transfer_total_price"`
	AuthInfoStillRequired		string			`xml:"authinfostillrequired"                   json:"auth_info_still_required"`
	EuInfoStillRequired		string			`xml:"euinfostillrequired"                     json:"eu_info_still_required"`
	CaInfoStillRequired		string			`xml:"cainfostillrequired"                     json:"ca_info_still_required"`
	DeInfoStillRequired		string			`xml:"deinfostillrequired"                     json:"de_info_still_required"`
	BeInfoStillRequired		string			`xml:"beinfostillrequired"                     json:"be_info_still_required"`
}

func (response TransferOrderReview) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

