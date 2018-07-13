package downcase

import "testing"

func toLower(s string) string {
  b := make([]byte, len(s))
  for i := range b {
	c := s[i]
	if c >= 'A' && c <= 'Z' {
		c += 'a' - 'A'
	}
    b[i] = c
  }
return string(b)
}
