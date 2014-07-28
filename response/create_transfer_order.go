package response

import (
	"encoding/xml"
)

type TPDetail struct {
	TransferOrderDetailID	int		`xml:"transferorderdetailid"             json:"transfer_order_detail_id"`
	Sld			string		`xml:"sld"                               json:"sld"`
	Tld			string		`xml:"tld"                               json:"tld"`
	StatusID		int		`xml:"statusid"                          json:"status_id"`
	StatusDesc		string		`xml:"statusdesc"                        json:"status_description"`
	Price			float32		`xml:"price"                             json:"price"`
	UseContacts		int		`xml:"usecontacts"                       json:"use_contacts"`
	PremiumDomain		string		`xml:"premiumdomain"                     json:"premium_domain"`
	CustomerSuppliedPrice	string		`xml:"customersuppliedprice"             json:"customer_supplied_price"` // TODO: float32 raises error with empty value
}

type CreateTransferOrder struct {
	Result

	TransferOrderID		int		`xml:"transferorder>transferorderid"     json:"transfer_order_id"`
	OrderDate		string		`xml:"transferorder>orderdate"           json:"order_date"`
	OrderTypeID		int		`xml:"transferorder>ordertypeid"         json:"order_type_id"`
	OrderTypeDesc		string		`xml:"transferorder>ordertypedesc"       json:"order_type_description"`
	StatusID		int		`xml:"transferorder>statusid"            json:"status_id"`
	StatusDesc		string		`xml:"transferorder>statusdesc"          json:"status_description"`
	AuthAmount		float32		`xml:"transferorder>authamount"          json:"auth_amount"`
	Version			int		`xml:"transferorder>version"             json:"version"`

	TransferOrderDetail	TPDetail	`xml:"transferorder>transferorderdetail" json:"transfer_order_detail"`
}

func (response CreateTransferOrder) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

