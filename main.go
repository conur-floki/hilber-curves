package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	width, heigth := int32(512), int32(512)
	order := 4
	n := int32(math.Pow(2, float64(order)))
	total := n * n
	path := []rl.Vector2{}

	for i := 0; i < int(total); i++ {
		path = append(path, hilbert(i, order))
		len := float32(width / n)
		path[i].X = path[i].X * len
		path[i].Y = path[i].Y * len
		path[i].X = path[i].X + len/2
		path[i].Y = path[i].Y + len/2
	}

	rl.InitWindow(width, heigth, "Hilbert Curve en Raylib con Go")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for i, v := range path {
			nextVec := i + 1
			if nextVec < len(path) {
				rl.DrawLineV(v, path[nextVec], rl.Black)
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func hilbert(i, order int) rl.Vector2 {
	vectorArr := []rl.Vector2{}

	vectorArr = append(vectorArr, rl.Vector2{X: 0, Y: 0})
	vectorArr = append(vectorArr, rl.Vector2{X: 0, Y: 1})
	vectorArr = append(vectorArr, rl.Vector2{X: 1, Y: 1})
	vectorArr = append(vectorArr, rl.Vector2{X: 1, Y: 0})

	index := i & 3
	v := vectorArr[index]
	for j := 1; j < order; j++ {
		i = i >> 2
		index = i & 3
		len := math.Pow(2, float64(j))
		switch index {
		case 0:
			temp := v.X
			v.X = v.Y
			v.Y = temp
		case 1:
			v.Y += float32(len)
		case 2:
			v.X += float32(len)
			v.Y += float32(len)
		case 3:
			temp := float32(len) - 1 - v.X
			v.X = float32(len) - 1 - v.Y
			v.Y = temp
			v.X += float32(len)
		}
	}
	return v
}
