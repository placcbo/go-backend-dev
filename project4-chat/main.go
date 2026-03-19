package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type Message struct {
	ID     int
	Text   string
	Sender User
}

type ChatRoom struct {
	Name     string
	Messages []Message
}

func NewMessage(id int, sender User, text string) (Message, error) {
	if text == "" {
		return Message{}, errors.New("message text cannot be empty")
	}

	message := Message{
		ID:     id,
		Text:   text,
		Sender: sender,
	}

	return message, nil
}

func (r ChatRoom) MessageCount() int {
	return len(r.Messages)
}

func (r *ChatRoom) AddMessage(message Message) {
	r.Messages = append(r.Messages, message)
}

func (r ChatRoom) PrintMessages() {
	fmt.Println("Chat history for room:", r.Name)

	for _, msg := range r.Messages {
		fmt.Println("[" + msg.Sender.Name + "] " + msg.Text)
	}
}

func main() {
	user1 := User{ID: 1, Name: "Kevin"}
	user2 := User{ID: 2, Name: "Alice"}

	message1, err := NewMessage(1, user1, "Hello Alice")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	message2, err := NewMessage(2, user2, "Hey Kevin")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	room := ChatRoom{
		Name: "General",
		Messages: []Message{
			message1,
			message2,
		},
	}

	fmt.Println("Room:", room.Name)
	fmt.Println("Message count:", room.MessageCount())

	message3, err := NewMessage(3, user1, "Now I understand errors better")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	room.AddMessage(message3)

	fmt.Println("Message count after adding:", room.MessageCount())
	room.PrintMessages()
}
