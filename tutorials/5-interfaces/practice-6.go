package main

// SendMessage demonstrates the use of an interface to format messages in different styles.
func messagesFormatting(){
	// Create a plain text message
	plainText := PlainText{message: "Hello, World!"}

	// Create a bold message
	boldMessage := Bold{message: "Hello, World!"}

	// Create a code message
	codeMessage := Code{message: "Hello, World!"}

	// Send messages using the Formatter interface
	println(SendMessage(plainText)) // Output: Hello, World!
	println(SendMessage(boldMessage)) // Output: **Hello, World!**
	println(SendMessage(codeMessage)) // Output: `Hello, World!`
}

func SendMessage(formatter Formatter) string {
	return formatter.Format() 
}

type PlainText struct{
	message string
}

type Bold struct{
	message string
}

type Code struct{
	message string
}

func (p PlainText) Format() string {
	return p.message
}
func (b Bold) Format() string {
	return "**" + b.message + "**"
}
func (c Code) Format() string {
	return "`" + c.message + "`"
}

type Formatter interface {
	Format() string
}

