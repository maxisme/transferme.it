package main

import (
	"github.com/joho/godotenv"
	"github.com/maxisme/appserver"
	"os"
)

func main() {
	// load .env if there
	_ = godotenv.Load()

	conf := appserver.ProjectConfig{
		Name:        "Transfer Me It",
		Host:        "transferme.it",
		DmgPath:     "Transfer Me It.dmg",
		KeyWords:    "transfer me it, transfermeit, transfer me, video, png, jpg, mp4, file, send, transfer, me, it, transfer, file transfer, send files, mac, osx, mac to mac",
		Description: "A minimal file transfer app for MacOS.",
		Recaptcha: appserver.Recaptcha{
			Pub:  os.Getenv("captchpub"),
			Priv: os.Getenv("captchpriv"),
		},
		Sparkle: appserver.Sparkle{
			Version:     "0.03",
			Description: "foo",
		},
		Email: appserver.Email{
			To:       "max@max.me.uk",
			Host:     os.Getenv("emailhost"),
			Port:     587,
			Username: os.Getenv("emailusername"),
			Password: os.Getenv("emailpassword"),
		},
		WebPort: 8080,
	}

	if err := appserver.Serve(conf); err != nil {
		panic(err)
	}
}
