package common

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"server/base-server/internal/conf"
	"server/common/errors"
)

func SendEmail(adminEmail *conf.AdminEmail, to string, subject string) error {
	e := email.NewEmail()
	e.From = adminEmail.Username
	e.To = []string{to}
	e.Subject = subject
	err := e.Send(fmt.Sprintf("%s:%d", adminEmail.SmtpHost, adminEmail.SmtpPort), smtp.PlainAuth("", adminEmail.Username, adminEmail.Password, adminEmail.SmtpHost))
	if err != nil {
		return errors.Errorf(err, errors.ErrorSendEmailFailed)
	}

	return nil
}
