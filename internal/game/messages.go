package game

// messageKey is used when defining keys for the messages map.
type messageKey int

const (
	introKey = messageKey(iota)
	beginRitualKey
)

// messages contains messages to be displayed to the player.
var messages = map[messageKey]string{
	introKey: "The corrupted data writhes its way out of the disk, a gateway to a hidden realm. You have entered the " +
		"Floppy Disk of Forbidden Creatures. And you know you have come here for a purpose - to summon a creature " +
		"beyond your comprehension.",
	beginRitualKey: "You begin the ritual...",
}
