package log

type Color string

const Reset Color = "\u001b[0m"

const (
	ForeBlack 	Color = "\u001b[30m"
	ForeRed 		Color = "\u001b[31m"
	ForeGreen 	Color = "\u001b[32m"
	ForeYellow 	Color = "\u001b[33m"
	ForeBlue	 	Color = "\u001b[34m"
	ForeMagenta	Color = "\u001b[35m"
	ForeCyan		Color = "\u001b[36m"
	ForeWhite		Color = "\u001b[37m"
)

const (
	BackBlack 	Color = "\u001b[40m"
	BackRed 		Color = "\u001b[41m"
	BackGreen 	Color = "\u001b[42m"
	BackYellow 	Color = "\u001b[43m"
	BackBlue	 	Color = "\u001b[44m"
	BackMagenta	Color = "\u001b[45m"
	BackCyan		Color = "\u001b[46m"
	BackWhite		Color = "\u001b[47m"
)