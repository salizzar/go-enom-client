package response

import (
	"encoding/xml"
)

type TransferOrders struct {
	Result

}

func (response TransferOrders) UnmarshalXml(body []byte) (ResponseParser, error) {
	err := xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

