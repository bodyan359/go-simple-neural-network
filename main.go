package main

import (
	"fmt"
	"math"
	"math/rand"
)

func sigmoid(x float64) float64{
	return 1 / (1 + math.Exp(x))
}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64() * (max - min)
	}
	return res
}

func dot(m1 [][]int, m2[][]int)  {
	m1ColLength := len(m1[0])
	m2RowLength := len(m2)
	if m1ColLength != m2RowLength {

	}
}

func main(){
	trainingInputs := [4][3]int{ 	{0, 0, 1},
									{1, 1, 1},
									{1, 0, 1},
									{0, 1, 1}	 }

	trainingOutputs := [4][1]int{	{0},
									{1},
									{1},
									{0}    }


	rand.Seed(1)
	synapticWeights := randFloats(-1, 1, 3)

	fmt.Println(trainingInputs)
	fmt.Println(trainingOutputs)
	fmt.Println(synapticWeights)
	// back propagation learning algorithm
	//input_layer := trainingInputs
	//outputs = sigmoid()
	//
	//fmt.Println(cmplx.Exp(trainingOutputs))
}
