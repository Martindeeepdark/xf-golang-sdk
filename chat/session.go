package chat

import "fmt"

type Session struct {
	svr *Server
	UID string
	Req *Request
}

func NewSession(svr *Server, uid string) *Session {
	session := &Session{
		svr: svr,
		UID: uid,
	}

	domain := getDomainFromRule(session.svr.hosturl)
	fmt.Println("Domain value:", domain) // 打印Domain的值

	session.Req = &Request{
		Header: &RequestHeader{
			AppID: session.svr.appID,
			UID:   session.UID,
		},
		Parameter: &RequestParameter{
			Chat: &SystemSetting{
				Domain: domain,
			},
		},
		Payload: &RequestPayload{
			Message: &RequestMessage{},
		},
	}

	return session
}

func getDomainFromRule(hostUrl string) string {
	//选择默认调用3.5模型接口，此后自行修改hostUrl
	fmt.Println("hostUrl value:", hostUrl)
	switch hostUrl {
	case "wss://spark-api.xf-yun.com/v3.5/chat":
		return "generalv3.5"
	case "wss://spark-api.xf-yun.com/v3.1/chat":
		return "generalv3"
	case "wss://spark-api.xf-yun.com/v2.1/chat":
		return "generalv2"
	case "wss://spark-api.xf-yun.com/v1.1/chat":
		return "general"
	default:
		return "general"
	}
}

func (s *Session) Send(question string) (string, error) {
	cmd := NewCommand(s)

	saveChatHistory(s.Req, RoleUser, question)

	answer, respErr := cmd.execute(s.Req)

	if respErr != nil {
		return "", respErr
	}

	saveChatHistory(s.Req, RoleAssistant, answer)

	return answer, nil
}

func saveChatHistory(req *Request, role string, content string) {
	history := req.Payload.Message.Text

	tokenCount := 0
	for _, text := range history {
		tokenCount += len(text.Content)
	}

	tokenCount += len(content)

	if tokenCount > MaxTokenSize {
		cutIndex := len(history)

		for index, text := range history {
			tokenCount -= len(text.Content)
			if tokenCount <= MaxTokenSize {
				cutIndex = index + 1
				break
			}
		}

		copy(history[:len(history)-cutIndex], history[cutIndex:])
	}

	history = append(history, &RequestText{Role: role, Content: content})

	req.Payload.Message.Text = history
}
