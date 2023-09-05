package constant

type ResponseStatus int
type Headers int
type General int

// Constant API
const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
)

func (r ResponseStatus) GetResponseStatus() int {
	return [...]int{200, 404, 500, 400, 401}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data not found", "Unknown error", "Invalid request", "Unauthorized"}[r-1]
}
