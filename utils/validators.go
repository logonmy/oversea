package utils

import (
	"regexp"
)


func IsEmail(email string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,})@(\w{1,}).([a-z]{2,4})$`, email); !m {
		return false
	}
	return true
}

func IsMobilePhone(phoneNumber string) bool {
	if m, _ := regexp.MatchString(`^(1[3|4|5|7|8|9][0-9]\d{4,8})$`, phoneNumber); !m {
		return false
	}
	return  true
}

func IsEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false
}

func IsIDCard(idcard string) bool {
	length := len(idcard)
	if length == 15 {
		if m, _ := regexp.MatchString(`^(\d{15})$`,idcard); !m {
			return false
		}
	} else if length == 18 {
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, idcard); !m {
			return false
		}
	} else {
		return false
	}
	return true
}