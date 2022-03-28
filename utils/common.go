package utils

import (
	"strings"
)

const (
	prefixAll = "/**"
)

var (
	Common Commons
)

func init() {
	Common = NewCommons()
}

type Commons struct{}

func NewCommons() Commons {
	return Commons{}
}

func (c *Commons) AntCheck(cl string, cr string) bool {
	s := strings.TrimSpace(cl)
	t := strings.TrimSpace(cr)
	if strings.EqualFold(s, t) {
		return true
	}
	if strings.EqualFold(s+prefixAll, t) {
		return true
	}
	return false
}
