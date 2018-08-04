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

	index int
	ID int
	service_completed bool

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

func sort (cylinders []Cylinder) (c []Cylinder) {

	c = make([]Cylinder, 0, len(cylinders))
	c = cylinders
	
	for i := 0; i < (len(c) - 1); i++ {

			minIndex := i

			for j := i+1; j < len(c); j++ {

				if (c[j].ID < c[minIndex].ID) {

						minIndex = j;
				}
				
				
			}
			//swap
			temp := c[minIndex]
			c[minIndex] = c[i]
			c[i] = temp
		} 

	return c
}

func getInitCycIndex (cylinders []Cylinder, initCYL_L int) (c []Cylinder, j int) {

	c = make([]Cylinder, 0, len(cylinders))
	c = cylinders
	
	c[0].index = 0;
	nextIndex := c[0].index
	
	for i := 1; i < len(c) ; i++ {
		nextIndex++
		c[i].index = nextIndex;	

		if(c[i].ID == initCYL_L) {
			j = c[i].index
		}
	} 

	return c, j
}

func remove (listOfCylinderRequests []Cylinder, cylinderToBeDeleted Cylinder) (c []Cylinder) {

	c = make([]Cylinder, 0, len(listOfCylinderRequests))
	c = listOfCylinderRequests

	for i:=0; i < len(c); i++ {

		if(c[i].ID == cylinderToBeDeleted.ID) {
			c = append(c[:i], c[i+1:]...)
			break
		}
	}

	return c
}
//read the input file from stdin and and assign global variables with the appropriate values
func parseInputFile (filename string) (lowerCYL_L int, upperCYL_L int, initCYL_L int, algorithmType_L string, cylinders_L []Cylinder){

	cylinders_L = make([]Cylinder, 0, 20)
	var tempCyclinder Cylinder

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

			tempCyclinder.ID, _ = strconv.Atoi(fileScanner.Text())
			
			cylinders_L = append(cylinders_L, tempCyclinder)
			

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


func scan(lowerCYL_L int, upperCYL_L int, initCYL_L int, cylinders_L []Cylinder) {

	//print initial outputs
	fmt.Printf("Seek algorithm: SCAN\n")
	fmt.Printf("\tLower cylinder: %5d\n", lowerCYL_L)
	fmt.Printf("\tUpper cylinder: %5d\n", upperCYL_L)
	fmt.Printf("\tInit cylinder: %5d\n", initCYL_L)
	fmt.Printf("\tCylinder requests:\n")

	for i:=0; i<len(cylinders_L); i++ {
		fmt.Printf("\t\tCylinder %5d\n", cylinders_L[i].ID)
	}

	//add the initially requested cylinder in the cylinder list
	var currentCylinder Cylinder
	currentCylinder.ID = initCYL_L
	cylinders_L = append(cylinders_L, currentCylinder)

	// sort the cylinders with respect to their Request ID and get index of the initial cylinder in the sorted slice
	cylinders_L, current_index := getInitCycIndex(sort(cylinders_L), initCYL_L)
	
	//remove initially requested cylinder from the slice and reset indices
	cylinders_L = remove(cylinders_L, currentCylinder)
	cylinders_L,_ = getInitCycIndex(sort(cylinders_L), initCYL_L)

	traversalDistance := 0
	start_processing := false
	cylinder_services_completed := 0
	previousRequest := initCYL_L
	totalRequests := len(cylinders_L)


	for i := lowerCYL_L; i < upperCYL_L; i++ {

		//current cylinder after servicing initial cylinder
		currentRequest := cylinders_L[current_index].ID

		//don't start servicing the cylinders until the head is at the initial cylinder
		if(i == initCYL_L) {
			start_processing = true
		}

		if (currentRequest > lowerCYL_L && currentRequest < upperCYL_L) {

			if(start_processing) {

				//if current cylinder has not been serviced then service it and calculate the traversal distance
				if(!cylinders_L[current_index].service_completed && cylinder_services_completed != totalRequests) {

					fmt.Printf("Servicing %d\n", currentRequest)

					traversalDistance += Abs(currentRequest - previousRequest)
					
					previousRequest = currentRequest

					cylinder_services_completed++
					cylinders_L[current_index].service_completed = true
				

					//enter this if statement when end of the cylinder list is reached but not all serivces are completed
					if(currentRequest == cylinders_L[totalRequests - 1].ID && cylinder_services_completed != totalRequests) {

						currentRequest = upperCYL_L

						traversalDistance += Abs(currentRequest - previousRequest)

						previousRequest = currentRequest

						//travel in reverse direction starting before initial cylinder
						for i := initCYL_L-1; i > lowerCYL_L; i-- {

							currentRequest = cylinders_L[current_index].ID

							//service the cylinders not serviced
							if(!cylinders_L[current_index].service_completed && currentRequest > lowerCYL_L) {	

								fmt.Printf("Servicing %d\n", currentRequest)

								traversalDistance += Abs(currentRequest - previousRequest)

								previousRequest = currentRequest
								
								cylinder_services_completed++
								cylinders_L[current_index].service_completed = true

							}

							if((current_index > 0) && cylinder_services_completed != totalRequests) {
								current_index--
							} else if (cylinder_services_completed == totalRequests) {
								break
							} else {
								current_index = 0
							}

						}


					}

					if((current_index != len(cylinders_L) - 1) && cylinder_services_completed != totalRequests) {
						current_index++
					} else if (cylinder_services_completed == totalRequests) {
						break
					} else {
						current_index = 0
					}

				} 

			}

		} else {

			//display error message

		}


	}

	//print traversal time
	fmt.Printf("SCAN traversal count = %d\n", traversalDistance)

}

func c_scan(lowerCYL_L int, upperCYL_L int, initCYL_L int, cylinders_L []Cylinder) {

	//print initial outputs
	fmt.Printf("Seek algorithm: C-SCAN\n")
	fmt.Printf("\tLower cylinder: %5d\n", lowerCYL_L)
	fmt.Printf("\tUpper cylinder: %5d\n", upperCYL_L)
	fmt.Printf("\tInit cylinder: %5d\n", initCYL_L)
	fmt.Printf("\tCylinder requests:\n")

	for i:=0; i<len(cylinders_L); i++ {
		fmt.Printf("\t\tCylinder %5d\n", cylinders_L[i].ID)
	}

	//add the initially requested cylinder in the cylinder list
	var currentCylinder Cylinder
	currentCylinder.ID = initCYL_L
	cylinders_L = append(cylinders_L, currentCylinder)

	// sort the cylinders with respect to their Request ID and get index of the initial cylinder in the sorted slice
	cylinders_L, current_index := getInitCycIndex(sort(cylinders_L), initCYL_L)
	
	//remove initially requested cylinder from the slice and reset indices
	cylinders_L = remove(cylinders_L, currentCylinder)
	cylinders_L,_ = getInitCycIndex(sort(cylinders_L), initCYL_L)

	traversalDistance := 0
	start_processing := false
	cylinder_services_completed := 0
	previousRequest := initCYL_L
	totalRequests := len(cylinders_L)


	for i := lowerCYL_L; i < upperCYL_L; i++ {

		//current cylinder after servicing initial cylinder
		currentRequest := cylinders_L[current_index].ID

		//don't start servicing the cylinders until the head is at the initial cylinder
		if(i == initCYL_L) {
			start_processing = true
		}

		if (currentRequest > lowerCYL_L && currentRequest < upperCYL_L) {

			if(start_processing) {

				//if current cylinder has not been serviced then service it and calculate the traversal distance
				if(!cylinders_L[current_index].service_completed && cylinder_services_completed != totalRequests) {

					fmt.Printf("Servicing %d\n", currentRequest)

					traversalDistance += Abs(currentRequest - previousRequest)
					previousRequest = currentRequest

					cylinder_services_completed++
					cylinders_L[current_index].service_completed = true
				

					//enter this if statement when end of the cylinder list is reached but not all serivces are completed
					if(currentRequest == cylinders_L[totalRequests - 1].ID && cylinder_services_completed != totalRequests) {

						currentRequest = upperCYL_L

						traversalDistance += Abs(currentRequest - previousRequest)
						previousRequest = currentRequest

						//starting from 0
						currentRequest = lowerCYL_L
						traversalDistance += Abs(currentRequest - previousRequest)
						previousRequest = currentRequest

						//travel in reverse direction starting before initial cylinder
						for i := lowerCYL_L; i < upperCYL_L; i++ {

							currentRequest = cylinders_L[current_index].ID

							//service the cylinders not serviced
							if(!cylinders_L[current_index].service_completed && currentRequest > lowerCYL_L) {	

								fmt.Printf("Servicing %d\n", currentRequest)

								traversalDistance += Abs(currentRequest - previousRequest)
								previousRequest = currentRequest
								
								cylinder_services_completed++
								cylinders_L[current_index].service_completed = true

							}

							if((current_index != len(cylinders_L) - 1) && cylinder_services_completed != totalRequests) {
								current_index++
							} else if (cylinder_services_completed == totalRequests) {
								break
							} else {
								current_index = 0
							}

						}


					}

					if((current_index != len(cylinders_L) - 1) && cylinder_services_completed != totalRequests) {
						current_index++
					} else if (cylinder_services_completed == totalRequests) {
						break
					} else {
						current_index = 0
					}

				} 

			}

		} else {

			//display error message

		}


	}

	//print traversal time
	fmt.Printf("C-SCAN traversal count = %d\n", traversalDistance)

}

// func sstf(lowerCYL_L int, upperCYL_L int, initCYL_L int, cylinders_L []Cylinder)  {

// 	//print initial outputs
// 	fmt.Printf("Seek algorithm: SSTF\n")
// 	fmt.Printf("\tLower cylinder: %5d\n", lowerCYL_L)
// 	fmt.Printf("\tUpper cylinder: %5d\n", upperCYL_L)
// 	fmt.Printf("\tInit cylinder: %5d\n", initCYL_L)
// 	fmt.Printf("\tCylinder requests:\n")

// 	//print list of cylinders
// 	for i:=0; i<len(cylinders_L); i++ {
// 		fmt.Printf("\t\tCylinder %5d\n", cylinders_L[i].ID)
// 	}

// 	//add the previous requested cylinder in the cylinder list as well
// 	var currentCylinder Cylinder
// 	currentCylinder.ID = initCYL_L
// 	cylinders_L = append(cylinders_L, currentCylinder)

// 	//sort Cylinders requests in the cylinder list with respect to ID and assign them indices in the slice
// 	cylinders_L = sort(cylinders_L)
// 	cylinders_L = resetIndex(cylinders_L)

// 	fmt.Printf("Cylinder Requests (Sorted): %v\n", cylinders_L)

// 	totalCylinderServiced := 0

// 	var nextCylinderLeft Cylinder
// 	var nextCylinderRight Cylinder
// 	var nextCylinder Cylinder
// 	var previousCylinder Cylinder

// 	for totalCylinderServiced != len(cylinders_L) {

// 		if(cylinders_L[i] == currentCylinder) {

// 			//get the left and right index


// 		}



// 	}



func main() {

	//get the input file and outfile names from CLI
	inputFile := os.Args[1]
	
	lowerCYL, upperCYL, initCYL, requestedAlgorithm, cylinders = parseInputFile(inputFile)

	if(requestedAlgorithm == "fcfs") {
		fcfs(lowerCYL, upperCYL, initCYL, cylinders)
	} else if (requestedAlgorithm == "sstf") {
	 	// sstf(lowerCYL, upperCYL, initCYL, cylinders)
	} else if (requestedAlgorithm == "scan") {
		scan(lowerCYL, upperCYL, initCYL, cylinders)
	} else if (requestedAlgorithm == "c-scan") {
		c_scan(lowerCYL, upperCYL, initCYL, cylinders)
	}
	 // else if (requestedAlgorithm == "look") {
	// 	look(lowerCYL, upperCYL, initCYL, cylinders)
	// } else if (requestedAlgorithm == "c-look") {
	// 	c_look(lowerCYL, upperCYL, initCYL, cylinders)
	// }	
   
 }
