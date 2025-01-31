package email

type SMTP struct {
	Host     string
	Port     int
	Auth     bool
	User     string
	Pass     string
	Secure   string
	SendFrom string
}

type Config struct {
	SiteTitle string
	SiteURL   string
	SMTP      SMTP
}

type Service struct {
	config Config
}

func New(c Config) *Service {
	s := Service{
		config: c,
	}

	return &s
}
