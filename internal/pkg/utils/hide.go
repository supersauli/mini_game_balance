package utils

import (
	"regexp"
	"strings"
)

func HidePhoneNumber(phoneNumber string) string {
	if phoneNumber == "" {
		return ""
	}
	const mobilePattern = `^(\d{3})\d{4}(\d{4})$`

	// 隐藏手机号中间四位
	mobileRegex, _ := regexp.Compile(mobilePattern)
	return mobileRegex.ReplaceAllString(phoneNumber, "$1****$2")
}

func HideMail(email string) string {
	if email == "" {
		return ""
	}

	//const emailPattern = `(?<=.{2}).+(?=.{2}@)|(?<=@.{1})/g`
	//// 隐藏邮箱账号中间的部分
	//emailRegex, _ := regexp.Compile(emailPattern)
	//return emailRegex.ReplaceAllString(mail, "*")

	//re := regexp.MustCompile(`(?i)[A-Z0-9._%+\-]+@([A-Z0-9.\-]+\.[A-Z]{2,})`)
	//output := re.ReplaceAllStringFunc(mail, func(s string) string {
	//	atSignIndex := len(s) - strings.Index(s, "@") - 1
	//	r := []rune(s)
	//	return fmt.Sprintf("%s***%c%s", string(r[0:2]), r[atSignIndex-1], string(r[atSignIndex:]))
	//})

	//re := regexp.MustCompile(`(\w)(\w+)(\w)@`)
	//replacement := "$1***$3@"
	//return re.ReplaceAllString(email, replacement)

	//return output

	atIndex := strings.Index(email, "@")
	if atIndex > 3 {
		visiblePart := email[:atIndex-4]
		hiddenPart := strings.Repeat("*", 3)
		return visiblePart + hiddenPart + email[atIndex-1:]
	}
	return email
}

func HideCardID(cardID string) string {

	if cardID == "" {
		return ""
	}
	return cardID[0:3] + "***********" + cardID[14:]
}

func HidName(name string) string {
	if name == "" {
		return ""
	}
	var ret []rune
	var last rune
	for _, v := range name {
		ret = append(ret, '*')
		last = v
	}
	ret[len(ret)-1] = last
	return string(ret)
}
