package tests

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
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
	rxImgTag                = regexp.MustCompile(`\$\{(img)\}`)
	rxUnsubTag              = regexp.MustCompile(`\$\{unsub\}`)

	ImgPlacementBytes   = []byte("${img}")
	httpImgReplPref     = []byte("http://img.")
	UnsubPlacementBytes = []byte("${unsub}")
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
<a src="${img}example/hello.gif" alt="" width="161" height="102" /> <br/>
<a href="${unsub}" target="_blank" style="color:#00addb; text-decoration: none;">Unsubscribe</a></span>
<a href="https://rm.regium.com/close?a=b&c=end" target="_blank" > close me </a>`

	tstTxt = `
boddy start
http://rm.regium.com/10101010 http://rm.regium.com/20202020 (http://rm.regium.com/30303030) http://rm.regium.com/40404040
www.wrong-uri.com
https://rm.regium.com/50505050
end http://rm.regium.com/60606060`

	tstUnsublink = "http://xr.xrtd/a101/b589/c894/d784"
)

// go test -v xr/snippets -bench ^Benchmark_RXLHFN_ -run ^$

// 65267 ns/op
func Benchmark_RXLHFN_RxFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setSimpleLinksToHtml_RX(tstHtml, "")
	}
}

// 14791 ns/op
func Benchmark_RXLHFN_UtFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setSimpleLinksToHtml_UT(tstHtml, "")
	}
}

func TestLHtmlFunc(t *testing.T) {
	fmt.Printf("setSimpleLinksToHtml_RX(): %v\n\n", setSimpleLinksToHtml_RX(tstHtml, ""))
	fmt.Printf("setSimpleLinksToHtml_UT(): %v\n\n", setSimpleLinksToHtml_UT(tstHtml, ""))
}

// go test -v xr/snippets -bench ^Benchmark_RXLTFN_ -run ^$

// 70231 ns/op
func Benchmark_RXLTFN_RxFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setSimpleLinksToText_RX(tstHtml, "")
	}
}

// 10769 ns/op
func Benchmark_RXLTFN_UtFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		setSimpleLinksToText_UT(tstHtml, "")
	}
}

func TestLTextFunc(t *testing.T) {
	fmt.Printf("setSimpleLinksToText_RX(): %v\n\n", setSimpleLinksToText_RX(tstTxt, ""))
	fmt.Printf("setSimpleLinksToText_UT(): %v\n\n", setSimpleLinksToText_UT(tstTxt, ""))
}

func TestImgFunc(t *testing.T) {
	h1, err := SetImgLinks_RX(tstHtml, "xrtd")
	fmt.Printf("SetImgLinks_RX(): %v\n\n err:%v\n\n", h1, err)

	h2, err := SetImgLinks_UT(tstHtml, "xrtd")
	fmt.Printf("SetImgLinks_UT(): %v\n\n err:%v\n\n", h2, err)
}

// go test -v xr/snippets -bench ^Benchmark_IMG_ -run ^$

// 2728 ns/op
func Benchmark_IMG_RxFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetImgLinks_RX(tstHtml, "xrts")
	}
}

// 1449 ns/op
func Benchmark_IMG_UtFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetImgLinks_UT(tstHtml, "xrts")
	}
}

func TestUnsubFunc(t *testing.T) {
	fmt.Printf("original: %v\n\n unsublink:%v\n\n", tstHtml, tstUnsublink)

	h1, err := SetUnsubLinks_RX(tstHtml, tstUnsublink)
	fmt.Printf("SetUnsubLinks_RX(): %v\n\n err:%v\n\n", h1, err)

	h2, err := SetUnsubLinks_UT(tstHtml, tstUnsublink)
	fmt.Printf("SetUnsubLinks_UT(): %v\n\n err:%v\n\n", h2, err)
}

// go test -v xr/snippets -bench ^Benchmark_UNS_ -run ^$

// 2001 ns/op
func Benchmark_UNS_RxFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetUnsubLinks_RX(tstHtml, tstUnsublink)
	}
}

