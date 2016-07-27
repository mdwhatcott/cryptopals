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

const base64Key = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

// See: https://en.wikipedia.org/wiki/Base64#Sample_Implementation_in_Java
func BytesToBase64(in []byte) string {
	out := new(bytes.Buffer)
	for i := 0; i < len(in); i += 3 {
		b := (in[i] & 0xFC) >> 2
		out.WriteByte(base64Key[b])
		b = (in[i] & 0x03) << 4
		if i+1 < len(in) {
			b |= (in[i+1] & 0xF0) >> 4
			out.WriteByte(base64Key[b])
			b = (in[i+1] & 0x0F) << 2
			if i+2 < len(in) {
				b |= (in[i+2] & 0xC0) >> 6
				out.WriteByte(base64Key[b])
				b = in[i+2] & 0x3F
				out.WriteByte(base64Key[b])
			} else {
				out.WriteByte(base64Key[b])
				out.WriteByte('=')
			}
		} else {
			out.WriteByte(base64Key[b])
			out.WriteString("==")
		}
	}
	return out.String()
}

/*
private static final String CODES = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";

private static String base64Encode(byte[] in)       {
	StringBuilder out = new StringBuilder((in.length * 4) / 3);
	int b;
	for (int i = 0; i < in.length; i += 3)  {
		b = (in[i] & 0xFC) >> 2;
		out.append(CODES.charAt(b));
		b = (in[i] & 0x03) << 4;
		if (i + 1 < in.length)      {
			b |= (in[i + 1] & 0xF0) >> 4;
			out.append(CODES.charAt(b));
			b = (in[i + 1] & 0x0F) << 2;
			if (i + 2 < in.length)  {
				b |= (in[i + 2] & 0xC0) >> 6;
				out.append(CODES.charAt(b));
				b = in[i + 2] & 0x3F;
				out.append(CODES.charAt(b));
			} else  {
				out.append(CODES.charAt(b));
				out.append('=');
			}
		} else      {
			out.append(CODES.charAt(b));
			out.append("==");
		}
	}

	return out.toString();
}
 */