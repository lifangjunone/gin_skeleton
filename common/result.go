package common

var (
	Resp *Result
)

type Result struct {
	Code int
	Msg  string
	Data interface{}
}

func newResult() *Result {
	return &Result{}
}

func (r *Result) SuccessNotData() *Result {
	return &Result{
		Code: 10000,
		Msg:  "successful",
		Data: map[string]string{},
	}
}

func (r *Result) Success(data interface{}) *Result {
	return &Result{
		Code: 10000,
		Msg:  "successful",
		Data: data,
	}
}

func (r *Result) Error(err string) *Result {
	return &Result{
		Code: 10002,
		Msg:  err,
		Data: map[string]string{},
	}
}

func init() {
	Resp = newResult()
}
