// Copyright(C) 2020 github.com/hidu  All Rights Reserved.
// Author: hidu (duv123+git@baidu.com)
// Date: 2020/1/4

package localmodule

import (
	"testing"

	"github.com/jzero-io/go_fmt/internal/common"
)

func Test_detectByGoMod(t *testing.T) {
	opt := common.Options{}
	got, err := detectByGoMod(opt, "")
	if err != nil {
		t.Fatalf("detectByGoMod() with error:%s", err)
	}
	want := "github.com/jzero-io/go_fmt"
	if got != want {
		t.Fatalf("detectByGoMod() =%q, want=%q", got, want)
	}
}
