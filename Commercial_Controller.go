package main

/*Modern Approach:
Method 1: RequestElevator (FloorNumber)
This method represents an elevator request on a floor or basement.
Method 2: AssignElevator (RequestedFloor)
This method will be used for the requests made on the first floor.
*/

import (
	"fmt"
	"math"
	"sort"
)

type FloorRequestButton struct {
	ID      int
	Pressed bool
}

type Button struct {
	direction string
	floor     int
	light     string
}

type Elevator struct {
	ID              int
	currentFloor    int
	direction       string
	requestList     []int
	floorList       []int
	door            string
	BufferDirection string
	BufferList      []int
}

type Column struct {
	ID             int
	nbOfElevator   int
	floorCall      int
	requestedFloor int
	elevatorList   []Elevator
	floorList      []int
	buttonList     []Button
}

type Battery struct {
	nbOfColumns   int
	nbOfElevators int
	nbOfFloors    int
	nbOfBasements int
	columnList    []Column
}

// Construction of a FloorRequestButton with its ID
// returning the object FloorRequestButton
func floorRequestButton(ID int) FloorRequestButton {
	m := new(FloorRequestButton)
	m.ID = ID
	m.Pressed = false
	return *m
}

// Construction of a Button USING direction and floor
// returning the object Button
func button(direction string, floor int) Button {
	m := new(Button)
	m.direction = direction
	m.floor = floor
	m.light = "OFF"
	return *m
}

// Construction of an elevator with its ID, the floorCall
// (currentFloor) and the requested Floor (User destination)
// returning the object Elevator
func elevator(ID, floorCall, requestedFloor int) Elevator {
	m := new(Elevator)
	m.ID = ID
	m.currentFloor = floorCall
	for i := floorCall; i <= requestedFloor; i++ {
		m.floorList = append(m.floorList, i)
	}
	return *m
}

/*	Construction of a column USING four parameters:
*	The ID of the column, the floorCall (currentFloor)
*	requested Floor (User destination) and the number
*	of Elevator in the column, returning the object column
 */
func column(ID, floorCall, requestedFloor, nbOfElevator int) Column {
	m := new(Column)
	m.ID = ID
	m.floorCall = floorCall
	m.requestedFloor = requestedFloor
	m.nbOfElevator = nbOfElevator
	m.floorList = append(m.floorList, 1)

	// The for loop is Creating Floor List for each Column
	for i := floorCall; i <= requestedFloor; i++ {
		m.floorList = append(m.floorList, i)
	}
	// Creating Elevator List for each Column in floors's range(floorCall to requestedFloor)
	for i := 1; i <= nbOfElevator; i++ {
		m.elevatorList = append(m.elevatorList, elevator(i, floorCall, requestedFloor))
	}
	// The IF clause is Creating CallButtons for Basements as columnID #1 = the basement
	// and the ELSE clause is creating for the others floors in floors's range(floorCall to requestedFloor)
	if ID == 1 {
		for i := requestedFloor; i <= floorCall; i++ {
			m.buttonList = append(m.buttonList, button("UP", i))
		}
	} else {
		for i := floorCall; i <= requestedFloor; i++ {
			m.buttonList = append(m.buttonList, button("DOWN", i))
		}
	}
	/*
		fmt.Print("ColumnID ", ID, " take these floors: ")
		for i := 0; i < len(m.floorList); i++ {
			fmt.Print("|  ", m.floorList[i])
		}
		fmt.Println("")
	*/
	return *m
}

/* Construction of a battery with four parameters :
*  The number of columns in the bulding(4) , the number
*  Of floors(66), the number of basement(6) and the
*  The number of elevator in each column, that mean 5
*  in this traitement
 */
