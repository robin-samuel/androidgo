package androidgo

import "fmt"

type DisplayMetrics struct {
	HeightPixels int
	WidthPixels  int
}

func (d *DisplayMetrics) HeightPixelsStr() string {
	return fmt.Sprintf("%d", d.HeightPixels)
}

func (d *DisplayMetrics) WidthPixelsStr() string {
	return fmt.Sprintf("%d", d.WidthPixels)
}
