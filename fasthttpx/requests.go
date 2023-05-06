package fasthttpx

import (
	"strings"

	"github.com/tzrd/go-antdv-core/errorx"
	"github.com/tzrd/go-antdv-core/fasthttpx/pathvar"
	"github.com/tzrd/go-antdv-core/mapping"
	"github.com/valyala/fasthttp"
)

// @enhance
const (
	formKey           = "form"
	pathKey           = "path"
	maxMemory         = 32 << 20 // 32MB
	maxBodyLen        = 8 << 20  // 8MB
	separator         = ";"
	tokensInAttribute = 2
)

var (
	formUnmarshaler = mapping.NewUnmarshaler(formKey, mapping.WithStringValues())
	pathUnmarshaler = mapping.NewUnmarshaler(pathKey, mapping.WithStringValues())
	xValidator      = NewValidator()
)

// Parse parses the request.
func Parse(ctx *fasthttp.RequestCtx, v any, isValidate bool) error {
	r := &ctx.Request
	if err := ParseJsonBody(r, v); err != nil {
		return err
	}

	if err := ParsePath(ctx, v); err != nil {
		return err
	}

	if err := ParseForm(ctx, v); err != nil {
		return err
	}

	if err := ParseHeaders(r, v); err != nil {
		return err
	}

	if isValidate {
		if errMsg := xValidator.Validate(v, string(r.Header.Peek("Accept-Language"))); errMsg != "" {
			return errorx.NewCodeInvalidArgumentError(errMsg)
		}
	}
	return nil
}

// ParseForm parses the form request.
func ParseForm(ctx *fasthttp.RequestCtx, v any) error {
	params, err := GetFormValues(ctx)
	if err != nil {
		return err
	}

	return formUnmarshaler.Unmarshal(params, v)
}

// ParseHeaders parses the headers request.
func ParseHeaders(r *fasthttp.Request, v any) error {
	v = r.Header

	return nil
}

// ParseHeader parses the request header and returns a map.
func ParseHeader(headerValue string) map[string]string {
	ret := make(map[string]string)
	fields := strings.Split(headerValue, separator)

	for _, field := range fields {
		field = strings.TrimSpace(field)
		if len(field) == 0 {
			continue
		}

		kv := strings.SplitN(field, "=", tokensInAttribute)
		if len(kv) != tokensInAttribute {
			continue
		}

		ret[kv[0]] = kv[1]
	}

	return ret
}

// ParseJsonBody parses the post request which contains json in body.
func ParseJsonBody(r *fasthttp.Request, v any) error {
	if withJsonBody(r) {
		return mapping.UnmarshalJsonBytes(r.Body(), v)
	}

	return mapping.UnmarshalJsonMap(nil, v)
}

func withJsonBody(r *fasthttp.Request) bool {
	return r.Header.ContentLength() > 0 && strings.Contains(string(r.Header.ContentType()), "application/json")
}

// ParsePath parses the symbols reside in url path.
// Like http://localhost/bag/:name
func ParsePath(ctx *fasthttp.RequestCtx, v any) error {
	vars := pathvar.Vars(ctx)
	m := make(map[string]any, len(vars))
	for k, v := range vars {
		m[k] = v
	}

	return pathUnmarshaler.Unmarshal(m, v)
}
