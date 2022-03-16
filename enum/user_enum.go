package enum

type ResponseType int

const (
	OperateOk ResponseType = 200
	OpreateFail ResponseType = 500
)

func (p ResponseType) String() string {
	switch p {
	case OperateOk:
		return "ok"
	case OpreateFail:
		return "fail"
	default:
		return "unknown"
	}
}