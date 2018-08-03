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

//Process struct to hold attributes of a Process
type ProcessInfo struct {
	ID int
   	name string

   	arrivalTime int
   	burstTime int
   	waitTime int

   	turnAroundTime int
   	completionTime int
   	selectionTime int

   	selected bool
   	completed bool
}

//Common variables to hold the information extracted from the input file
var (
	processCount, totalTime, quantum int
	schedulingAlgorithm string

	process []ProcessInfo	
)

//returns the processCount, totalTime, schedulingAlgorithm, and quantum from the input file
func getProcessInfo (filename string) (pc int, tT int, sa string, q int) {

	//create a file pointer
	file, _ := os.Open(filename)

	//make a scan object to iterate through the file using the file pointer
	fileScanner := bufio.NewScanner(file)

	//split the file into words
	fileScanner.Split(bufio.ScanWords)

	
	for fileScanner.Scan() {

		//get the first word as a string and set it to word variable 
		word := fileScanner.Text()

		//if the string encountered is "processcount" then increment the pointer to point to next word,
		// which is the process count value and set it equal to pc variable
		// else if the string encountered is "runfor" then increment the pointer to point to next word,
		// which is the total time to run the algorithm for and set it equal to tT variable
		// else if the string encountered is "use" then increment the pointer to point to next word,
		// which tells what algorithm to use and set it equal to sa variable
		// else if the string encountered is "quantum" then increment the pointer to point to next word,
		// which is the quantum number and set it equal to q variable
		if(word == "processcount") {

			//increment scanner
			fileScanner.Scan()
			//convert the string to integer
		 	pc,_ = strconv.Atoi(fileScanner.Text())

		} else if(word == "runfor") {

        	fileScanner.Scan()
        	tT,_ = strconv.Atoi(fileScanner.Text())

        } else if(word == "use") {

        	fileScanner.Scan()
        	sa = fileScanner.Text()

        	//break the scanning if the algorithm is not round robin since because then it will not have a quantum value
        	if (sa != "rr") {
        		break
        	}

        } else if(word == "quantum"){

        	fileScanner.Scan()

        	//check if the scanned string is an integer to prevent quantum = -
        	temp, err := strconv.Atoi(fileScanner.Text()); 

        	if (err == nil) {
    			q = temp
			}

			//break the scanning after the quantum number is found
			break
        } 

	}

	return pc, tT, sa, q 
}

func getListOfProcesses (filename string, pc int) (p []ProcessInfo) {

	//create a file pointer
	file, _ := os.Open(filename)

	//make a scan object to iterate through the file using the file pointer
	fileScanner := bufio.NewScanner(file)

	//split the file into words
	fileScanner.Split(bufio.ScanWords)

	//allocate memory to p slice and set initial length to 0 and capacity equal to the number of processes
	p = make([]ProcessInfo, 0, pc)

	//struct variable
	var process ProcessInfo
	var count int = 0

	//iterate through the input file to get Process Information and then add them to the p slice
	for fileScanner.Scan() {

		word := fileScanner.Text()

		if(word == "end") {

			break

		} else if(word == "name") {

        	fileScanner.Scan()
        	process.name = fileScanner.Text()
        	process.ID++
        	count++

        } else if(word == "arrival") {

        	fileScanner.Scan()
        	process.arrivalTime,_ = strconv.Atoi(fileScanner.Text())
        	count++
        	
        } else if(word == "burst") {

        	fileScanner.Scan()
        	process.burstTime,_ = strconv.Atoi(fileScanner.Text())
        	count++

        }

        //when all three entries have been added to the struct put it in the slice
        if(count == 3) {
        	p = append(p, process)
        	count = 0
        }
	}

	//return the slice
	return p
}

//selection sort to sort the slices with respect to the 
func sort (process []ProcessInfo, typeOfSort string) (p []ProcessInfo) {

	p = make([]ProcessInfo, len(process))
	p = process

	/*
		 Sort processes with respect to following attributes:
		 AT = Arrival Time
		 BT = Burst Time
		 ID = Process ID
	*/
	for i := 0; i < (len(p) - 1); i++ {

			minIndex := i

			for j := i+1; j < len(p); j++ {

				if(typeOfSort == "AT" && p[j].arrivalTime < p[minIndex].arrivalTime) {

						minIndex = j;

				} else if (typeOfSort == "BT" && p[j].burstTime < p[minIndex].burstTime) {

						minIndex = j;

				} else if (typeOfSort == "ID" && p[j].ID < p[minIndex].ID) {

						minIndex = j;
				}
				
				
			}
			//swap
			temp := p[minIndex]
			p[minIndex] = p[i]
			p[i] = temp
		} 

	return p
}

