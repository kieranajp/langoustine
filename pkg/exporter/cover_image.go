package exporter

import (
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/kieranajp/langoustine/pkg/domain"
	"github.com/pkg/errors"
)

type CoverImage struct {
	Image image.Image
}

const coverWidth = 1600
const coverHeight = 2560
const backgroundImageUrl = "https://picsum.photos"

//go:embed "res/PlayfairDisplay-Bold.ttf"
var fontFace []byte

func NewCoverImage() *CoverImage {
	return &CoverImage{}
}

func (c *CoverImage) Generate(recipe *domain.Recipe) error {
	dc := gg.NewContext(coverWidth, coverHeight)
	if err := c.drawBackgroundImage(dc); err != nil {
		return err
	}

	c.drawOverlay(dc)
	if err := c.drawTitle(dc, recipe.Name); err != nil {
		return err
	}
	c.Image = dc.Image()

	return nil
}

func (c *CoverImage) Write(w io.Writer) (int64, error) {
	if err := png.Encode(w, c.Image); err != nil {
		return 0, err
	}

	// TODO: Probably wrong, but whatever.
	return int64(c.Image.Bounds().Dx()) * int64(c.Image.Bounds().Dy()), nil
}

func (c *CoverImage) drawBackgroundImage(dc *gg.Context) error {
	response, err := http.Get(fmt.Sprintf("%s/%d/%d", backgroundImageUrl, coverWidth, coverHeight))
	if err != nil {
		return errors.Wrap(err, "Failed to load background image")
	}
	backgroundImage, _, err := image.Decode(response.Body)
	if err != nil {
		return errors.Wrap(err, "Failed to decode background image")
	}

	dc.DrawImage(backgroundImage, 0, 0)
	return nil
}

func (c *CoverImage) drawOverlay(dc *gg.Context) {
	x := 80.0
	y := 80.0
	w := float64(dc.Width()) - (2.0 * x)
	h := float64(dc.Height()) - (2.0 * y)
	dc.SetColor(color.RGBA{0, 0, 0, 204})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()
}

func (c *CoverImage) drawTitle(dc *gg.Context, recipeTitle string) error {
	f, err := truetype.Parse(fontFace)
	if err != nil {
		return errors.Wrap(err, "Failed to load font")
	}
	dc.SetFontFace(truetype.NewFace(f, &truetype.Options{
		Size: 160,
	}))

	textShadowColor := color.Black
	textColor := color.White
	x := 180.0
	y := 220.0

	maxWidth := float64(dc.Width()) - (2 * x)
	dc.SetColor(textShadowColor)
	dc.DrawStringWrapped(recipeTitle, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	dc.SetColor(textColor)
	dc.DrawStringWrapped(recipeTitle, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	return nil
}
