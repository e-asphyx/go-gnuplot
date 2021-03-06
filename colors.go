package gnuplot

import "fmt"

type Color interface {
	Color() string
}

type NamedColor string

func (n NamedColor) Color() string {
	return string(n)
}

type RGB struct {
	R uint8
	G uint8
	B uint8
}

func (r RGB) Color() string {
	return fmt.Sprintf("#%02x%02x%02x", r.R, r.G, r.B)
}

type RGBA struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (r RGBA) Color() string {
	return fmt.Sprintf("#%02x%02x%02x%02x", r.A, r.R, r.G, r.B)
}

var (
	ColorWhite           NamedColor = "white"
	ColorBlack           NamedColor = "black"
	ColorDarkGrey        NamedColor = "dark-grey"
	ColorRed             NamedColor = "red"
	ColorWebGreen        NamedColor = "web-green"
	ColorWebBlue         NamedColor = "web-blue"
	ColorDarkMagenta     NamedColor = "dark-magenta"
	ColorDarkCyan        NamedColor = "dark-cyan"
	ColorDarkOrange      NamedColor = "dark-orange"
	ColorDarkYellow      NamedColor = "dark-yellow"
	ColorRoyalblue       NamedColor = "royalblue"
	ColorGoldenrod       NamedColor = "goldenrod"
	ColorDarkSpringGreen NamedColor = "dark-spring-green"
	ColorPurple          NamedColor = "purple"
	ColorSteelblue       NamedColor = "steelblue"
	ColorDarkRed         NamedColor = "dark-red"
	ColorDarkChartreuse  NamedColor = "dark-chartreuse"
	ColorOrchid          NamedColor = "orchid"
	ColorAquamarine      NamedColor = "aquamarine"
	ColorBrown           NamedColor = "brown"
	ColorYellow          NamedColor = "yellow"
	ColorTurquoise       NamedColor = "turquoise"
	ColorGrey0           NamedColor = "grey0"
	ColorGrey10          NamedColor = "grey10"
	ColorGrey20          NamedColor = "grey20"
	ColorGrey30          NamedColor = "grey30"
	ColorGrey40          NamedColor = "grey40"
	ColorGrey50          NamedColor = "grey50"
	ColorGrey60          NamedColor = "grey60"
	ColorGrey70          NamedColor = "grey70"
	ColorGrey            NamedColor = "grey"
	ColorGrey80          NamedColor = "grey80"
	ColorGrey90          NamedColor = "grey90"
	ColorGrey100         NamedColor = "grey100"
	ColorLightRed        NamedColor = "light-red"
	ColorLightGreen      NamedColor = "light-green"
	ColorLightBlue       NamedColor = "light-blue"
	ColorLightMagenta    NamedColor = "light-magenta"
	ColorLightCyan       NamedColor = "light-cyan"
	ColorLightGoldenrod  NamedColor = "light-goldenrod"
	ColorLightPink       NamedColor = "light-pink"
	ColorLightTurquoise  NamedColor = "light-turquoise"
	ColorGold            NamedColor = "gold"
	ColorGreen           NamedColor = "green"
	ColorDarkGreen       NamedColor = "dark-green"
	ColorSpringGreen     NamedColor = "spring-green"
	ColorForestGreen     NamedColor = "forest-green"
	ColorSeaGreen        NamedColor = "sea-green"
	ColorBlue            NamedColor = "blue"
	ColorDarkBlue        NamedColor = "dark-blue"
	ColorMidnightBlue    NamedColor = "midnight-blue"
	ColorNavy            NamedColor = "navy"
	ColorMediumBlue      NamedColor = "medium-blue"
	ColorSkyblue         NamedColor = "skyblue"
	ColorCyan            NamedColor = "cyan"
	ColorMagenta         NamedColor = "magenta"
	ColorDarkTurquoise   NamedColor = "dark-turquoise"
	ColorDarkPink        NamedColor = "dark-pink"
	ColorCoral           NamedColor = "coral"
	ColorLightCoral      NamedColor = "light-coral"
	ColorOrangeRed       NamedColor = "orange-red"
	ColorSalmon          NamedColor = "salmon"
	ColorDarkSalmon      NamedColor = "dark-salmon"
	ColorKhaki           NamedColor = "khaki"
	ColorDarkKhaki       NamedColor = "dark-khaki"
	ColorDarkGoldenrod   NamedColor = "dark-goldenrod"
	ColorBeige           NamedColor = "beige"
	ColorOlive           NamedColor = "olive"
	ColorOrange          NamedColor = "orange"
	ColorViolet          NamedColor = "violet"
	ColorDarkViolet      NamedColor = "dark-violet"
	ColorPlum            NamedColor = "plum"
	ColorDarkPlum        NamedColor = "dark-plum"
	ColorDarkOlivegreen  NamedColor = "dark-olivegreen"
	ColorOrangered4      NamedColor = "orangered4"
	ColorBrown4          NamedColor = "brown4"
	ColorSienna4         NamedColor = "sienna4"
	ColorOrchid4         NamedColor = "orchid4"
	ColorMediumpurple3   NamedColor = "mediumpurple3"
	ColorSlateblue1      NamedColor = "slateblue1"
	ColorYellow4         NamedColor = "yellow4"
	ColorSienna1         NamedColor = "sienna1"
	ColorTan1            NamedColor = "tan1"
	ColorSandybrown      NamedColor = "sandybrown"
	ColorLightSalmon     NamedColor = "light-salmon"
	ColorPink            NamedColor = "pink"
	ColorKhaki1          NamedColor = "khaki1"
	ColorLemonchiffon    NamedColor = "lemonchiffon"
	ColorBisque          NamedColor = "bisque"
	ColorHoneydew        NamedColor = "honeydew"
	ColorSlategrey       NamedColor = "slategrey"
	ColorSeagreen        NamedColor = "seagreen"
	ColorAntiquewhite    NamedColor = "antiquewhite"
	ColorChartreuse      NamedColor = "chartreuse"
	ColorGreenyellow     NamedColor = "greenyellow"
	ColorGray            NamedColor = "gray"
	ColorLightGray       NamedColor = "light-gray"
	ColorLightGrey       NamedColor = "light-grey"
	ColorDarkGray        NamedColor = "dark-gray"
	ColorSlategray       NamedColor = "slategray"
	ColorGray0           NamedColor = "gray0"
	ColorGray10          NamedColor = "gray10"
	ColorGray20          NamedColor = "gray20"
	ColorGray30          NamedColor = "gray30"
	ColorGray40          NamedColor = "gray40"
	ColorGray50          NamedColor = "gray50"
	ColorGray60          NamedColor = "gray60"
	ColorGray70          NamedColor = "gray70"
	ColorGray80          NamedColor = "gray80"
	ColorGray90          NamedColor = "gray90"
	ColorGray100         NamedColor = "gray100"
)
