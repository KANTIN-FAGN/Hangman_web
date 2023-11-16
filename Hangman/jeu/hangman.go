package HangmanWeb

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
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

var win = false
var isInWord = false
var start = 1
var end = 8
var c int
var alreadyEnteredLetter []string
var alreadyEnteredWords []string
var mistake = false

func (d *Display) Nickname() {
	var pseudo string
	fmt.Println("Entrez votre Pseudo en sah")
	fmt.Scan(&pseudo)
	var choix int
	fmt.Println("\033[H\033[2J\nVoulez-vous vous appeler : ", pseudo, "?\n1. Oui\n2. Non")
	fmt.Scan(&choix)
	switch choix {
	case 1:
		d.Pseudo = pseudo
	case 2:
		Disp.Nickname()
	}
}

func (d *Display) ChooseDico() {
	Disp.Nickname()
	var choose int
	fmt.Println("Choisissez le thème avec lequel vous allez jouer :\n1. Easy\n2. Moyen\n3.Legend")
	fmt.Scan(&choose)
	fmt.Scan()
	switch choose {
	case 1:
		mot := WriteWord("Hangman/jeu/dico_easy.txt")
		Disp.Mode = "Easy"
		ModeDeJeu(mot)
	case 2:
		mot := WriteWord("Hangman/jeu/dico_moyen.txt")
		Disp.Mode = "Moyen"
		ModeDeJeu(mot)
	case 3:
		mot := WriteWord("Hangman/jeu/dico_legend.txt")
		Disp.Mode = "Legende"
		ModeDeJeu(mot)
	}
}

func ModeDeJeu(mot string) {
	var choixmode int
	fmt.Println("Choisissez votre mode de jeu : \n\n1. On vous donne la lettre random dans le mot\n2. Ca joue en mode vaillant et tu trouve tout tout seul")
	fmt.Scan(&choixmode)
	fmt.Scan()
	switch choixmode {
	case 1:
		mode := "lettres KDO"
		Jeu(mot, mode)
	case 2:
		mode := "vrai homme"
		fmt.Println("Vrai homme joue sans indice 🔥🔥🔥")
		Jeu(mot, mode)
	}
}

// Fonction centrale qui regroupe toutes les autres

