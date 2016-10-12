package tests

import (
	"fmt"
	"testing"

	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"
	"strings"
	"xr/xutor/utils"
)

const ()

var (
	crs = []string{
		"9c77268378a4758cdcd3c1bc2965a3b7972da64322567c18928262298a266c98c233beaae257b0cde175e10fcbf0bc", // old uns
		"4467a47ab38fa22468015d5182258d7de6c88e987517c627c80da7f384cc23f52ec83ef35a2e5640a21a621ace4642e454722b103e868ed7ad06db5f8cc4a93065b18a9a9df4f015758802b6",
		"9129e151874679b4715f81d2f2bd1b7568c759f268544d975d802fb80d3430adb9ea18e4f7d8362b46aed88c18b0976cf07433ea587aac1b2d71f0f053ef02a247ce503f79c83b9ca1bec92d",
		"da7cabd343427935dccb9c38f79d24df1f8e0bc797f5e25656ee07a847a9f696b704ec8d319be5f06dcb38955704545ae6ac63b4c7ff03e84efe13adc818177c559b4c1ceae73221c6261d12ee9fa7c2c3000c9c44134198c6079af6f5c7aa98932fa9472030a81da26c4dfae296a553b2d48581b8a5a5629ec07454c68e31c4f0a977cafdadd459eb896f2698435a899653498fa29345bf559fd655ecfd864bc7f8e9c310ca517e",

		"7909bb5f956a9bfaccc2aea71af6c9cec5439ea08b11675cabe126909692fd462ad11e0706985a49906409eda6f007135b55bdbbf99b5652c77d80b40b794e87", // old/new uns
		"37ff342bdc83cd4e330365c3566fe4189eb3d5e395ca5f8d9294a3dd346797e815d5de6949526b6142efbf57d472c9d506b09a33bada3c8aaf65257f35b3bed45ebba63b421b4d68ed774352af2b39909c0d4949400b3c331e35b11ac5",
		"de905eb636dcbba2a83cd62f4185e436fdff07370eb42bf9149a35508265b262022d665b0ff3425ea97df7f2d1cddb853174a21b4ed17050110f989e824c82ab12fb7f27217b10dbf75f52e62271a35dd50640e8fbbe628475327e4276",
		"95c30101d36e44367b4f0e5505b633305c24706efa6f05c9e4c796c5d896ecee2fe1c4fe6eeb2ea3c44a8becddaa7e80683da2999e4705e34cb80ba50923afb3c4efbf89e490f2948d7a3b541488e938cd205806f2072db5829298f676586d7acfd6df5bd6ef078cab969e29019e2643dbfc1688a886561506ae2a3d81df0a34a1825f25d47b24ebf1a608f046344c543f4cb1df8e8656603f296cd3fd4cb1b0d0431b90d6753b39e05abb852180ed6440a2e99667ede7fe8d",

		"fdca25408a51671d789e0e5f74de0090fb93591ce1674ec9fea36766989cd745c08a9a97966907fcbadfb549c4eda58e07a948b11b1c96976c134c663b29ff6e", // new uns
		"f2315d7f066df0dba0cefe2d600add468d4a99614abeacec3a1b40b014a02a34cba6b00542642715da27a80385f9b70626854fb90cc57588be343c64f82691c18f8ac3cc206cd1a9948861d2f5c5d9c294fb447005155737796780447c",
		"294ad1b18677272339790cd5ecc43cc02659b56ed05e53545e0b0b3886049e07442cd6f3afeea1eb3e39c59a30775fd3a3e15f41c69b8919f32bb77fe20a6843fce6e2e55fcb03eaae7d78b7b26588fc5d2d0dc4b9a2dcf3a974a3f61afbb8d7cb1e21141c8fac203eec912015dca54c1e00c1b9732b119b12fb5a7b1380be4823ebdbe578b37bf427b030e1f0a5d450fff4994a6b4bf272accad39af356473df4e6842c373387a2dc3ffdbedc6729de78177dde9d9aaf55b0",

		`49b51fc91d2668a475eb7bb5fdb7f923337f69c2028eade6f573a457927b3b74787768255a393ea67fe27d3765775b6fbeb5fe282a42046d3b27c45d35533a9ead9a8e90b79f1c3ed9d9e05dd0e4d9ad67ddfd8597366969559bf3191565d349cf32686d7b7defd7b510b3a45ba7e13da19abb6435002273bc1bfca5adfd85a198d87450735774e9ba976dd8595bcd39a877cec127e619247dd65c2618bc0a67c336533632f0ba1257c9de0f2daa5db54bc7e200de50e1a6139d94ccce0e8c9fd06d1c04a56852be539ab9cbab4f0096b4b690c85ac08c563f71489e896cd41f0642ba72122114f5d7c2e4d0859558a3dd90d2f687870e5a5a944fd0456fcbf0ccb04c82e581e34272eeaaba8f61a241140bae2702c1ac688b50d2ed11e420117911b82a7949482dff8ab5900a121ba73f59e9af03521c84b0b9`,
	}
)

func TestGetMur(t *testing.T) {
	eMl := "alikbinsk@yahoo.com"
	fmt.Println(utils.GetMurmur3Int64Pk(eMl))
}

func TestDecr(t *testing.T) {
	for _, x := range crs {
		s, e := utils.AesDecryptHex(x)
		fmt.Printf("%v\n\t[%s] (err:%v)\n", x, s, e)
	}
}

