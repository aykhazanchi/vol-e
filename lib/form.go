package lib

import (
	"github.com/charmbracelet/huh"
)

func BuildForm() (answer string) {

	huh.NewSelect[string]().
		Title("Hi, I am Vol-E, your friendly Swedish study buddy. What would you like to do today?").
		Options(
			huh.NewOption("Learn Svenska", "learn"),
			huh.NewOption("Just practice", "practice"),
			huh.NewOption("Have some fun", "fun"),
			huh.NewOption("Quit", "quit"),
		).
		Value(&answer).
		Run()

	return answer
}

func Learn() (answer string) {

	huh.NewSelect[string]().
		Title("What would you like to learn today?").
		Options(
			huh.NewOption("Nouns", "nouns"),
			huh.NewOption("Verbs", "verbs"),
			huh.NewOption("Tense", "tense"),
			huh.NewOption("Article", "article"),
			huh.NewOption("Common Phrases", "phrases"),
			huh.NewOption("Singular and Plural", "singularplural"),
		).
		Value(&answer).
		Run()

	return answer

}

func Practice() (answer string) {

	huh.NewSelect[string]().
		Title("Ok, let's start an easy chat and see how it goes along the way. Ready?").
		Options(
			huh.NewOption("Ja, vi kör!!", "beginner"),
			huh.NewOption("Nej, ge mig något svåra!!", "advanced"),
		).
		Value(&answer).
		Run()

	return answer
}

func Fun() (answer string) {

	huh.NewSelect[string]().
		Title("What would you like to do for fun?").
		Options(
			huh.NewOption("Play a game", "game"),
			huh.NewOption("Learn some fun facts", "facts"),
		).
		Value(&answer).
		Run()

	return answer
}
