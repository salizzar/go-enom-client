package response

type ResponseParser interface {
	UnmarshalXml(body []byte) (ResponseParser, error)
	UnmarshalCfg(body []byte) (ResponseParser, error)
}

