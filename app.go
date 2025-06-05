package main

import (
	"context"
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime" // Import the runtime package
	"golang.design/x/hotkey"
)

// App speichert den Kontext, damit wir später WindowShow/WindowHide aufrufen können.
type App struct {
	ctx       context.Context
	isVisible bool
}

func NewApp() *App {
	return &App{
		isVisible: false, // Da wir StartHidden=true setzen, ist es initial unsichtbar
	}
}

func (a *App) startup(ctx context.Context) {

	a.ctx = ctx

	go func() {
		hk := hotkey.New([]hotkey.Modifier{hotkey.ModCmd}, hotkey.KeySpace)
		if err := hk.Register(); err != nil {
			log.Fatalf("Hotkey-Registrierung fehlgeschlagen: %v", err)
		}
		log.Println("Hotkey registriert")

		for {
			<-hk.Keydown()
			log.Println("Hotkey gedrückt!")
			if a.isVisible {
				runtime.WindowHide(a.ctx) // Use runtime.WindowHide
				a.isVisible = false
				log.Println("Fenster versteckt")
			} else {
				runtime.WindowShow(a.ctx) // Use runtime.WindowShow
				a.isVisible = true
				log.Println("Fenster angezeigt")
			}
		}
	}()

}
