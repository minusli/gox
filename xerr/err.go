package xerr

type Error struct {
	Code int64
	Msg  string
}

func (e Error) Error() string {
	return e.Msg
}

//goland:noinspection GoUnusedGlobalVariable
var (
	BadRequest          = Error{Code: 400, Msg: "bad request"}
	Forbidden           = Error{Code: 403, Msg: "forbidden"}
	NotFound            = Error{Code: 404, Msg: "not found"}
	InternalServerError = Error{Code: 500, Msg: "internal server error"}
)
