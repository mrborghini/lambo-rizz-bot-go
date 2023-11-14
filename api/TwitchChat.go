package api

import (
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type TwitchChat struct {
	nickname string
	oAuth    string
	channels []string
	log      *Logger
	conn     *websocket.Conn
	mu       sync.Mutex
}

const (
	twitchIRCServer = "irc-ws.chat.twitch.tv"
	twitchIRCPort   = 443
)

func NewTwitchChat(Nickname string, OAuth string) *TwitchChat {
	return &TwitchChat{
		nickname: Nickname,
		oAuth:    OAuth,
		channels: make([]string, 0),
		log:      NewLogger("TwitchChat"),
	}
}

func (tc *TwitchChat) JoinChat(Channel string) {
	tc.channels = append(tc.channels, Channel)
}

func (tc *TwitchChat) ConnectAsync() {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	u := url.URL{
		Scheme: "wss",
		Host:   fmt.Sprintf("%s:%d", twitchIRCServer, twitchIRCPort),
		Path:   "/",
	}
	tc.log.Info("Attempting to connect...")
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		tc.log.Error(fmt.Sprintf("Error connecting to Twitch IRC: %s", err))
		return
	}

	tc.conn = conn

	if tc.conn == nil {
		tc.log.Error("WebSocket connection is nil")
	}

	sendOAuth := fmt.Sprintf("PASS oauth:%s", tc.oAuth)
	if err := tc.conn.WriteMessage(websocket.TextMessage, []byte(sendOAuth)); err != nil {
		tc.log.Error(fmt.Sprintf("Error sending OAuth credentials: %s", err))
		return
	}

	sendNickname := fmt.Sprintf("NICK %s", tc.nickname)
	if err := tc.conn.WriteMessage(websocket.TextMessage, []byte(sendNickname)); err != nil {
		tc.log.Error(fmt.Sprintf("Error sending nickname: %s", err))
		return
	}

	for _, channel := range tc.channels {
		joinChannel := fmt.Sprintf("JOIN #%s", channel)
		if err := tc.conn.WriteMessage(websocket.TextMessage, []byte(joinChannel)); err != nil {
			tc.log.Error(fmt.Sprintf("Error joining channel %s: %s", channel, err))
			return
		}
	}

	tc.log.Info(fmt.Sprintf("Connected successfully %s", tc.nickname))
}

func (tc *TwitchChat) SendPrivMSG(message string, channel string) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	sendmessage := fmt.Sprintf("PRIVMSG %s :%s", channel, message)
	if err := tc.conn.WriteMessage(websocket.TextMessage, []byte(sendmessage)); err != nil {
		tc.log.Error(fmt.Sprintf("Error joining channel %s: %s", channel, err))
		return
	}
}

func (tc *TwitchChat) SendTwitchFunc(message string) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	sendmessage := message
	if err := tc.conn.WriteMessage(websocket.TextMessage, []byte(sendmessage)); err != nil {
		tc.log.Error(fmt.Sprintf("Error joining channel %s", err))
		return
	}
}

func (tc *TwitchChat) ReceiveSingleMessage() (string, error) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	_, msg, err := tc.conn.ReadMessage()

	if err != nil {
		tc.log.Error(fmt.Sprintf("%s", err))
		return "", nil
	} else {
		return string(msg), nil
	}

}

func (tc *TwitchChat) StartListeningForMessages() {
	for {
		message, err := tc.ReceiveSingleMessage()
		if err != nil {
			tc.log.Error(fmt.Sprintf("Error receiving message: %s", err))
			return
		} else {
			switch true {
			case strings.Contains(message, "PING"):
				tc.SendTwitchFunc("PONG")
			case strings.Contains(message, "PRIVMSG"):

				cleanmessage := GetMessage(message)
				channel := GetChannel(message)
				username := GetUsername(message)

				if strings.Contains(strings.ToLower(cleanmessage), "rizz") {
					result := Rizz(username)
					tc.SendPrivMSG(result, channel)
					tc.log.Info(fmt.Sprintf("%s in channel: %s message: %s", result, channel, cleanmessage))
				}
			}
		}
	}
}
