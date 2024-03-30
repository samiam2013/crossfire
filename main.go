package main

import (
	"os"
	"time"

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

	debateTopic := "'vanilla' Javascript is better than React" // simulation theory suggested by copilot

	talkingStick := make(chan struct{})
	go func(ts chan struct{}) {
		voice.Say("Welcome to crossfire, I'm your host Daniel representing Open AI's Chat GPT", voice.WithVoice(voice.Daniel))
		time.Sleep(500 * time.Millisecond)
		voice.Say("And I'm your co-host Karen, representing Anthropic's Claude. The topic of today's debate?: ", voice.WithVoice(voice.Karen))
		time.Sleep(500 * time.Millisecond)
		voice.Say(debateTopic, voice.WithVoice(voice.Daniel))
		voice.Say("I'll be taking the affirmative position", voice.WithVoice(voice.Karen))
		voice.Say("And I'll be taking the opposing position", voice.WithVoice(voice.Daniel))
		ts <- struct{}{}
	}(talkingStick)
	oAPI := openai.NewAPI(oAPIkey)
	anthAPI := anthropic.NewAPI(anthAPIkey)

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

	<-talkingStick
	go func(ts chan struct{}) {
		voice.Say(contentStr, voice.WithVoice(voice.Karen))
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
			voice.Say(oAPIFirstContent, voice.WithVoice(voice.Daniel))
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
			voice.Say(contentStr, voice.WithVoice(voice.Karen))
			ts <- struct{}{}
		}(talkingStick)
	}
}
