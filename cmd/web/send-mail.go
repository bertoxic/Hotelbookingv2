package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/bertoxic/bert/models"

)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
		SendMail(msg)
		}
	}()
}

func SendMail(m models.MailData) {

	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.SendTimeout = 10 * time.Second
	server.ConnectTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println("no client kkk",err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		data, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
		if err != nil {
            app.ErrorLog.Println(err)
        }
        mailTemplate := string(data)
        msgToSend :=strings.Replace(mailTemplate,"[%body%]",m.Content,1)
        email.SetBody(mail.TextHTML, msgToSend)
	}

	err = email.Send(client)
	if err != nil {
		log.Println("xzzzzzzzzzzzz",err)
	} else {
		log.Println("email sent")
	}
}
