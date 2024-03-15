package life

type Cell struct {
	pos    Vector2
	radius float64

	// for neural net (3rd neuron input)
	bumper float64

	Brain NeuralNet
}

func CreateCell(p Vector2, rad float64) *Cell {
	return &Cell{
		pos:    p,
		radius: rad,
	}
}

// type CellInterface interface {
// 	MovePosAbs(n_pos Vector2) Vector2
// 	MovePosDelta(delta Vector2) Vector2
// 	GetPos() Vector2

// 	SetRadius(rad float64)
// 	GetRadius() float64
// }

func (cell *Cell) MovePosAbs(n_pos Vector2) Vector2 {
	cell.pos = n_pos
	return cell.pos
}

func (cell *Cell) MovePosDelta(delta Vector2) Vector2 {
	cell.pos.X += delta.X
	cell.pos.Y += delta.Y
	return cell.pos
}

func (cell Cell) GetPos() Vector2 {
	return cell.pos
}

func (cell *Cell) SetRadius(rad float64) {
	cell.radius = rad
}

func (cell *Cell) GetRadius() float64 {
	return cell.radius
}

func (cell *Cell) Act() {

	d_x := cell.Brain.OutNeurons[0].output * 2
	d_y := cell.Brain.OutNeurons[1].output * 2

	if cell.pos.X+d_x < 953 && cell.pos.X+d_x > 7 {
		cell.MovePosDelta(Vector2{d_x, 0})
	} else {
		cell.bumper = 1.
	}

	if cell.pos.Y+d_y < 533 && cell.pos.Y+d_y > 7 {
		cell.MovePosDelta(Vector2{0, d_y})
	} else {
		cell.bumper = 1.
	}
}

func (cell *Cell) Update() {

	// randomized movement

	// v := Vector2{
	// 	float64(rand.IntN(3) - 1),
	// 	float64(rand.IntN(3) - 1),
	// }
	// cell.MovePosDelta(v)

	// movement with NeuralNet

	if cell.bumper > 0 {
		cell.bumper -= 1. / 120.
	}

	if cell.bumper < 0 {
		cell.bumper = 0.
	}

	cell.Brain.SetInput(cell.pos, cell.bumper)

	cell.Brain.Forward()

	cell.Act()
}
