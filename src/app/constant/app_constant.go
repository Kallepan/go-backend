package constant

type ResponseStatus int
type Headers int
type General int

// Constant API
const (
	Success ResponseStatus = iota + 1
	InvalidRequest
	Unauthorized
	DataNotFound
	UnknownError
)

func (r ResponseStatus) GetResponseStatus() int {
	return [...]int{
		200,
		400,
		401,
		404,
		500,
	}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{
		"Success",
		"Invalid request",
		"Unauthorized",
		"Data not found",
		"Unknown error",
	}[r-1]
}
