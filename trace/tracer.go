package trace

import (
	"fmt"
	"io"
)

// Tracerはコード内での出来事を記録するオブジェクトを表すインターフェース
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
