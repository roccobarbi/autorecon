package network

type GetRequest[T any] interface {
	SetBaseUrl(string)
	SetQueryKeyValue(string, string)
	Request() []T
}
