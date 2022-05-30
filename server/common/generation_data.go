package common

type Data struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func NewData(code int, data interface{}, msg string) Data {
	return Data{
		code,
		data,
		msg,
	}
}
