package controller

import (
	h "HangmanWeb/Hangman/jeu"
	initTemplate "HangmanWeb/temp"
	"net/http"
)

func DisplayInit(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "init", nil)
}

func InitInit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/init", http.StatusPermanentRedirect)
	}
	h.Test.Pseudo = r.FormValue("pseudo")
	http.Redirect(w, r, "/choose", http.StatusPermanentRedirect)

}

func DisplayChoose(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "choose", h.Test)
}

func InitChoose(w http.ResponseWriter, r *http.Request) {
	h.Test.Mode = r.FormValue("choix")
	h.Test.Mot = "lol"
	// if h.Test.Mode == "EASY"{
	// 	h.Test.Mot = h.WriteWord(dico_easy)
	// } else if h.Test.Mode == "MEDIUM"{
	// 	h.Test.Mot = h.WriteWord(dico_moyen)
	// } else{
	// 	h.Test.Mot = h.WriteWord(dico_legend)
	// }
	h.Test.InitTableau()
	h.Test.Image = "../static/img/hangman/hangman_base.png"
	http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
}

func DisplayJeu(w http.ResponseWriter, r *http.Request) {
	if h.Test.Mode == "" {
		http.Redirect(w, r, "/choose", http.StatusPermanentRedirect)
	}
	h.Test.Erreur = 0
	initTemplate.Temp.ExecuteTemplate(w, "jeu", h.Test)
}

func InitJeu(w http.ResponseWriter, r *http.Request) {
	h.Test.Val = r.FormValue("lettre") // on récupère la lettre entrée par l'utilisateur
	if len(h.Test.Val) == 1 {
		for _, i := range h.Test.AEL {
			if i == h.Test.Val {
				h.Test.Erreur = 1
				http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
				return
			}
		}
	} else {
		for _, i := range h.Test.AEW {
			if i == h.Test.Val {
				h.Test.Erreur = 2
				http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
				return
			}
		}
	}

	h.Test.CheckVal() // on vérifie si la lettre est dans le mot ou si le mot est bon
	//et on affecte AlreadyEnteredLetter ou AlreadyEnteredWord
	if h.Test.Win {
		http.Redirect(w, r, "/win", http.StatusMovedPermanently)
	} else if h.Test.Cpt >= 10 {
		http.Redirect(w, r, "/loose", http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)

	}
}

func DisplayWin(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "win", h.Test)
}

func DisplayLoose(w http.ResponseWriter, r *http.Request) {
	initTemplate.Temp.ExecuteTemplate(w, "loose", h.Test)
}

func Restart(w http.ResponseWriter, r *http.Request) {
	h.Test.Restart()
	http.Redirect(w, r, "choose", http.StatusMovedPermanently)
}
