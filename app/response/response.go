package response

type Response interface {
	GetMessage() string
	GetStatus() int
	GetData() interface{}
}
