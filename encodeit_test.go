package main

import (
	"bytes"
	"testing"
)

var value0 = []byte("")
var value1 = []byte("hash_me")
var value2 = []byte("hash me too")
var value3 = []byte(`hash me too
please?
`)

var res0 = ""
var res1 = "aGFzaF9tZQ=="
var res2 = "aGFzaCBtZSB0b28="
var res3 = "aGFzaCBtZSB0b28="
var res33 = "hash me too"

func TestSource(t *testing.T) {

}

func TestFetchInput0(t *testing.T) {
	buf := bytes.NewBuffer(value0)
	res, err := fetchInput(buf)
	if err != nil {
		t.Error(err)
	}
	if res != string(value0) {
		t.Errorf("Should be %s was %s", value0, res)
	}
}

func TestFetchInput1(t *testing.T) {
	buf := bytes.NewBuffer(value1)
	res, err := fetchInput(buf)
	if err != nil {
		t.Error(err)
	}
	if res != string(value1) {
		t.Errorf("Should be %s was %s", value1, res)
	}
}

func TestFetchInput2(t *testing.T) {
	buf := bytes.NewBuffer(value2)
	res, err := fetchInput(buf)
	if err != nil {
		t.Error(err)
	}
	if res != string(value2) {
		t.Errorf("Should be %s was %s", value2, res)
	}
}

func TestFetchInput3(t *testing.T) {
	buf := bytes.NewBuffer(value3)
	res, err := fetchInput(buf)
	if err != nil {
		t.Error(err)
	}
	if res != res33 {
		t.Errorf("Should be %s was %s", res33, res)
	}
}

// ---

func TestEncodeBase64NoValue(t *testing.T) {

	res, err := encodeValue(Base64, value0)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if res != res0 {
		t.Errorf("Should be %s was %s", res0, res)
	}
}
func TestEncodeBase64Value1(t *testing.T) {

	res, err := encodeValue(Base64, value1)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if res != res1 {
		t.Errorf("Should be %s was %s", res1, res)
	}
}
func TestEncodeBase64Value2(t *testing.T) {

	res, err := encodeValue(Base64, value2)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if res != res2 {
		t.Errorf("Should be %s was %s", res2, res)
	}
}