func battery(nbOfColumns, nbOfFloors, nbOfBasements, nbOfElevators int) Battery {
	m := new(Battery)
	m.nbOfColumns = nbOfColumns
	m.nbOfFloors = nbOfFloors
	m.nbOfBasements = nbOfBasements
	m.nbOfElevators = nbOfElevators
	// BuldingFloor includes all the top floors in the building
	// but excluding the Ground Floor and the basements of the building
	BuldingFloor := nbOfFloors - nbOfBasements
	floorPerColumn := int(math.Floor(float64(BuldingFloor / (nbOfColumns - 1))))
	floorCall := 2
	requestedFloor := floorPerColumn
	currentColumnID := 1
	for currentColumnID <= nbOfColumns {
		// instantiation of the 1st column for the basements excluding the top floor of the bulding from -6 to -1
		if currentColumnID == 1 {
			m.columnList = append(m.columnList, column(1, -nbOfBasements, -1, nbOfElevators))
		} else if currentColumnID < nbOfColumns {

			// instantiation of the other columns of the top bulding excluding the basements
			m.columnList = append(m.columnList, column(currentColumnID, floorCall, requestedFloor, nbOfElevators))
			floorCall = requestedFloor + 1
			requestedFloor = requestedFloor + floorPerColumn
		} else {
			// instantiation of the last columns
			m.columnList = append(m.columnList, column(currentColumnID, floorCall, requestedFloor, nbOfElevators))
		}
		currentColumnID++
	}
	fmt.Println("1- Showing the parameters of constructor Battery")
	for i := 0; i < len(m.columnList); i++ {
		fmt.Print("column ", m.columnList[i].ID, ", number of elevator: ", len(m.columnList[i].elevatorList), ", floorLists: from ", " to ")
		for j := 0; j < len(m.columnList[i].floorList); j++ {
			fmt.Print("| ", m.columnList[i].floorList[j])
		}
		fmt.Println("")
	}

	return *m
}

func (bat Battery) UpdateList(elevator Elevator, List []int, Position int) []int {
	check := true
	for i := 0; i < len(List); i++ {
		stop := List[i]
		if stop == Position {
			check = false
		}
	}
	if check {
		List = append(List, Position)
		sort.Ints(List)
	}
	return List
}

func (bat Battery) calculateGap(elevator Elevator, UserCurrentFloor int, UserDirection string) int {
	if elevator.direction != "IDLE" || len(elevator.requestList) != 0 {
		if elevator.direction == UserDirection {
			if elevator.direction == "UP" && elevator.currentFloor <= UserCurrentFloor {

				return int(math.Abs(float64(elevator.currentFloor - UserCurrentFloor)))
			} else if elevator.direction == "DOWN" && elevator.currentFloor >= UserCurrentFloor {

				return int(math.Abs(float64(elevator.currentFloor - UserCurrentFloor)))
			} else {

				return int(math.Abs(float64(elevator.requestList[len(elevator.requestList)-1]-elevator.currentFloor))) + int(math.Abs(float64(elevator.requestList[len(elevator.requestList)-1]-UserCurrentFloor)))
			}
		} else {

			return int(math.Abs(float64(elevator.requestList[len(elevator.requestList)-1]-elevator.currentFloor))) + int(math.Abs(float64(elevator.requestList[len(elevator.requestList)-1]-UserCurrentFloor)))
		}
	} else {

		return int(math.Abs(float64(elevator.currentFloor - UserCurrentFloor)))
	}
}

/*  Method 1: RequestElevator (FloorNumber)
 *   This method represents an elevator request on a floor or basement.
 *   Example: someone is at 54th floor and requests the 1st floor, so an
 *   elevator should be expected to pick the user up at his currentFloor
 *   and bring him back to the 1st floor.
 */
