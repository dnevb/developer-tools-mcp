package converter

import (
	"context"
	"fmt"
	"image/color"
	"regexp"
	"strings"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/muesli/gamut"
	"github.com/muesli/gamut/palette"
)

// tabs, newlines, whitespace, etc
var re = regexp.MustCompile(`\s+`)

func getColorName(c color.Color) string {
	colorMatchList, _ := palette.Wikipedia.Name(c)
	if len(colorMatchList) <= 0 {
		return ""
	}
	return colorMatchList[0].Name
}

func FormatColor(c color.Color) string {
	formated := fmt.Sprintf("code: %s", gamut.ToHex(c))
	formated = fmt.Sprintf("%s name: %s", formated, getColorName(c))

	return formated
}
func ParseColor(colorstr string) (color.Color, error) {
	// remove all whitespace
	colorstr = re.ReplaceAllString(colorstr, "")
	colorstr = strings.ToLower(strings.TrimSpace(colorstr))

	if strings.HasPrefix(colorstr, "#") {
		return colorful.Hex(colorstr)
	} else if strings.HasPrefix(colorstr, "rgb") {
		var r, g, b uint8
		_, err := fmt.Sscanf(colorstr, "rgb(%d,%d,%d)", &r, &g, &b)
		if err != nil {
			return nil, err
		}
		return color.RGBA{R: r, G: g, B: b, A: 255}, nil
	} else if strings.HasPrefix(colorstr, "hsl") {
		var h, s, l float64
		_, err := fmt.Sscanf(colorstr, "hsl(%f,%f%%,%f%%)", &h, &s, &l)
		if err != nil {
			return nil, err
		}
		return colorful.Hsl(h, s/100, l/100), nil
	} else if strings.HasPrefix(colorstr, "hsv") {
		var h, s, v float64
		_, err := fmt.Sscanf(colorstr, "hsv(%f,%f%%,%f%%)", &h, &s, &v)
		if err != nil {
			return nil, err
		}
		return colorful.Hsv(h, s/100, v/100), nil
	} else if strings.HasPrefix(colorstr, "cmyk") {
		var c, m, y, k uint8
		_, err := fmt.Sscanf(colorstr, "cmyk(%d%%,%d%%,%d%%,%d%%)", &c, &m, &y, &k)
		if err != nil {
			return nil, err
		}
		return color.CMYK{C: c, M: m, Y: y, K: k}, nil
	}
	color, err := colorful.Hex("#" + colorstr)
	if err != nil {
		return nil, fmt.Errorf("the given color %s format is not supported", colorstr)
	}

	return color, nil
}

type ColorConverterParams = mcp.CallToolParamsFor[struct {
	Color string `json:"color" jsonschema:"Color string to convert. Supported formats: hex, rgb, hsl, hsv, cmyk."`
}]

func ColorConverter(
	ctx context.Context,
	session *mcp.ServerSession,
	params *ColorConverterParams,
) (*mcp.CallToolResultFor[any], error) {
	c, err := ParseColor(params.Arguments.Color)
	if err != nil {
		return nil, fmt.Errorf("format color %w is not supported", err)
	}

	rgba := color.RGBAModel.Convert(c).(color.RGBA)

	colorfulColor, _ := colorful.MakeColor(rgba)
	h, s, l := colorfulColor.Hsl()
	hsv_h, hsv_s, hsv_v := colorfulColor.Hsv()
	cmyk_c, cmyk_m, cmyk_y, cmyk_k := color.RGBToCMYK(rgba.R, rgba.G, rgba.B)

	var resultBuilder strings.Builder
	resultBuilder.WriteString(fmt.Sprintf("hex: #%02x%02x%02x\n", rgba.R, rgba.G, rgba.B))
	resultBuilder.WriteString(fmt.Sprintf("rgb: rgb(%d, %d, %d)\n", rgba.R, rgba.G, rgba.B))
	resultBuilder.WriteString(fmt.Sprintf("hsl: hsl(%.2f, %.2f%%, %.2f%%)\n", h, s*100, l*100))
	resultBuilder.WriteString(fmt.Sprintf("hsv: hsv(%.2f, %.2f%%, %.2f%%)\n", hsv_h, hsv_s*100, hsv_v*100))
	resultBuilder.WriteString(fmt.Sprintf("cmyk: cmyk(%d%%, %d%%, %d%%, %d%%)\n", uint32(cmyk_c)*100/0xffff, uint32(cmyk_m)*100/0xffff, uint32(cmyk_y)*100/0xffff, uint32(cmyk_k)*100/0xffff))
	resultBuilder.WriteString(fmt.Sprintf("name: %s\n", getColorName(c)))

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{
			Text: strings.TrimSpace(resultBuilder.String()),
		}},
	}, nil
}
