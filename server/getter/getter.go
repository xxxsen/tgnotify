package getter

import (
	"context"
	"tgnotify/constant"

	"github.com/xxxsen/common/errs"
	"github.com/xxxsen/common/naivesvr"
)

func MustGetUserList(ctx context.Context) map[string]string {
	v, ok := naivesvr.GetAttachKey(ctx, constant.KeyUserList)
	if !ok {
		panic(errs.New(errs.ErrParam, "user list key not found"))
	}
	return v.(map[string]string)
}
