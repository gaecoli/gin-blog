package utils

import "testing"

func TestSendEmail(t *testing.T) {
	err := SendEmail("dj <xx@qq.com>", []string{"xxxx@gmail.com"}, "Test", []byte("Test"), "text")
	if err != nil {
		t.Errorf("SendEmail() error = %v", err)
		return
	}
}
