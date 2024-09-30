package utils

import (
	"fmt"
	"regexp"
)

func IsValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	return re.MatchString(email)
}

func IsValidPassword(password string) error {
	if len(password) < 6 {
		return fmt.Errorf("密码少于 6 位，请检查！")
	}

	// TODO: 增强密码校验
	// feat: 密码必须包含至少一个大写字母、一个小写字母、一个数字和一个特殊字符
	passwordPattern := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{6,}$`
	re := regexp.MustCompile(passwordPattern)
	if !re.MatchString(password) {
		return fmt.Errorf("密码不符合要求，请检查！")
	}

	return nil
}
