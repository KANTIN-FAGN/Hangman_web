package controller

import (
	h "HangmanWeb/Hangman/jeu"
	initTemplate "HangmanWeb/temp"
	"net/http"
)

func DisplayInit(w http.ResponseWriter, r *http.Request) {
	if h.Test.URL != "" {
		http.Redirect(w, r, h.Test.URL, http.StatusMovedPermanently)
	}
	h.Test.URL = "/init"
	initTemplate.Temp.ExecuteTemplate(w, "init", nil)
}

func InitInit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/init", http.StatusPermanentRedirect)
	}
	h.Test.Pseudo = r.FormValue("pseudo")
	h.Test.URL = "/choose"
	http.Redirect(w, r, "/choose", http.StatusPermanentRedirect)

}

func DisplayChoose(w http.ResponseWriter, r *http.Request) {

	if h.Test.URL != "/choose" {
		http.Redirect(w, r, h.Test.URL, http.StatusMovedPermanently)
	}
	h.Test.URL = "/jeu"
	initTemplate.Temp.ExecuteTemplate(w, "choose", h.Test)
}

func InitChoose(w http.ResponseWriter, r *http.Request) {
	h.Test.Mode = r.FormValue("choix")
	if h.Test.Mode == "EASY" { // le barème des points gagnés ou perdu en fct du mode
		h.Test.BonneLettre = 1
		h.Test.BonMot = 3
		h.Test.MauvaiseLettre = 0
		h.Test.MauvaisMot = -2
		h.Test.Mot = h.WriteWord("Hangman/dico/dico_easy.txt") // choix du dico en fct du mode
	} else if h.Test.Mode == "MEDIUM" {
		h.Test.BonneLettre = 3
		h.Test.BonMot = 5
		h.Test.MauvaiseLettre = -1
		h.Test.MauvaisMot = -3
		h.Test.Mot = h.WriteWord("Hangman/dico/dico_moyen.txt")
	} else {
		h.Test.BonneLettre = 5
		h.Test.BonMot = 9
		h.Test.MauvaiseLettre = -3
		h.Test.MauvaisMot = -6
		h.Test.Mot = h.WriteWord("Hangman/dico/dico_legend.txt")
	}

	h.Test.InitTableau()
	h.Test.Image = "../static/img/hangman/hangman_base.png"
	h.Test.URL = "/jeu"
	http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
}

func DisplayJeu(w http.ResponseWriter, r *http.Request) {
	if h.Test.URL != "/jeu" {
		http.Redirect(w, r, h.Test.URL, http.StatusMovedPermanently)
	}
	h.Test.URL = "/jeu"
	if h.Test.Mode == "" {
		http.Redirect(w, r, "/choose", http.StatusPermanentRedirect)
	}
	initTemplate.Temp.ExecuteTemplate(w, "jeu", h.Test)
	h.Test.Erreur = 0
}

func InitJeu(w http.ResponseWriter, r *http.Request) {

	h.Test.Val = r.FormValue("lettre") // on récupère la lettre entrée par l'utilisateur
	if len(h.Test.Val) == 1 {          //si c'est que une lettre :
		for _, i := range h.Test.AEL {
			if i == h.Test.Val {
				h.Test.Erreur = 1
				http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
				return
			}
		}
	} else { //si c'est un mot (plus de 2 lettres) :
		for _, i := range h.Test.AEW {
			if i == h.Test.Val {
				h.Test.Erreur = 2
				http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
				return
			}
		}
	}
	for _, i := range h.Test.Val {
		if i < 97 || i > 122 {
			h.Test.Erreur = 3
			http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
			return
		}
	}
	h.Test.CheckVal() // on vérifie si la lettre est dans le mot ou si le mot est bon
	//et on affecte AlreadyEnteredLetter ou AlreadyEnteredWord
	if h.Test.Win {
		h.Test.URL = "/win"
		http.Redirect(w, r, "/win", http.StatusMovedPermanently)
	} else if h.Test.Cpt >= 10 {
		h.Test.URL = "/loose"
		http.Redirect(w, r, "/loose", http.StatusMovedPermanently)
	} else {
		h.Test.URL = "/jeu"
		http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)

	}
}

func DisplayWin(w http.ResponseWriter, r *http.Request) {
	if h.Test.URL != "/win" {
		http.Redirect(w, r, h.Test.URL, http.StatusMovedPermanently)
	}
	initTemplate.Temp.ExecuteTemplate(w, "win", h.Test)
}

func DisplayLoose(w http.ResponseWriter, r *http.Request) {
	if h.Test.URL != "/loose" {
		http.Redirect(w, r, h.Test.URL, http.StatusMovedPermanently)
	}
	if h.Test.PtsUser < 0 {
		h.Test.PtsUser = 0
	}
	initTemplate.Temp.ExecuteTemplate(w, "loose", h.Test)

}

func Restart(w http.ResponseWriter, r *http.Request) {
	h.Test.Restart()
	h.Test.URL = "/choose"
	http.Redirect(w, r, "choose", http.StatusMovedPermanently)
}
