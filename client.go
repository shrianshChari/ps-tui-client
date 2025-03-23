// client.go
package main

import (
	"bufio"
	"charm-psclient/commands"
	"charm-psclient/utils"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	// Autoloads .env variables into os.Getenv
	_ "github.com/joho/godotenv/autoload"
)

var done chan bool
var interrupt chan os.Signal

var inputChannel chan string

func processInput(inputChan <-chan string, done chan<- bool, connection *websocket.Conn) {
	for input := range inputChan {
		log.Printf("Input received: '%s'\n", input)

		if !strings.Contains(input, "|") {
			input = "|" + input
		}

		log.Printf("Sending message: ")
		connection.WriteMessage(websocket.TextMessage, []byte(input))
	}
	done <- true
}

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		msgStr := string(msg)
		log.Printf("Received: %s\n\n", msgStr)

		var room string = "lobby"
		lines := strings.Split(msgStr, "\n")
		for _, line := range lines {
			if len(line) == 0 {
				log.Println("Empty line, skipping.")
				continue
			} else if strings.HasPrefix(line, ">") {
				// Room ID
				room = line[1:]
			} else if strings.HasPrefix(msgStr, "|") {
				// |TYPE|DATA
				split := strings.SplitN(line, "|", 3)
				log.Println(split, len(split))
				var messageType string
				var messageData string
				if len(split) == 3 {
					messageType = split[1]
					messageData = split[2]
				} else {
					messageType = split[1]
					messageData = ""
				}

				switch strings.ToLower(messageType) {
				case "challstr":
					data, err := chats.ChallStr(messageData)
					if err != nil {
						log.Printf("Error in challstr: %v\n", err)
					} else {
						trn := fmt.Sprintf("|/trn %s,0,%s", data.Curuser.Username, data.Assertion)
						log.Printf("Sending: %s\n", trn)
						connection.WriteMessage(websocket.TextMessage, []byte(trn))
					}
				case "chat", "c":
					chatMsg, err := chats.Chat(messageData, room)
					if err != nil {
						log.Printf("Error in chat: %v\n", err)
					} else {
						color := utils.UsernameToColor(chatMsg.Username)
						log.Printf("%s (%s): %s", chatMsg.Username, color, chatMsg.Message)
					}
				case "chat:", "c:":
					chatMsg, err := chats.ChatTimestamp(messageData, room)
					if err != nil {
						log.Printf("Error in chat: %v\n", err)
					} else {
						color := utils.UsernameToColor(chatMsg.Username)
						log.Printf("[%s] %s (%s): %s", chatMsg.Time, chatMsg.Username, color, chatMsg.Message)
					}
				}
				log.Printf("%s,%s\n", messageType, messageData)
			} else {
				log.Println(line)
			}
		}
		log.Println(room)
	}
}

func main() {
	done = make(chan bool)           // Channel to indicate that the receiverHandler is done
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully

	inputChannel = make(chan string)

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT

	var u url.URL

	server := os.Getenv("PS_SERVER")
	serverUrl := utils.SelectServer(server)
	if len(serverUrl) == 0 {
		serverUrl = server
	}

	if strings.Contains(serverUrl, "localhost") {
		u = url.URL{Scheme: "ws", Host: serverUrl, Path: "showdown/websocket"}
	} else {
		u = url.URL{Scheme: "wss", Host: serverUrl, Path: "showdown/websocket"}
	}
	log.Printf("Attempting to connect to %s\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()
	go receiveHandler(conn)
	go processInput(inputChannel, done, conn)

	inputRead := make(chan string)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Println("Error reading input:", err)
				return
			}
			input = strings.TrimSpace(input)
			inputRead <- input
		}
	}()

	// Our main loop for the client
	// We send our relevant packets here
	for {
		select {
		case <-time.After(time.Duration(10) * time.Millisecond * 1000):
			// Send an echo packet every second
			err := conn.WriteMessage(websocket.TextMessage, []byte("|/cmd rooms"))
			if err != nil {
				log.Println("Error during writing to websocket:", err)
				return
			}

		case msg := <-inputRead:
			inputChannel <- msg

		case <-interrupt:
			// We received a SIGINT (Ctrl + C). Terminate gracefully...
			log.Println("Received SIGINT interrupt signal. Closing all pending connections")

			// Close our websocket connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Error during closing websocket:", err)
				return
			}

			select {
			case <-done:
				log.Println("Receiver Channel Closed! Exiting....")
			case <-time.After(time.Duration(1) * time.Second):
				log.Println("Timeout in closing receiving channel. Exiting....")
			}
			return
		}
	}
}