func Jeu(mot string, mode string) {

	stock := make([]string, len(mot))
	var cpt int
	for _, k := range mot {
		if k == 45 {
			stock[cpt] = "-"
		} else if k == 32 {
			stock[cpt] = " "
		} else if k == 44 {
			stock[cpt] = ","
		} else if k == 39 {
			stock[cpt] = "'"
		} else if k == 95 {
			stock[cpt] = " "
		}
		cpt++
	}
	if mode != "vrai homme" {
		var lettreal string
		var compteur int
		random := rand.Intn(len(mot))
		for _, i := range mot {
			if compteur == random {
				stock[compteur] = string(i)
				break
			}
			compteur++
		}
		compteur = 0
		for _, j := range mot {
			if lettreal == string(j) {
				stock[compteur] = lettreal
			}
			compteur++
		}
	}
	for i := 1; c < 10; i++ {
		Input(mot, stock)
		if c == 10 {
			fmt.Println("Vous avez perdu !\nLe mot était : ", mot)
			DisplayResult(19, 34)
			Replay()
			return
		}
		if win {
			DisplayResult(1, 17)
			Replay()
			return
		}
	}
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

func WriteWord(dico string) string {
	f, err := ReadLines(dico)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	motJeu := rand.Intn(len(f))

	return f[motJeu]
}

// Print la chaine avec les _ ou les lettres si trouvées

func Displaystock(mot string, stock []string) {
	for _, i := range stock {
		if i == "" {
			fmt.Print("_ ")
		} else {
			fmt.Print(i, " ")
		}
	}
}

// nombre de caractères présents dans le mot pour comparer dans InputMot

func LenMot(stock []string) int {
	var lenmo int
	for _, i := range stock {
		if i != "" {
			lenmo++
		}
	}
	return lenmo
}

// Fonction pour entrer une lettre

func Inputletter(mot string, stock []string) {
	fmt.Println("\nEntrez une lettre : ")
	var letter string
	fmt.Scan(&letter)
	fmt.Scan()
	Check(mot, stock, letter, alreadyEnteredLetter)
	alreadyEnteredLetter = append(alreadyEnteredLetter, letter)
	GoodLetter(mot, stock, letter)
	if !isInWord && !mistake {
		DisplayPendu(start, end)
		start += 8
		end += 8
		c++
	}
	isInWord = false
	mistake = false
}

func GoodLetter(mot string, stock []string, letter string) {
	letter = string(ToLower(letter))
	for i := 0; i <= len(mot)-1; i++ {
		if letter == string(mot[i]) {
			stock[i] = letter
			isInWord = true
		}
	}
}

// Fonction pour entrer directement le mot

func Inputword(mot string, stock []string) {
	fmt.Println("Entrez le mot que vous pensez bon")
	var guess string
	fmt.Scanln()
	for {
		var inputChar byte
		_, err := fmt.Scanf("%c", &inputChar)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if inputChar == 13 {
			break
		} else {
			guess += string(inputChar)
		}
	}
	if guess == mot {
		fmt.Println("C'est le bon mot !")
		win = true
		Affectstock(mot, stock)

	} else {
		fmt.Println("Ce n'est pas le bon mot...")
		start += 8
		end += 8
		c++
	}
}

func Input(mot string, stock []string) {
	Displaystock(mot, stock)

	fmt.Println("\nListe des lettres que vous avez déjà entrées : ", alreadyEnteredLetter)
	fmt.Println("Liste des mots que vous avez déjà entrées : ", alreadyEnteredWords)
	fmt.Println("\nEntrez une lettre ou un mot")
	var val string
	fmt.Scan()
	fmt.Scan(&val)
	if len(val) == 1 {
		Check(mot, stock, val, alreadyEnteredLetter)
		alreadyEnteredLetter = append(alreadyEnteredLetter, val)
		Disp.AEL = alreadyEnteredLetter
		GoodLetter(mot, stock, val)
		if !isInWord && !mistake {
			DisplayPendu(start, end)
			start += 8
			end += 8
			c++
		}
		isInWord = false
		mistake = false
	} else {
		if val == mot {
			fmt.Println("C'est le bon mot !")
			win = true
			Affectstock(mot, stock)

		} else {
			fmt.Println("Ce n'est pas le bon mot...")
			start += 8
			end += 8
			c++
			alreadyEnteredWords = append(alreadyEnteredWords, val)
			Disp.AEW = alreadyEnteredWords

		}
	}

}

// Fonction pour compléter le stock si le mot est trouvé

func Affectstock(mot string, stock []string) {
	var count int
	for _, i := range mot {
		stock[count] = string(i)
		count++
	}

}

// Fonction pour choisir entre rentrer une lettre ou directement tout le mot

func Choose(mot string, stock []string) {

	var choix int
	var a int
	fmt.Println("\nChoisissez une option : \n1. Emettre une hypothèse sur une lettre présente dans le mot\n2. Entrer directement le mot")
	fmt.Scan(&choix)
	fmt.Scan()
	switch choix {
	case 1:
		//fmt.Println("\033[H\033[2J", "Liste des lettres que vous avez déjà entrées : ", alreadyEnteredLetter)
		Inputletter(mot, stock)
		a = LenMot(stock)
		if a == len(mot) {
			fmt.Println("\nBien joué ! Vous avez trouvé le mot : ", mot)
			win = true
			return
		}

	case 2:
		//fmt.Println("\033[H\033[2J")
		Inputword(mot, stock)
		if win {
			return
		}
		Choose(mot, stock)

	default:
		//fmt.Println("\033[H\033[2J")
		fmt.Println("Choix invalide, Veuillez choisir une option valide")
		Choose(mot, stock)

	}
}

func Check(mot string, stock []string, letter string, alreadyEntered []string) {
	for _, i := range alreadyEntered {
		if i == letter {
			fmt.Println("Vous avez déjà entré cette lettre !")
			mistake = true
			Inputletter(mot, stock)
		}
	}
	for _, j := range letter {
		if j < 64 || j > 91 && j < 97 || j > 122 {
			fmt.Println("T'es serieux à mettre", letter, "? On veut que des lettres nous !")
			mistake = true
			Inputletter(mot, stock)
		}
	}
	if len(letter) > 1 {
		fmt.Println("Pas plus d'une lettre dindon des iles !!!")
		mistake = true
		Inputletter(mot, stock)
	}
}

func Replay() {
	var replay int
	fmt.Println("\nCa te dis de rejouer ?\n1. OUI c'est trop bien le pendu \n2. NON il est guez ton jeu...")
	fmt.Scan(&replay)
	fmt.Scan()
	switch replay {
	case 1:
		fmt.Println("Let's go mon gaté !")
		Init()
		Disp.ChooseDico()
	case 2:
		fmt.Println("T'es trop nul de toute façon...")
		return
	}
}

func Init() {
	win = false
	isInWord = false
	start = 1
	end = 8
	c = 0
	alreadyEnteredLetter = nil
	alreadyEnteredWords = nil
	mistake = false
}

func ToLower(letter string) rune {
	var fin rune
	for _, i := range letter {
		if i > 64 && i < 91 {
			fin = rune(i) + 32
		} else {
			fin = i
		}
	}
	return fin
}
