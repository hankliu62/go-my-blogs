package codes

import (
	"hankliu.com.cn/go-my-blog/errors"
	tp "hankliu.com.cn/go-my-blog/types"
)

const (
	CommonMissingRequiredFields tp.Code = 1000001
	OldPasswordNotMatched       tp.Code = 1000002
	InvalidPasswordFormat       tp.Code = 1000003
)

var codeMessage = map[tp.Code]string{
	CommonMissingRequiredFields: "Missing required fields",
	OldPasswordNotMatched:       "Current passworld not matched",
	InvalidPasswordFormat:       "Invalid password format",
}

// NewError create error without extra
func NewError(code tp.Code) error {
	return errors.NewBServerError(code, codeMessage[code])
}

// NewErrorWithExtra create error with extra
func NewErrorWithExtra(code tp.Code, extra tp.StringMap) error {
	return errors.NewBServerErrorWithExtra(code, codeMessage[code], extra)
}
