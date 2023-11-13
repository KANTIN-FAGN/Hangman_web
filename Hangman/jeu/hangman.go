package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var win = false
var isInWord = false
var start = 1
var end = 8
var c int
var alreadyEntered []string
var mistake = false

func ChooseDico() {
	var choose int
	fmt.Println("Choisissez le thème avec lequel vous allez jouer :\n1. Animaux\n2. Films\n3. Personnages de Films\n4. Voitures\n5. Villes\n6. Dictionnaire officiel du Scrabble 2012")
	fmt.Scan(&choose)
	fmt.Scan()
	switch choose {
	case 1:
		mot := WriteWord("jeu/dico_animaux.txt")
		ModeDeJeu(mot)
	case 2:
		mot := WriteWord("jeu/dico_films.txt")
		ModeDeJeu(mot)
	case 3:
		mot := WriteWord("jeu/dico_pfilm.txt")
		ModeDeJeu(mot)
	case 4:
		mot := WriteWord("jeu/dico_voiture.txt")
		ModeDeJeu(mot)
	case 5:
		mot := WriteWord("jeu/dico-villes.txt")
		ModeDeJeu(mot)
	case 6:
		mot := WriteWord("dictionnaire.txt")
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
		Displaystock(mot, stock)
		Choose(mot, stock)
		if c == 10 {
			fmt.Println("Vous avez perdu !\nLe mot était : ", mot)
			DisplayResult(19, 34)
			mot = WriteWord("dictionnaire.txt")
			Replay(mot)
			return
		}
		if win {
			mot = WriteWord("dictionnaire.txt")
			DisplayResult(1, 17)
			Replay(mot)
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
	Check(mot, stock, letter, alreadyEntered)
	alreadyEntered = append(alreadyEntered, letter)
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
		fmt.Println("\033[H\033[2J", "Liste des lettres que vous avez déjà entrées : ", alreadyEntered)
		Inputletter(mot, stock)
		a = LenMot(stock)
		if a == len(mot) {
			fmt.Println("\nBien joué ! Vous avez trouvé le mot : ", mot)
			win = true
			return
		}

	case 2:
		fmt.Println("\033[H\033[2J")
		Inputword(mot, stock)
		if win {
			return
		}
		Choose(mot, stock)

	default:
		fmt.Println("\033[H\033[2J")
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

func Replay(mot string) {
	var replay int
	fmt.Println("\nCa te dis de rejouer ?\n1. OUI c'est trop bien le pendu \n2. NON il est guez ton jeu...")
	fmt.Scan(&replay)
	fmt.Scan()
	switch replay {
	case 1:
		fmt.Println("Let's go mon gaté !")
		Init()
		ChooseDico()
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
	alreadyEntered = nil
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