func (bat Battery) RequestElevator(FloorNumber int) {

	// Here is Finding the best column in columnList
	currentColumn := bat.columnList[0]
	for i := 0; i < len(bat.columnList); i++ {

		if FloorNumber >= bat.columnList[i].floorList[1] && FloorNumber <= bat.columnList[i].floorList[len(bat.columnList[i].floorList)-1] {
			currentColumn = bat.columnList[i]
		}
	}

	// The if clause determine the USER Direction
	var direction string
	if FloorNumber < 1 {
		direction = "UP"
	} else {
		direction = "DOWN"
	}
	fmt.Println("Current columnID ", currentColumn.ID, " which is from floor ", currentColumn.floorList[1], " is goind to floor ", currentColumn.floorList[len(currentColumn.floorList)-1])
	// finding the nearest elevator by comparing destination AND elevator down:
	//      1- the moving elevator which is arriving to the user
	//      2- the IDLE elevator
	//      3- other elevators
	gap := 1000
	var distance int
	/*  Define the best elevator which down and ready to serve in the
	 *   Elevator List and get the best one in the column of elevatorList
	 */
	var FirstElevatorDown []Elevator
	var SecondElevatorDown []Elevator
	var ThirdElevatorDown []Elevator
	BestElevator := currentColumn.elevatorList[0]
	var currentDestination int
	for i := 0; i < len(currentColumn.elevatorList); i++ {
		elevator := currentColumn.elevatorList[i]
		if elevator.direction == "IDLE" {
			currentDestination = 0
		} else {
			currentDestination = elevator.requestList[len(elevator.requestList)-1]
		}
		fmt.Println("ElevatorID : ", elevator.ID, ", Elevator Position = ", elevator.currentFloor, ", Elevator direction = ", elevator.direction, " and current destination is ", currentDestination)
		distance = bat.calculateGap(elevator, FloorNumber, direction)
		if elevator.direction == direction {
			if (direction == "UP" && elevator.currentFloor <= FloorNumber) || (direction == "DOWN" && elevator.currentFloor >= FloorNumber) {
				FirstElevatorDown = append(FirstElevatorDown, elevator)
			}
		} else if elevator.direction == "IDLE" {
			SecondElevatorDown = append(SecondElevatorDown, elevator)
		} else {
			ThirdElevatorDown = append(ThirdElevatorDown, elevator)
		}
	}
	if len(FirstElevatorDown) > 0 {
		for i := 0; i < len(FirstElevatorDown); i++ {
			elevator := FirstElevatorDown[i]
			distance = bat.calculateGap(elevator, FloorNumber, direction)
			if distance <= gap {
				gap = distance
				BestElevator = elevator
			}
		}
	} else if len(SecondElevatorDown) > 0 {
		for i := 0; i < len(SecondElevatorDown); i++ {
			elevator := SecondElevatorDown[i]
			distance = bat.calculateGap(elevator, FloorNumber, direction)
			if distance <= gap {
				gap = distance
				BestElevator = elevator
			}
		}
	} else {
		for i := 0; i < len(ThirdElevatorDown); i++ {
			elevator := ThirdElevatorDown[i]
			distance = bat.calculateGap(elevator, FloorNumber, direction)
			if distance <= gap {
				gap = distance
				BestElevator = elevator
			}
		}
	}
	fmt.Println("The best ElevatorID is ", BestElevator.ID)
	//Updating the RequestList of the selected elevator
	if BestElevator.direction == direction || BestElevator.direction == "IDLE" {
		if BestElevator.direction == "DOWN" && BestElevator.currentFloor >= FloorNumber {
			fmt.Println("Take the column ", currentColumn.ID, " and ElevatorID: ", BestElevator.ID, " which is currently at floor ", BestElevator.currentFloor)
			BestElevator.requestList = bat.UpdateList(BestElevator, BestElevator.requestList, FloorNumber)
			BestElevator.requestList = bat.UpdateList(BestElevator, BestElevator.requestList, 1)
			fmt.Println("Go and Take the column ", currentColumn.ID, ", and the nearest elevator ", BestElevator.ID)
			bat.moveElevator(BestElevator)
		} else if BestElevator.direction == "UP" && BestElevator.currentFloor <= FloorNumber {
			fmt.Println("Take the column ", currentColumn.ID, " and ElevatorID: ", BestElevator.ID, " which is currently at floor ", BestElevator.currentFloor)
			BestElevator.requestList = bat.UpdateList(BestElevator, BestElevator.requestList, FloorNumber)
			BestElevator.requestList = bat.UpdateList(BestElevator, BestElevator.requestList, 1)
			fmt.Println("Go and Take the column ", currentColumn.ID, ", and the nearest elevator ", BestElevator.ID)
			bat.moveElevator(BestElevator)
		} else if BestElevator.direction == "IDLE" {
			fmt.Println("Take the column ", currentColumn.ID, " and ElevatorID: ", BestElevator.ID, " which is currently at floor ", BestElevator.currentFloor)
			BestElevator.requestList = bat.UpdateList(BestElevator, BestElevator.requestList, FloorNumber)
			BestElevator.requestList = bat.UpdateList(BestElevator, BestElevator.requestList, 1)
			fmt.Println("Go and Take the column ", currentColumn.ID, ", and the nearest elevator ", BestElevator.ID)
			bat.moveElevator(BestElevator)
		}
		// Updating the BUFFERLIST  of the selected elevator
	} else {
		fmt.Println("Take the column ", currentColumn.ID, " and ElevatorID: ", BestElevator.ID, " which is currently at floor ", BestElevator.currentFloor)
		BestElevator.BufferList = bat.UpdateList(BestElevator, BestElevator.BufferList, FloorNumber)
		BestElevator.requestList = bat.UpdateList(BestElevator, BestElevator.requestList, 1)
		fmt.Println("Go and Take the column ", currentColumn.ID, ", and the nearest elevator ", BestElevator.ID)
		if FloorNumber > 1 {
			BestElevator.BufferDirection = "DOWN"
			bat.moveElevator(BestElevator)

		} else {
			BestElevator.BufferDirection = "UP"
			bat.moveElevator(BestElevator)
		}
	}
}

