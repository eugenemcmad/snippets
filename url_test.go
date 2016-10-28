package tests

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrlQueryEscape(t *testing.T) {
	url0 := `utmsource=&email=kozlovaa62+test2@gmail.com&ip=94.231.116.43&status=11&apikey=os6peow476hmf53t&utmmedium=&sourceurl=https://drivingtests.us/l/1-subscribe&fname=test&token=68fffdf7dfcac7ce2c84092a3297fbff&regdate=2016-10-24T07:03:25-04:00`
	fmt.Println(url0)
	fmt.Println(url.QueryEscape(url0))
}
