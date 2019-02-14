package format

import (
	"regexp"
)

var (
	phoneRegCheck   = regexp.MustCompile("^(\\+7|8)\\d+$")
	phoneRegReplace = regexp.MustCompile("[^+0-9]")
	defaultCode     = "+7"
)

func getCode(code ...string) string {
	if len(code) == 0 {
		return defaultCode
	} else {
		return code[0]
	}
}

func Phone(source string, code ...string) (phone string, check bool) {
	phone = phoneRegReplace.ReplaceAllString(source, "")
	if check = phoneRegCheck.MatchString(phone); check && phone[0] != '+' {
		phone = getCode(code...) + phone[1:]
	}
	return
}

func PhoneLength(source string, length int, code ...string) (phone string, check bool) {
	if phone, check = Phone(source, code...); check {
		check = len(phone) == length
	}
	return
}

func PhoneMinMax(source string, min, max int, code ...string) (phone string, check bool) {
	if phone, check = Phone(source, code...); check {
		check = len(phone) > min && len(phone) < max
	}
	return
}
