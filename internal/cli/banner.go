package cli

import (
	"strings"
	"time"
)

var (
	// asciiBeer is a slice of strings representing the ASCII art of a beer mug.
	asciiBeer = strings.Split(`
  _.._..,_,_
 (          )
  ]~,"-.-~~[
  ])' (;  ([=,
  ]:: '  ::[ ||
  ]): .) :([ ||
  ])' '   ([='
   ~~----~~
`, "\n")
	// messages is a slice of strings representing the startup messages.
	messages = []string{
		"Weaponized irony,\nresponsibly contained.",
		"The decentralized app store for chaotic,\nfun or hyper-specific terminal utilities and vibes.",
		"Because sometimes you just need a tool that does one thing,\nand does it badly.",
		"Where the command line meets absurdity.",
		"For when you want to run a command that makes no sense,\nbut is hilarious.",
		"The app store for the terminal,\nwhere every command is a gamble.",
		"Because the best way to learn is to break things,\nand then fix them.",
		"Where the terminal meets the absurd,\nand the absurd meets the terminal.",
		"If you can dream it,\nwe probably have a vibe for it.\nIf not, submit one!",
		"How to run a command that does nothing,\nbut makes you laugh.",
		"Where the command line meets chaos,\nand chaos meets the command line.",
		"When you are tired of the same old commands,\nand want to try something new.",
		"Just because it doesn't make sense,\ndoesn't mean it isn't useful.",
		"Share your weirdest, most niche command line tools,\nand find others who appreciate them.",
		"Easy to use,\nhard to explain,\nimpossible to forget.",
		"You never know what you'll find in the next pub,\nbut you can be sure it will be interesting.",
		"Keep your beer cold and your terminal hot.",
		"Shipped it during a deploy freeze because the vibes had quorum.",
		"The test failed locally,\nwhich means production deserves a vote.",
		"If the script wanted a sandbox,\nit should have brought snacks.",
		"This worked once in a terminal named definitely-final.",
		"Rollback plans are just fan fiction with better indentation.",
	}
)

// beerArt outputs the HoldMyBeer signature with a pseudo-random message.
func beerArt() string {
	// Select pseudo-random message from the messages slice (seeded with the current time).
	i := time.Now().UnixNano() % int64(len(messages))

	asciiBeerWidth := 0
	for _, line := range asciiBeer {
		if len(line) > asciiBeerWidth {
			asciiBeerWidth = len(line)
		}
	}
	asciiBeerWidth += 2 // Add padding for the message

	// Split the message into lines that fit within 80 chars minus the width of the ASCII art.
	msg := messages[i]
	beerArt := strings.Builder{}
	appendText := func(line, text string, offset int) string {
		spaces := strings.Repeat(" ", offset-len(line))
		return line + spaces + text
	}
	messageLines := []string{
		"",
		"",
		"   Hold My Beer",
		"<-=+*^#@~~@#^*+=->",
		"",
	}
	messageLines = append(messageLines, strings.Split(msg, "\n")...)
	for i, line := range asciiBeer {
		if i < len(messageLines) && messageLines[i] != "" {
			beerArt.WriteString(appendText(line, messageLines[i], asciiBeerWidth))
		} else {
			beerArt.WriteString(line)
		}
		beerArt.WriteString("\n")
	}
	return beerArt.String()
}
