package HangmanWeb

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

var Test = Choose{}

func (t *Choose) CheckVal() {
	var isInWord bool
	if len(t.Val) == 1 {
		t.AEL = append(Test.AEL, t.Val) // On affecte le tableau des lettres qu'on à déjà entrée
		for c, i := range t.Mot {
			if string(i) == t.Val {
				t.DispMot[c] = t.Val
				isInWord = true //si la lettre est dans le mot on ajoute les points en fct du mode
				if t.Mode == "EASY" {
					t.PtsUser += 1
				} else if t.Mode == "MEDIUM" {
					t.PtsUser += 3
				} else {
					t.PtsUser += 5
				}
			}
		}
		if !isInWord {
			t.Cpt++
			if t.Mode == "MEDIUM" {
				t.PtsUser -= 1
			} else if t.Mode == "HARD" {
				t.PtsUser -= 3
			}
			t.Image = "../static/img/hangman/hangman" + string(t.Cpt+47) + ".png" // affichage de l'image du hangman
		}
		isInWord = false
		t.DisplayStock() // on affiche le mot avec les underscores
		t.WordComplete() // si toutes les lettres sont dans le mot on a gagné
	} else {
		t.AEW = append(Test.AEW, t.Val) // on affecte la liste des mots deja entrée
		if t.Val == t.Mot {
			if t.Mode == "EASY" {
				t.PtsUser += 3
			} else if t.Mode == "MEDIUM" {
				t.PtsUser += 5
			} else {
				t.PtsUser += 9
			}
			t.Win = true
			return
		} else {
			t.Cpt++
			if t.Mode == "EASY" {
				t.PtsUser -= 2
			} else if t.Mode == "MEDIUM" {
				t.PtsUser -= 3
			} else {
				t.PtsUser -= 6
			}
			t.Image = "../static/img/hangman/hangman" + string(t.Cpt+47) + ".png"
		}
	}
}

func (t *Choose) WordComplete() {
	var c int
	for _, i := range t.DispMot {
		if i != "_" {
			c++
		}
	}
	if c == len(t.Mot) {
		Test.Win = true
	}
}

func (t *Choose) DisplayStock() {
	for c, i := range t.DispMot {

		if i == "" {
			Test.DispMot[c] = "_"
		} else if i == "_" {
			Test.DispMot[c] = "_"
		} else {
			Test.DispMot[c] = i + ""
		}
	}
}

func (t *Choose) InitTableau() {
	t.DispMot = make([]string, len(t.Mot))
	t.DisplayStock()
}

// lis le fichier texte dans lequel se trouve le dico

func ReadLines(dico string) ([]string, error) {
	file, err := os.Open(dico)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// cherche un mot random dans le dico

func WriteWord(dico string) string { // choix du mot aléatoire dans un dico
	f, err := ReadLines(dico)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	motJeu := rand.Intn(len(f))

	return f[motJeu]
}

func (t *Choose) Restart() {
	t.AEL = nil
	t.AEW = nil
	t.Win = false
	t.Mode = ""
	t.DispMot = nil
	t.Cpt = 0
}
