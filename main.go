package main

import (
	routeur "HangmanWeb/routeur"
	initTemplate "HangmanWeb/temp"
)

func main() {
	initTemplate.InitTemplate()
	routeur.InitServe()
}
