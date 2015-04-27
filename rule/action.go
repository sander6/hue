package rule

type (
	Action struct {
		Address string
		Method  Method
		Body    interface{}
	}

	Method string
)

const (
	GET Method = "GET"
	PUT        = "PUT"
)
