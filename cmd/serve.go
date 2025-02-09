package cmd

import (
	"context"

	"github.com/softwr-skullclown/azeroth-registration/internal/db"
	"github.com/softwr-skullclown/azeroth-registration/internal/db/auth"
	"github.com/softwr-skullclown/azeroth-registration/internal/db/realm"
	"github.com/softwr-skullclown/azeroth-registration/internal/email"
	"github.com/softwr-skullclown/azeroth-registration/internal/http"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the example webapp on https://localhost:8080",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func serve() {
	emailService := email.New(c.UseOSFilesystem, email.Config{
		SiteTitle: c.SiteTitle,
		SMTP: email.SMTP{
			Auth:     c.SMTP.Auth,
			Secure:   c.SMTP.Secure,
			Host:     c.SMTP.Host,
			Port:     c.SMTP.Port,
			User:     c.SMTP.User,
			Pass:     c.SMTP.Pass,
			SendFrom: c.SMTP.SendFrom,
		},
	})
	authDBSvc := auth.Service{DB: db.New(&db.Config{
		Host: c.AuthDatabase.Host,
		Port: c.AuthDatabase.Port,
		User: c.AuthDatabase.User,
		Pass: c.AuthDatabase.Pass,
		Name: c.AuthDatabase.Name,
	}).DB}

	realmSvcs := map[int]http.RealmDBService{}
	realmIds := make([]int, 0)

	for _, r := range c.Realms {
		realmIds = append(realmIds, r.Id)
		realmSvcs[r.Id] = &realm.Service{DB: db.New(&db.Config{
			Host: r.CharacterDatabase.Host,
			Port: r.CharacterDatabase.Port,
			User: r.CharacterDatabase.User,
			Pass: r.CharacterDatabase.Pass,
			Name: r.CharacterDatabase.Name,
		}).DB}
	}

	h := http.New(http.Config{
		ListenAddress:         c.ListenAddress,
		RealmIds:              realmIds,
		UseOSFilesystem:       c.UseOSFilesystem,
		AllowMultipleAccounts: c.AllowMultipleAccounts,
	}, &authDBSvc, realmSvcs, emailService)

	err := h.ListenAndServe(context.Background())
	if err != nil {
		panic(err)
	}
}