func fcfs (process []ProcessInfo, processCount int, usefor int, outputFile string)  {
	
	//output stream to print result to the output file
	output, _ := os.Create(outputFile)

	//sort processes with respect to arrival times
	process = sort(process,"AT")

	//Print the number of processes and the algorithm being used
	fmt.Fprintf(output, "%3d processes\n", processCount)
	fmt.Fprintf(output, "Using First-Come First-Served\n")

	var arrivalQueue []ProcessInfo
 	arrivalQueue = make([]ProcessInfo, processCount)
 	arrivalQueueCapacity := 0
		
	time := 0
	index := 0

	for time < usefor {

		//if time = arrival time of a process, then add it to arrival queue and increase queue capacity
		for i:=0; i < processCount; i++ {
			if(process[i].arrivalTime == time) {

				fmt.Fprintf(output, "Time %3d : %s arrived\n", time, process[i].name)
				arrivalQueue[i] = process[i]
				arrivalQueueCapacity++
			}
		}

		//if there is nothing in the arrival queue, then the CPU is idle
		 if (arrivalQueueCapacity == 0) {
		 	fmt.Fprintf(output, "Time %3d : Idle\n", time)
		 }

		 //if there is something in the arrival queue, then select a process and run it
		if(arrivalQueueCapacity > 0 ) {


			//if the selelcted process is completed, then flag it as completed and decerease queue capacity
			if(arrivalQueue[index].selected && ((arrivalQueue[index].selectionTime + arrivalQueue[index].burstTime) == time)) {

				arrivalQueue[index].completed = true
				arrivalQueue[index].completionTime = time

				arrivalQueue[index].selected = false
				arrivalQueueCapacity--;

				fmt.Fprintf(output, "Time %3d : %s finished\n", time, process[index].name)

				//increment the index to move on to the next process in the arrival queue 
				if(index < (processCount-1)) {
					index++
				}
				
			}

			//if none of the processes is selected in the queue then select one and mark it as selected
			if(!arrivalQueue[index].selected && !arrivalQueue[index].completed && arrivalQueueCapacity > 0) {

				arrivalQueue[index].selected = true
				arrivalQueue[index].selectionTime = time

				fmt.Fprintf(output, "Time %3d : %s selected (burst %3d)\n", time, process[index].name, process[index].burstTime)

			} else if (arrivalQueueCapacity == 0) {
				fmt.Fprintf(output, "Time %3d : Idle\n", time)
			}

		} 

		time++
	}

	//set the arrival queue to process to update the wait and turn around times
	process = arrivalQueue

	for i:=0; i<processCount; i++ {

		process[i].turnAroundTime = process[i].completionTime - process[i].arrivalTime
		process[i].waitTime = process[i].turnAroundTime - process[i].burstTime
		
	}

 	//print how long was the system supposed to run for
 	fmt.Fprintf(output, "Finished at time  %d\n\n", usefor)

 	//sort processes with respect to process IDs
 	process = sort(process,"ID")

	//print the wait and turn around times
	for i:=0; i<processCount; i++ {
		fmt.Fprintf(output, "%s wait %3d turnaround %3d\n", process[i].name, process[i].waitTime, process[i].turnAroundTime)
	}
		
}

