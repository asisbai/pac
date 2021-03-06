package blackwhite

import (
	"strings"
	"testing"
)

func TestIsWhite(t *testing.T) {
	ok := IsWhite("211.94.114.48")
	t.Log(ok)

	ok = IsWhite("www.google.com")
	t.Log(ok)

	ok = IsWhite("music.163.com")
	t.Log(ok)

	ok = IsWhite("www.txthinking.com")
	t.Log(ok)

	ok = IsWhite("121.196.205.97")
	t.Log(ok)

	ok = IsWhite("a.cn")
	t.Log(ok)

	ok = IsWhite("txthinking.com")
	t.Log(ok)
}

func TestGetWhiteAPP(t *testing.T) {
	s := string(strings.TrimSpace(white_app))
	ss := strings.Split(s, "\n")
	t.Log(":" + ss[0] + ":")
	t.Log(":" + ss[len(ss)-1] + ":")
}
