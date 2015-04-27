package rule

type Operator string

const (
	Equals      Operator = "eq"
	GreaterThan          = "gt"
	LessThan             = "lt"
	Changed              = "dx"
)

type Condition struct {
	Address  string
	Operator Operator
	Value    uint32
}
