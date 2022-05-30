package response

type CaptchaResp struct {
	CaptchaId     string `json:"captchaId"`
	CaptchaLength int    `json:"captchaLength"`
	PicPath       string `json:"picPath"`
}
