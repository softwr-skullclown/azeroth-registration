package config

type Config struct {
	ListenAddress string `mapstructure:"listen_address"`
	// SiteTitle is the title of the webpage e.g. My Private Server Registration
	SiteTitle string `mapstructure:"site_title"`
	// RealmList is the Realmlist of your server aka the IP or DOMAIN to access your server e.g. logon.myserver.com
	RealmList string `mapstructure:"realm_list"`
	// GameVersion is the version of the game that your server is running. e.g. 3.3.5a
	GameVersion string `mapstructure:"game_version"`
	// AuthDatabase is used for the registration and account management features
	AuthDatabase DatabaseConnection `mapstructure:"auth_database"`
	// Realms is the list of Realms to display associated to the Auth Server above
	Realms []Realm
	// SMTP server is required to send out emails
	SMTP SMTPConfig
	// UseOSFilesystem allows the application to use the filesystem for UI and Email template files (intended for local development speedup)
	UseOSFilesystem bool `mapstructure:"use_os_filesystem"`
}

// DatabaseConnection is a MySQL/Maria DB connection used by the auth and realm character databases
type DatabaseConnection struct {
	Name string
	User string
	Pass string
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
	CharacterDatabase DatabaseConnection `mapstructure:"character_database"`
}

// SMTP SMTPConfig is required for sending out emails (e.g., password recovery)
type SMTPConfig struct {
	Host string
	Port int
	// Auth method to use for smtp server (defaults to 'PLAIN')
	Auth   string
	User   string
	Pass   string
	Secure bool
	// SkipTLSVerify whether or not to skip TLS verification
	SkipTLSVerify bool `mapstructure:"skip_tls_verify"`
	// SendFrom is the email address emails are sent from e.g. a no-reply
	SendFrom string `mapstructure:"send_from"`
}
