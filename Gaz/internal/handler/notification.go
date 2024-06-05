package handler

import (
	"fmt"
	"log/slog"
	"net/smtp"
	"os"
)

const (
	mailHost = "smtp.gmail.com"
	mailAddr = "smtp.gmail.com:587"
)

func (h *Handler) SendMessage() {
	births, err := h.service.SendMessage()
	if err != nil {
		slog.Error("smth wrong", "error", err)
	}
	slog.Info("send message", "mails", births)
	for _, birth := range births {
		err := sendMessage(birth.UserWithBirth, birth.Email)
		if err != nil {
			slog.Error("smth wrong", "error", err)
		} else {
			slog.Info("send message", "mail", birth.Email)
		}

	}
}

func sendMessage(name string, mail string) error {
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_SENDER_ADDRESS"), os.Getenv("EMAIL_SENDER_PASSWORD"), mailHost)
	msg := fmt.Sprintf("This message was send for your, because you subscribe to %s, today is his birthday, so please, don't forget it", name)
	err := smtp.SendMail(mailAddr, auth, os.Getenv("EMAIL_SENDER_ADDRESS"), []string{mail}, []byte(msg))
	return err
}
