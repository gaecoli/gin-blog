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

	return nil
}