/*  Method 2: AssignElevator (RequestedFloor)
 *   This method will be used for the requests made on the first floor.
 *   Example: someone is at 1st floor and requests the 20th floor, so an
 *   elevator should be expected to be sent to the user position
 *   and get him up to the 20th floor.
 */
func (bat Battery) AssignElevator(RequestedFloor int) {

	// Here is Finding the best column in columnList
	currentColumn := bat.columnList[0]
	fmt.Println("2- Finding the best columnID and choose the best on which is goind to the same direction as the User")
	for i := 0; i < len(bat.columnList); i++ {
		fmt.Println("Current columnID ", currentColumn.ID, " which is from floor ", currentColumn.floorList[1], " is goind to floor ", currentColumn.floorList[len(currentColumn.floorList)-1])
		if RequestedFloor >= bat.columnList[i].floorList[1] && RequestedFloor <= bat.columnList[i].floorList[len(bat.columnList[i].floorList)-1] {
			currentColumn = bat.columnList[i]
		}
	}

	fmt.Println()
	fmt.Println("3- Here the column has been choosen. The first one which match to the same direction as the User")
	fmt.Println("the columnID for floor ", RequestedFloor, " is ", currentColumn.ID, "\n")
	gap := 1000
	var distance int
	BestElevator := currentColumn.elevatorList[0]

	var UserDirection string
	if RequestedFloor > 1 {
		UserDirection = "UP"
	} else {
		UserDirection = "DOWN"
	}

	// Finding the best elevator
	fmt.Println("4- Here the method findBestElevator is deployed, findind de best Elevator in column's elevatorList ")
	for i := 0; i < len(currentColumn.elevatorList); i++ {
		elevator := currentColumn.elevatorList[i]

		var currentDestination int
		if elevator.direction == "IDLE" {
			currentDestination = 0
		} else {
			currentDestination = elevator.requestList[len(elevator.requestList)-1]
		}
		fmt.Println("ElevatorID : ", elevator.ID, ", Elevator Position = ", elevator.currentFloor, ", Elevator direction = ", elevator.direction, " and current destination is ", currentDestination)
		distance = bat.calculateGap(elevator, 1, UserDirection)
		if distance < gap {
			gap = distance
			BestElevator = elevator
		}
	}

	fmt.Println()
	fmt.Println("5- The best Column && the nearest Elevator has been found")
	fmt.Println("Go and Take the column ", currentColumn.ID, ", and the nearest elevator ", BestElevator.ID)

	if BestElevator.currentFloor == 1 {
		fmt.Println("The nearest ElevatorID is ", BestElevator.ID)
		BestElevator.requestList = bat.UpdateList(BestElevator, BestElevator.requestList, RequestedFloor)
	} else {
		fmt.Println("The nearest ElevatorID is ", BestElevator.ID)
		BestElevator.requestList = bat.UpdateList(BestElevator, BestElevator.BufferList, RequestedFloor)
		// Setting Buffer direction For Basements
		if RequestedFloor > 1 {
			BestElevator.BufferDirection = "DOWN"
			// For other floors
		} else {
			BestElevator.BufferDirection = "UP"
		}
	}
	bat.moveElevator(BestElevator)
}

