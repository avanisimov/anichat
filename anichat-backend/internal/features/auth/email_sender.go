package auth

import (
	"log"

	"github.com/resend/resend-go/v3"
)

type EmailSender struct {
	APIKey string
}

func NewEmailSender(apiKey string) *EmailSender {
	return &EmailSender{APIKey: apiKey}
}

func (c *EmailSender) SendOtp(to, code string) error {
	client := resend.NewClient(c.APIKey)

    params := &resend.SendEmailRequest{
        From:    "AniChat <onboarding@resend.dev>",
        To:      []string{to},
        Subject: "Your AniChat login code",
        Html:    "<h1>Your code: " + code + "</h1><p>Expires in 10 minutes</p>",
    }

    sent, err := client.Emails.Send(params)
    if err != nil {
        log.Fatalf("failed to send email: %v", err)
    }

    log.Printf("email sent: %s", sent.Id)

	return err
}