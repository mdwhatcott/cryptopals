package cryptopals

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestHexToBytes(t *testing.T) {
	actual := HexToBytes("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	expected, _ := hex.DecodeString("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if !bytes.Equal(expected, actual) {
		t.Errorf("\nExpected: %q\nActual:   %q", expected, actual)
	}
}

func TestHexToBase64(t *testing.T) {
	actual := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if actual != expected {
		t.Errorf("\nExpected: %q\nActual:   %q", expected, actual)
	}
}

func TestXOR(t *testing.T) {
	a := HexToBytes("1c0111001f010100061a024b53535009181c")
	b := HexToBytes("686974207468652062756c6c277320657965")
	c := make([]byte, len(a))
	XOR(a, b, c)
	expected := HexToBytes("746865206b696420646f6e277420706c6179")
	if !bytes.Equal(expected, c) {
		t.Errorf("Fail")
	}
}