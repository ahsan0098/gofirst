package main

import "fmt"

func L6() {
	messages := []Message{
		TextMessage{"Alice", "Hello, World!"},
		MediaMessage{"Bob", "image", "A beautiful sunset"},
		LinkMessage{"Charlie", "http://example.com", "Example Domain"},
		TextMessage{"Dave", "Another text message"},
		MediaMessage{"Eve", "video", "Cute cat video"},
		LinkMessage{"Frank", "https://boot.dev", "Learn Coding Online"},
	}

	var arr []Message

	arr = filterMessages(messages, "text")
	fmt.Println(arr)

	arr = filterMessages(messages, "media")
	fmt.Println(arr)

	arr = filterMessages(messages, "link")
	fmt.Println(arr)
}

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	var arr []Message
	for _, msg := range messages {
		if msg.Type() == filterType {

			arr = append(arr, msg)
		}
	}

	return arr

}
