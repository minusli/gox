package xerr

//goland:noinspection GoUnusedGlobalVariable
var (
	BadRequest          = &Error{code: 400, msg: "bad request"}
	Forbidden           = &Error{code: 403, msg: "forbidden"}
	NotFound            = &Error{code: 404, msg: "not found"}
	InternalServerError = &Error{code: 500, msg: "internal server error"}
)

type Error struct {
	code int64
	msg  string
}

func (e *Error) Code() int64 {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) With(msg string) *Error {
	return &Error{
		code: e.code,
		msg:  msg,
	}
}
