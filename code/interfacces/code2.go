package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}

	return d.priorityLevel
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch nt := n.(type) {
	case directMessage:
		return nt.senderUsername, nt.importance()
	case groupMessage:
		return nt.groupName, nt.importance()
	case systemAlert:
		return nt.alertCode, nt.importance()
	default:
		return "", 0
	}
}
