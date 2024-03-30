package voice

import (
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// say either uses exec to say the text or fatally logs the error
func Say(text string, voiceSetting ...int) {
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
