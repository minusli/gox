package xerr

type XErr struct {
	Code int64
	Msg  string
}

func (e XErr) Error() string {
	return e.Msg
}

//goland:noinspection GoUnusedGlobalVariable
var (
	BadRequest          = XErr{Code: 400, Msg: "bad request"}
	Forbidden           = XErr{Code: 403, Msg: "forbidden"}
	NotFound            = XErr{Code: 404, Msg: "not found"}
	InternalServerError = XErr{Code: 500, Msg: "internal server error"}
)
