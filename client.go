// client.go
package main

import (
	"bufio"
	"charm-psclient/commands"
	"charm-psclient/datastructs"
	"charm-psclient/utils"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	// Autoloads .env variables into os.Getenv
	_ "github.com/joho/godotenv/autoload"
)

var done chan bool
var interrupt chan os.Signal

var inputChannel chan string

var serverState datastructs.Server
var fileLogger *log.Logger

func processInput(inputChan <-chan string, done chan<- bool, connection *websocket.Conn) {
	for input := range inputChan {
		fileLogger.Printf("Input received: '%s'\n", input)

		if !strings.Contains(input, "|") {
			input = "|" + input
		}

		fileLogger.Printf("Sending message: ")
		connection.WriteMessage(websocket.TextMessage, []byte(input))
	}
	done <- true
}

func receiveHandler(connection *websocket.Conn) {
	defer close(done)
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			fileLogger.Println("Error in receive:", err)
			return
		}
		msgStr := string(msg)
		fileLogger.Printf("Received: %s\n", msgStr)

		var roomName string = "lobby"
		lines := strings.Split(msgStr, "\n")
		for _, line := range lines {
			if len(line) == 0 {
				fileLogger.Println("Empty line, skipping.")
				continue
			} else if strings.HasPrefix(line, ">") {
				// Room ID
				roomName = line[1:]
			} else if strings.HasPrefix(line, "|") {
				// |TYPE|DATA
				split := strings.SplitN(line, "|", 3)
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
				case "init":
					serverState.Rooms[roomName] = datastructs.Room{}
					fileLogger.Printf("Initializing room %s\n", roomName)
				case "title":
					room := serverState.Rooms[roomName]
					room.RoomName = messageData
					serverState.Rooms[roomName] = room
					fileLogger.Printf("Title of room %s: %s\n", roomName, messageData)
				case "users":
					room := serverState.Rooms[roomName]
					room.Users = make(map[string]datastructs.User)
					users := commands.Users(messageData)
					for _, user := range users {
						room.Users[user.Id] = user
					}
					serverState.Rooms[roomName] = room
					fileLogger.Printf("Users in room %s: %v\n", roomName, users)
				case "deinit":
					room := serverState.Rooms[roomName]
					room.Users = nil
					serverState.Rooms[roomName] = room
					delete(serverState.Rooms, roomName)
					fileLogger.Printf("Deinitializing room %s\n", roomName)
				case "challstr":
					data, err := commands.ChallStr(messageData, fileLogger)
					if err != nil {
						fileLogger.Printf("Error in challstr: %v\n", err)
					} else {
						trn := fmt.Sprintf("|/trn %s,0,%s", data.Curuser.Username, data.Assertion)
						fileLogger.Printf("Sending: %s\n", trn)
						connection.WriteMessage(websocket.TextMessage, []byte(trn))
					}
				case "chat", "c":
					chatMsg, err := commands.Chat(messageData, roomName)
					if err != nil {
						fileLogger.Printf("Error in chat: %v\n", err)
					} else {
						color := utils.UsernameToColor(chatMsg.Username.Username)
						fmt.Printf("(%s) %s%s: %s\n", color, chatMsg.Username.Group.Symbol,
							chatMsg.Username.Username, chatMsg.Message)
						fileLogger.Printf("New message in room %s: %v\n", roomName, chatMsg)
					}
					room, ok := serverState.Rooms[roomName]
					if ok {
						room.ChatMessages = append(room.ChatMessages, chatMsg)
					}
				case "chat:", "c:":
					chatMsg, err := commands.ChatTimestamp(messageData, roomName)
					if err != nil {
						fileLogger.Printf("Error in chat: %v\n", err)
					} else {
						color := utils.UsernameToColor(chatMsg.Username.Username)
						fmt.Printf("[%s] (%s) %s%s: %s\n", chatMsg.Time, color, chatMsg.Username.Group.Symbol,
							chatMsg.Username.Username, chatMsg.Message)
						fileLogger.Printf("New message in room %s: %v\n", roomName, chatMsg)
					}
					room, ok := serverState.Rooms[roomName]
					if ok {
						room.ChatMessages = append(room.ChatMessages, chatMsg)
					}
				case "queryresponse":
					responseSplit := strings.SplitN(messageData, "|", 2)
					queryType := responseSplit[0]
					queryJson := responseSplit[1]

					if queryType == "rooms" {
						roomData, err := commands.QueryresponseRooms(queryJson)
						if err != nil {
							fileLogger.Printf("Error in queryresponse rooms: %v\n", err)
						}
						serverState.RoomsInfo = roomData
						fileLogger.Printf("Roomdata: %v\n", roomData)
					} else {
						fileLogger.Printf("Unknown querytype %s\n", queryType)
					}
				default:
					fileLogger.Printf("Message type: %s\n", messageType)
					fileLogger.Printf("Message data: %s\n", messageData)
				}
			} else {
				fileLogger.Println(line)
			}
		}
	}
}

func main() {
	info, err := os.Stat("logs")
	if err != nil && os.IsNotExist(err) {
		err := os.Mkdir("logs", 0750)
		if err != nil {
			log.Fatalf("Error when trying to create logs folder: %v\n", err)
		}
	} else if !info.IsDir() {
		log.Fatal("logs already exists as file; please delete before rerunning.")
	}

	logfileName := fmt.Sprintf("logs/%s.log", time.Now().Format(time.DateTime))
	logfile, _ := os.OpenFile(logfileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer logfile.Close()

	// This will ensure that my logger will write UTF-8 to my log file
	utf8Encoder := unicode.UTF8.NewEncoder()
	writer := transform.NewWriter(logfile, utf8Encoder)

	fileLogger = log.New(writer, "INFO: ", log.LstdFlags)

	serverState = datastructs.Server{}
	serverState.Rooms = make(map[string]datastructs.Room)
	serverState.RoomsInfo = datastructs.RoomResponse{}

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
	fileLogger.Printf("Attempting to connect to %s\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fileLogger.Fatal("Error connecting to Websocket Server:", err)
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
				fileLogger.Println("Error reading input:", err)
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
				fileLogger.Println("Error during writing to websocket:", err)
				return
			}

		case msg := <-inputRead:
			inputChannel <- msg

		case <-interrupt:
			// We received a SIGINT (Ctrl + C). Terminate gracefully...
			fileLogger.Println("Received SIGINT interrupt signal. Closing all pending connections")

			// Close our websocket connection
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fileLogger.Println("Error during closing websocket:", err)
				return
			}

			select {
			case <-done:
				fileLogger.Println("Receiver Channel Closed! Exiting....")
			case <-time.After(time.Duration(1) * time.Second):
				fileLogger.Println("Timeout in closing receiving channel. Exiting....")
			}
			return
		}
	}
}
