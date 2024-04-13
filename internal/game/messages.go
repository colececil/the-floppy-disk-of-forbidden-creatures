package game

type messageKey int

const (
	introKey = messageKey(iota)
)

var messages = map[messageKey]string{
	introKey: "The corrupted data writhes its way out of the disk, a gateway to a hidden realm. You have entered the " +
		"Floppy Disk of Forbidden Creatures.",
}
