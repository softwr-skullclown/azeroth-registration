package email

import (
	"context"
	"fmt"
)

// SendWelcome sends the welcome email after account registration
func (s *Service) SendWelcome(ctx context.Context, email string, username string) error {
	type Data struct {
		Email     string
		Username  string
		SiteTitle string
		SiteURL   string
	}

	sub := fmt.Sprintf("%s welcome to %s", username, s.config.SiteTitle)
	data := Data{
		Email:     email,
		Username:  username,
		SiteTitle: s.config.SiteTitle,
		SiteURL:   s.config.SiteURL,
	}

	err := s.send(ctx, username, email, sub, "welcome", data)
	return err
}

// SendPasswordReset sends the forgot password reset email
func (s *Service) SendPasswordReset(ctx context.Context, email string, username string, token string) error {
	return nil
}

// SendPasswordUpdated sends the password updated email
func (s *Service) SendPasswordUpdated(ctx context.Context, email string, username string) error {
	return nil
}