func sjf (process []ProcessInfo, processCount int, usefor int, outputFile string)  {

	//output stream to print result to the output file
	output, _ := os.Create(outputFile)
	
	//sort processes with respect to arrival times
	process = sort(process,"AT")

	fmt.Fprintf(output, "%3d processes\n", processCount)
	fmt.Fprintf(output, "Using preemptive Shortest Job First\n")

	//create an arrival queue with capacity of processCount and length 0
 	var arrivalQueue []ProcessInfo = make([]ProcessInfo, 0, processCount)
 	
	time := 0
	arrivalQueueCapacity := 0
	
	for time < usefor {	

		//add processes to arrival queue as they arrive
		for i:=0; i < processCount; i++ {
			if(process[i].arrivalTime == time) {

				//print the processes that have arrived along with at what time did they arrive
				fmt.Fprintf(output, "Time %3d : %s arrived\n", time, process[i].name)

				//add the process into arrival queue
				arrivalQueue = append(arrivalQueue, process[i])
				arrivalQueueCapacity++
				 
			}
		}

		//if the arrival queue is empty then the CPU is Idle
		if (arrivalQueueCapacity == 0){
			fmt.Fprintf(output, "Time %3d : Idle\n", time)
		}

		//keep track of what process in the arrival queue are present and being used
		index := 0
		previousProcessID := 0

		if(arrivalQueueCapacity > 0){

			//if the process is completed then move on to the next process in the queue
			for (arrivalQueue[index].completed) {
				index++
			}

			//save previous processes ID if it was not completed so that it can be scheduled after process with shorter burst
			if((index < processCount) &&arrivalQueue[index].selected) {
				previousProcessID = arrivalQueue[index].ID
			}

			//deselect all the processes in the arrival queue
			for i:=0; i < len(arrivalQueue); i++ {
				arrivalQueue[i].selected = false
			}

			//sort the arrival queue based on its burst times
			arrivalQueue = sort(arrivalQueue,"BT")

			//increment the index until an incomplete process is encountered
			for ((index < processCount) && arrivalQueue[index].completed) {
				index++
			}

			//if the process available in the arrival queue is same as the process in previous burst then select that process
			if((index < processCount) && (arrivalQueue[index].ID == previousProcessID)) {
				arrivalQueue[index].selected = true
			}

			//if the selelcted process is completed, then flag it as completed and decerease queue capacity signifying that number of processes to be excuted has decreased
			if((index < processCount) && arrivalQueue[index].selected && (arrivalQueue[index].burstTime == 0) && !arrivalQueue[index].completed) {

				//unselect the process, mark it completed and note its completion time for turn around and wait time calcualtions
				arrivalQueue[index].selected = false
				arrivalQueue[index].completed = true
				arrivalQueue[index].completionTime = time
				arrivalQueueCapacity--;

				fmt.Fprintf(output, "Time %3d : %s finished\n", time, arrivalQueue[index].name)
				
			}

			//if the encountered process in the arrival queue is completed and there are more processes present in the arrival queue then increment the index to point to the next process in the queue
			if ((index < processCount) && arrivalQueue[index].completed && arrivalQueueCapacity > 0) {
				index++
			}

			
			//if none of the processes is selected then select one from the arrival queue and start processes it.
			//Mark the selected processes as selected by changing its boolean value to true
			if((index < processCount) && !arrivalQueue[index].selected && !arrivalQueue[index].completed && arrivalQueueCapacity > 0) {

				arrivalQueue[index].selected = true
				arrivalQueue[index].selectionTime = time

				fmt.Fprintf(output, "Time %3d : %s selected (burst %3d)\n", time, arrivalQueue[index].name, arrivalQueue[index].burstTime)

			} 

			//if the arrival queue is empty then there is nothing to process and therefore print Idle
			if (arrivalQueueCapacity == 0) {
				fmt.Fprintf(output, "Time %3d : Idle\n", time)
			}

		}  

		 time++


		if((index < processCount) && arrivalQueue[index].burstTime > 0){
			arrivalQueue[index].burstTime--
		}

	}

	//calculate the turn around time of the processes
	for i:=0; i<processCount; i++ {
		arrivalQueue[i].turnAroundTime = arrivalQueue[i].completionTime - arrivalQueue[i].arrivalTime
	}

	//sort processes with respect to process IDs
 	arrivalQueue = sort(arrivalQueue,"ID")
 	process = sort(process,"ID")

 	//calculate the wait time of the processes
	for i:=0; i<processCount; i++ {
		arrivalQueue[i].waitTime = arrivalQueue[i].turnAroundTime - process[i].burstTime
	}

 	//print how long was the system supposed to run for
 	fmt.Fprintf(output, "Finished at time  %d\n\n", usefor)

	//print the wait and turn around times
	for i:=0; i<processCount; i++ {
		fmt.Fprintf(output,"%s wait %3d turnaround %3d\n", arrivalQueue[i].name, arrivalQueue[i].waitTime, arrivalQueue[i].turnAroundTime)
	}

}

