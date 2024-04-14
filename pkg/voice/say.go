package voice

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type VoiceChoice string

const (
	Karen  VoiceChoice = "Karen"
	Daniel VoiceChoice = "Daniel"
)

type VoiceConfig struct {
	voice  VoiceChoice
	volume float64
}

type Option func(*VoiceConfig)

func WithVoice(v VoiceChoice) Option {
	return func(vc *VoiceConfig) {
		vc.voice = v
	}
}

// say either uses exec to say the text or fatally logs the error
func Say(text string, opts ...Option) {
	vc := &VoiceConfig{}
	for _, opt := range opts {
		opt(vc)
	}
	if vc.voice == "" {
		vc.voice = Karen
	}
	if vc.volume == 0.0 {
		vc.volume = 1.0
	} else if vc.volume > 1.0 || vc.volume < 0.0 {
		vc.volume = 1.0
		log.Warnf("volume must be between 0.0 and 1.0, defaulting to 1.0")
	}
	fmt.Println(text + "\n\n")

	command := exec.Command("say", "-v", string(vc.voice), fmt.Sprintf("\"[[volm %.2f]] %s\"", vc.volume, text))
	if err := command.Run(); err != nil {
		log.WithError(err).Fatalf("failed to run say command")
	}
}
