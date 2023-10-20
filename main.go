package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	width, heigth := int32(512), int32(512)
	order := 1
	n := int32(math.Pow(2, float64(order)))
	total := n * n
	path := []rl.Vector2{}
	for i := 0; i < int(total); i++ {
		path = append(path, hilbert(i))
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

func hilbert(i int) rl.Vector2 {
	vectorArr := []rl.Vector2{}

	vectorArr = append(vectorArr, rl.Vector2{X: 0, Y: 0})
	vectorArr = append(vectorArr, rl.Vector2{X: 0, Y: 1})
	vectorArr = append(vectorArr, rl.Vector2{X: 1, Y: 1})
	vectorArr = append(vectorArr, rl.Vector2{X: 1, Y: 0})

	return vectorArr[i]
}
