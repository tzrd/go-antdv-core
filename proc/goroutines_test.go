package proc

import (
	"strings"
	"testing"

	"github.com/tzrd/go-antdv-core/logx"

	"github.com/stretchr/testify/assert"
)

func TestDumpGoroutines(t *testing.T) {
	var buf strings.Builder
	w := logx.NewWriter(&buf)
	o := logx.Reset()
	logx.SetWriter(w)
	defer func() {
		logx.Reset()
		logx.SetWriter(o)
	}()

	dumpGoroutines()
	assert.True(t, strings.Contains(buf.String(), ".dump"))
}
