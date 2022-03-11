package ui

import "app/dom"

var App = dom.Component{

	Action:func() {},
	Model: func()string{
		return `<h1>Hola mundo </h1>`
	},
}