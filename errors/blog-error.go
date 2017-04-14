package errors

import (
	"encoding/json"
	"fmt"

	tp "hankliu.com.cn/go-my-blog/types"
)

const (
	SEPARATOR string = "\n\n"
)

// BServerError blog project custom error
type BServerError struct {
	Code    tp.Code
	Message string
	Extra   map[string]interface{}
}

func (e BServerError) Error() string {
	errMsg := fmt.Sprintf("%d%s%s", e.Code, SEPARATOR, e.Message)

	if e.Extra != nil {
		buf, err := json.Marshal(e.Extra)

		if err == nil {
			errMsg = fmt.Sprintf("%s%s%s", errMsg, SEPARATOR, string(buf))
		}
	}

	return errMsg
}

// NewBServerError create blog server error without extra
func NewBServerError(code tp.Code, message string) error {
	return &BServerError{Code: code, Message: message}
}

// NewBServerErrorWithExtra create blog server error with extra
func NewBServerErrorWithExtra(code tp.Code, message string, extra tp.StringMap) error {
	bServerError := NewBServerError(code, message).(*BServerError)
	bServerError.Extra = extra
	return bServerError
}
