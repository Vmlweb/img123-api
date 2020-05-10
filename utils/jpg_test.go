package utils

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestJPGtoPNG(t *testing.T) {
	inputJpg, err := ioutil.ReadFile("res/cat.jpg")
	if err != nil {
		t.Error(err)
	}

	inputPng, err := ioutil.ReadFile("res/cat.png")
	if err != nil {
		t.Error(err)
	}

	outputPng, err := JPGtoPNG(inputJpg)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(outputPng, inputPng) {
		t.Error("input png does not equal output png")
	}
}