func rr (process []ProcessInfo, processCount int, usefor int, q int, outputFile string)  {
	
	//output stream to print result to the output file
	output, _ := os.Create(outputFile)

	//sort the processes with respect to their arrival times
	process = sort(process,"AT")

	//print the number of processes being processed, the algorithm being used and the quantu number
	fmt.Fprintf(output,"%3d processes\n", processCount)
	fmt.Fprintf(output,"Using Round-Robin\n")
	fmt.Fprintf(output,"Quantum %3d\n\n", q)

	//create an arrival queue with capacity of processCount and length 0
	//create a completed queue to keep track of which processes have been completed
 	var arrivalQueue []ProcessInfo = make([]ProcessInfo, 0, processCount)
 	var completedQueue []ProcessInfo = make([]ProcessInfo, 0, processCount)
 	var scheduledProcess ProcessInfo
 	
 	
	time := 0
	arrivalQueueCapacity := 0
	index := 0

	for time < usefor {	

		//reset quantum
		quantum := 0

		//add processes to the arrival queue as they arrive
		if((index < processCount) && (process[index].arrivalTime == time)) {

			fmt.Fprintf(output,"Time %3d : %s arrived\n", time, process[index].name)
			arrivalQueue = append(arrivalQueue, process[index])
			arrivalQueueCapacity++

			//increment the index to point to next process in the arrival queue
			index++
		}


		if (len(arrivalQueue) == 0){
			fmt.Fprintf(output, "Time %3d : Idle\n", time)
			time++
		}

		if(len(arrivalQueue) > 0){

			//if arrival queue has elements
			// 1. Put it in scheduled process variable and mark it selected to show it is running
			// 2. Remove it from arrival queue and decrease queue capacity
			// 3. Print the selected process
			// 4. Run it through quantum:- decrement burst time, increment quantum, and increment time. 
			//	  At. each time increment check if a process has arrived
			// 5. If a process has a arrived then add it to the arrival queue
			// 6. Add the selected process back to arrival queue at the end, run the process, and unselect it

			scheduledProcess = arrivalQueue[0]
			scheduledProcess.selected = true

			//remove first element from arrival queue
			if(scheduledProcess.selected) {
				arrivalQueueCapacity--
				if(len(arrivalQueue) > 0) {
					arrivalQueue = arrivalQueue[1:]
				} else {
					arrivalQueue = nil
				}
			}
			
			if(scheduledProcess.selected && !scheduledProcess.completed) {

				// fmt.Printf("Time %3d : %s selected (burst %3d)\n", time, scheduledProcess.name, scheduledProcess.burstTime)
				fmt.Fprintf(output, "Time %3d : %s selected (burst %3d)\n", time, scheduledProcess.name, scheduledProcess.burstTime)

				for quantum < q {

					scheduledProcess.burstTime--
					quantum++
					time++

					if((index < processCount) && (process[index].arrivalTime == time)) {

						//add processes to the arrival queue as they arrive
						fmt.Fprintf(output, "Time %3d : %s arrived\n", time, process[index].name)
						arrivalQueue = append(arrivalQueue, process[index])
						arrivalQueueCapacity++

						//increment the index to point to next process in the queue
						index++
					}

					if(scheduledProcess.burstTime == 0) {

						scheduledProcess.completed = true
						scheduledProcess.selected = false
						scheduledProcess.completionTime = time
						arrivalQueueCapacity--

						fmt.Fprintf(output, "Time %3d : %s finished\n", time, scheduledProcess.name)

						completedQueue = append(completedQueue, scheduledProcess)

						break
					}

					
				}

				//add the process to back of arrival queue after running it upto quantum
				scheduledProcess.selected = false
				arrivalQueue = append(arrivalQueue, scheduledProcess)
				arrivalQueueCapacity++
			}

			if (len(arrivalQueue) == 0){
				
				fmt.Fprintf(output, "Time %3d : Idle\n", time)
				time++

				//add processes to the arrival queue as they arrive
				if((index < processCount) && (process[index].arrivalTime == time)) {

						fmt.Fprintf(output, "Time %3d : %s arrived\n", time, process[index].name)
						arrivalQueue = append(arrivalQueue, process[index])
						arrivalQueueCapacity++

						//increment the index to point to next process in the queue 
						index++

						
					}
			}

		}	

	}


 	//print how long was the system supposed to run for
 	fmt.Fprintf(output, "Finished at time  %d\n\n", usefor)

 	//sort processes with respect to process IDs
 	completedQueue = sort(completedQueue,"ID")
 	process = sort(process,"ID")

 	//calculate the turn around times and wait times of the processes
	for i:=0; i<len(completedQueue); i++ {
		completedQueue[i].turnAroundTime = completedQueue[i].completionTime - completedQueue[i].arrivalTime
		completedQueue[i].waitTime = completedQueue[i].turnAroundTime - process[i].burstTime
	}

 	//print the wait and turn around times
	for i:=0; i<len(completedQueue); i++ {
		fmt.Fprintf(output, "%s wait %3d turnaround %3d\n", completedQueue[i].name, completedQueue[i].waitTime, completedQueue[i].turnAroundTime)

	}

	
}

func main() {

	//get the input file and outfile names from CLI
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	//get process count, time to run the algorithm for, type of algorithm to be used and quantum from the input file
	//make a slice of all the processes and add their respective information from the input file
	processCount, totalTime, schedulingAlgorithm, quantum = getProcessInfo(inputFile)
	process = getListOfProcesses(inputFile, processCount)

	//run the algorithm depending on which one is mentioned in the input file
	if(schedulingAlgorithm == "fcfs") {
		fcfs(process, processCount, totalTime, outputFile)
	} else if (schedulingAlgorithm == "sjf") {
		sjf(process, processCount, totalTime, outputFile)
	} else if (schedulingAlgorithm == "rr") {
		rr(process, processCount, totalTime, quantum, outputFile)
	}

	
   
 }
