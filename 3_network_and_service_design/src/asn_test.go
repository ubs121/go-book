package inout

import (
	"encoding/asn1"
	"fmt"
	"testing"
)

func TestAsn(t *testing.T) {
	mdata, err := asn1.Marshal(13)
	if err != nil {
		t.Error(err)
	}

	var n int
	_, err1 := asn1.Unmarshal(mdata, &n)
	if err1 != nil {
		t.Error(err1)
	}

	fmt.Println("After marshal/unmarshal: ", n)
}
