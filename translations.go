package main

import (
	"embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/lang"
)

//go:embed translations/*.json
var translationsFS embed.FS

func setupTranslations() {
	if err := lang.AddTranslationsFS(translationsFS, "translations"); err != nil {
		fyne.LogError("failed to load translations", err)
	}
}
