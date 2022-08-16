package startthread

import (
  "github.com/konapun/flozart/flow"
)

type Flow struct {

}

func (f *Flow) Start(ctx *flow.Context) {
  go func() {
    ctx.Next()
  }()
}

func (f *Flow) Build() flow.Flow {
  return &Flow{}

}
