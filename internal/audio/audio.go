package audio

import (
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/gopxl/beep/v2/wav"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const speakerSampleRate = beep.SampleRate(24000)

var currentlyPlaying = map[SoundEffectFilename]bool{}
var currentlyPlayingMutex sync.RWMutex
var lastPlayedSegment = map[SoundEffectFilename]int{}
var lastPlayedSegmentMutex sync.RWMutex

// init initializes the audio.
func init() {
	err := speaker.Init(speakerSampleRate, speakerSampleRate.N(time.Second/10))
	if err != nil {
		panic(err)
	}
}

// Play plays the given sound effect if it's not already playing. If fileSegmentIndex is not nil, it will be used to
// replace "%d" in the filename, in order to select the correct segment of the sound effect. If allowOverlap is true,
// it will play the sound effect even if another instance of the same sound effect is already playing.
func Play(filename SoundEffectFilename, fileSegmentIndex *int, allowOverlap bool) error {
	originalFilename := filename
	if fileSegmentIndex != nil {
		filename = filename.Segment(*fileSegmentIndex)
	}

	if getCurrentlyPlaying(originalFilename) && !allowOverlap {
		return nil
	}

	setCurrentlyPlaying(originalFilename, true)
	if fileSegmentIndex != nil {
		setLastSegmentPlayed(originalFilename, *fileSegmentIndex)
	}

	pathToExecutable, err := os.Executable()
	if err != nil {
		setCurrentlyPlaying(originalFilename, false)
		return err
	}
	dirOfExecutable := filepath.Dir(pathToExecutable)

	path := filepath.Join(dirOfExecutable, "assets", "audio", string(filename))
	file, err := os.Open(path)
	if err != nil {
		setCurrentlyPlaying(originalFilename, false)
		return err
	}

	streamer, format, err := wav.Decode(file)
	if err != nil {
		setCurrentlyPlaying(originalFilename, false)
		return err
	}

	// Make sure the streamer's sample rate matches the speaker's sample rate.
	resampledStreamer := beep.Resample(4, format.SampleRate, speakerSampleRate, streamer)

	//speaker.Play(resampledStreamer)
	speaker.Play(beep.Seq(resampledStreamer, beep.Callback(func() {
		setCurrentlyPlaying(originalFilename, false)
		_ = streamer.Close()
	})))

	return err
}

// LastSegmentPlayed returns the last segment played for the given sound effect. If the sound effect has not been played
// yet, it returns -1.
func LastSegmentPlayed(filename SoundEffectFilename) int {
	return getLastSegmentPlayed(filename)
}

// ResetLastSegmentPlayed resets the last segment played for the given sound effect.
func ResetLastSegmentPlayed(filename SoundEffectFilename) {
	setLastSegmentPlayed(filename, -1)
}

// setCurrentlyPlaying sets whether the given sound effect is currently playing.
func setCurrentlyPlaying(filename SoundEffectFilename, value bool) {
	currentlyPlayingMutex.Lock()
	defer currentlyPlayingMutex.Unlock()
	currentlyPlaying[filename] = value
}

// getCurrentlyPlaying returns whether the given sound effect is currently playing.
func getCurrentlyPlaying(filename SoundEffectFilename) bool {
	currentlyPlayingMutex.RLock()
	defer currentlyPlayingMutex.RUnlock()
	return currentlyPlaying[filename]
}

// setLastSegmentPlayed sets the last segment played for the given sound effect.
func setLastSegmentPlayed(filename SoundEffectFilename, value int) {
	lastPlayedSegmentMutex.Lock()
	defer lastPlayedSegmentMutex.Unlock()
	lastPlayedSegment[filename] = value
}

// getLastSegmentPlayed returns the last segment played for the given sound effect. If the sound effect has not been
// played yet, it returns -1.
func getLastSegmentPlayed(filename SoundEffectFilename) int {
	lastPlayedSegmentMutex.RLock()
	defer lastPlayedSegmentMutex.RUnlock()
	_, ok := lastPlayedSegment[filename]
	if !ok {
		return -1
	}
	return lastPlayedSegment[filename]
}
