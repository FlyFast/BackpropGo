package main

import "time"
import "fmt"
import "os"
import "bufio"
import "math/rand"

var trainExamplesPOS[] string // TODO - eliminate this global
var trainExamplesNEG[] string // TODO - eliminate this global
var testExamplesPOS[]  string // TODO - eliminate this global
var testExamplesNEG[]  string // TODO - eliminate this global

// TODO - Can we eliminate these and use counters that are local
//        to the read functions? We should be able to get these
//        values from counts.
var numTrainPositives int = 0
var numTrainNegatives int = 0
var numTestPositives  int = 0
var numTestNegatives  int = 0

const numInputUnits   int = 40
const numHiddenUnits  int = 10
const numOutputUnits  int = 1

var weightsLayerOne[numInputUnits][numHiddenUnits] float64
var weightsLayerTwo[numHiddenUnits][numOutputUnits] float64
var weightsHiddenUnitsBias[numHiddenUnits] float64
var weightsOutputUnitsBias[numOutputUnits] float64
var hiddenLayerOutput[numHiddenUnits] float64
var outputLayerOutput[numOutputUnits] float64

func main() {

	var randomSeed int64 = 0

	pTime := fmt.Println
	startTime := time.Now()

	fmt.Printf("\n\n")
	fmt.Printf("                           Backprop.go\n\n")

	// TODO - Would it be faster to feed records into train() and test() as we read them
	//	      rather than reading them all in and then processing them?
	readTrainingSets()
	readTestingSets()

	//TODO: Make this configurable at runtime
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
		// fmt.Printf("P %4d: [%s]\n", numTrainPositives, trainExamplesPOS[numTrainPositives])
		numTrainPositives++
	}

	fmt.Printf("Num train POS: %d\n", numTrainPositives)


	// Read negative training examples
	negFilePath := "./trainNEG.txt" // TODO - pass as parameter and get from commandline arg

	negFile, _ := os.Open(negFilePath) // TODO- error handling
	defer negFile.Close()
	negScanner := bufio.NewScanner(negFile)
		negScanner.Split(bufio.ScanLines)

	for negScanner.Scan() {
		// TODO - pass trainExamplesNEG into this function instead of global
		trainExamplesNEG = append(trainExamplesNEG, negScanner.Text())
		// fmt.Printf("N %4d: [%s]\n", numTrainNegatives, trainExamplesNEG[numTrainNegatives])
		numTrainNegatives++
	}

	fmt.Printf("Num train NEG: %d\n", numTrainNegatives)


	// return the number of positives and negatives rather than using globals
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
		//fmt.Printf("P %4d: [%s]\n", numTestPositives, testExamplesPOS[numTestPositives])
		numTestPositives++
	}

	fmt.Printf("Num test POS: %d\n", numTestPositives)

	// Read negative testing examples
	negFilePath := "./testNEG.txt" // TODO - pass as parameter and get from commandline arg

	negFile, _ := os.Open(negFilePath) // TODO- error handling
	defer negFile.Close()
	negScanner := bufio.NewScanner(negFile)
		negScanner.Split(bufio.ScanLines)

	for negScanner.Scan() {
		// TODO - pass trainExamplesNEG into this function instead of global
		testExamplesNEG = append(testExamplesNEG, negScanner.Text())
		//fmt.Printf("N %4d: [%s]\n", numTestNegatives, testExamplesNEG[numTestNegatives])
		numTestNegatives++
	}

	fmt.Printf("Num test NEG: %d\n", numTestNegatives)


	// return the number of positives and negatives rather than using globals
}

func reset(seed int64) {
	rand.Seed(seed)
	//var r float64

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

func train() {
	// TODO - START HERE
}


func test() {
}
