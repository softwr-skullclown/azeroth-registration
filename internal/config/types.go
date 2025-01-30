package config

type Config struct {
	ListenAddress string `mapstructure:"LISTEN_ADDRESS"`
	// SiteTitle is the title of the webpage e.g. My Private Server Registration
	SiteTitle string
	// RealmList is the Realmlist of your server aka the IP or DOMAIN to access your server e.g. logon.myserver.com
	RealmList string
	// GameVersion is the version of the game that your server is running. e.g. 3.3.5a
	GameVersion string
	// AuthDatabase is used for the registration and account management features
	AuthDatabase DatabaseConnection
	// Realms is the list of Realms to display associated to the Auth Server above
	Realms []Realm
	// SMTP server is required to send out emails
	SMTP SMTPConfig
}

// DatabaseConnection is a MySQL/Maria DB connection used by the auth and realm character databases
type DatabaseConnection struct {
	Name string
	User string
	Host string
	Port int
}

// Realm is used for Online stats like Number of Online players
type Realm struct {
	// Id corresponds to the ID in the auth database for realmlist
	Id int
	// Name is the visible display name of the Realm in the registration UI
	Name string
	// CharacterDatabase is used to display number of online players
	CharacterDatabase DatabaseConnection
}

// SMTP SMTPConfig is required for sending out emails (e.g., password recovery)
type SMTPConfig struct {
	Host string
	Port int
	Auth bool
	User string
	Pass string
	// Secure is the encryption method: 'tls' or 'ssl'
	Secure string
	// SendFrom is the email address emails are sent from e.g. a no-reply
	SendFrom string
}
