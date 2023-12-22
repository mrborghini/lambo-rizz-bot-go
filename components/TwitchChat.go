package api

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"github.com/gorilla/websocket"
)

// Get the required data to start the bot
type TwitchChat struct {
	nickname string
	oAuth    string
	channels []string
	log      *Logger
	conn     *websocket.Conn
	mu       sync.Mutex
}

// The server it will be connecting to
const (
	twitchIRCServer = "irc-ws.chat.twitch.tv"
	twitchIRCPort   = 443
)

// Constructor for a new object for twitch
func NewTwitchChat(Nickname string, OAuth string) *TwitchChat {
	return &TwitchChat{
		nickname: Nickname,
		oAuth:    OAuth,
		channels: make([]string, 0),
		log:      NewLogger("TwitchChat"),
	}
}

// Add channels to the Twitch channel list
func (tc *TwitchChat) JoinChat(Channel string) {
	tc.channels = append(tc.channels, Channel)
}

// Connect to twitch servers
func (tc *TwitchChat) ConnectAsync() {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	// Create a new websocket url
	u := url.URL{
		Scheme: "wss",
		Host:   fmt.Sprintf("%s:%d", twitchIRCServer, twitchIRCPort),
		Path:   "/",
	}

	tc.log.Info("Attempting to connect...")

	// Connect to the websocket
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		tc.log.Error(fmt.Sprintf("Error connecting to Twitch IRC: %s", err))
		os.Exit(1)
	}

	// Add connection to the property
	tc.conn = conn

	// Check if the connection has been established correctly
	if tc.conn == nil {
		tc.log.Error("WebSocket connection is nil")
		os.Exit(1)
	}

	// Send the OAuth token to Twitch
	sendOAuth := fmt.Sprintf("PASS oauth:%s", tc.oAuth)
	if err := tc.conn.WriteMessage(websocket.TextMessage, []byte(sendOAuth)); err != nil {
		tc.log.Error(fmt.Sprintf("Error sending OAuth credentials: %s", err))
		os.Exit(1)
	}

	// Send the bot's name to Twitch
	sendNickname := fmt.Sprintf("NICK %s", tc.nickname)
	if err := tc.conn.WriteMessage(websocket.TextMessage, []byte(sendNickname)); err != nil {
		tc.log.Error(fmt.Sprintf("Error sending nickname: %s", err))
		os.Exit(1)
	}

	// Join all the channels in the channel list
	for _, channel := range tc.channels {
		joinChannel := fmt.Sprintf("JOIN #%s", channel)
		tc.log.Info(fmt.Sprintf("Joined channel: %s", channel));
		if err := tc.conn.WriteMessage(websocket.TextMessage, []byte(joinChannel)); err != nil {
			tc.log.Error(fmt.Sprintf("Error joining channel %s: %s", channel, err))
			os.Exit(1)
		}
	}

	tc.log.Info(fmt.Sprintf("Connected successfully %s", tc.nickname))
}

// Send a message in someone's twitch channel
func (tc *TwitchChat) SendPrivMSG(message string, channel string) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	sendmessage := fmt.Sprintf("PRIVMSG %s :%s", channel, message)
	if err := tc.conn.WriteMessage(websocket.TextMessage, []byte(sendmessage)); err != nil {
		tc.log.Error(fmt.Sprintf("Error joining channel %s: %s", channel, err))
		os.Exit(1)
	}
}

// Send Twitch functions shown here: https://dev.twitch.tv/docs/irc/#supported-irc-messages

func (tc *TwitchChat) SendTwitchFunc(message string) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	sendmessage := message
	if err := tc.conn.WriteMessage(websocket.TextMessage, []byte(sendmessage)); err != nil {
		tc.log.Error(fmt.Sprintf("Error joining channel %s", err))
		os.Exit(1)
	}
}

// Get a single message from Twitch

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

// Wait for messages to get sent

func (tc *TwitchChat) StartListeningForMessages() {
	for {
		message, err := tc.ReceiveSingleMessage()

		// Check if oauth token is valid
		if strings.TrimSpace(message) == ":tmi.twitch.tv NOTICE * :Login authentication failed" {
			tc.log.Error("Invalid oauth token");
			os.Exit(1)
		}

		if err != nil {
			tc.log.Error(fmt.Sprintf("Error receiving message: %s", err))
			os.Exit(1)
		}
		if tc.conn == nil || message == "" {
			tc.log.Error("WebSocket connection is nil")
			os.Exit(1)
		}

		switch true {
		// When twitch sends PING respond with PONG
		case strings.Contains(message, "PING"):
			tc.SendTwitchFunc("PONG")
		// Check if a message gets send in someones twitch chat
		case strings.Contains(message, "PRIVMSG"):
			// Only get the message
			cleanmessage := GetMessage(message)

			// Only get the channel where the message got sent
			channel := GetChannel(message)

			// Get the username of the person that sent the message
			username := GetUsername(message)

			// Check if a message in someones chat contains rizz. If so respond with their rizz levels en log it
			if strings.Contains(strings.ToLower(cleanmessage), "rizz") {
				result := Rizz(username)
				tc.SendPrivMSG(result, channel)
				tc.log.Info(fmt.Sprintf("%s in channel: %s message: %s", result, channel, cleanmessage))
			}
		}
	}
}
