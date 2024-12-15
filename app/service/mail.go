package service

import (
	"net/smtp"
)

func SendEmail(to string, token string) error {
	from := ""
	password := ""

	smtpHost := "smtp.mail.ru"
	smtpPort := "587"

	message := []byte("Subject: Подтверждение почты" + "\r\n" +
		"\r\n" + "Для того, чтобы подтвердить почту перейдите по следующей ссылке: http://localhost:8000/confirm?token=" + token)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
}
