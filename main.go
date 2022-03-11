package main

import (
	"app/ui"

	"app/dom"
)

func main() {
 
	dom.New(ui.App , dom.Window{
							Width: 800, 
							Height:300,
							PosX: 100,
							PosY: 100,
							Icon:"./src/logo.png",
							Title: "App",
	})
	dom.OnWait()
}