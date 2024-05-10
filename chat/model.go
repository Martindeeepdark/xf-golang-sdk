package chat

type RequestHeader struct {
	AppID string `json:"app_id,omitempty"`
	UID   string `json:"uid,omitempty"`
}

type SystemSetting struct {
	Domain      string `json:"domain,omitempty"`
	Temperature string `json:"temperature,omitempty"`
	MaxToken    int    `json:"max_token,omitempty"`
	TopK        int    `json:"top_k,omitempty"`
	ChatID      string `json:"chat_id,omitempty"`
}

type XfapiStruct struct {
	AppID     string `json:"appid,omitempty"`
	APIKey    string `json:"apiKey,omitempty"`
	APISecret string `json:"apiSecret,omitempty"`
	HostUrl   string `json:"hostUrl,omitempty"`
}

type RequestParameter struct {
	Chat  *SystemSetting `json:"chat,omitempty"`
	Xfapi *XfapiStruct   `json:"xfapi,omitempty"`
}

type RequestText struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestMessage struct {
	Text []*RequestText `json:"text,omitempty"`
}

type RequestPayload struct {
	Message *RequestMessage `json:"message,omitempty"`
}

type Request struct {
	Header    *RequestHeader    `json:"header,omitempty"`
	Parameter *RequestParameter `json:"parameter,omitempty"`
	Payload   *RequestPayload   `json:"payload,omitempty"`
}

type ResponseHeader struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Sid     string `json:"sid,omitempty"`
	Status  int    `json:"status,omitempty"`
}

type ResponseChoicesText struct {
	Content string `json:"content"`
}

type ResponseChoices struct {
	Text   []ResponseChoicesText `json:"text"`
	Seq    int                   `json:"seq,omitempty"`
	Status int                   `json:"status,omitempty"`
}

type ResponsePayload struct {
	Choices *ResponseChoices `json:"choices,omitempty"`
}

type Response struct {
	Header  *ResponseHeader  `json:"header,omitempty"`
	Payload *ResponsePayload `json:"payload,omitempty"`
}
