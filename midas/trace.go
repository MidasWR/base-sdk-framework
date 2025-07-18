package midas

import (
	"github.com/sony/sonyflake"
	"strconv"
)

type SonyConfig struct {
	Sf sonyflake.Sonyflake
}

func (s *SonyConfig) GenerateUUID() (string, error) {
	id, err := s.Sf.NextID()
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(id, 10), nil
}
