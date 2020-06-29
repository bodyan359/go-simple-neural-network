package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main(){
	trainingInputs := [][]int{
		{0, 0, 1},
		{1, 1, 1},
		{1, 0, 1},
		{0, 1, 1}	 }

	trainingOutputs := [4][1]float64{
		{0},
		{1},
		{1},
		{0}    }

	rand.Seed(1)
	synapticWeights := randFloats(-1, 1, 3)
	fmt.Println("Weights before training:\n", synapticWeights)

	// back propagation learning algorithm
	var inputLayer [][]int
	var outputs [4][1]float64

	for i := 0; i < 20000; i++ {
		inputLayer = trainingInputs
		outputs = sigmoid(dot(inputLayer, synapticWeights))

		err := arrayMinusArray(trainingOutputs, outputs)
		outputsMinus := numMinusArray(1, outputs)
		outputsSum := dotFloats(outputs, outputsMinus)
		errOutputsSum := dotFloats(err, outputsSum)
		adjustments := dotAdj(transposeArray(inputLayer), errOutputsSum)
		synapticWeights = prepareWeights(synapticWeights, adjustments)
	}

	fmt.Println("Weights after training:\n", synapticWeights)
	fmt.Println("Result:\n", outputs)

	// new situation
	newInputs := []int{ 1, 1, 0}
	output := newSigmoid(newDot(newInputs, synapticWeights))
	fmt.Println("Result for new situation:", output)
}

// math operations with matrix, generally built on NumPy (Python library) functions
func prepareWeights(array1 [3][1]float64, array2 [4][1]float64) [3][1]float64 {
	for i := 0; i < len(array1); i++ {
		array1[i][0] = array1[i][0] + array2[i][0]
	}
	return array1
}

func dotAdj(array1 [3][4]int, array2 [4][1]float64) [4][1]float64{
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

func newSigmoid(num float64) float64 {
	return 1 / (1 + math.Exp(-num))
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

func newDot(array1 []int, array2 [3][1]float64) float64 {
	var result float64
	for i := 0; i < len(array2); i++ {
		result += array2[i][0] * float64(array1[i])
	}

	return result
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

func numMinusArray(num int, array [4][1]float64) [4][1]float64{
	for i := 0; i < len(array); i++ {
		array[i][0] = float64(1) - array[i][0]
	}

	return array
}

func dotFloats(array1 [4][1]float64, array2 [4][1]float64) [4][1]float64{
	count := 0
	var result float64
	var resultArray [4][1]float64

	for i := 0; i < len(array1); i++ {
		result = 0
		count = 0
		for j := 0; j < len(array1[i]); j++ {
			result += array2[count][0] * array1[i][j]
			count++
		}
		resultArray[i][0] = result
	}

	return resultArray
}