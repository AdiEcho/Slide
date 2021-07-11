package Encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"github.com/gookit/color"
	"math/big"
	mathRand "math/rand"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

type Encrypt struct {
}

// CreateSecretKey 创建随机16位16进制随机字符串
func CreateSecretKey() []byte {
	rnd := mathRand.New(mathRand.NewSource(time.Now().Unix()))
	temp := strconv.FormatUint(rnd.Uint64(), 16)

	// 以下是强转换，如果不想使用unsafe包的话可以注释这部分然后使用被注释的return
	sh := (*reflect.StringHeader)(unsafe.Pointer(&temp))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
	//return []byte(temp)
}

// RSAEncrypt 将传入参数进行RSA加密
func RSAEncrypt(text []byte) string {
	n := new(big.Int)
	n, _ = n.SetString("00C1E3934D1614465B33053E7F48EE4EC87B14B95EF88947713D25EECBFF7E74C7977D02DC1D9451F79DD5D1C10C29ACB6A9B4D6FB7D0A0279B6719E1772565F09AF627715919221AEF91899CAE08C0D686D748B20A3603BE2318CA6BC2B59706592A9219D0BF05C9F65023A21D2330807252AE0066D59CEEFA5F2748EA80BAB81", 16) // 这边网上的base是10，不知道会不会有影响
	PublicKey := rsa.PublicKey{
		N: n,
		E: 65537,
	}
	byteText, err := rsa.EncryptPKCS1v15(rand.Reader, &PublicKey, text)
	if err != nil {
		color.Errorln("RSAEncrypt Error: ", err)
		panic(err)
	}
	return hex.EncodeToString(byteText)
}

// AESEncrypt 对传入参数进行AES加密
func AESEncrypt(text []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		color.Errorln("NewCipher Error: ", err)
		panic(err)
	}
	blockSize := block.BlockSize()

	padding := blockSize - len(text)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	origData := append(text, padText...)
	iv := []byte("0000000000000000")
	blocMode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(origData))
	blocMode.CryptBlocks(encrypted, origData)

	return encrypted
}

// BytesToString 将[]byte经过解密转换为string
func BytesToString(array []byte) string {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789()"

	var o func(int) string
	var e func(int, int) int

	o = func(t int) string {
		if t < 0 || t > len(s) {
			return "."
		} else {
			return string(s[t])
		}
	}

	e = func(t int, e int) int {
		n := 0
		for r := 24; r >= 0; r-- {
			if 1 == (e >> r & 1) {
				n = (n << 1) + (t >> r & 1)
			}
		}
		return n
	}

	n := ""
	r := ""
	for a := 0; a < len(array); a += 3 {
		if a+2 < len(array) {
			u := (int(array[a]) << 16) + (int(array[a+1]) << 8) + int(array[a+2])
			n += o(e(u, 7274496)) + o(e(u, 9483264)) + o(e(u, 19220)) + o(e(u, 235))
		} else {
			c := len(array) % 3
			if c == 2 {
				u := (int(array[a]) << 16) + (int(array[a+1]) << 8)
				n += o(e(u, 7274496)) + o(e(u, 9483264)) + o(e(u, 19220))
				r = "."
			} else if c == 1 {
				u := int(array[a]) << 16
				n += o(e(u, 7274496)) + o(e(u, 9483264))
				r = ".."
			}
		}
	}
	return n + r

}
