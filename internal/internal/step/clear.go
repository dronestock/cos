package step

import (
	"context"

	"github.com/dronestock/cos/internal/internal/config"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type Clear struct {
	source *config.Source
	cos    *cos.Client
}

func NewClear(source *config.Source, cos *cos.Client) *Clear {
	return &Clear{
		source: source,
		cos:    cos,
	}
}

func (s *Clear) Runnable() bool {
	return true
}

func (s *Clear) Run(_ *context.Context) (err error) {
	return
}
