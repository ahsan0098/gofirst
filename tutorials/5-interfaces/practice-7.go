package main

func processNotification() {
	// Create a direct message
	dm := directMessage{
		senderUsername: "john_doe",
		messageContent: "Hello, how are you?",
		priorityLevel:  10,
		isUrgent:       true,
	}

	// Create a group message
	gm := groupMessage{
		groupName:      "developers",
		messageContent: "New project update available.",
		priorityLevel:  20,
	}

	// Create a system alert
	sa := systemAlert{
		alertCode:      "SYS001",
		messageContent: "System maintenance scheduled for tonight.",
	}

	// Process notifications
	receiver, importance := processing(dm)
	println("Direct Message Receiver:", receiver, "Importance Level:", importance)
	receiver, importance = processing(gm)
	println("Group Message Receiver:", receiver, "Importance Level:", importance)
	receiver, importance = processing(sa)
	println("System Alert Receiver:", receiver, "Importance Level:", importance)
}

func processing(n notification) (string, int) {
	switch n := n.(type) {
	case directMessage:
		return n.senderUsername, n.importance()
	case groupMessage:
		return n.groupName, n.importance()
	case systemAlert:
		return n.alertCode, n.importance()
	default:
		return "Unknown", 0
	}
}

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return 50
	}
	return dm.priorityLevel
}

func (dm groupMessage) importance() int {
	return dm.priorityLevel
}

func (dm systemAlert) importance() int {
	return 100
}