// Here's the steps to move the elevator once the user is in the elevator
func (bat Battery) moveElevator(elevator Elevator) {
	i := 0
	for i <= len(elevator.requestList) {
		if elevator.requestList[0] > elevator.currentFloor {
			elevator.direction = "UP"
			fmt.Println()
			fmt.Println("6- The Elevator is moving to pick up the User and go to his destination")
			j := elevator.currentFloor
			for j < elevator.requestList[0] {
				elevator.currentFloor++
				if elevator.currentFloor != 0 {
					fmt.Println("Elevator ", elevator.ID, " is at floor ", elevator.currentFloor)
				}
				if elevator.currentFloor == elevator.floorList[len(elevator.floorList)-1] {
					elevator.direction = "IDLE"
				}
				j++
			}
			elevator.door = "OPEN"
			fmt.Println("Door is open")
			elevator.requestList = elevator.requestList[1:]
		} else {
			elevator.direction = "DOWN"
			fmt.Println()
			fmt.Println("7- The Elevator is moving to pick up the User and go to his destination")
			j := elevator.currentFloor
			for j > elevator.requestList[len(elevator.requestList)-1] {
				elevator.currentFloor--
				if elevator.currentFloor != 0 {
					fmt.Println("Elevator ", elevator.ID, " is at floor ", elevator.currentFloor)
				}
				if elevator.currentFloor == elevator.floorList[0] {
					elevator.direction = "IDLE"
				}
				j--
			}
			elevator.door = "OPEN"
			elevator.requestList = elevator.requestList[:len(elevator.requestList)-1]
		}
		i++
		elevator.door = "CLOSED"
		fmt.Println("Door is closed")
		elevator.direction = "IDLE"
	}
	if len(elevator.BufferList) > 0 {
		elevator.requestList = elevator.BufferList
		elevator.direction = elevator.BufferDirection
		bat.moveElevator(elevator)
	} else {
		elevator.direction = "IDLE"
	}

}

