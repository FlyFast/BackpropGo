package main

import "time"
import "fmt"
import "os"
import "bufio"

func main() {

	pTime := fmt.Println
	startTime := time.Now()

	fmt.Printf("\n\n")
	fmt.Printf("                           Backprop.go\n")

	readTrainingSets()
	readTestingSets()

	//TODO: Make this configurable at runtime
	for randomSeed := 102; randomSeed < 200; randomSeed += 10 {
	
		reset()
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

	numTrainPositives := 0

	posFile, _ := os.Open(posFilePath) // TODO - error handling
	defer posFile.Close()
	posScanner := bufio.NewScanner(posFile)
		posScanner.Split(bufio.ScanLines)

	for posScanner.Scan() {
		// TODO - add code to add the training case to trainingExamplesPOS
		numTrainPositives++
	}

	fmt.Println("Number of training positives: %d", numTrainPositives)

/*
	// Read negative training examples
	outFilePath := "./trainNEG.txt" // TODO - pass as parameter and get from commandline arg

	numTrainExamples := 0

	inFile, err := os.Open(inFilePath)
	defer inFile.Close()
	inScanner := bufio.NewScanner(inFile)
		inScanner.Split(bufio.ScanLines)

	for inScanner.Scan() {
		// TODO - add code to add the training case to trainingExamplesPOS
		numTrainPositives++
	}

*/


	// return the number of positives and negatives rather than using globals
}

func readTestingSets() {
}

func reset() {
}

func train() {
}

func test() {
}
