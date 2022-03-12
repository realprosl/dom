package ui

import (
	"app/dom"
)
var Message *dom.State

var App = dom.NewComponent(
		//Action
		func(){

			Message = dom.NewState("message","Hello world")
		},
		//Model
		func()string{

			dom.AddChilds(&Botonera)
			
			return `
				<div class='app'>
					<h1>$message</h1>
					</Botonera>
				</div>
			`
		},
)