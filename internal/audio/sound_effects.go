package audio

import (
	"fmt"
	"strings"
)

type SoundEffectFilename string

const (
	ClickSoundEffect                   SoundEffectFilename = "click.wav"
	DialupModemSoundEffect             SoundEffectFilename = "dialup_modem.wav"
	DoubleBeepSoundEffect              SoundEffectFilename = "double_beep.wav"
	DoubleBuzzSoundEffect              SoundEffectFilename = "double_buzz_%d.wav"
	HighPitchedBeepSoundEffect         SoundEffectFilename = "high_pitched_beep.wav"
	LongBuzzSoundEffect                SoundEffectFilename = "long_buzz.wav"
	LongLowPitchedBeepSoundEffect      SoundEffectFilename = "long_low_pitched_beep.wav"
	QuietTapSoundEffect                SoundEffectFilename = "quiet_tap.wav"
	RhythmicClicksAndBuzzesSoundEffect SoundEffectFilename = "rhythmic_clicks_and_buzzes_%d.wav"
	ShortBuzzSoundEffect               SoundEffectFilename = "short_buzz.wav"
	ShortLowPitchedBeepSoundEffect     SoundEffectFilename = "short_low_pitched_beep.wav"
	TapSoundEffect                     SoundEffectFilename = "tap.wav"
)

// Segment returns a new SoundEffectFilename with the given segment index, if the filename contains "%d". Otherwise, it
// returns the original SoundEffectFilename.
func (f SoundEffectFilename) Segment(index int) SoundEffectFilename {
	if !strings.Contains(string(f), "%d") {
		return f
	}
	return SoundEffectFilename(fmt.Sprintf(string(f), index))
}
