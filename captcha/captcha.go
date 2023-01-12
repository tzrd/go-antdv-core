package captcha

import (
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

// configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

// SetStore 设置store
func SetStore(s base64Captcha.Store) {
	base64Captcha.DefaultMemStore = s
}

func DriverDigitFunc(height int, width int, length int, maxSkew float64, dotCount int) (id, b64s string, err error) {
	e := configJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.NewDriverDigit(height, width, length, maxSkew, dotCount)
	driver := e.DriverDigit
	cap := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	return cap.Generate()
}

// Verify 校验验证码
func Verify(id, code string, clear bool) bool {
	return base64Captcha.DefaultMemStore.Verify(id, code, clear)
}
