package ui

import (
	"app/dom"
)

var Message *dom.State 
var Otro *dom.State
var App = dom.NewComponent(
		//Action
		func(){

			Message = dom.NewState("Hello world")
			Otro = dom.NewState("hola state")
		},
		//Model
		func()string{

			dom.AddChilds(&Botonera)
			
			return `
				<div class='app'>
					<h1>$Message</h1>
					<h2>$Otro</h2>
					</Botonera>
				</div>
			`
		},
)