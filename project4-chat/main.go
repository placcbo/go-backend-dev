package main

import "fmt"

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

func NewMessage(id int, sender User, text string) Message {
	if text == "" {
		fmt.Println("warning: empty message text")
	}

	return Message{
		ID:     id,
		Text:   text,
		Sender: sender,
	}
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

	message1 := NewMessage(1, user1, "Hello Alice")
	message2 := NewMessage(2, user2, "Hey Kevin")

	room := ChatRoom{
		Name: "General",
		Messages: []Message{
			message1,
			message2,
		},
	}

	fmt.Println("Room:", room.Name)
	fmt.Println("Message count:", room.MessageCount())

	message3 := NewMessage(3, user1, "Today I am learning functions and methods")
	room.AddMessage(message3)

	fmt.Println("Message count after adding:", room.MessageCount())
	room.PrintMessages()

	emptyMessage := NewMessage(4, user2, "")
	fmt.Println(emptyMessage)
}
