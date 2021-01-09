package encoding

import "testing"

func TestBase64EncodeDecode(t *testing.T) {
	od := []byte("abc")

	ec := Base64Encode(od)
	t.Log(string(ec))

	dc := Base64Decode(ec)
	t.Log(string(dc))

	if string(dc) != string(od) {
		t.Error("invalid")
	}
}
