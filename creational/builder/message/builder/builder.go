package builder

import "fmt"

//Product
type Message struct {
	Body        string
	VoteOptions []string
	Type        string
	To          string
	Cc          string
	Secure      bool
}

type MessageSender interface {
	Send() MessageSender
	AddVote() MessageSender
	Build() (*Message, error)
}

type EmailBuilder struct {
	To          string
	Cc          string
	secure      bool
	MailText    string
	voteOptions []string
}

func (e *EmailBuilder) Send() MessageSender {
	e.secure = true
	fmt.Println("email is encrypted")

	return e
}

func (e *EmailBuilder) AddVote() MessageSender {
	e.voteOptions = []string{"yes", "no"}

	return e
}

func (e *EmailBuilder) Build() (*Message, error) {
	return &Message{
		To:          e.To,
		Cc:          e.Cc,
		Type:        "Email",
		Body:        e.MailText,
		Secure:      true,
		VoteOptions: []string{"yes, no"},
	}, nil
}
