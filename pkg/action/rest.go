package action

type Method string

const (
	GET    = Method("get")
	POST   = Method("post")
	PUT    = Method("put")
	PATCH  = Method("patch")
	DELETE = Method("delete")
)

func RestRequest(url string, method Method, payload any) Action {

}
