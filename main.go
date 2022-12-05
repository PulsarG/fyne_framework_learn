package main

import (
	"fyne.io/fyne/v2/app"

	"test/crypt"
)

func main() {
	app := app.New()

	Cryptor := crypt.NewCryptor(app)
	/* Cryptor.SetWidgetsInCryptor() */

	windowMain := Cryptor.CreateMainWindow()

	windowMain.Show()
	windowMain.SetMaster()
	app.Run()
}

/*
__      _________      __
\ \    / /_   _\ \    / /\
 \ \  / /  | |  \ \  / /  \
  \ \/ /   | |   \ \/ / /\ \
   \  /   _| |_   \  / ____ \
 _  \/   |_____|  _\/_/    \_\     _   _
| | | |          |  _ \           | | | |
| |_| |__   ___  | |_) | ___  __ _| |_| | ___  ___
| __| '_ \ / _ \ |  _ < / _ \/ _` | __| |/ _ \/ __|
| |_| | | |  __/ | |_) |  __/ (_| | |_| |  __/\__ \
 \__|_| |_|\___| |____/ \___|\__,_|\__|_|\___||___/


*/
