package colors

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	Options  func() func(*Color)
	ColorId  func(uint8) string
	RGBColor func(string) string
	RGB      struct {
		R, G, B uint8
	}
	Color struct {
		bold          bool
		underline     bool
		dim           bool
		italic        bool
		blink         bool
		inverse       bool
		invisible     bool
		strikethrough bool
		background    bool
		isFgCustomId  bool
		isFgRGB       bool
		isBgCustomId  bool
		isBgRGB       bool
		fgCode        colorCode
		fgRGB         RGB
		bgCode        colorCode
		bgRGB         RGB
	}
)

func New(code colorCode, options ...Options) *Color {
	c := &Color{}
	for _, option := range options {
		option()(c)
	}

	c.fgCode = code

	return c
}

func NewCustomId(id uint8, options ...Options) *Color {
	c := &Color{}
	for _, option := range options {
		option()(c)
	}

	c.fgCode = colorCode(id)
	c.isFgCustomId = true

	return c
}

func NewRGB(rgb RGB, options ...Options) *Color {
	c := &Color{}
	for _, option := range options {
		option()(c)
	}

	c.fgRGB = rgb
	c.isFgRGB = true

	return c
}

func (c Color) String(s string) string {
	if len(s) == 0 {
		return s
	}

	return c.createColor(s)
}

func (c *Color) createColor(s string) string {
	codes := []string{}

	if c.background {
		switch {
		case c.isBgRGB:
			codes = append(codes, fmt.Sprintf("48;2;%d;%d;%d", c.bgRGB.R, c.bgRGB.G, c.bgRGB.B))
		case c.isBgCustomId:
			codes = append(codes, fmt.Sprintf("48;5;%d", c.bgCode))
		default:
			codes = append(codes, strconv.Itoa(int(c.fgCode)))
		}
	}

	if c.isFgRGB {
		codes = append(codes, fmt.Sprintf("38;2;%d;%d;%d", c.fgRGB.R, c.fgRGB.G, c.fgRGB.B))
	} else if c.isFgCustomId {
		codes = append(codes, fmt.Sprintf("38;5;%d", c.fgCode))
	} else {
		codes = append(codes, strconv.Itoa(int(c.fgCode)))
	}

	c.setFontTypes(&codes)

	sb := strings.Builder{}

	joinedCodes := strings.Join(codes, ";")

	sb.Grow(len("\033[") + len(joinedCodes) + 1 + len(s) + len(resetAll()))

	sb.WriteString("\033[")
	sb.WriteString(joinedCodes)
	sb.WriteString("m")

	sb.WriteString(s)
	sb.WriteString(resetAll())

	return sb.String()
}

func (c *Color) EnableBold() {
	c.bold = true
}

func (c *Color) EnableUnderline() {
	c.bold = true
}

func (c *Color) EnableDim() {
	c.dim = true
}

func (c *Color) EnableItalic() {
	c.italic = true
}

func (c *Color) EnableBlink() {
	c.blink = true
}

func (c *Color) EnableInverse() {
	c.inverse = true
}

func (c *Color) EnableInvisible() {
	c.invisible = true
}

func (c *Color) EnableStrikethrough() {
	c.strikethrough = true
}

func (c *Color) DisableBold() {
	c.bold = false
}

func (c *Color) DisableUnderline() {
	c.bold = false
}

func (c *Color) DisableDim() {
	c.dim = false
}

func (c *Color) DisableItalic() {
	c.italic = false
}

func (c *Color) DisableBlink() {
	c.blink = false
}

func (c *Color) DisableInverse() {
	c.inverse = false
}

func (c *Color) DisableInvisible() {
	c.invisible = false
}

func WithBold() func(c *Color) {
	return func(c *Color) {
		c.bold = true
	}
}

func WithUnderline() func(c *Color) {
	return func(c *Color) {
		c.underline = true
	}
}

func WithDim() func(c *Color) {
	return func(c *Color) {
		c.dim = true
	}
}

func WithItalic() func(c *Color) {
	return func(c *Color) {
		c.italic = true
	}
}

func WithBlink() func(c *Color) {
	return func(c *Color) {
		c.blink = true
	}
}

func WithInverse() func(c *Color) {
	return func(c *Color) {
		c.inverse = true
	}
}

func WithInvisible() func(c *Color) {
	return func(c *Color) {
		c.invisible = true
	}
}

func WithStrikethrough() func(c *Color) {
	return func(c *Color) {
		c.strikethrough = true
	}
}

func WithBackground(code colorCode) func() func(c *Color) {
	return func() func(c *Color) {
		return func(c *Color) {
			if !c.isBgCustomId && !c.isBgRGB {
				c.bgCode = code
				c.background = true
			}
		}
	}
}

func WithBackgroundId(code colorCode) func() func(c *Color) {
	return func() func(c *Color) {
		return func(c *Color) {
			if !c.isBgCustomId && !c.isBgRGB {
				c.bgCode = code
				c.background = true
				c.isBgCustomId = true
			}
		}
	}
}

func WithBackgroundRGB(rgb RGB) func() func(c *Color) {
	return func() func(c *Color) {
		return func(c *Color) {
			if !c.isBgCustomId && !c.isBgRGB {
				c.bgRGB = rgb
				c.background = true
				c.isBgRGB = true
			}
		}
	}
}

func (c *Color) setFontTypes(codes *[]string) {
	if c.bold {
		*codes = append(*codes, "1")
	}

	if c.dim {
		*codes = append(*codes, "2")
	}

	if c.italic {
		*codes = append(*codes, "3")
	}

	if c.underline {
		*codes = append(*codes, "4")
	}

	if c.blink {
		*codes = append(*codes, "5")
	}

	if c.inverse {
		*codes = append(*codes, "7")
	}

	if c.invisible {
		*codes = append(*codes, "8")
	}

	if c.strikethrough {
		*codes = append(*codes, "9")
	}
}

func resetAll() string {
	return "\033[0m"
}
