package routeur

import (
	controller "HangmanWeb/controller"
	"net/http"
	"os"
)

func InitServe() {
	http.HandleFunc("/init", controller.DisplayInit)
	http.HandleFunc("/choose", controller.DisplayChoose)
	http.HandleFunc("/jeu", controller.DisplayJeu)
	http.HandleFunc("/init/treatment", controller.InitInit)
	http.HandleFunc("/choose/treatment", controller.InitChoose)
	http.HandleFunc("/jeu/treatment", controller.InitJeu)

	http.HandleFunc("/win", controller.DisplayWin)
	http.HandleFunc("/loose", controller.DisplayLoose)
	http.HandleFunc("/restart", controller.Restart)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)
}
