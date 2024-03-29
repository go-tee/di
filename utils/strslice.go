package utils

import (
	"fmt"
)

type StringSlice []string

func (s *StringSlice) String() string {
	return fmt.Sprintf("%s", *s)
}

func (s *StringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}
