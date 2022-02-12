package solution

import "github.com/kyokomi/emoji"

// GetMessage returns an invitation string with emoji
func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}
