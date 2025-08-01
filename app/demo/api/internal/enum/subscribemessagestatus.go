package enum

//go:generate enumer -type=SubscribeMessageStatus -linecomment -json -sql

type SubscribeMessageStatus int

const (
	// SubscribeMessageUnsent 未发送
	SubscribeMessageUnsent SubscribeMessageStatus = iota // unsent
	// SubscribeMessageSent 已发送
	SubscribeMessageSent // sent
	// SubscribeMessageFailed 发送失败
	SubscribeMessageFailed // failed
)
