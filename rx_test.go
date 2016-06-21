package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"xr/xutor/global"
	"xr/xutor/utils"
)

func TestRx(t *testing.T) {
	var rx = regexp.MustCompile(global.EMAIL_PTRN)
	fmt.Println(rx.MatchString("troitskiy.evgeniy+0106@gmail.com"))
}

func TestRxHttpzUrl(t *testing.T) {
	var validID = regexp.MustCompile(`^https?://[0-9a-zA-Z]+`)

	fmt.Println(validID.MatchString("https://aaaaaaaaaaaaaaa"))
	fmt.Println(validID.MatchString("http://aaaaaaaaaaaaaaa"))
	fmt.Println(validID.MatchString("http://a"))
	fmt.Println()
	fmt.Println(validID.MatchString("http://"))
	fmt.Println(validID.MatchString("http:/"))
	fmt.Println(validID.MatchString("h"))
	fmt.Println()
}

func TestRx1(t *testing.T) {
	ptrn := `^[1-2]{1}[0-9]{3}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}(-|\+)[0-9]{2}:[0-9]{2}$`
	var rx = regexp.MustCompile(ptrn)

	fmt.Println(rx.MatchString("2015-05-27T23:42:11-04:00"))
	fmt.Println(rx.MatchString("2015-09-02T17:03:54+03:00"))
	fmt.Println(rx.MatchString("2015-09-02T17:03:54*03:00"))
}

func TestRxEmail(t *testing.T) { // #65403
	var validID = regexp.MustCompile(global.EMAIL_PTRN)

	fmt.Println(validID.MatchString("liliana_benavides@hotmail.com"))
	fmt.Println(validID.MatchString("liliana_benavides_@hotmail.com"))
	fmt.Println(validID.MatchString("liliana.benavides@hotmail.com"))
	fmt.Println(validID.MatchString("liliana`benavides@hotmail.com"))
	fmt.Println()
	fmt.Println(validID.MatchString("Lorraine_concepcion@ymail.com"))
	fmt.Println()
	fmt.Println(validID.MatchString("http://"))
	fmt.Println(validID.MatchString("http:/"))
	fmt.Println(validID.MatchString("z"))
	fmt.Println(validID.MatchString("@"))
	fmt.Println()
}

func TestRxReplace1(t *testing.T) { // #66628
	str := `<div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div><a href="http://questandrest.com/?question=173" target="_blank" style="text-decoration: none;">`
	ptrn := `(<a href=)"https?://([a-z0-9+-_/\s.:;=#%$@&]{5,500})"`
	fmt.Println(str)
	rx := regexp.MustCompile(ptrn)
	fmt.Println(rx.ReplaceAllString(str, `$1"${redir:$2}"`))
}

func TestRxReplace101(t *testing.T) { // #66628
	str := `<div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div><a href="http://questandrest.com/?question=173" target="_blank" style="text-decoration: none;">`
	ptrn := `(<a href=)"(https?://[a-z0-9+-_/\s.:;=#%$@&]{5,500})"`
	fmt.Println(str)
	rx := regexp.MustCompile(ptrn)
	fmt.Println(rx.ReplaceAllString(str, `$1"${redir:$2}"`))
}

func TestRxReplace2(t *testing.T) { // #66628
	str := `<div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div><a href="http://questandrest.com/?question=173" target="_blank" style="text-decoration: none;">`
	ptrn := `(<a href=)"https?://([a-z0-9+-_/\s.:;=#%$@&]{5,500})"`
	fmt.Println(str)
	rx := regexp.MustCompile(ptrn)
	fmt.Println(rx.ReplaceAllString(str, `$1"http://xr.questandrest.net/[[any codes include $2]]"`))
}

func TestRxReplace3(t *testing.T) { // #66628
	str := `<div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div><a href="http://questandrest.com/?question=173" target="_blank" style="text-decoration: none;">`
	ptrn := `(<a href=)"https?://([a-z0-9+-_/\s.:;=#%$@&]{5,500})"`
	fmt.Println(str)
	rx := regexp.MustCompile(ptrn)

	str = rx.ReplaceAllStringFunc(str, func(src string) string {
		s0 := rx.ReplaceAllString(src, `$1"`)
		s1 := rx.ReplaceAllString(src, `http://xr.questandrest.net/`)
		s2 := rx.ReplaceAllString(src, `[[--$2--]]"`)
		return s0 + s1 + s2

	})
	fmt.Println(str)
}

