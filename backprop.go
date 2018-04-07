package main

import "time"
import "fmt"
import "os"
import "bufio"
import "math/rand"
import "math"

var trainExamplesPOS []string
var trainExamplesNEG []string
var testExamplesPOS []string
var testExamplesNEG []string

// TODO - Can we eliminate these and use counters that are local
//        to the read functions? We should be able to get these
//        values from counts.
var numTrainPositives int = 0
var numTrainNegatives int = 0
var numTestPositives int = 0
var numTestNegatives int = 0

const numInputUnits int = 40
const numHiddenUnits int = 10
const numOutputUnits int = 1
const rate float64 = 0.25
const numEpocs int = 10000

var randomSeed int64 = 0

var inputs [numInputUnits]float64

var weightsLayerOne [numInputUnits][numHiddenUnits]float64
var weightsLayerTwo [numHiddenUnits][numOutputUnits]float64
var weightsHiddenUnitsBias [numHiddenUnits]float64
var weightsOutputUnitsBias [numOutputUnits]float64
var hiddenLayerOutput [numHiddenUnits]float64
var outputLayerOutput [numOutputUnits]float64

func main() {

	pTime := fmt.Println
	startTime := time.Now()

	fmt.Printf("\n\n")
	fmt.Printf("                           Backprop.go\n\n")

	readTrainingSets()
	readTestingSets()

	for randomSeed = 102; randomSeed < 200; randomSeed += 10 {

		reset(randomSeed)
		train()
		test()
	}

	endTime := time.Now()
	diffTime := endTime.Sub(startTime)
	fmt.Printf("Runtime (secs): ")
	pTime(diffTime.Seconds())

}

// Read the training files. There must be two files, one positive and one negative
func readTrainingSets() {

	//TODO - Make the file reading generic and reuse the code for both positive and
	//			negatives by passing parameters.

	// Read positive training examples
	posFilePath := "./trainPOS.txt" // TODO - pass as parameter and get from commandline arg

	posFile, _ := os.Open(posFilePath) // TODO - error handling
	defer posFile.Close()
	posScanner := bufio.NewScanner(posFile)
	posScanner.Split(bufio.ScanLines)

	for posScanner.Scan() {
		// TODO - pass trainExamplesPOS into this function instead of global
		trainExamplesPOS = append(trainExamplesPOS, posScanner.Text())
		numTrainPositives++
	}

	// Read negative training examples
	negFilePath := "./trainNEG.txt" // TODO - pass as parameter and get from commandline arg

	negFile, _ := os.Open(negFilePath) // TODO- error handling
	defer negFile.Close()
	negScanner := bufio.NewScanner(negFile)
	negScanner.Split(bufio.ScanLines)

	for negScanner.Scan() {
		// TODO - pass trainExamplesNEG into this function instead of global
		trainExamplesNEG = append(trainExamplesNEG, negScanner.Text())
		numTrainNegatives++
	}

}

func readTestingSets() {

	//TODO - Make the file reading generic and reuse the code for both positive and
	//			negatives by passing parameters.

	// Read positive testing examples
	posFilePath := "./testPOS.txt" // TODO - pass as parameter and get from commandline arg

	posFile, _ := os.Open(posFilePath) // TODO - error handling
	defer posFile.Close()
	posScanner := bufio.NewScanner(posFile)
	posScanner.Split(bufio.ScanLines)

	for posScanner.Scan() {
		// TODO - pass trainExamplesPOS into this function instead of global
		testExamplesPOS = append(testExamplesPOS, posScanner.Text())
		numTestPositives++
	}

	// Read negative testing examples
	negFilePath := "./testNEG.txt" // TODO - pass as parameter and get from commandline arg

	negFile, _ := os.Open(negFilePath) // TODO- error handling
	defer negFile.Close()
	negScanner := bufio.NewScanner(negFile)
	negScanner.Split(bufio.ScanLines)

	for negScanner.Scan() {
		// TODO - pass trainExamplesNEG into this function instead of global
		testExamplesNEG = append(testExamplesNEG, negScanner.Text())
		numTestNegatives++
	}

}

