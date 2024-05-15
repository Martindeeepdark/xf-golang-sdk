package main

import (
	"bufio"
	"fmt"
	"github.com/Martindeeepdark/xf-golang-sdk/chat"
	"os"
)

func main() {
	appID := chat.EnvString("APP_ID")
	apiKey := chat.EnvString("API_KEY")
	apiSecret := chat.EnvString("API_SECRET")
	hostUrl := chat.EnvString("HOST_URL")
	if appID == "" || apiKey == "" || apiSecret == "" {
		panic("APP_ID, API_KEY, API_SECRET,HOST_URL are required")
	}
	s := chat.NewServer(appID, apiKey, apiSecret, hostUrl)
	session, sessionErr := s.GetSession("123456789")

	if sessionErr != nil {
		panic(sessionErr)
	}

	answer := ""
	var err error
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You: ")
		question, _ := reader.ReadString('\n')
		answer, err = session.Send(question)

		if err != nil {
			panic(err)
		}

		fmt.Println("AI: ", answer)
	}
}