// Scenario 1
func Scenario1() {
	fmt.Println()
	fmt.Println("*******************      Scenario 1     *******************")
	fmt.Println("******************* ******************* *******************")
	fmt.Println()
	fmt.Println("******** User at floor 1. He goes UP to floor 20 ********")
	fmt.Println("*********** Elevator 5 from Column 2 is expected **********")
	fmt.Println("******************* ******************* *******************")
	Battery1 := battery(4, 66, 6, 5)
	// Initializing Elevator 1 of Column 2
	Battery1.columnList[1].elevatorList[0].currentFloor = 20
	Battery1.columnList[1].elevatorList[0].direction = "DOWN"
	Battery1.columnList[1].elevatorList[0].requestList = append(Battery1.columnList[1].elevatorList[0].requestList, 5)
	// Initializing Elevator 2 of Column 2
	Battery1.columnList[1].elevatorList[1].currentFloor = 2
	Battery1.columnList[1].elevatorList[1].direction = "UP"
	Battery1.columnList[1].elevatorList[1].requestList = append(Battery1.columnList[1].elevatorList[1].requestList, 15)
	// // Initializing Elevator 3 of Column 2
	Battery1.columnList[1].elevatorList[2].currentFloor = 13
	Battery1.columnList[1].elevatorList[2].direction = "DOWN"
	Battery1.columnList[1].elevatorList[2].requestList = append(Battery1.columnList[1].elevatorList[2].requestList, 1)
	// // Initializing Elevator 4 of Column 2
	Battery1.columnList[1].elevatorList[3].currentFloor = 15
	Battery1.columnList[1].elevatorList[3].direction = "DOWN"
	Battery1.columnList[1].elevatorList[3].requestList = append(Battery1.columnList[1].elevatorList[3].requestList, 2)
	// // Initializing Elevator 5 of Column 2
	Battery1.columnList[1].elevatorList[4].currentFloor = 6
	Battery1.columnList[1].elevatorList[4].direction = "DOWN"
	Battery1.columnList[1].elevatorList[4].requestList = append(Battery1.columnList[1].elevatorList[4].requestList, 1)

	Battery1.AssignElevator(20)
}

// Scenario 2
func Scenario2() {
	fmt.Println()
	fmt.Println("*******************      Scenario 2     *******************")
	fmt.Println("******************* ******************* *******************")
	fmt.Println()
	fmt.Println("******** User at floor 1. He goes UP to floor 36 ********")
	fmt.Println("*********** Elevator 1 from Column 3 is expected **********")
	fmt.Println("******************* ******************* *******************")
	Battery2 := battery(4, 66, 6, 5)
	// Initializing Elevator 1 of Column 3
	Battery2.columnList[2].elevatorList[0].currentFloor = 1
	Battery2.columnList[2].elevatorList[0].direction = "UP"
	Battery2.columnList[2].elevatorList[0].requestList = append(Battery2.columnList[2].elevatorList[0].requestList, 21)
	// Initializing Elevator 2 of Column 3
	Battery2.columnList[2].elevatorList[1].currentFloor = 23
	Battery2.columnList[2].elevatorList[1].direction = "UP"
	Battery2.columnList[2].elevatorList[1].requestList = append(Battery2.columnList[2].elevatorList[1].requestList, 28)
	// // Initializing Elevator 3 of Column 3
	Battery2.columnList[2].elevatorList[2].currentFloor = 33
	Battery2.columnList[2].elevatorList[2].direction = "DOWN"
	Battery2.columnList[2].elevatorList[2].requestList = append(Battery2.columnList[2].elevatorList[2].requestList, 1)
	// // Initializing Elevator 4 of Column 3
	Battery2.columnList[2].elevatorList[3].currentFloor = 40
	Battery2.columnList[2].elevatorList[3].direction = "DOWN"
	Battery2.columnList[2].elevatorList[3].requestList = append(Battery2.columnList[2].elevatorList[3].requestList, 24)
	// // Initializing Elevator 5 of Column 3
	Battery2.columnList[2].elevatorList[4].currentFloor = 39
	Battery2.columnList[2].elevatorList[4].direction = "DOWN"
	Battery2.columnList[2].elevatorList[4].requestList = append(Battery2.columnList[2].elevatorList[4].requestList, 1)
	Battery2.AssignElevator(36)
}

