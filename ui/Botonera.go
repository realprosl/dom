package ui

import "app/dom"


var Botonera = dom.NewComponent(
		//Action
		func(){

			dom.Bind("handleBtn",func(){
				Message.Set("Bye world")
			})
		},
		//Model
		func()string{
			return `
				<button onclick=handleBtn()>Press</button>
			`
		},
)