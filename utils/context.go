package utils

import (
	"context"
	"time"
)

func GetContext(second int) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(second))
	return ctx
}

func GetDefaultContext() context.Context {
	return GetContext(5)
}
