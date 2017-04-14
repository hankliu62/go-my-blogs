package constants

import (
	"regexp"
)

const (
	// MAX_LIMIT db query max limiter
	MAX_LIMIT uint32 = 999999

	// PASSWORD_REG password regexp string
	PASSWORD_REG = "^[a-zA-Z][0-9a-zA-Z_]{7,19}$"

	// Pagination
	DEFAULT_PAGE_SIZE  uint32 = 20
	DEFAULT_PAGE_INDEX uint32 = 1

	// OrderBy
	ASC  string = "asc"
	DESC string = "desc"
)

var (
	// PasswordReg password regexp
	PasswordReg, _ = regexp.Compile(PASSWORD_REG)
)
