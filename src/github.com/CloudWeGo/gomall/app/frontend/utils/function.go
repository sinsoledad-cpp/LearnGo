package utils

import (
	"context"
)

func GetUserIdFromCtx(ctx context.Context) int64 {
	// fmt.Println(SessionUserId)
	userId := ctx.Value(SessionUserId)
	// fmt.Println("GetUserIdFromCtx1:  ")
	if userId == nil {
		return 0
	}
	// fmt.Println(userId.(int64))
	return userId.(int64)
}
