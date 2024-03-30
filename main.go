package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	"github.com/samiam2013/crossfire/anthropic"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.WithError(err).Fatalf("failed to load .env file")
	}
	oAPIkey := os.Getenv("OPENAI_API_KEY")
	if oAPIkey == "" {
		log.Fatalf("OPENAI_API_KEY env var is not set")
	}
	anthAPIkey := os.Getenv("ANTHROPIC_API_KEY")
	if anthAPIkey == "" {
		log.Fatalf("ANTHROPIC_API_KEY env var is not set")
	}
	// run the 'say' command with the text "hello world"
	// say("welcome to crossfire")
	// oAPI := openai.NewAPI(oAPIkey)
	anthAPI := anthropic.NewAPI(anthAPIkey)

	debateTopic := "Why intent matters"
	// say("Let's debate the topic: " + debateTopic)

	// compResp, err := oAPI.GetCompletion("Make your best case for why '" + debateTopic + "' is true")
	// if err != nil {
	// 	log.WithError(err).Fatalf("failed to get completion")
	// }
	// if len(compResp.Choices) == 0 {
	// 	log.Fatalf("no choices in completion response")
	// }
	// fmt.Println(compResp.Choices[0].Message.Content)
	// say(compResp.Choices[0].Message.Content)

	anthResp, err := anthAPI.GetMessageResponse("Make your best case for why '" + debateTopic + "' is true")
	if err != nil {
		log.WithError(err).Fatalf("failed to get message response")
	}

	fmt.Printf("%+v\n", anthResp)
}

// say either uses exec to say the text or fatally logs the error
func say(text string) {
	command := exec.Command("say", fmt.Sprintf("\"[[volm 0.1]] %s\"", text))
	if err := command.Run(); err != nil {
		log.WithError(err).Fatalf("failed to run say command")
	}
}
