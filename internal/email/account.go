package email

// SendWelcome sends the welcome email after account registration
func (s *Service) SendWelcome(email string, username string) error {
	return nil
}

// SendPasswordReset sends the forgot password reset email
func (s *Service) SendPasswordReset(email string, username string, token string) error {
	return nil
}

// SendPasswordUpdated sends the password updated email
func (s *Service) SendPasswordUpdated(email string, username string) error {
	return nil
}
