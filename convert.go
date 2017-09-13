package cryptopals

import (
	"bytes"
	"log"
)

func HexToBase64(hex string) string {
	value := HexToBytes(hex)
	return BytesToBase64(value)
}

func HexToBytes(hex string) []byte {
	value := make([]byte, len(hex)/2)
	for x := 0; x < len(hex)/2; x++ {
		first := hexChar(hex[x*2+0])
		second := hexChar(hex[x*2+1])
		value[x] = (first << 4) | second
	}
	return value
}

func hexChar(c byte) byte {
	if c >= '0' && c <= '9' {
		return c - '0'
	} else if c >= 'a' && c <= 'f' {
		return c - 'a' + 10
	} else if c >= 'A' && c <= 'F' {
		return c - 'A' + 10
	} else {
		log.Panicf("Invalid hex byte: %d", c)
	}
	return 0
}

func BytesToBase64(in []byte) string {
	if len(in) == 0 {
		return ""
	}

	out := new(bytes.Buffer)
	x := 0

	for ; x+3 < len(in); x += 3 {
		triplet := in[x : x+3]
		out.Write(tripletToBase64(triplet))
	}

	if x < len(in) {
		remainder := in[x:]
		out.Write(tripletToBase64(remainder))
	}
	return out.String()
}

func tripletToBase64(in []byte) (out []byte) {
	out = make([]byte, 4)

	out[0] = in[0] >> 2
	out[0] = base64Key[out[0]]

	if len(in) == 1 {
		in = append(in, 0)
		out[2] = base64Padding
	}

	out[1] = in[0]
	out[1] = out[1] << 6
	out[1] = out[1] >> 2
	out[1] = out[1] + (in[1] >> 4)
	out[1] = base64Key[out[1]]

	if len(in) == 2 {
		in = append(in, 0)
		out[3] = base64Padding
	}

	if out[2] == 0 {
		out[2] = in[1] << 4
		out[2] = out[2] >> 2
		out[2] = out[2] + in[2]>>6
		out[2] = base64Key[out[2]]
	}

	if out[3] == 0 {
		out[3] = in[2]
		out[3] = out[3] << 2
		out[3] = out[3] >> 2
		out[3] = base64Key[out[3]]
	}

	return out
}

const base64Key = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
const base64Padding = '='

func XOR(a, b, c []byte) {
	for x := 0; x < len(a); x++ {
		c[x] = a[x] ^ b[x]
	}
}
