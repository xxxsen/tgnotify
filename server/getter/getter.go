package getter

import (
	"context"
	"tgnotify/constant"

	"github.com/xxxsen/common/cgi"
	"github.com/xxxsen/common/errs"
)

func MustGetUserList(ctx context.Context) map[string]string {
	v, ok := cgi.GetAttachKey(ctx, constant.KeyUserList)
	if !ok {
		panic(errs.New(errs.ErrParam, "user list key not found"))
	}
	return v.(map[string]string)
}
