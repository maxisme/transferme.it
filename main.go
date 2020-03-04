package main

import (
	"github.com/maxisme/appserver"
	"os"
	"github.com/joho/godotenv" 
)

func main() {
	godotenv.Load()
	conf := appserver.ProjectConfig{
		Name: "Transfer Me It",
		Host: "transferme.it",
		DmgPath: "Transfer Me It.dmg",
		KeyWords: "transfer me it, transfermeit, file, send, transfer, me, it, transfer, file transfer, send files, mac, osx, mac to mac",
		Description: "A minimal file transfer app for MacOS.",
		Recaptcha: appserver.Recaptcha{
			Pub: os.Getenv("captchpub"),
			Priv: os.Getenv("captchpriv"),
		},
		Sparkle: appserver.Sparkle{
			Version:"0.03",
			Description:"foo",
		},
		Email: appserver.Email{
			To: "max@max.me.uk",
			Host: os.Getenv("emailhost"),
			Port: 587,
			Username: os.Getenv("emailusername"),
			Password: os.Getenv("emailpassword"),
		},
	}

	if err := appserver.Serve(conf); err != nil {
		panic(err)
	}
}
