package main

import (
	"fmt"

	"Go-Life/life"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func DrawCells(sim *(life.Simulation), imd *imdraw.IMDraw) {
	for _, cell := range sim.Cells {
		imd.Push(pixel.V(cell.GetPos().X, cell.GetPos().Y))
		imd.Circle(cell.GetRadius(), 2)
	}
}

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Lights",
		Bounds: pixel.R(0, 0, 960, 540),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Color = colornames.Aqua

	// imd.Push(pixel.V(480, 270))
	// imd.Circle(200, 20)

	simulation := life.Simulation{}

	for !win.Closed() {
		win.Clear(pixel.RGB(0.3, 0, 0))
		imd.Draw(win)

		if win.JustPressed(pixelgl.KeySpace) {
			imd.Clear()
			simulation.StartSim(50)
		}

		if win.Pressed(pixelgl.MouseButtonLeft) {
			simulation.PullAllInRad(
				life.Vector2(win.MousePosition()),
				100,
				1,
			)
		}

		if len(simulation.Cells) > 0 {
			fmt.Println(simulation.Cells[0].GetPos())
		}

		imd.Clear()
		DrawCells(&simulation, imd)
		simulation.UpdateAll()
		win.Update()
	}
}

func main() {

	fmt.Println("Hello, Go!")

	pixelgl.Run(run)

}
