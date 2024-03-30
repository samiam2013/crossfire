package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/samiam2013/crossfire/anthropic"
	"github.com/samiam2013/crossfire/openai"
	"github.com/samiam2013/crossfire/pkg/history"
	"github.com/samiam2013/crossfire/pkg/voice"
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

	voice.Say("welcome to crossfire", voice.WithVoice(voice.Daniel))
	oAPI := openai.NewAPI(oAPIkey)
	anthAPI := anthropic.NewAPI(anthAPIkey)

	debateTopic := "multi-level polymorphic inheritance is bad in software code" // simulation theory suggested by copilot
	voice.Say("The topic of today's debate is: "+debateTopic, voice.WithVoice(voice.Karen))

	hist := history.NewMessageHistory()
	anthResp, err := anthAPI.GetMessageResponse("Make your best case for '"+debateTopic+"' is true", hist)
	if err != nil {
		log.WithError(err).Fatalf("failed to get message response")
	}
	contentStr, err := anthResp.FirstContent()
	if err != nil {
		log.WithError(err).Fatalf("failed to get first content from first message")
	}
	hist.Add(history.AuthorClaude, contentStr)

	talkingStick := make(chan struct{})
	go func(ts chan struct{}) {
		voice.Say("The first argument by claude for '"+debateTopic+"' being true: "+contentStr, voice.WithVoice(voice.Karen))
		ts <- struct{}{}
	}(talkingStick)

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
		hist.Add(history.AuthorOpenAI, oAPIFirstContent)
		<-talkingStick
		go func(ts chan struct{}) {
			voice.Say("The response argument by open ai: "+oAPIFirstContent, voice.WithVoice(voice.Daniel))
			ts <- struct{}{}
		}(talkingStick)

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
		hist.Add(history.AuthorClaude, contentStr)

		<-talkingStick
		go func(ts chan struct{}) {
			voice.Say("The response argument by claude: "+contentStr, voice.WithVoice(voice.Karen))
			ts <- struct{}{}
		}(talkingStick)

	}
}