// Scenario 3
func Scenario3() {
	fmt.Println()
	fmt.Println("*******************      Scenario 3     *******************")
	fmt.Println("******************* ******************* *******************")
	fmt.Println()
	fmt.Println("******** User at floor 54. He goes DOWN to floor 1 ********")
	fmt.Println("*********** Elevator 1 from Column 4 is expected **********")
	fmt.Println("******************* ******************* *******************")
	Battery3 := battery(4, 66, 6, 5)
	// Initializing Elevator 1 of Column 3
	Battery3.columnList[3].elevatorList[0].currentFloor = 58
	Battery3.columnList[3].elevatorList[0].direction = "DOWN"
	Battery3.columnList[3].elevatorList[0].requestList = append(Battery3.columnList[3].elevatorList[0].requestList, 1)
	// Initializing Elevator 2 of Column 3
	Battery3.columnList[3].elevatorList[1].currentFloor = 50
	Battery3.columnList[3].elevatorList[1].direction = "UP"
	Battery3.columnList[3].elevatorList[1].requestList = append(Battery3.columnList[3].elevatorList[1].requestList, 60)
	// // Initializing Elevator 3 of Column 3
	Battery3.columnList[3].elevatorList[2].currentFloor = 46
	Battery3.columnList[3].elevatorList[2].direction = "UP"
	Battery3.columnList[3].elevatorList[2].requestList = append(Battery3.columnList[3].elevatorList[2].requestList, 58)
	// // Initializing Elevator 4 of Column 3
	Battery3.columnList[3].elevatorList[3].currentFloor = 1
	Battery3.columnList[3].elevatorList[3].direction = "UP"
	Battery3.columnList[3].elevatorList[3].requestList = append(Battery3.columnList[3].elevatorList[3].requestList, 54)
	// // Initializing Elevator 5 of Column 3
	Battery3.columnList[3].elevatorList[4].currentFloor = 60
	Battery3.columnList[3].elevatorList[4].direction = "DOWN"
	Battery3.columnList[3].elevatorList[4].requestList = append(Battery3.columnList[3].elevatorList[4].requestList, 1)
	Battery3.RequestElevator(54)
}

// Scenario 4
func Scenario4() {
	fmt.Println()
	fmt.Println("*******************      Scenario 4     *******************")
	fmt.Println("******************* ******************* *******************")
	fmt.Println()
	fmt.Println("********* User at floor -3. He goes UP to floor 1 *********")
	fmt.Println("*********** Elevator 4 from Column 1 is expected **********")
	fmt.Println("******************* ******************* *******************")
	Battery4 := battery(4, 66, 6, 5)
	// Initializing Elevator 1 of Column 4
	Battery4.columnList[0].elevatorList[0].currentFloor = -4
	Battery4.columnList[0].elevatorList[0].direction = "IDLE"
	// Battery4.ColumnList[0].ElevatorList[0].StopList = append(Battery4.ColumnList[0].ElevatorList[0].StopList, 1)
	// Initializing Elevator 2 of Column 4
	Battery4.columnList[0].elevatorList[1].currentFloor = 1
	Battery4.columnList[0].elevatorList[1].direction = "IDLE"
	// Battery4.ColumnList[0].ElevatorList[1].StopList = append(Battery4.ColumnList[0].ElevatorList[1].StopList, );
	// // Initializing Elevator 3 of Column 4
	Battery4.columnList[0].elevatorList[2].currentFloor = -3
	Battery4.columnList[0].elevatorList[2].direction = "DOWN"
	Battery4.columnList[0].elevatorList[2].requestList = append(Battery4.columnList[0].elevatorList[2].requestList, -5)
	// // Initializing Elevator 4 of Column 4
	Battery4.columnList[0].elevatorList[3].currentFloor = -6
	Battery4.columnList[0].elevatorList[3].direction = "UP"
	Battery4.columnList[0].elevatorList[3].requestList = append(Battery4.columnList[0].elevatorList[3].requestList, 1)
	// // Initializing Elevator 5 of Column 4
	Battery4.columnList[0].elevatorList[4].currentFloor = -1
	Battery4.columnList[0].elevatorList[4].direction = "DOWN"
	Battery4.columnList[0].elevatorList[4].requestList = append(Battery4.columnList[0].elevatorList[4].requestList, -6)
	Battery4.RequestElevator(-3)
}
func main() {

	Scenario1()
	Scenario2()
	Scenario3()
	Scenario4()

}
