package life

import (
	"math/rand/v2"
)

type Simulation struct {
	Cells []Cell
}

func (sim *Simulation) StartSim(cell_num int) {

	sim.Cells = nil

	for i := 0; i < cell_num; i++ {

		v := Vector2{
			float64(rand.IntN(960)),
			float64(rand.IntN(540)),
		}

		sim.Cells = append(sim.Cells, *CreateCell(v, 5))
	}
}

func (sim *Simulation) UpdateAll() {
	for i, _ := range sim.Cells {
		sim.Cells[i].Update()
	}
}

func (sim *Simulation) PullAllInRad(mouse_pos Vector2, rad float64, force float64) {
	for i, _ := range sim.Cells {
		if GetLenBetweenVectors(
			&(sim.Cells[i].pos),
			&mouse_pos,
		) > rad {
			continue
		}

		momentum := Vector2{
			mouse_pos.X - sim.Cells[i].pos.X,
			mouse_pos.Y - sim.Cells[i].pos.Y,
		}

		momentum.Normalize()

		momentum.X *= force
		momentum.Y *= force

		sim.Cells[i].MovePosDelta(momentum)
	}
}
