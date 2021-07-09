package main

import (
  "code.google.com/p/go-tour/pic"
  "image"
  "image/color"
)

type Image struct{
  Height, Width int
}

func (i Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, i.Height, i.Width)
}

func (i Image) At(x, y int) color.Color {
  c := uint8(x ^ y)
  return color.RGBA{c, c, 255, 255}
}

func main() {
  m := Image{256, 256}
  pic.ShowImage(m)
}
