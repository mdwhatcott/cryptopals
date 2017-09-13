package cryptopals

import (
	"encoding/hex"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestConvertFixture(t *testing.T) {
	gunit.Run(new(ConvertFixture), t)
}

type ConvertFixture struct {
	*gunit.Fixture
}

func (this *ConvertFixture) TestHexToBytes() {
	const input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	actual := HexToBytes(input)
	expected, _ := hex.DecodeString(input)
	this.So(actual, should.Resemble, expected)
}

func (this *ConvertFixture) assertBytesToBase64(input []byte, expected string) {
	this.So(BytesToBase64(input), should.Equal, expected)
}
func (this *ConvertFixture) TestBytesToBase64_NothingIn_NothingOut() {
	this.So(BytesToBase64(nil), should.Equal, "")
}
func (this *ConvertFixture) TestBytesToBase64_ThreeOctetsIn_FourCharactersOut() {
	this.So(BytesToBase64([]byte("Man")), should.Equal, "TWFu")
}
func (this *ConvertFixture) TestBytesToBase64_TwoOctetsIn_FourthCharacterIsPadding() {
	this.So(BytesToBase64([]byte("Ma")), should.Equal, "TWE=")
}
func (this *ConvertFixture) TestBytesToBase64_OneOctetIn_ThirdAndFourthCharactersArePadding() {
	this.So(BytesToBase64([]byte("M")), should.Equal, "TQ==")
}
func (this *ConvertFixture) TestHexToBase64() {
	const input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	this.So(HexToBase64(input), should.Equal, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")
}

func (this *ConvertFixture) TestXOR() {
	a := HexToBytes("1c0111001f010100061a024b53535009181c")
	b := HexToBytes("686974207468652062756c6c277320657965")
	c := make([]byte, len(a))
	XOR(a, b, c)
	expected := HexToBytes("746865206b696420646f6e277420706c6179")
	this.So(c, should.Resemble, expected)
}
