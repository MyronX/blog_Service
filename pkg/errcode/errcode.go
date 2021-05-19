package errcode

type Error struct {
	code	int			`json:"code"`
	msg		string		`json:"msg"`
	details	[]string	`json:""details""`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code];ok {
		panic("fmt.")
	}

	return nil
}
