package functions

import "github.com/segmentio/ksuid"

func NewTraceId() string {
	return ksuid.New().String()
}