// Reset the neural network to random values (0-1) and set random seed.
func reset(seed int64) {

	rand.Seed(seed)

	for j := 0; j < numHiddenUnits; j++ {
		for i := 0; i < numInputUnits; i++ {
			weightsLayerOne[i][j] = rand.Float64()
		}

		for k := 0; k < numOutputUnits; k++ {
			weightsLayerTwo[j][k] = rand.Float64()
		}

		weightsHiddenUnitsBias[j] = rand.Float64()
	}

	for k := 0; k < numOutputUnits; k++ {
		weightsOutputUnitsBias[k] = rand.Float64()
	}
}

// Train the neural network with the training data
func train() {

	// It would be better to remove this look and to put numEpocs into the
	// calls below as the third parameter.
	for num := 0; num < numEpocs; num++ {

		for p := 0; p < len(trainExamplesPOS); p++ {

			loadInputs(trainExamplesPOS[p])

			trainOneOutputUnitOnOneExampleForMultipleEpochs(0, 1, 1)

			loadInputs(trainExamplesNEG[p])

			trainOneOutputUnitOnOneExampleForMultipleEpochs(0, 0, 1)
		}
	}
}

func loadInputs(s string) {

	//fmt.Printf("inputs: [")

	for c := 0; c < len(s); c++ {
		inputs[c] = (float64)(s[c] - 48) // convert char value to numeric 1 or 0
	}
}

func trainOneOutputUnitOnOneExampleForMultipleEpochs(k int, label float64, numEpochs int) {

	for i := 0; i < numEpochs; i++ {
		trainOneOutputUnitOnOneExampleForOneEpoch(k, label)
	}
}

func trainOneOutputUnitOnOneExampleForOneEpoch(k int, d float64) {

	runNet() // Run the neural net then update values

	f := outputLayerOutput[k]

	deltaK := (d - f) * f * (1 - f)

	var deltaJ float64

	// First update the weight on the connection from output unit k's bias
	weightsOutputUnitsBias[k] += rate * deltaK

	// Next update the weight of the bias going into output unit k from hidden units
	for j := 0; j < numHiddenUnits; j++ {
		// First back prop the error from output unit k into hidden layer j.
		// Calculate the back prop error delta

		deltaJ = hiddenLayerOutput[j] * (1 - hiddenLayerOutput[j]) * (deltaK * weightsLayerTwo[j][k])

		// Then update the weight of the bias going into hidden unit j.
		weightsHiddenUnitsBias[j] += rate * deltaJ

		// Then update the weights on the connections from the input units into
		// hidden unit j
		for i := 0; i < numInputUnits; i++ {
			weightsLayerOne[i][j] += rate * deltaJ * (float64)(inputs[i])
		}

		// Last go back and update the weight going from hidden unit j to output unit k
		weightsLayerTwo[j][k] += rate * deltaK * hiddenLayerOutput[j]
	}
}

func runNet() { // TODO - change to parameters rather than globals

	var summedInput float64 = 0.0

	for j := 0; j < numHiddenUnits; j++ {

		summedInput = 0.0

		for i := 0; i < numInputUnits; i++ {

			summedInput += weightsLayerOne[i][j] * (inputs[i])
		}
		summedInput += weightsHiddenUnitsBias[j]
		hiddenLayerOutput[j] = sigmoid(summedInput)
	}

	for k := 0; k < numOutputUnits; k++ {

		summedInput = 0.0
		for j := 0; j < numHiddenUnits; j++ {

			summedInput += weightsLayerTwo[j][k] * hiddenLayerOutput[j]
		}
		summedInput += weightsOutputUnitsBias[k]
		outputLayerOutput[k] = sigmoid(summedInput)
	}
}

func sigmoid(input float64) float64 {

	return 1.0 / (1.0 + math.Pow(math.E, -input))
}

