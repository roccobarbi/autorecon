package network

type GetRequest interface {
	SetBaseUrl(string)
	SetQueryKeyValue(string, string)
	Request() []byte
}
