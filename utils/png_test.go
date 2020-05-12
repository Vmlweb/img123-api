package utils

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestPNG(t *testing.T) {
	testPng, err := ioutil.ReadFile("res/cat.png")
	if err != nil {
		t.Error(err)
	}

	png, err := DecodePNG(testPng)
	if err != nil {
		t.Error(err)
	}

	output, err := EncodePNG(png)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(testPng, output) {
		t.Error("input and output png do not match")
	}
}
