package colors

type colorCode uint8

const (
	BlackFg   colorCode = 30
	RedFg     colorCode = 31
	GreenFg   colorCode = 32
	YellowFg  colorCode = 33
	BlueFg    colorCode = 34
	MagentaFg colorCode = 35
	CyanFg    colorCode = 36
	WhiteFg   colorCode = 37
	DefaultFg colorCode = 39

	BlackBg   colorCode = 40
	RedBg     colorCode = 41
	GreenBg   colorCode = 42
	YellowBg  colorCode = 43
	BlueBg    colorCode = 44
	MagentaBg colorCode = 45
	CyanBg    colorCode = 46
	WhiteBg   colorCode = 47
	DefaultBg colorCode = 49

	BrightBlackFg   colorCode = 90
	BrightRedFg     colorCode = 91
	BrightGreenFg   colorCode = 92
	BrightYellowFg  colorCode = 93
	BrightBlueFg    colorCode = 94
	BrightMagentaFg colorCode = 95
	BrightCyanFg    colorCode = 96
	BrightWhiteFg   colorCode = 97

	BrightBlackBg   colorCode = 100
	BrightRedBg     colorCode = 101
	BrightGreenBg   colorCode = 102
	BrightYellowBg  colorCode = 103
	BrightBlueBg    colorCode = 104
	BrightMagentaBg colorCode = 105
	BrightCyanBg    colorCode = 106
	BrightWhiteBg   colorCode = 107
)
