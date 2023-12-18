package main

import (
	"bufio"
	"context"
	"os"
	"strings"

	"github.com/aykhazanchi/vol-e/lib"
	"github.com/sashabaranov/go-openai"
)

func main() {

	ctx := context.Background()
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)

	lib.PrintAsciiArt()

	req := lib.InitGpt(&ctx)

	var setup, answer string
	for {
		setup = lib.BuildForm()

		if setup == "quit" {
			return
		}
		if setup == "learn" {
			answer = lib.Learn()

			var prompt string

			learnPrompts := map[string]string{
				"nouns":          "Teach: Teach me any 3 new nouns. Don't use proper nouns.",
				"verbs":          "Teach: Teach me any 3 new verbs.",
				"tense":          "Teach: Teach me the tense forms of any 3 new verbs.",
				"phrases":        "Teach: Teach me any 3 new most common phrases.",
				"article":        "Teach: Teach me 3 new words with their 'en/ett' form.",
				"singularplural": "Teach: Teach me singular and plurals of any 3 words.",
			}

			prompt, exists := learnPrompts[answer]

			if exists {
				lib.MakeChatCompletionRequest(prompt, &req, *client, &ctx)
			}

		} else if setup == "practice" {
			prompt := lib.Practice()
			if prompt == "advanced" {
				prompt = "Converse: Hi, I'm an advanced Swedish speaker. Converse normally in Swedish and keep the conversation flowing."
			} else {
				prompt = "Converse: Hi, I'm a really beginner level language learner. Use small sentences and simple words and begin a conversation in Swedish."
			}
			lib.MakeChatCompletionRequest(prompt, &req, *client, &ctx)
			// Continue the conversation
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				if scanner.Text() == "quit" || scanner.Text() == "exit" || scanner.Text() == "stop" {
					return
				}
				lib.MakeChatCompletionRequest(scanner.Text(), &req, *client, &ctx)
			}
		} else if setup == "fun" {
			answer = lib.Fun()
			var prompt string

			if answer == "game" {
				scanner := bufio.NewScanner(os.Stdin)
				for {
					prompt = "Game: play a multiple-choice trivia game related to Scandinavia. Do not repeat information previously shared. End the game if an answer is wrong and output the final score."
					resp, err := lib.MakeChatCompletionRequest(prompt, &req, *client, &ctx)
					if err != nil || strings.Contains(resp.Choices[0].Message.Content, "Final score") {
						break
					}
					scanner.Scan()
				}
			} else {
				prompt = "Facts: Give me 3 random fun facts about Scandinavia. Do not repeat information previously shared. The output should be nicely formatted as a numbered list."
				lib.MakeChatCompletionRequest(prompt, &req, *client, &ctx)
			}
		}
	}
}
