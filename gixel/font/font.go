package font

import (
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

type GxlFont struct {
	tt      *sfnt.Font
	presets map[float64]*GxlFontPreset
}

func NewFont(path string) *GxlFont {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var font GxlFont
	font.tt, err = opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	font.presets = make(map[float64]*GxlFontPreset)

	return &font
}

func (f *GxlFont) GetPreset(size float64) *GxlFontPreset {
	preset, ok := f.presets[size]
	if ok {
		return preset
	}

	face, err := opentype.NewFace(f.tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72, // TODO: Support high dpi displays
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	preset = &GxlFontPreset{
		size: size,
		face: &face,
	}
	// Add to cached presets
	f.presets[size] = preset

	return preset
}

type GxlFontPreset struct {
	size float64
	face *font.Face
}

func (fp *GxlFontPreset) GetSize() float64 {
	return fp.size
}

func (fp *GxlFontPreset) GetFace() font.Face {
	return *fp.face
}
