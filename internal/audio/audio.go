package audio

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type AudioPanel struct {
	sampleRate beep.SampleRate
	streamer   beep.StreamSeeker
	ctrl       *beep.Ctrl
	resampler  *beep.Resampler
	volume     *effects.Volume
	stopCh     chan bool
}

func NewAudioPanel(sampleRate beep.SampleRate, streamer beep.StreamSeeker) *AudioPanel {
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer)}
	resampler := beep.ResampleRatio(4, 1, ctrl)
	volume := &effects.Volume{Streamer: resampler, Base: 2}
	return &AudioPanel{sampleRate, streamer, ctrl, resampler, volume, nil}
}

func (ap *AudioPanel) Play(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}
	ap.stopCh = make(chan bool)
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	for {
		select {
		case <-ap.stopCh:
			// Arrêter la lecture et sortir de la boucle
			ap.ctrl.Paused = true
			ap.ctrl.Streamer = nil
			return
		default:
			// Continuer la lecture
		}
	}

}

func (ap *AudioPanel) Stop() {
	if ap.stopCh != nil {
		ap.stopCh <- true
	}
}
