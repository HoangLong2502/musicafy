package mail

import (
	"github.com/rs/zerolog/log"
	"github.com/wneessen/go-mail"
)

type Mailer struct {
	sender   string
	password string
	address  string
	m        *mail.Msg
}

func NewMailer(sender string, password string, address string) Mailer {
	m := mail.NewMsg()
	return Mailer{sender: sender, password: password, address: address, m: m}
}

func (ml *Mailer) SendMail(to string, subject string, body string) error {
	m := ml.m
	if err := m.From(ml.address); err != nil {
		return err
	}

	if err := m.To(to); err != nil {
		return err
	}

	m.Subject(subject)
	m.SetBodyString(mail.TypeTextHTML, body)

	// Secondly the mail client
	c, err := mail.NewClient(ml.sender,
		mail.WithPort(25), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(ml.address), mail.WithPassword(ml.password),
	)
	if err != nil {
		log.Fatal().Msgf("failed to create mail client: %s", err)
	}

	// Finally let's send out the mail
	if err := c.DialAndSend(m); err != nil {
		log.Fatal().Msgf("failed to send mail: %s", err)
	}
	return nil
}
