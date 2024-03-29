package messages

import (
	"discord-metrics-server/v2/users"
)

type DiscordMessageID struct {
	MessageID string `uri:"id" binding:"required"`
}

type DiscordMessageQuery struct {
	PageSize   int    `form:"page_size"`
	PageNumber int    `form:"page_number"`
	UserID     string `form:"user_id"`
	ChannelID  string `form:"channel_id"`
}

func (q DiscordMessageQuery) validate() *string {
	if q.PageSize > 100 {
		ret := "Page size cannot be greater than 100"
		return &ret
	}
	if q.PageSize < 1 {
		ret := "Page size must be at least 1"
		return &ret
	}
	if q.PageNumber < 1 {
		ret := "Page number must be at least 1"
		return &ret
	}
	return nil
}

type DiscordMessage struct {
	UserID    string `json:"user_id"`
	Contents  string `json:"contents"`
	SentAt    string `json:"sent_at"`
	MessageID string `json:"message_id"`
	ChannelID string `json:"channel_id"`
	InReplyTo string `json:"in_reply_to"`
}

type UpdateDiscordMessage struct {
	Contents *string `json:"contents,omitempty"`
}

type DiscordMessageUserMention struct {
	MessageID string `json:"message_id"`
	UserID    string `json:"user_id"`
}

type DiscordMessageResponse struct {
	ID             int               `json:"id"`
	User           users.DiscordUser `json:"user"`
	Contents       string            `json:"contents"`
	SentAt         string            `json:"sent_at"`
	MessageID      string            `json:"message_id"`
	ChannelID      string            `json:"channel_id"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	InReplyTo      string            `json:"in_reply_to"`
	UsersMentioned []string          `json:"users_mentioned"`
}
