package main

import "time"
import "fmt"
import "os"
import "bufio"

func main() {

	pTime := fmt.Println
	startTime := time.Now()

	fmt.Printf("\n\n")
	fmt.Printf("                           Backprop.go\n\n")

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

	numPositives := 0

	posFile, _ := os.Open(posFilePath) // TODO - error handling
	defer posFile.Close()
	posScanner := bufio.NewScanner(posFile)
		posScanner.Split(bufio.ScanLines)

	for posScanner.Scan() {
		// TODO - add code to add the training case to trainingExamplesPOS
		numPositives++
	}

	fmt.Printf("Num train POS: %d\n", numPositives)


	// Read negative training examples
	negFilePath := "./trainNEG.txt" // TODO - pass as parameter and get from commandline arg

	numNegatives := 0

	negFile, _ := os.Open(negFilePath) // TODO- error handling
	defer negFile.Close()
	negScanner := bufio.NewScanner(negFile)
		negScanner.Split(bufio.ScanLines)

	for negScanner.Scan() {
		// TODO - add code to add the training case to trainingExamplesNEG
		numNegatives++
	}

	fmt.Printf("Num train NEG: %d\n", numNegatives)


	// return the number of positives and negatives rather than using globals
}

func readTestingSets() {

	//TODO - Make the file reading generic and reuse the code for both positive and
	//			negatives by passing parameters.

	// Read positive testing examples
	posFilePath := "./testPOS.txt" // TODO - pass as parameter and get from commandline arg

	numPositives := 0

	posFile, _ := os.Open(posFilePath) // TODO - error handling
	defer posFile.Close()
	posScanner := bufio.NewScanner(posFile)
		posScanner.Split(bufio.ScanLines)

	for posScanner.Scan() {
		// TODO - add code to add the training case to trainingExamplesPOS
		numPositives++
	}

	fmt.Printf("Num test POS: %d\n", numPositives)


	// Read negative testing examples
	negFilePath := "./testNEG.txt" // TODO - pass as parameter and get from commandline arg

	numNegatives := 0

	negFile, _ := os.Open(negFilePath) // TODO- error handling
	defer negFile.Close()
	negScanner := bufio.NewScanner(negFile)
		negScanner.Split(bufio.ScanLines)

	for negScanner.Scan() {
		// TODO - add code to add the training case to trainingExamplesNEG
		numNegatives++
	}

	fmt.Printf("Num test NEG: %d\n", numNegatives)


	// return the number of positives and negatives rather than using globals
}

func reset() {
}

func train() {
}

func test() {
}
