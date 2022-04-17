package factories

import (
	"context"
	"go-todos/internal/utils/config"
	"time"
)

func NewContext() (context.Context, context.CancelFunc) {
	conf := config.New()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.AppTimeout)*time.Second)
	return ctx, cancel
}
