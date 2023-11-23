package HangmanWeb

type Choose struct {
	Pseudo         string
	Level          int
	Cpt            int
	Mode           string
	Val            string
	Mot            string
	AEL            []string
	AEW            []string
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
