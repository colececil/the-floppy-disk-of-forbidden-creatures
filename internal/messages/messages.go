package messages

// MessageKey is used when defining keys for the messages map.
type MessageKey int

const (
	IntroMessage MessageKey = iota
	BeginRitualMessage
	AwaitingAcknowledgementMessage
	SummoningMessage
	EndingMessage
)

// messages contains messages to be displayed to the player.
var messages = map[MessageKey]string{
	IntroMessage: "The corrupted data writhes its way out of the disk, a gateway to a hidden realm. You have entered " +
		"the Floppy Disk of Forbidden Creatures. And you know you have come here for a purpose - to summon a " +
		"creature beyond your comprehension.",
	BeginRitualMessage: "You begin the ritual...",
	SummoningMessage:   "Summoning in progress",
	EndingMessage: "Your summoning complete, you may now return to your own world. But will you regret what you have " +
		"unleashed upon it?",
	AwaitingAcknowledgementMessage: "<Press Enter to continue.>",
}
