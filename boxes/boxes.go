package boxes

import (
	"strings"
)

const (
	horizontal        = "─"
	vertical          = "│"
	cornerTopLeft     = "┌"
	cornerTopRight    = "┐"
	cornerBottomLeft  = "└"
	cornerBottomRight = "┘"
)

type Box struct {
	Width  int
	Height int
}

type BoxWithContent struct {
	Box     Box
	Content string
}

func (b Box) String() string {
	boxString := ""
	for i := 0; i < b.Height; i++ {
		// the leftmost characters in a row
		switch i {
		case 0:
			boxString += cornerTopLeft
		case b.Height - 1:
			boxString += cornerBottomLeft
		default:
			boxString += vertical
		}

		// the in-between characters in a row (except first and last)
		switch i {
		case 0, b.Height - 1:
			boxString += strings.Repeat(horizontal, b.Width-2)
		default:
			boxString += strings.Repeat(" ", b.Width-2)
		}

		// the rightmost characters in a row
		switch i {
		case 0:
			boxString += cornerTopRight
		case b.Height - 1:
			boxString += cornerBottomRight
		default:
			boxString += vertical
		}

		boxString += "\n"
	}

	return boxString
}

func (b BoxWithContent) String() string {
	boxString := b.Box.String()

	boxLines := strings.Split(boxString, "\n")
	contentLines := strings.Split(b.Content, "\n")

	firstContentRow := (len(boxLines) - len(contentLines)) / 2

	for i, line := range contentLines {
		row := firstContentRow + i
		boxLines[row] = addTextToRow(boxLines[row], line)
	}

	return strings.Join(boxLines, "\n")
}

func addTextToRow(container, text string) string {
	begin := (len(container) - len(text)) / 2
	end := begin + len(text)

	return container[:begin] + text[:end-begin] + container[end:]
}