func TestRxReplace4(t *testing.T) { // #66628
	str := `<div style="height: 20px; line-height:20px; font-size:18px;">&nbsp;</div>
	<a href="http://questandrest.com/?question=173" target="_blank" style="text-decoration: none;">
	<img border="0" height="6" src="http://li.interactive95.com/imp?s=123902100&sz=2x1&li=zadavakanew&m=_tomd5&p=_placementid" width="2"/>`
	ptrn := ` (href|src)="https?://([a-z0-9+-_/\s.:;=#%$@&]{5,500})"`
	fmt.Println(str)
	rx := regexp.MustCompile(ptrn)

	str = rx.ReplaceAllStringFunc(str, func(src string) string {
		from := ""
		switch rx.ReplaceAllString(src, `$1`) {
		case `src`:
			from = "(SRC)"
		case `href`:
			from = "(HREF)"
			fmt.Println("")
		}

		s0 := rx.ReplaceAllString(src, ` $1="`)
		s1 := from + `http://xr.zazazaka.net/`
		s2 := rx.ReplaceAllString(src, `((getActLink($2))"`)
		return s0 + s1 + s2

	})
	fmt.Println(`*******************************************************************************************************`)
	fmt.Println(str)
} // ` (href|src)="https?://([a-z0-9+-_/\s.:;=#%$@&]{5,500})"`

func TestRxReplace5(t *testing.T) { // #100896
	str := `Hi ${fname}!
Don't stop now. Keep climbing to the top of the leaderboard!
Play - http://zadavaka.us/?token=1234567890&utm_source=zadavaka_subscribers&utm_medium=email&utm_campaign=zadavaka_daily_broadcast
Your Stats:
Level ${userLevel},
${countWP} WisePoints
Answer today's trivia - http://zadavaka.us/?token=${token}&utm_source=zadavaka_subscribers&utm_medium=email&utm_campaign=zadavaka_daily_broadcast

Interactive95
207 E O‌hio, #158
C‌hicago, IL 60611

Unsubscribe ${unsub}`

	ptrn := `\shttps?://([a-z0-9+-_/.:;=#%$@&{}]{5,500})(\n|\s)`

	fmt.Println(str)
	rx := regexp.MustCompile(ptrn)

	str = rx.ReplaceAllStringFunc(str, func(src string) string {

		url := rx.ReplaceAllString(src, `$1`)
		end := rx.ReplaceAllString(src, `$2`)
		fmt.Printf("url:%s\n", url)

		s2 := `http://xr.zazazaka.net/`
		s3 := rx.ReplaceAllString(src, `((getActLink($1))`)

		return s2 + s3 + end

	})
	fmt.Println(`*******************************************************************************************************`)
	fmt.Println(str)
} // `\shttps?://([a-z0-9+-_/.:;=#%$@&{}]{5,500})(\n|\s)`

func TestRxReplaceZ5(t *testing.T) { // #100065
	str := `http://zadavaka.us/feed/rating?token=${token}`
	ptrn := `\$\{(token)\}`
	fmt.Println(str)
	rx := regexp.MustCompile(ptrn)

	str = rx.ReplaceAllStringFunc(str, func(src string) string {
		replKey := string(src[2 : len(src)-1])
		fmt.Println(replKey)

		return `NEKOT`

	})
	fmt.Println(str)
}

func TestFindAllString1(t *testing.T) {
	str := `http://questandrest.com/feed/user?userid=${XxHash64Hex16}`
	rx := regexp.MustCompile(`\$\{(XxHash64Hex16)\}`)

	res := rx.ReplaceAllStringFunc(str, func(src string) string {
		codetype := string(src[2 : len(src)-1])
		switch {
		case strings.HasPrefix(codetype, "XxHash64Hex16"):
			return utils.GetXxhash64Hex("a@b.com", 16)

		default:
			fmt.Printf("code type='%v' not supported.\n", codetype)
			return "error"
		}
	})
	fmt.Printf("result: %s\n", res)

	fmt.Println()
}