func TestGetXxhashHex(t *testing.T) {
	for i := 0; i < 20; i++ {
		fmt.Printf("%v\t[%s]\n", i, utils.GetXxhash64Hex(strconv.Itoa(i), i+10))
	}
}

func TestGetXxhashHexSimple(t *testing.T) {
	fmt.Printf("%s\n", utils.GetXxhash64Hex("bartleymonifa@yahoo.com", 16))
}

func TestGetXxhashHexFromEmails_Dbg(t *testing.T) {

	emails := []string{
		"candi5710@yahoo.com", "macdade77@yahoo.com", "amotheriam@gmail.com", "ashleypeck8@gmail.com", "pdbturton@yahoo.com",
		"j.smith2605@yahoo.com", "pjsoler180@yahoo.com", "butterfly91603@yahoo.com", "coachterrel@yahoo.com terry", "aded22963@yahoo.com",
		"philbeck35@gmail.com", "bdawn54@yahoo.com Belinda", "phatharris@yahoo.com", "ken.knott18@gmail.com", "anihil8@aol.com",
		"lorijaeger5720@yahoo.com", "nightcat1231@gmail.com", "rkremen@yahoo.com", "jsnyder6473@yahoo.com", "mestallsmith@gmail.com",
	}

	for _, eml := range emails {
		fmt.Printf("%v,", utils.GetXxhash64Hex(eml, 16))
	}
}

func TestGetXxhashHexFromEmails_Test(t *testing.T) {

	emails := []string{
		"mrjony11@yahoo.com", "mrjony33@yahoo.com", "eugenemcmad@gmail.com", "brain545@yandex.ru",
		"eugene.makushkin@regium.com", "denis.uspenskiy@regium.com",
	}

	for _, eml := range emails {
		fmt.Printf("%v,", utils.GetXxhash64Hex(eml, 16))
	}
}

func TestEncodeLen(t *testing.T) {
	for i := 1; i <= 10; i++ {

		params_a := fmt.Sprintf("%d/%d/%d/%d/%d/%d/%d/%s", i, i, i, i, i, i, i, "azazazazazazazzazazzazazazzzazaza")
		params_b := fmt.Sprintf("&a=%d&a=%d&a=%d&a=%d&a=%d&a=%d&a=%d&a=%s", i, i, i, i, i, i, i, "azazazazazazazzazazzazazazzzazaza")
		a, _ := utils.AesEncrypt(utils.Message{[]byte(params_a)})
		b, _ := utils.AesEncrypt(utils.Message{[]byte(params_b)})
		fmt.Printf("%d %d (%s %s)\n", len(a), len(b), a, b)

	}
}

func TestGetMd5(t *testing.T) {

	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	fmt.Printf("%x\n", h.Sum(nil))
	m, err := utils.GetMd5Str("The fog is getting thicker!")
	fmt.Printf("%s (%v)\n", m, err)
}

func TestGetMd5MT(t *testing.T) {

	anydata := []interface{}{

		int32(1234567890),
		"1234567890",
		[]byte("1234567890"),
	}

	for _, nd := range anydata {
		m5, err := utils.GetMd5Str(nd)
		fmt.Printf("%s (%v)\n", m5, err)
	}
}

func TestGetMurmur3Int64Pk(t *testing.T) {
	eMl := "Xr_Rx@Aol.com"
	eml := strings.ToLower(eMl)
	if utils.GetMurmur3Int64Pk(eMl) != utils.GetMurmur3Int64Pk(eml) {
		t.Fail()
	}
	if utils.GetMurmur3Int64PkStr(eMl) != utils.GetMurmur3Int64PkStr(eml) {
		t.Fail()
	}
	if utils.GetMurmur3Int64(eMl) == utils.GetMurmur3Int64(eml) {
		t.Fail()
	}
	if utils.GetMurmur3Int64Str(eMl) == utils.GetMurmur3Int64Str(eml) {
		t.Fail()
	}
}

func TestBlowfishDecrypt(t *testing.T) {
	encoded := `9fba1144122a222b72f3bc6b380dbbf1bf0ec8a9c3ca338c95a20b5d7fccca12c4525f18157d3cf6`
	fmt.Printf("encoded `s`:`%s`\n", encoded)
	b, err := hex.DecodeString(encoded)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := utils.BlowfishDecrypt(b)
	fmt.Printf("decoded `s`:`%s`, (err: %v)\n", decoded, err)
	decFin := strings.Trim(string(decoded), ` `)
	fmt.Printf("dec fin `s`:`%s`, (err: %v)\n", decFin, err)
	if decFin != `{"email":"eugene.makushkin@regium.com"}` {
		t.FailNow()
	}
}

func TestBlowfish(t *testing.T) {
	orig := `{"email":"test-test-test-ox_coc2@yahoo.com"}`
	fmt.Printf("original `s`: `%s`\n", orig)
	encoded, err := utils.BlowfishEncrypt([]byte(orig))
	fmt.Printf("encoded `x`:`%x`, (err: %v)\n", encoded, err)
	decoded, err := utils.BlowfishDecrypt(encoded)
	fmt.Printf("decoded `s`:`%s`, (err: %v)\n", decoded, err)
	decFin := strings.Trim(string(decoded), ` `)
	fmt.Printf("dec fin `s`:`%s`, (err: %v)\n", decFin, err)
	if orig != decFin {
		t.FailNow()
	}
}
