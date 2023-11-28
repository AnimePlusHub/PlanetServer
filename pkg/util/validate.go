package util

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

func ValidateEmail(email string) bool {
	// str := "^[w.-]+@[a-zA-Z0-9]+.[a-zA-Z]{2,4}$"
	// regex := regexp.MustCompile(str)
	// return regex.MatchString(email)
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
