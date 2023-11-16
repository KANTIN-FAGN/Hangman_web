package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Display struct {
	Pseudo  string
	Level   int
	Cpt     int      // compteur pour le display
	Mot     string   // mot à trouver
	Lettre  string   // lettre entrée par l'utilisateur
	DispMot []string // les underscores
	Mode    string   //Pour le display dans le titre
	AEL     []string //tableau des lettres
	AEW     []string //tableau des mots
}

var Disp = Display{}

func main() {
	temp, err := template.ParseGlob("./temp/*.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/choose", func(w http.ResponseWriter, r *http.Request) {

		Disp.Pseudo = "wokay"
		temp.ExecuteTemplate(w, "choose", Disp)
	})
	http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) {

		Disp.Lettre = "l"
		temp.ExecuteTemplate(w, "jeu", Disp)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)

}
