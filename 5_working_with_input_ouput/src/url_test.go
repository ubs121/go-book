package inout

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURL(t *testing.T) {
	u, err := url.Parse("http://news.zone.mn")
	if err == nil {
		fmt.Println(u.Host)
	} else {
		fmt.Println(err)
	}
}
