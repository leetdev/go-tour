// https://tour.golang.org/moretypes/18

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for y := 0; y < dy; y++ {
		var row []uint8
		for x := 0; x < dx; x++ {
			row = append(row, uint8(x*x+y*y+x*y))
		}
		pic = append(pic, row)
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
