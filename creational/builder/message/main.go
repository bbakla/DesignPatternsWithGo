package main

import (
	"fmt"
	"github.com/bbakla/DesignPatternsWithGo/creational/builder/message/builder"
)

func main() {

	emailBuilder := &builder.EmailBuilder{
		To:       "bb",
		Cc:       "",
		MailText: "test email",
	}

	director := builder.Director{}
	message, _ := director.ConstructMessage(emailBuilder)
	fmt.Printf("message is %v\n", message)

}
