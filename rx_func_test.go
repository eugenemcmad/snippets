package tests

import (
	"fmt"
	"regexp"
	"testing"
	"xr/xutor/global"
)

const (
	HREF_ATTR = `href`
	SRC_ATTR  = `src`

	ActionRedirect = global.ActionRedirect
	ActionNo       = global.ActionNo
)

var (
	rxTypeHttpUrlPtrnInHtml = regexp.MustCompile(global.TYPE_HTTP_URL_PTRN_IN_HTML)
	rxHttpUrlPtrnInText     = regexp.MustCompile(global.HTTP_URL_PTRN_IN_TEXT)
)
var (
	tstHtml = `
	<b>Hello world ${email}</b> <br/>
<a href="http://rm.regium.com" target="_blank" > zhmyak </a>  your isp:${edg}? <br/>
<a href="rm.regium.com" target="_blank" > invalid-link </a> <br/>
<div class="image" style="font-size: 12px;font-style: normal;font-weight: 400;" align="center">
          <img style="display: block;border: 0;max-width: 161px;" src="${img}/questandrest/img/logo.png" alt="" width="161" height="102" />
        </div> <br/>
<a href="${unsub}" target="_blank" style="color:#00addb; text-decoration: none;">Unsubscribe</a></span>`
)

// 3321 ns/op
func Benchmark_RXFN_RxFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setSimpleLinksToHtml_RX(tstHtml, "", true)
	}
}

func TestRcFunc(t *testing.T) {

	fmt.Printf("setSimpleLinksToHtml_RX(): %v \n", setSimpleLinksToHtml_RX(tstHtml, "", true))

	fmt.Printf("setSimpleLinksToHtml_UT(): %v \n", setSimpleLinksToHtml_UT(tstHtml, "", true))

}

// Set open/simple redirect & src links to html content. TODO: OPTIMIZ 0000 WINNER
// ` (href|src)="https?://([a-z0-9+-~_/.:;=#%$@&{}]{5,1000})"`
func setSimpleLinksToHtml_RX( /*trackParams trackParams,*/ str, trackDomain string, https bool) string {
	return rxTypeHttpUrlPtrnInHtml.ReplaceAllStringFunc(str, func(src string) string {
		typeAttr := rxTypeHttpUrlPtrnInHtml.ReplaceAllString(src, `$1`) // (href|src)
		url := rxTypeHttpUrlPtrnInHtml.ReplaceAllString(src, `$2`)      // pure url

		//actionId := global.ActionNo
		//switch typeAttr {
		//case HREF_ATTR:
		//	actionId = ActionRedirect
		//case SRC_ATTR:
		//	actionId = ActionNo
		//}

		start := rxTypeHttpUrlPtrnInHtml.ReplaceAllString(src, ` $1="`)

		//newLink, err := getActLink(trackParams, actionId, trackDomain, https, url)
		//if err != nil {
		//	return ``
		//}
		newLink := "HTTP://" + url + "-AAAA-" + typeAttr + "-ZZZZ-"

		repLink := start + newLink + `"`

		return repLink
	})
}

// ` (href|src)="https?://([a-z0-9+-~_/.:;=#%$@&{}]{5,1000})"`
func setSimpleLinksToHtml_UT( /*trackParams trackParams,*/ str, trackDomain string, https bool) string {

	// find start: (href|src)="https?://(
	// find end:   "
	// get action
	// get url origin
	// get url result
	// replace
	var res string
	var l, actionId int

	for i := 0; i < len(str)-10; i++ {
		if str[i:i+11] == ` href="http` {
			actionId = ActionRedirect
		} else if str[i:i+10] == ` src="http` {
			actionId = ActionNo
		} else {
			continue
		}
		l = i
		fmt.Printf("%v l:%d, aid:%v \n", str[i:i+5], l, actionId)
	}

	return res
}
