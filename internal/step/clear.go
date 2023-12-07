package step

import (
	"context"

	"github.com/dronestock/cos/internal/config"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type Clear struct {
	config *config.Wrapper
	cos    *cos.Client
}

func NewClear(config *config.Wrapper, cos *cos.Client) *Clear {
	return &Clear{
		config: config,
		cos:    cos,
	}
}

func (s *Clear) Runnable() bool {
	return true
}

func (s *Clear) Run(_ context.Context) (err error) {
	return
}
