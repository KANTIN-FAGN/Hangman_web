package HangmanWeb

type Choose struct {
	Pseudo         string
	Level          int
	Cpt            int // chances pour le hangman
	Mode           string
	Val            string //valeur entrée par l'utilisateur
	Mot            string
	AEL            []string
	AEW            []string
	URL            string
	DispMot        []string
	Win            bool
	Image          string
	Erreur         int
	PtsUser        int
	BonneLettre    int
	BonMot         int
	MauvaiseLettre int
	MauvaisMot     int
}
