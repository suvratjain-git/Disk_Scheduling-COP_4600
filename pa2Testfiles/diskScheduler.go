/*
I Suvrat Jain (su956999) affirm that this program is entirely my own work and that 
I have neither developed my code together with any another person, 
nor copied any code from any other person, nor permitted my code to be copied or otherwise used by any other person, 
nor have I copied, modified, or otherwise used programs created by others. 
I acknowledge that any violation of the above terms will be treated as academic dishonesty.
*/
package main

import (
     "fmt"
     "os"
     "bufio"
     "strconv"
)

//Process struct to hold attributes of a Cylinder
type Cylinder struct {

	ID int
}

//parse the input file and assign the information to the variables
var (
	lowerCYL int
	upperCYL int
	initCYL int
	requestedAlgorithm string
	cylinders []Cylinder
)

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

//read the input file from stdin and and assign global variables with the appropriate values
func parseInputFile (filename string) (lowerCYL_L int, upperCYL_L int, initCYL_L int, algorithmType_L string, cylinders_L []Cylinder){

	cylinders_L = make([]Cylinder, 0, 20)
	var cylinderID Cylinder

	//open the input file and get a pointer to it
	file, _ := os.Open(filename)

	//make a pointer for the text inside the file
	fileScanner := bufio.NewScanner(file)

	//split the file into words
	fileScanner.Split(bufio.ScanWords)

	//iterate through the words in the file
	for fileScanner.Scan() {

		//get the first word as a string and set it to word variable 
		word := fileScanner.Text()

		//break the loop if "end" is encountered in the input file
		if(word == "end"){
			break
		}

		if(word == "cylreq") {

			//increment pointer
			fileScanner.Scan()

			cylinderID.ID, _ = strconv.Atoi(fileScanner.Text())

			cylinders_L = append(cylinders_L, cylinderID)
			
			

		} else if(word == "use") {
			//increment pointer
			fileScanner.Scan()
			//assign the algorithm name to algorithmType variable
			algorithmType_L = fileScanner.Text()

		} else if(word == "lowerCYL") {

			//increment pointer
        	fileScanner.Scan()
        	//assign lower cylinder value to lowerCYL_L variable
        	lowerCYL_L,_ = strconv.Atoi(fileScanner.Text())

        } else if(word == "upperCYL") {

        	//increment pointer
        	fileScanner.Scan()
        	//assign upper cylinder value to upperCYL_L variable
        	upperCYL_L,_ = strconv.Atoi(fileScanner.Text())

        } else if(word == "initCYL"){

        	//increment pointer
        	fileScanner.Scan()
        	//assign initial cylinder value to initCYL_L variable
        	initCYL_L,_ = strconv.Atoi(fileScanner.Text())
        } 

	}

	return lowerCYL_L, upperCYL_L, initCYL_L, algorithmType_L, cylinders_L
}

func fcfs(lowerCYL_L int, upperCYL_L int, initCYL_L int, cylinders_L []Cylinder)  {

	//print initial outputs
	fmt.Printf("Seek algorithm: FCFS\n")
	fmt.Printf("\tLower cylinder: %5d\n", lowerCYL_L)
	fmt.Printf("\tUpper cylinder: %5d\n", upperCYL_L)
	fmt.Printf("\tInit cylinder: %5d\n", initCYL_L)
	fmt.Printf("\tCylinder requests:\n")

	//print list of cylinders
	for i:=0; i<len(cylinders_L); i++ {
		fmt.Printf("\t\tCylinder %5d\n", cylinders_L[i].ID)
	}

	totalRequests  := len(cylinders_L)
	traversalDistance := 0
	previousRequest := initCYL_L

	for i := 0; i < totalRequests; i++ {

		//get the current requested cylinder
		currentRequest := cylinders_L[i].ID

		//if the requested cylinder is in bounds then process it, else generate error and continue
		if ((currentRequest > lowerCYL_L) && (currentRequest < upperCYL_L)) {

			//Display current cylinder under service
			fmt.Printf("Servicing %5d\n", currentRequest)

			//calculate the traversal distance
			traversalDistance += Abs(currentRequest - previousRequest)

			//update the previous request to current request
			previousRequest = currentRequest		

		} else {
			//generate error message 
		}

	}

	//print traversal time
	fmt.Printf("FCFS traversal count = %d\n", traversalDistance)
		
}



func main() {

	//get the input file and outfile names from CLI
	inputFile := os.Args[1]
	
	lowerCYL, upperCYL, initCYL, requestedAlgorithm, cylinders = parseInputFile(inputFile)

	if(requestedAlgorithm == "fcfs") {
		fcfs(lowerCYL, upperCYL, initCYL, cylinders)
	}
	// else if (requestedAlgorithm == "sstf") {
	// 	sstf(lowerCYL, upperCYL, initCYL, cylinders)
	// } else if (requestedAlgorithm == "scan") {
	// 	scan(lowerCYL, upperCYL, initCYL, cylinders)
	// } else if (requestedAlgorithm == "c-scan") {
	// 	c_scan(lowerCYL, upperCYL, initCYL, cylinders)
	// } else if (requestedAlgorithm == "look") {
	// 	look(lowerCYL, upperCYL, initCYL, cylinders)
	// } else if (requestedAlgorithm == "c-look") {
	// 	c_look(lowerCYL, upperCYL, initCYL, cylinders)
	// }	
   
 }
