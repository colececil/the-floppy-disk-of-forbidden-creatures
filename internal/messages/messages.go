package messages

// MessageKey is used when defining keys for the messages map.
type MessageKey int

const (
	IntroMessage MessageKey = iota
	BeginRitualMessage
	AwaitingAcknowledgementMessage
	SummoningMessage
	SummoningErrorMessage
	CreatureDescriptionPrompt
	EndingMessage
)

// messages contains messages to be displayed to the player.
var messages = map[MessageKey]string{
	IntroMessage: "The corrupted data writhes its way out of the disk, a gateway to a hidden realm. You have entered " +
		"the Floppy Disk of Forbidden Creatures. And you know you have come here for a purpose - to summon a " +
		"creature beyond your comprehension.",
	BeginRitualMessage: "You begin the ritual...",
	SummoningMessage:   "Summoning in progress",
	SummoningErrorMessage: "You expect to see a monstrous creature appear from the summoning circle, but you only " +
		"see a small poof of smoke. Something has clearly gone wrong, but what? Cursing to yourself, you decide to " +
		"cast the blame on technology.",
	CreatureDescriptionPrompt: "You are the narrator for a game about summoning monsters. Your task is to " +
		"generate a description of the monster being summoned, based on several responses given by the player. The " +
		"description should be a single paragraph, which both narrates the appearance of the monster from the " +
		"summoning circle, and describes what the monster is like. It should also end with a narration explaining " +
		"what becomes of the player (who should be addressed as \"you\") once the monster they summoned has appeared." +
		"\n\n" +
		"The responses given by the player may be things that can directly apply to the monster's appearance, or " +
		"they indirectly provide an attribute of the monster. Please be creative and unpredictable in how the " +
		"player's responses influence what the monster is like. Also, it's better if the description brings up the " +
		"things influenced by the player responses in a different order than they are provided to you. It's also " +
		"better if the description doesn't include the exact wording of the player responses, but applies them in a " +
		"more subtle manner." +
		"\n\n" +
		"Please use descriptive language that paints a mental picture, and keep in mind that the game has a " +
		"foreboding and Lovecraftian tone. Your response should be a single paragraph no longer than 8 sentences. Do " +
		"not include anything other than the description in your response. The player responses are provided below, " +
		"separated by commas:" +
		"\n\n",
	EndingMessage: "Your summoning complete, you may now return to your own world. But will you regret what you have " +
		"unleashed upon it?",
	AwaitingAcknowledgementMessage: "<Press Enter to continue.>",
}
