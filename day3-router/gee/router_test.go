package gee

import "testing"

func TestParsePattern(t *testing.T) {
	t.Log(parsePattern("/"))
	t.Log(parsePattern("//"))
	t.Log(parsePattern("index"))
	t.Log(parsePattern("index.html"))
	t.Log(parsePattern("/index.html"))
	t.Log(parsePattern("*"))
	t.Log(parsePattern("/*"))
	t.Log(parsePattern("*/"))
	t.Log(parsePattern("/*/"))
	t.Log(parsePattern("**"))
}
