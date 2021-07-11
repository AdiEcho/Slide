package Function

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// CalA 返回A
func CalA(c []int64, s, gt, challenge string) string {
	tt := time.Now().UnixNano() / 1e6
	ee := [][]interface{}{
		{"move", 303, 418, 1563888260111, "pointermove"},
		{"move", 302, 419, 1563888260129, "pointermove"},
		{"move", 302, 418, 1563888260130, "mousemove"},
		{"move", 301, 419, 1563888260157, "pointermove"},
		{"move", 301, 418, 1563888260158, "mousemove"},
		{"move", 301, 419, 1563888260162, "pointermove"},
		{"move", 301, 419, 1563888260218, "pointermove"},
		{"move", 301, 419, 1563888260258, "pointermove"},
		{"move", 300, 419, 1563888260259, "mousemove"},
		{"move", 300, 419, 1563888260293, "pointermove"},
		{"move", 300, 420, 1563888260300, "pointermove"},
		{"move", 299, 420, 1563888260642, "pointermove"},
		{"move", 299, 421, 1563888260650, "pointermove"},
		{"move", 298, 421, 1563888260663, "pointermove"},
		{"move", 297, 421, 1563888260676, "pointermove"},
		{"move", 296, 423, 1563888260685, "pointermove"},
		{"move", 295, 423, 1563888260696, "pointermove"},
		{"move", 294, 425, 1563888260709, "pointermove"},
		{"move", 294, 424, 1563888260710, "mousemove"},
		{"move", 293, 427, 1563888260719, "pointermove"},
		{"move", 292, 427, 1563888260730, "pointermove"},
		{"move", 291, 430, 1563888260746, "pointermove"},
		{"move", 291, 431, 1563888260752, "pointermove"},
		{"move", 290, 432, 1563888260763, "pointermove"},
		{"move", 289, 434, 1563888260777, "pointermove"},
		{"move", 289, 435, 1563888260785, "pointermove"},
		{"move", 287, 440, 1563888260797, "pointermove"},
		{"move", 287, 440, 1563888260798, "mousemove"},
		{"move", 287, 441, 1563888260809, "pointermove"},
		{"move", 287, 441, 1563888260810, "mousemove"},
		{"move", 286, 443, 1563888260819, "pointermove"},
		{"move", 284, 449, 1563888260831, "pointermove"},
		{"move", 284, 448, 1563888260832, "mousemove"},
		{"move", 283, 450, 1563888260844, "pointermove"},
		{"move", 281, 453, 1563888260853, "pointermove"},
		{"move", 280, 453, 1563888260854, "mousemove"},
		{"move", 280, 454, 1563888260865, "pointermove"},
		{"move", 279, 455, 1563888260877, "pointermove"},
		{"move", 279, 455, 1563888260887, "pointermove"},
		{"down", 279, 455, 1563888261343, "pointerdown"},
		{"focus", 1563888261344},
		{"up", 279, 455, 1563888261450, "pointerup"},
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(ee); i++ {
		if len(ee[i]) >= 3 {
			ee[i][3] = tt
		} else {
			ee[i][1] = tt
		}
		tt += rand.Int63n(20)
	}
	fp, r := CalFPR(ee)
	n := CalN(r)
	t1233 := "6322!!7608!!CSS1Compat!!1!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!2!!3!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!1!!-1!!-1!!-1!!10!!44!!0!!0!!737!!784!!1687!!888!!zh-CN!!zh-CN,zh!!-1!!1.5!!24!!Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36!!1!!1!!1706!!960!!1707!!960!!1!!1!!1!!-1!!Linux x86_64!!0!!-8!!0a93728bbc5be4241e024b729f6c5e5d!!3f4cf59bdc7d0da206835d0e495e258a!!internal-pdf-viewer,mhjfbmdgcfjbbpaeojofohoefgiehjai!!0!!-1!!0!!8!!Arial,BitstreamVeraSansMono,Courier,CourierNew,Helvetica,Monaco,Times,TimesNewRoman,Wingdings,Wingdings2,Wingdings3!!1563801637024!!-1,-1,9,4,11,0,17,0,50,2,10,10,30,95,95,98,99,99,99,-1!!-1!!-1!!12!!-1!!-1!!-1!!5!!false!!false"
	// 这里有一部分特别麻烦的正则替换，不写了，不影响输出结果

	a := map[string]interface{}{
		"lang":          "zh-cn",
		"type":          "fullpage",
		"tt":            CalTT(n, s, c),
		"light":         "SPAN_0",
		"s":             "c7c3e21112fe4f741921cb3e4ff9f7cb",
		"h":             "779d8a2a74e1810ceed8cbec54cd2341",
		"hh":            "b32e52186851ec15ceaf3a01ded29473",
		"hi":            Hash(t1233),
		"ep":            CalEP(fp, gt, challenge),
		"captcha_token": "bboy",
		"passtime":      RandInt64(100, 300),
	}
	a["rp"] = Hash(gt + challenge + strconv.FormatInt(a["passtime"].(int64), 10))
	dataType, _ := json.Marshal(a)
	dataString := string(dataType)
	return dataString
}

// CalC 返回C
func CalC(e [][]interface{}) [][]interface{} {
	r := "pointer"
	s := e[:]
	size := len(s)
	a := []string{"move", "down", "up"}
	for c := size - 1; c >= 0; c-- {
		u := s[c]
		m := u[0].(string)
		for _, v := range a {
			if find := strings.Contains(v, m); find {
				if len(u) >= 5 && strings.Contains(u[4].(string), r) {
					s = append(s[:c], s[c+1:]...)
				}
			}
		}
	}
	return s
}

// CalT 返回T
func CalT(e int) int {
	t := 32767
	if t < e {
		e = t
	} else if e < -t {
		e = -t
	}
	return e
}

// CalFPR 返回fp和r
func CalFPR(e [][]interface{}) ([]interface{}, [][]interface{}) {
	var i = 0
	var r [][]interface{}
	var a []interface{}
	c := CalC(e)
	u := len(c)
	o := []string{"down", "move", "up", "scroll"}
	m := 0
	for idx := m; i <= u; idx++ {
		l := c[idx]
		h := l[0].(string)
		for _, v := range o {
			if find := strings.Contains(v, h); find {
				if len(a) == 0 {
					a = l
				}
				var appendR []interface{}
				appendR = append(appendR, h)
				appendR = append(appendR, l[1])
				appendR = append(appendR, l[2])
				appendR = append(appendR, CalT(i))
				r = append(r, appendR)
				i = int(l[3].(int64))
			}
		}
	}
	return a, r
}

// CalN 返回N
func CalN(e [][]interface{}) string {
	var t []interface{}
	var n, r, o []int
	for i := 0; i < len(e); i++ {
		s := e[i]
		c := len(s)
		t = append(t, s[0])
		if c == 2 {
			n = append(n, s[1].(int))
		} else {
			n = append(n, s[3].(int))
		}
		if c == 3 {
			r = append(r, s[1].(int))
			o = append(o, s[2].(int))
		}
	}
	u := "10000000001000101000111110001100000101000010100000000000000010000000000000001000000000000000"
	lenU := 92
	u += convertToBin(0, 6-lenU%6)
	s := "()*,-./0123456789:?@ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"
	c := ""
	d := lenU / 6
	for q := 0; q < d; q++ {
		side, _ := strconv.ParseInt(u[6*q:6*(q+1)], 10, 2)
		c += string(s[side])
	}
	return c
}

// convertToBin 将输入的十进制数字转换为二进制，fill左边补充0，返回字符串
func convertToBin(num int, fill int) string {
	s := ""

	if num == 0 {
		if fill != 0 {
			return strings.Repeat("0", fill)
		} else {
			return "0"
		}
	}

	for ; num > 0; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	if fill != 0 {
		lenS := len(s)
		if lenS >= fill {
			return s
		} else {
			return strings.Repeat("0", fill-lenS) + s
		}
	} else {
		return s
	}
}

// CalTT 返回tt
func CalTT(t, n string, e []int64) string {
	if t == "" || n == "" {
		return ""
	}
	o := t
	s := e[0]
	a := e[2]
	u := e[4]
	for i := 0; i < len(n); i += 2 {
		r := n[i : i+2]
		c, _ := strconv.ParseInt(r, 16, 10)
		m := string(int32(c))
		l := (s*c*c + a*c + u) % int64(len(t))
		o = o[:l] + m + o[l:]
	}
	return o
}

// CalEP 返回ep
func CalEP(fp []interface{}, gt, challenge string) map[string]interface{} {
	a := time.Now().UnixNano() / 1e6
	f := a + RandInt64(2, 8)
	b := a + RandInt64(50, 80)
	l := a + RandInt64(3, 9)
	m := l + RandInt64(30, 50)
	n := m + RandInt64(1, 5)
	o := n + RandInt64(10, 50)
	p := o + RandInt64(70, 90)
	r := p + RandInt64(10, 100)
	s := r + RandInt64(1, 2)
	tm := map[string]int64{
		"a": a,
		"b": b,
		"c": b,
		"d": 0,
		"e": 0,
		"f": f,
		"g": f,
		"h": f,
		"i": f,
		"j": f,
		"k": 0,
		"l": l,
		"m": m,
		"n": n,
		"o": o,
		"p": p,
		"q": p,
		"r": a,
		"s": s,
		"t": s,
		"u": s,
	}
	em := map[string]int{
		"cp": 0,
		"ek": 11,
		"nt": 0,
		"ph": 0,
		"sc": 0,
		"si": 0,
		"wd": 0,
	}
	e := map[string]interface{}{
		"ts": time.Now().UnixNano() / 1e6, "v": "8.7.9", "ip": "192.168.0.5,112.232.49.88", "f": Hash(gt + challenge),
		"de": false, "te": false, "me": true, "ven": "Intel Open Source Technology Center",
		"ren": "Intel(R) Iris Xe MAX Graphics", "ac": "32fyic1fps0s6ptcilehil6au4r3l9ir",
		"pu": false, "ph": false, "ni": false, "se": false, "fp": fp, "lp": fp, "em": em,
		"tm": tm, "by": 0,
	}
	return e
}

// RandInt64 返回输入区间的随机数
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

// Hash 对传入字符串进行md5加密
func Hash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
