package factories

import (
	"context"
	"go-todos/internal/utils/config"
	"time"
)

func NewContext() (context.Context, context.CancelFunc) {
	/*
		Creates a context with a timeout to tell blocking functions that it should  abandon its work after the timeout elapses.
		Returns a new context
	*/
	// Pull configuration: for this case pull appTimeout from conf.AppTimout
	conf := config.New()
	// Create a context with a timeout, context.WithTimeout returns a new context and a cancel function
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.AppTimeout)*time.Second)
	return ctx, cancel
}
