package avutil_test

import (
	"github.com/asticode/goav/avutil"
	"testing"
)

func TestVersionMajor(t *testing.T) {
	v := avutil.Version()
	t.Log(avutil.Major(v), avutil.Minor(v), avutil.Micro(v))
	t.Log(avutil.AvVersion(v))
	t.Log(avutil.VersionInfo())
}

func TestAvGetPictureTypeChar(t *testing.T) {
	c := avutil.AvGetPictureTypeChar(avutil.AvPictureTypeI)
	if c != 'I' {
		t.FailNow()
	}
	t.Logf("%c", c)
}
