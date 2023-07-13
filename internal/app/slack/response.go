package slack

type Response struct {
	OK               bool             `json:"ok"`
	Channel          string           `json:"channel,omitempty"`
	Message          Message          `json:"message,omitempty"`
	Error            string           `json:"error,omitempty"`
	Errors           []string         `json:"errors,omitempty"`
	Warning          string           `json:"warning,omitempty"`
	ResponseMetadata ResponseMetadata `json:"response_metadata,omitempty"`
}

type Message struct {
	BotID      string     `json:"bot_id"`
	Type       string     `json:"type"`
	Text       string     `json:"text"`
	User       string     `json:"user"`
	AppID      string     `json:"app_id"`
	Team       string     `json:"team"`
	BotProfile BotProfile `json:"bot_profile"`
}

type BotProfile struct {
	ID     string `json:"bot_profile"`
	AppID  string `json:"app_id"`
	Name   string `json:"name"`
	TeamID string `json:"team_id"`
}

type ResponseMetadata struct {
	Messages []string `json:"messages,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}
