package main

import (
	"context"
	"os"
	"time"

	"hospital-app/config"
	"hospital-app/internal"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
	os.Mkdir(config.AppDataFolder, 0755)
	os.Mkdir(config.FilesFolder, 0755)
	go internal.WebServer()
}

// domReady is called after the front-end dom has been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
	runtime.WindowMaximise(ctx)
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	config.SERVER_HANDLE.Shutdown(nil)
	time.Sleep(2 * time.Second)
}
