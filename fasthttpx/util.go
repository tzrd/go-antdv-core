package fasthttpx

import (
	"errors"
	"strings"

	"github.com/tzrd/go-antdv-core/logx"
	"github.com/valyala/fasthttp"
	"golang.org/x/text/language"

	enLang "github.com/go-playground/locales/en"
	zhLang "github.com/go-playground/locales/zh_Hans"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

const xForwardedFor = "X-Forwarded-For"

var supportLang map[string]string

// GetFormValues returns the form values.
func GetFormValues(ctx *fasthttp.RequestCtx) (map[string]any, error) {
	var args *fasthttp.Args
	method := string(ctx.Method())
	if method == "GET" {
		args = ctx.QueryArgs()
	} else if method == "POST" {
		args = ctx.PostArgs()
	} else {
		return nil, errors.New("Not get or post request")
	}

	params := make(map[string]any, args.Len())
	args.VisitAll(func(key, value []byte) {
		if len(value) > 0 {
			name := string(key)
			params[name] = string(value)
		}
	})

	return params, nil
}

// GetRemoteAddr returns the peer address, supports X-Forward-For.
func GetRemoteAddr(ctx *fasthttp.RequestCtx) string {
	v := ctx.Request.Header.Peek(xForwardedFor)
	if len(v) > 0 {
		return string(v)
	}

	return ctx.RemoteAddr().String()
}

type Validator struct {
	Validator *validator.Validate
	Uni       *ut.UniversalTranslator
	Trans     map[string]ut.Translator
}

func NewValidator() *Validator {
	v := Validator{}
	en := enLang.New()
	zh := zhLang.New()
	v.Uni = ut.New(zh, en, zh)
	v.Validator = validator.New()
	enTrans, _ := v.Uni.GetTranslator("en")
	zhTrans, _ := v.Uni.GetTranslator("zh")
	v.Trans = make(map[string]ut.Translator)
	v.Trans["en"] = enTrans
	v.Trans["zh"] = zhTrans
	// add support languages
	initSupportLanguages()

	err := enTranslations.RegisterDefaultTranslations(v.Validator, enTrans)
	if err != nil {
		logx.Errorw("register English translation failed", logx.Field("detail", err.Error()))
		return nil
	}
	err = zhTranslations.RegisterDefaultTranslations(v.Validator, zhTrans)
	if err != nil {
		logx.Errorw("register Chinese translation failed", logx.Field("detail", err.Error()))

		return nil
	}

	return &v
}

func ParseAcceptLanguage(lang string) (string, error) {
	tags, _, err := language.ParseAcceptLanguage(lang)
	if err != nil {
		return "", errors.New("fail to parse accept language")
	}

	for _, v := range tags {
		if val, ok := supportLang[v.String()]; ok {
			return val, nil
		}
	}

	return "zh", nil
}

func (v *Validator) Validate(data any, lang string) string {
	err := v.Validator.Struct(data)
	if err == nil {
		return ""
	}

	targetLang, parseErr := ParseAcceptLanguage(lang)
	if parseErr != nil {
		return parseErr.Error()
	}

	errs, ok := err.(validator.ValidationErrors)

	if ok {
		transData := errs.Translate(v.Trans[targetLang])
		s := strings.Builder{}
		for _, v := range transData {
			s.WriteString(v)
			s.WriteString(" ")
		}
		return s.String()
	}

	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		return invalid.Error()
	}

	return ""
}

func initSupportLanguages() {
	supportLang = make(map[string]string)
	supportLang["zh"] = "zh"
	supportLang["zh-CN"] = "zh"
	supportLang["en"] = "en"
	supportLang["en-US"] = "en"
}

// RegisterValidation registers the validation function to validator
func RegisterValidation(tag string, fn validator.Func) {
	if err := xValidator.Validator.RegisterValidation(tag, fn); err != nil {
		logx.Must(errors.Join(err, errors.New("failed to register the validation function, tag is "+tag)))
	}
}

// RegisterValidationTranslation regiters the validation translation for validator
func RegisterValidationTranslation(tag string, trans ut.Translator, registerFn validator.RegisterTranslationsFunc,
	translationFn validator.TranslationFunc) {
	if err := xValidator.Validator.RegisterTranslation(tag, trans, registerFn, translationFn); err != nil {
		logx.Must(errors.Join(err, errors.New("failed to register the validation translation, tag is "+tag)))
	}
}