func test() {

	var testPOSScore int = 0
	var testNEGScore int = 0
	var trainPOSScore int = 0
	var trainNEGScore int = 0

	var testPOSScorePCT int = 0
	var testNEGScorePCT int = 0
	var trainPOSScorePCT int = 0
	var trainNEGScorePCT int = 0

	testPOSScore = 0
	for p := 0; p < numTestPositives; p++ {
		loadInputs(testExamplesPOS[p])
		runNet()

		if LTU(outputLayerOutput[0]) == 1 {
			testPOSScore++
		}
	}
	testPOSScorePCT = (int)((100.0 * testPOSScore) / numTestPositives)

	testNEGScore = 0
	for p := 0; p < numTestNegatives; p++ {
		loadInputs(testExamplesNEG[p])
		runNet()
		if LTU(outputLayerOutput[0]) == 0 {
			testNEGScore++
		}
	}
	testNEGScorePCT = (int)((100.0 * testNEGScore) / numTestNegatives)

	fmt.Printf("TEST POS:  %3d\n", testPOSScorePCT)
	fmt.Printf("TEST NEG:  %3d\n", testNEGScorePCT)

	trainPOSScore = 0
	for p := 0; p < numTrainPositives; p++ {
		loadInputs(trainExamplesPOS[p])
		runNet()
		if LTU(outputLayerOutput[0]) == 1 {
			trainPOSScore++
		}
	}
	trainPOSScorePCT = (int)((100.0 * trainPOSScore) / numTrainPositives)

	trainNEGScore = 0
	for p := 0; p < numTrainNegatives; p++ {
		loadInputs(trainExamplesNEG[p])
		runNet()
		if LTU(outputLayerOutput[0]) == 0 {
			trainNEGScore++
		}
	}
	trainNEGScorePCT = (int)((100.0 * trainNEGScore) / numTrainNegatives)

	trainScore := trainPOSScore + trainNEGScore
	testScore := testPOSScore + testNEGScore
	POSScore := trainPOSScore + testPOSScore
	NEGScore := trainNEGScore + testNEGScore
	trainScorePCT := (int)((100.0 * trainScore) / (numTrainPositives + numTrainNegatives))
	testScorePCT := (int)((100.0 * testScore) / (numTestPositives + numTestNegatives))
	POSScorePCT := (int)((100.0 * POSScore) / (numTrainPositives + numTestPositives))
	NEGScorePCT := (int)((100.0 * NEGScore) / (numTrainNegatives + numTestNegatives))
	allExamplesScore := trainScore + testScore
	allExamplesScorePCT := (int)((100.0 * allExamplesScore) / (numTrainPositives + numTrainNegatives + numTestPositives + numTestNegatives))

	fmt.Printf("\n\n")
	fmt.Printf("                                RANDOM SEED:   %d\n\n", randomSeed)
	fmt.Printf("                      NUMBER OF INPUT UNITS:   %d\n", numInputUnits)
	fmt.Printf("                     NUMBER OF HIDDEN UNITS:   %d\n", numHiddenUnits)
	fmt.Printf("                     NUMBER OF OUTPUT UNITS:   %d\n", numOutputUnits)
	fmt.Printf("                        TRAINING ITERATIONS:   %d\n", numEpocs)
	fmt.Printf("\n")
	fmt.Printf("                      |   POSITIVES    |    NEGATIVES    |   ALL EXAMPLES   |\n")
	fmt.Printf("                      |-----------------------------------------------------|\n")
	fmt.Printf("    TRAINING SET      |      %3d%%      |       %3d%%      |      %3d%%        |\n", trainPOSScorePCT, trainNEGScorePCT, trainScorePCT)
	fmt.Printf("    TEST     SET      |      %3d%%      |       %3d%%      |      %3d%%        |\n", testPOSScorePCT, testNEGScorePCT, testScorePCT)
	fmt.Printf(" ---------------------|----------------|-----------------|------------------|\n")
	fmt.Printf("    OVER ALL EXAMPLES |      %3d%%      |       %3d%%      |      %3d%%        |\n", POSScorePCT, NEGScorePCT, allExamplesScorePCT)
	fmt.Printf("\n\n")
}

// Linear threshold unit to determine if value is closer to 1 or zero
func LTU(input float64) int {

	if (input) > 0.5 {
		return 1
	} else {
		return 0
	}
}
