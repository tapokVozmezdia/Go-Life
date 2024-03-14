package life

import "fmt"

/*
	Hard-coded for neural network with 3 input neurons,
	2 output neurons and 1 hidden layer
*/

type NeuralNet struct {
	InpNeurons [3]InputNeuron
	HidNeurons [hiddenNeuronNumber]HiddenNeuron
	OutNeurons [2]OutputNeuron
}

func (nn *NeuralNet) SetInput(inp Vector2, neighbor_num uint8) {
	nn.InpNeurons[0].input = (inp.X - 480.) / 480.
	nn.InpNeurons[1].input = (inp.Y - 270.) / 270.
	nn.InpNeurons[2].input = float64(neighbor_num) / 20.
}

func (nn *NeuralNet) RandomizeWeights() {
	for i := 0; i < 3; i++ {
		nn.InpNeurons[i].RandomizeWeights()
	}
	for i := 0; i < hiddenNeuronNumber; i++ {
		nn.HidNeurons[i].RandomizeWeights()
	}
}

func (nn *NeuralNet) Forward() {
	for i := 0; i < 3; i++ {
		fmt.Println(nn.InpNeurons[i].output)
		nn.InpNeurons[i].Activate()
		fmt.Println(nn.InpNeurons[i].output)
	}

	for i := 0; i < hiddenNeuronNumber; i++ {
		nn.HidNeurons[i].input = 0
		for j := 0; j < 3; j++ {
			nn.HidNeurons[i].input += nn.InpNeurons[j].output
		}
		nn.HidNeurons[i].Activate()
	}

	for i := 0; i < 2; i++ {
		nn.OutNeurons[i].input = 0
		for j := 0; j < hiddenNeuronNumber; j++ {
			nn.OutNeurons[i].input += nn.HidNeurons[j].output
		}
		nn.OutNeurons[i].Activate()
	}
}
