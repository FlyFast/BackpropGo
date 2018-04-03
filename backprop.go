package main

import "time"
import "fmt"
import "os"
import "bufio"

var trainExamplesPOS[] string // TODO - eliminate this global
var trainExamplesNEG[] string // TODO - eliminate this global
var testExamplesPOS[]  string // TODO - eliminate this global
var testExamplesNEG[]  string // TODO - eliminate this global


func main() {

	pTime := fmt.Println
	startTime := time.Now()

	fmt.Printf("\n\n")
	fmt.Printf("                           Backprop.go\n\n")

	// TODO - Would it be faster to feed records into train() and test() as we read them
	//	      rather than reading them all in and then processing them?
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
		// TODO - pass trainExamplesPOS into this function instead of global
		trainExamplesPOS = append(trainExamplesPOS, posScanner.Text())
		// fmt.Printf("P %4d: [%s]\n", numPositives, trainExamplesPOS[numPositives])
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
		// TODO - pass trainExamplesNEG into this function instead of global
		trainExamplesNEG = append(trainExamplesNEG, negScanner.Text())
		// fmt.Printf("N %4d: [%s]\n", numNegatives, trainExamplesNEG[numNegatives])
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
		// TODO - pass trainExamplesPOS into this function instead of global
		testExamplesPOS = append(testExamplesPOS, posScanner.Text())
		//fmt.Printf("P %4d: [%s]\n", numPositives, testExamplesPOS[numPositives])
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
		// TODO - pass trainExamplesNEG into this function instead of global
		testExamplesNEG = append(testExamplesNEG, negScanner.Text())
		//fmt.Printf("N %4d: [%s]\n", numNegatives, testExamplesNEG[numNegatives])
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
