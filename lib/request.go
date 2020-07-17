package lib

import (
	"encoding/base64"
	"strings"
)

func IsPublicEndpoint(path string) bool {
	listPublic := []string{
		"/auth/token",
		"/sys/ping",
	}

	for _, e := range listPublic {
		if strings.HasPrefix(path, e) {
			return true
		}
	}

	return false
}

func ParseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}
