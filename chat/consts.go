package chat

import "errors"

const (
	chatAPIAddr = "https://spark-api.xf-yun.com/v3.5/chat"
)

const (
	RoleAssistant = "assistant"
	RoleUser      = "user"
)

const (
	StatusFirst  = 0 // 首次结果
	StatusMiddle = 1 // 中间结果
	StatusLast   = 2 // 最后一个结果
)

const (
	defaultTimeoutSeconds = 10
)

const (
	MaxTokenSize = 8192
)

var ErrTimeout = errors.New("timeout")
