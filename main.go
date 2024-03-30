package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	"github.com/samiam2013/crossfire/anthropic"
	"github.com/samiam2013/crossfire/history"
	"github.com/samiam2013/crossfire/openai"
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

	say("welcome to crossfire")
	oAPI := openai.NewAPI(oAPIkey)
	anthAPI := anthropic.NewAPI(anthAPIkey)

	debateTopic := "determinism is true rather than free will (whether or not it subjectively matters)" // simulation theory suggested by copilot
	say("The topic of today's debate is: " + debateTopic)

	hist := history.NewMessageHistory()
	anthResp, err := anthAPI.GetMessageResponse("Make your best case for '"+debateTopic+"' is true", hist)
	if err != nil {
		log.WithError(err).Fatalf("failed to get message response")
	}
	contentStr, err := anthResp.FirstContent()
	if err != nil {
		log.WithError(err).Fatalf("failed to get first content from first message")
	}
	say("The first argument by claude for '" + debateTopic + "' being true: " + contentStr)

	hist.Add(history.AuthorClaude, contentStr)
	for {
		oAPIPrompt := "Make your best case for why the following argument for '" +
			debateTopic + "' being false in response to this argument: " + contentStr
		oAPIResp, err := oAPI.GetCompletion(oAPIPrompt, hist)
		if err != nil {
			log.WithError(err).Fatalf("failed to get completion response from openai")
		}
		oAPIFirstContent, err := oAPIResp.FirstContent()
		if err != nil {
			log.WithError(err).Fatalf("failed to get first content from openai response")
		}
		say("The response argument by openai: "+oAPIFirstContent, 1)
		hist.Add(history.AuthorOpenAI, oAPIFirstContent)

		// now get the response from anthropic and put in contentStr
		anthResp, err = anthAPI.GetMessageResponse("Make your best case for why the following argument for '"+
			debateTopic+"' is true in response to this argument: "+oAPIFirstContent, hist)
		if err != nil {
			log.WithError(err).Fatalf("failed to get message response")
		}
		contentStr, err = anthResp.FirstContent()
		if err != nil {
			log.WithError(err).Fatalf("failed to get first content from first message")
		}
		say("The response argument by claude: " + contentStr)
		hist.Add(history.AuthorClaude, contentStr)
	}
}

// say either uses exec to say the text or fatally logs the error
func say(text string, voiceSetting ...int) {
	fmt.Println(text + "\n\n")
	var voice string
	if len(voiceSetting) == 0 {
		voiceSetting = []int{0}
	}
	if len(voiceSetting) != 1 {
		panic("invalid number of voice settings (should be one or none)")
	}
	switch voiceSetting[0] {
	case 0:
		voice = "Karen"
	case 1:
		voice = "Daniel"
	default:
		panic("invalid voice")
	}

	command := exec.Command("say", "-v", voice, fmt.Sprintf("\"[[volm 0.1]] %s\"", text))
	if err := command.Run(); err != nil {
		log.WithError(err).Fatalf("failed to run say command")
	}
}
