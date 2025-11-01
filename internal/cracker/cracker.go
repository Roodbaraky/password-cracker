package cracker

import (
	"crypto/md5"
	"strings"
)

type Config struct {
	Min          int
	Max          int
	SpecialChars bool
	Digits       bool
	Capitals     bool
}

func Crack(data [16]byte, passwordConfig Config) []string {
	charset := "abcdefghijklmnopqrstuvwxyz"
	matchingHashes := []string{}
	if passwordConfig.Capitals {
		charset += strings.ToUpper(charset)
	}
	if passwordConfig.Digits {
		charset += "0123456789"
	}
	if passwordConfig.SpecialChars {
		charset += "!@#$%^&*()_+"
	}

	charsetLen := len(charset)

	for length := passwordConfig.Min; length <= passwordConfig.Max; length++ {
		indices := make([]int, length)
		for {
			combination := make([]byte, length)
			for i := 0; i < length; i++ {
				combination[i] = charset[indices[i]]
			}

			if md5.Sum(combination) == data {
				matchingHashes = append(matchingHashes, string(combination))
			}

			pos := length - 1
			for pos >= 0 {
				indices[pos]++
				if indices[pos] < charsetLen {
					break
				}
				indices[pos] = 0
				pos--
			}

			if pos < 0 {
				break
			}
		}
	}
	return matchingHashes
}
