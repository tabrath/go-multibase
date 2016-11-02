package multibase

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestBaseRoundTrips(t *testing.T) {
	testBaseRoundTrip(t, Base16)
	testBaseRoundTrip(t, Base58BTC)
	testBaseRoundTrip(t, Base58Flickr)
}

func testBaseRoundTrip(t *testing.T, base int) {
	buf := make([]byte, 16)
	rand.Read(buf)

	enc, err := Encode(base, buf)
	if err != nil {
		t.Fatal(err)
	}

	e, out, err := Decode(enc)
	if err != nil {
		t.Fatal(err)
	}

	if e != base {
		t.Fatal("got wrong encoding out")
	}

	if !bytes.Equal(buf, out) {
		t.Fatal("input wasnt the same as output", buf, out)
	}

	_, _, err = Decode("")
	if err == nil {
		t.Fatal("shouldnt be able to decode empty string")
	}
}
