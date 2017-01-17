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

	ActionRedirect = global.ActionRedirect // 0
	ActionNo       = global.ActionNo       // 9
)

var (
	rxTypeHttpUrlPtrnInHtml = regexp.MustCompile(global.TYPE_HTTP_URL_PTRN_IN_HTML)
	rxHttpUrlPtrnInText     = regexp.MustCompile(global.HTTP_URL_PTRN_IN_TEXT)
)

var (
	tstHtml = `<br/>
<b>Hello world ${email}</b> <br/>
<a href="http://rm.regium.com" target="_blank" > zhmyak </a>  your isp:${edg}? <br/>
<a href="rm.regium.com" target="_blank" > invalid-link </a> <br/>
<a src="http://rm.regium.com/questandrest/img/logo.png" alt="" width="161" height="102" /> <br/>
<div class="image" style="font-size: 12px;font-style: normal;font-weight: 400;" align="center">
          <img style="display: block;border: 0;max-width: 161px;" src="${img}/questandrest/img/logo.png" alt="" width="161" height="102" />
        </div> <br/>
<a href="${unsub}" target="_blank" style="color:#00addb; text-decoration: none;">Unsubscribe</a></span>
<a href="https://rm.regium.com/close?a=b&c=end" target="_blank" > close me </a>`
)

// go test -v xr/snippets -bench ^Benchmark_RXLHFN_ -run ^$

// 3321 ns/op
func Benchmark_RXLHFN_RxFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setSimpleLinksToHtml_RX(tstHtml, "")
	}
}

// 3321 ns/op
func Benchmark_RXLHFN_UtFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setSimpleLinksToHtml_UT(tstHtml, "")
	}
}

func TestRcFunc(t *testing.T) {

	fmt.Printf("setSimpleLinksToHtml_RX(): %v\n\n", setSimpleLinksToHtml_RX(tstHtml, ""))

	fmt.Printf("setSimpleLinksToHtml_UT(): %v\n\n", setSimpleLinksToHtml_UT(tstHtml, ""))

}

// Set open/simple redirect & src links to html content. TODO: OPTIMIZ 0000 WINNER
// ` (href|src)="https?://([a-z0-9+-~_/.:;=#%$@&{}]{5,1000})"`
func setSimpleLinksToHtml_RX( /*trackParams trackParams,*/ str, trackDomain string) string {
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
		newLink := "HTTP://" + url + "-ATTR-IS-" + typeAttr + "-END"

		repLink := start + newLink + `"`

		return repLink
	})
}

// ` (href|src)="https?://([a-z0-9+-~_/.:;=#%$@&{}]{5,1000})"`
func setSimpleLinksToHtml_UT(str, trackDomain string) string {

	// find start: (href|src)="https?://(
	// find end:   "
	// get action
	// get url origin
	// get url result
	// replace
	var res, url string
	var l0, l1, r0, actionId int
	var in, isSrc bool
	for i := 0; i < len(str)-14; i++ {

		switch {
		case str[i:i+15] == ` href="https://`:
			l0 = i
			i += 14
			l1 = l0 + 7
			in = true
			isSrc = false
			actionId = ActionRedirect
			break
		case str[i:i+14] == ` href="http://`:
			l0 = i
			i += 13
			l1 = l0 + 7
			in = true
			isSrc = false
			actionId = ActionRedirect
			break
		case str[i:i+14] == ` src="https://`:
			l0 = i
			i += 13
			l1 = l0 + 6
			in = true
			isSrc = true
			actionId = ActionNo
			break
		case str[i:i+13] == ` src="http://`:
			l0 = i
			i += 12
			l1 = l0 + 6
			in = true
			isSrc = true
			actionId = ActionNo
			break
		case in && str[i] == ' ': // Error - skip pattern
			in = false
			actionId = 0
			//fmt.Printf("[ERR] '%v' l:%d, %d aid:%v \n", str[l0:l0+20], l0, l1, actionId)
			break
		case in && str[i:i+2] == `" `: // Pattern complete
			in = false
			url = str[l1:i]
			if isSrc {
				res += str[r0:l0] + ` src="http://` + "[START[" + url + "]END]"
			} else {
				res += str[r0:l0] + ` href="http://` + "[START[" + url + "]END]"
			}

			r0 = i
			fmt.Printf("'%v' l:%d, %d aid:%v \n", url, l0, l1, actionId)
			break
		default:
			continue
		}

		//fmt.Printf("[%v] l:%d, aid:%v \n", str[l0:l0+20], l0, actionId)
	}

	res += str[r0:]
	actionId++ // test compatible

	return res
}

// Set open/simple redirect & src links to text content.
// `([\s\(]|^)https?://([a-z0-9+-~_/.:;=#%$@&{}]{5,1000})([\n\s\)]|$)`
func setSimpleLinksToText_RX( /*trackParams trackParams,*/ str, trackDomain string) string {
	return rxHttpUrlPtrnInText.ReplaceAllStringFunc(str, func(src string) string {

		start := rxHttpUrlPtrnInText.ReplaceAllString(src, `$1`) // pre ([\s\(]|^)  ==       [' ', '(', '{START_STRING}']
		url := rxHttpUrlPtrnInText.ReplaceAllString(src, `$2`)   // pure url
		end := rxHttpUrlPtrnInText.ReplaceAllString(src, `$3`)   // end ([\n\s\)]|$)== ['\n', ' ', ')', '{END_STRING}']

		//newLink, err := getActLink(trackParams, ActionRedirect, trackDomain, url)
		newLink := "HTTP://" + url + "-ATTR-IS-" + "ActionRedirect" + "-END"

		repLink := start + newLink + end

		return repLink
	})
}

// Set open/simple redirect & src links to text content.
// `([\s\(]|^)https?://([a-z0-9+-~_/.:;=#%$@&{}]{5,1000})([\n\s\)]|$)`
func setSimpleLinksToText_UT( /*trackParams trackParams,*/ str, trackDomain string) string {
	return rxHttpUrlPtrnInText.ReplaceAllStringFunc(str, func(src string) string {

		start := rxHttpUrlPtrnInText.ReplaceAllString(src, `$1`) // pre ([\s\(]|^)  ==       [' ', '(', '{START_STRING}']
		url := rxHttpUrlPtrnInText.ReplaceAllString(src, `$2`)   // pure url
		end := rxHttpUrlPtrnInText.ReplaceAllString(src, `$3`)   // end ([\n\s\)]|$)== ['\n', ' ', ')', '{END_STRING}']

		//newLink, err := getActLink(trackParams, ActionRedirect, trackDomain, url)
		newLink := "HTTP://" + url + "-ATTR-IS-" + "ActionRedirect" + "-END"

		repLink := start + newLink + end

		return repLink
	})
}
