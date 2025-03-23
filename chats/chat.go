package chats

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ChatMessage struct {
	Room      string
	Username  string
	Message   string
	Timestamp int64
	Time      string
}

func Chat(chatData string, room string) (message ChatMessage, e error) {
	if room == "" {
		room = "lobby"
	}

	chatMsg := ChatMessage{}

	data := strings.SplitN(chatData, "|", 2)
	if len(data) < 2 {
		return chatMsg, fmt.Errorf("Expected %d arguments in chatData, received %d.", 2, len(data))
	}

	user := data[0]
	msg := data[1]

	chatMsg.Room = room
	chatMsg.Username = user
	chatMsg.Message = msg

	chatMsg.Timestamp = -1
	chatMsg.Time = ""

	return chatMsg, nil
}

func ChatTimestamp(chatData string, room string) (message ChatMessage, e error) {
	if room == "" {
		room = "lobby"
	}

	chatMsg := ChatMessage{}

	data := strings.SplitN(chatData, "|", 3)
	if len(data) < 3 {
		return chatMsg, fmt.Errorf("Expected %d arguments in chatData, received %d.", 3, len(data))
	}

	timestamp := data[0]
	user := data[1]
	msg := data[2]

	timestampInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return chatMsg, fmt.Errorf("Error when converting timestamp to integer: %s", err)
	}

	chatMsg.Room = room
	chatMsg.Username = user
	chatMsg.Message = msg

	chatMsg.Timestamp = timestampInt
	chatMsg.Time = time.Unix(timestampInt, 0).Local().Format(time.TimeOnly)

	return chatMsg, nil
}

type ChatMessagesSortable []ChatMessage

func (m ChatMessagesSortable) Len() int           { return len(m) }
func (m ChatMessagesSortable) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m ChatMessagesSortable) Less(i, j int) bool { return m[i].Timestamp < m[j].Timestamp }