// 1423 ns/op
func Benchmark_UNS_UtFn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetUnsubLinks_UT(tstHtml, tstUnsublink)
	}
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
	var res, url string
	var l0, l1, r0, actionId int
	var in bool
	for i := 0; i < len(str)-14; i++ {

		switch {
		case str[i:i+15] == ` href="https://`:
			l0 = i
			i += 14
			l1 = l0 + 7
			in = true
			actionId = ActionRedirect
			break
		case str[i:i+14] == ` href="http://`:
			l0 = i
			i += 13
			l1 = l0 + 7
			in = true
			actionId = ActionRedirect
			break
		case str[i:i+14] == ` src="https://`:
			l0 = i
			i += 13
			l1 = l0 + 6
			in = true
			actionId = ActionNo
			break
		case str[i:i+13] == ` src="http://`:
			l0 = i
			i += 12
			l1 = l0 + 6
			in = true
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
			res += str[r0:l1] + "[START[" + url + "]END]"

			r0 = i
			//fmt.Printf("'%v' l:%d, %d aid:%v \n", url, l0, l1, actionId)
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
func setSimpleLinksToText_UT(str, trackDomain string) string {
	// `([\s\(]|^)https?://([a-z0-9+-~_/.:;=#%$@&{}]{5,1000})([\n\s\)]|$)`
	//start := rxHttpUrlPtrnInText.ReplaceAllString(src, `$1`) // pre ([\s\(]|^)  ==       [' ', '(', '{START_STRING}']
	//url := rxHttpUrlPtrnInText.ReplaceAllString(src, `$2`)   // pure url
	//end := rxHttpUrlPtrnInText.ReplaceAllString(src, `$3`)   // end ([\n\s\)]|$)== ['\n', ' ', ')', '{END_STRING}']
	var res, url string
	var l0, r0 int
	var in bool
	for i := 0; i < len(str); i++ {
		switch {
		case i < len(str)-8 && str[i:i+8] == `https://`:
			l0 = i
			i += 7
			in = true
			break
		case i < len(str)-7 && str[i:i+7] == `http://`:
			l0 = i
			i += 6
			in = true
			break
		case in && (str[i] == ' ' || str[i] == ')' || str[i] == '\n'): // Pattern complete
			in = false
			url = str[l0:i]
			res += str[r0:l0] + "[START[" + url + "]END]"
			r0 = i
			//fmt.Printf("[%v] l:%d from %d\n", url, l0, len(str))
			break
		case in && i == len(str)-1: // Pattern complete
			in = false
			url = str[l0:]
			res += str[r0:l0] + "[START[" + url + "]END]"
			r0 = len(str)
			//fmt.Printf("[%v] l:%d from %d\n", url, l0, len(str))
			break
		default:
			continue
		}
	}

	res += str[r0:]

	return res
}

// Set specific src links.
//  `src="${img}example/hello.gif"` e.g.
//  return IError.
func SetImgLinks_RX(in, trackdomain string) (string, error) {
	if in == "" {
		return "", fmt.Errorf("set img link fail - input html is empty")
	}
	if trackdomain == "" {
		return "", fmt.Errorf("set img link fail - tracking domain is empty")
	}

	s := rxImgTag.ReplaceAllStringFunc(in, func(src string) string {
		pref := string(src[2 : len(src)-1])
		switch {
		case strings.HasPrefix(pref, "img"): // `img`
			imgLink := fmt.Sprintf(`http://%s.%s/`, `img`, trackdomain)
			return imgLink
		default:
			return ""
		}
	})

	return s, nil
}

// Set specific src links.
//  `src="${img}example/hello.gif"` e.g.
//  return IError.
func SetImgLinks_UT(in, trackdomain string) (string, error) {
	if in == "" {
		return "", fmt.Errorf("set img link fail - input html is empty")
	}
	if trackdomain == "" {
		return "", fmt.Errorf("set img link fail - tracking domain is empty")
	}

	var buf bytes.Buffer
	buf.Write(httpImgReplPref)
	buf.WriteString(trackdomain)
	buf.WriteByte('/')

	var res = string(bytes.Replace([]byte(in), ImgPlacementBytes, buf.Bytes(), -1))

	return res, nil
}

// Set specific unsubscribe url.
//  Make specific link and replace `${unsub}`.
//  return IError.
func SetUnsubLinks_RX(in, unsublink string) (string, error) {
	if in == "" {
		return "", fmt.Errorf("set unsub link fail - input text/html is empty")
	}
	if unsublink == "" {
		return "", fmt.Errorf("set unsub link fail - input unsublink text/html is empty")
	}

	s := rxUnsubTag.ReplaceAllStringFunc(in, func(src string) string {
		pref := string(src[2 : len(src)-1])
		switch {

		case strings.HasPrefix(pref, `unsub`): // `unsub`
			return unsublink

		default:
			return ""
		}
	})

	return s, nil
}

// Set specific unsubscribe url.
//  Make specific link and replace `${unsub}`.
//  return IError.
func SetUnsubLinks_UT(in, unsublink string) (string, error) {
	if in == "" {
		return "", fmt.Errorf("set unsub link fail - input text/html is empty")
	}
	if unsublink == "" {
		return "", fmt.Errorf("set unsub link fail - input unsublink text/html is empty")
	}

	var res = string(bytes.Replace([]byte(in), UnsubPlacementBytes, []byte(unsublink), -1))

	return res, nil
}
