package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main(){
	trainingInputs := [][]int{ 	{0, 0, 1},
								{1, 1, 1},
								{1, 0, 1},
								{0, 1, 1}	 }

	trainingOutputs := [4][1]float64{	{0},
										{1},
										{1},
										{0}    }

	rand.Seed(1)
	synapticWeights := randFloats(-1, 1, 3)

	// back propagation learning algorithm
	var inputLayer [][]int
	var outputs [4][1]float64


	inputLayer = trainingInputs
	outputs = sigmoid(dot(inputLayer, synapticWeights))

	err := arrayMinusArray(trainingOutputs, outputs)
	//adjustments := dot(transposeArray(inputLayer), synapticWeights)


	fmt.Println(err)
	fmt.Println(inputLayer)
	fmt.Println(outputs)
}

func sigmoid(array [4][1]float64) [4][1]float64 {
	for i := 0; i < len(array); i++ {
		array[i][0] = 1 / (1 + math.Exp(-array[i][0]))
	}

	return array
}

func randFloats(min, max float64, n int) [3][1]float64 {
	var res [3][1]float64
	for i := 0; i < n; i++ {
		res[i][0] = min + rand.Float64() * (max - min)
	}

	return res
}

func dot(array1 [][]int, array2 [3][1]float64) [4][1]float64{
	count := 0
	var result float64
	var resultArray [4][1]float64

	for i := 0; i < len(array1); i++ {
		result = 0
		count = 0
		for j := 0; j < len(array1[i]); j++ {
			result += array2[count][0] * float64(array1[i][j])
			count++
		}
		resultArray[i][0] = result
	}

	return resultArray
}

//func errorDot(array1 [3][4]int, array2 [3][1]float64) [4][1]float64{
//
//}

func arrayMinusArray(array1 [4][1]float64, array2 [4][1]float64) [4][1]float64{
	var resultArray [4][1]float64

	for i := 0; i < len(array1); i++ {
		resultArray[i][0] = array1[i][0] - array2[i][0]
	}

	return resultArray
}

func transposeArray(array [][]int) [3][4]int{
	var transposedArray [3][4]int
	for i := 0; i < len(transposedArray); i++ {
		for j := 0; j < len(transposedArray[i]); j++ {
			transposedArray[i][j] = array[j][i]
		}
	}

	return transposedArray
}