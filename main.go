package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	// l := log.New(log.Writer(), log.Prefix(), log.Flags())
	a, err := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
		AppName: "clock",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer a.Close()
	// Handle signals
	a.HandleSignals()
	// Start
	if err = a.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// New window
	var w, _ = a.NewWindow("res/index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(300),
		Width:  astikit.IntPtr(600),
	})
	// Create windows
	_ = w.Create()

	// Blocking pattern
	a.Wait()
}
