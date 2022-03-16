package resp

type Entity struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Total     int         `json:"total"`
	TotalPage int         `json:"total_page"`
	Data      interface{} `json:"data"`
}