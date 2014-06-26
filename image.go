package gotourexercises

import (
    "image"
    "image/color"
)

type Image struct{
    w, h int
}

func (img Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, img.w, img.h)
}

func (img Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (img Image) At(x int, y int) color.Color {
    v := uint8((x+y)/2)
    return color.RGBA{v, v, 255, 255}
}