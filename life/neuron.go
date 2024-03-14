package life

/*
	Hard-coded for neural network with 3 input neurons,
	2 output neurons and 1 hidden layer
*/

import (
	"math"
	"math/rand/v2"
)

const hiddenNeuronNumber = 4

type InputNeuron struct {
	input  float64
	output float64

	weights [hiddenNeuronNumber]float64
}

type HiddenNeuron struct {
	input  float64
	output float64

	weights [2]float64
}

type OutputNeuron struct {
	input  float64
	output float64
}

type NeuronInterface interface {
	RandomizeWeights()
	Activate()

	GetWeights() []float64
	GetInfo() Vector2
}

func (n *InputNeuron) RandomizeWeights() {
	for i := 0; i < hiddenNeuronNumber; i++ {
		n.weights[i] = rand.Float64()*2. - 1.
	}
}

func (n *InputNeuron) Activate() {
	n.output = n.input
}

func (n *HiddenNeuron) RandomizeWeights() {
	for i := 0; i < 2; i++ {
		n.weights[i] = rand.Float64()*2. - 1.
	}
}

func (n *HiddenNeuron) Activate() {
	n.output = math.Tanh(n.input)
}

func (n *OutputNeuron) Activate() {
	n.output = math.Tanh(n.input)
}

func (n *InputNeuron) GetInfo() Vector2 {
	return Vector2{
		n.input,
		n.output,
	}
}
func (n *HiddenNeuron) GetInfo() Vector2 {
	return Vector2{
		n.input,
		n.output,
	}
}
func (n *OutputNeuron) GetInfo() Vector2 {
	return Vector2{
		n.input,
		n.output,
	}
}

func (n *InputNeuron) GetWeights() []float64 {
	return n.weights[:]
}

// func (n *Neuron) Activate() {
// 	n.output = math.Tanh(n.input)
// }
