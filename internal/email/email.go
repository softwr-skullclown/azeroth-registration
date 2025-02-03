package email

import (
	"context"
	"crypto/tls"
	"fmt"
	"html/template"
	"io/fs"

	"github.com/wneessen/go-mail"
)

type SMTP struct {
	Host              string
	Port              int
	Auth              string
	User              string
	Pass              string
	Secure            bool
	SendFrom          string
	SmtpSkipTLSVerify bool
}

type Config struct {
	SiteTitle string
	SiteURL   string
	SMTP      SMTP
}

type Service struct {
	config    Config
	tlsConfig *tls.Config
	authType  mail.SMTPAuthType
	templates fs.FS
}

func New(useOS bool, c Config) *Service {
	s := Service{
		config: c,
	}

	s.templates = EmbedTemplates(useOS)

	s.authType = mail.SMTPAuthType(s.config.SMTP.Auth)
	s.tlsConfig = &tls.Config{
		InsecureSkipVerify: s.config.SMTP.SmtpSkipTLSVerify || !s.config.SMTP.Secure,
		ServerName:         s.config.SMTP.Host,
	}

	return &s
}

// send - utility function to send emails
func (s *Service) send(ctx context.Context, userName string, userEmail string, subject string, templateName string, data interface{}) error {
	var err error
	var c *mail.Client

	fromEmail := s.config.SMTP.SendFrom

	m := mail.NewMsg()
	if err = m.From(fromEmail); err != nil {
		return fmt.Errorf("failed to set From address %s error: %v", fromEmail, err)
	}
	if err = m.To(userEmail); err != nil {
		return fmt.Errorf("failed to set To address %s error: %v", userEmail, err)
	}

	tplName := fmt.Sprintf("%s.html", templateName)
	tpl, err := template.New(tplName).ParseFS(s.templates, tplName)
	if err != nil {
		return fmt.Errorf("error loading welcome email template: %w", err)
	}

	m.Subject(subject)
	err = m.SetBodyHTMLTemplate(tpl, data)
	if err != nil {
		return fmt.Errorf("error setting body template html: %w", err)
	}
	m.SetAddrHeaderIgnoreInvalid(mail.HeaderFrom, fmt.Sprintf("%s <%s>", s.config.SiteTitle, fromEmail))
	m.SetAddrHeaderIgnoreInvalid(mail.HeaderTo, fmt.Sprintf("%s <%s>", userName, userEmail))

	if s.config.SMTP.Secure {
		c, err = mail.NewClient(s.config.SMTP.Host, mail.WithPort(s.config.SMTP.Port), mail.WithSMTPAuth(s.authType),
			mail.WithUsername(s.config.SMTP.User), mail.WithPassword(s.config.SMTP.Pass), mail.WithTLSConfig(s.tlsConfig))
	} else {
		c, err = mail.NewClient(s.config.SMTP.Host, mail.WithPort(s.config.SMTP.Port), mail.WithTLSConfig(s.tlsConfig),
			mail.WithTLSPolicy(mail.TLSOpportunistic))
	}
	if err != nil {
		return fmt.Errorf("failed to create mail client: %v", err)
	}

	if err = c.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send mail: %v", err)
	}

	return err
}
