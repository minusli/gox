package xerr

type err struct {
	Code int64
	Msg  string
}

func (e err) Error() string {
	return e.Msg
}

//goland:noinspection GoUnusedGlobalVariable
var (
	BadRequest          = err{Code: 400, Msg: "bad request"}
	Forbidden           = err{Code: 403, Msg: "forbidden"}
	NotFound            = err{Code: 404, Msg: "not found"}
	InternalServerError = err{Code: 500, Msg: "internal server error"}
)
