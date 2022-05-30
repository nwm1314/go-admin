package request

type Login struct {
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}


