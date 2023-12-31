package main_test

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	res := count(b, false, false)
	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}

}
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 3
	res := count(b, true, false)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
func TestCountbytes(t *testing.T) {
	b := bytes.NewBufferString("gopher")
	exp := 6
	res := count(b, false, true)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestMain(m *testing.M) {
	exp := "hello World"
	res := main()
	if res != exp {
		t.Errorf("Expected %s, got %s", exp, res)
	}
}
