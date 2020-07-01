using System;
using System.Collections.Generic;
using System.Linq;

/*Modern Approach:
Method 1: RequestElevator (FloorNumber)
This method represents an elevator request on a floor or basement.
Method 2: AssignElevator (RequestedFloor)
This method will be used for the requests made on the first floor.
*/
namespace Commercial_Controller
{
    class Battery{
        public int numberOfColumns;
        public int numberOfElevators;
        public int numberOfFloors;
        public int numberOfBasements;
        public List<Column> columnList;

        /* Construction of a battery with four parameters :
        *  The number of columns in the bulding(4) , the number 
        *  Of floors(66), the number of basement(6) and the 
        *  The number of elevator in each column, that mean 5
        *  in this traitement
        */
        public Battery(int numberOfColumns, int numberOfFloors, int numberOfBasements, int numberOfElevators){
            this.numberOfColumns = numberOfColumns;
            this.numberOfFloors = numberOfFloors;
            this.numberOfBasements = numberOfBasements;
            this.numberOfElevators = numberOfElevators;
            // BuldingFloor includes all the top floors in the building
            // but excluding the Ground Floor and the basements of the building
            double BuldingFloor = numberOfFloors - numberOfBasements;
            int floorPerColumn = (int)Math.Floor(BuldingFloor / (numberOfColumns - 1));
            int floorCall = 2; 
            int requestedFloor = floorPerColumn;
            int currentColumnID = 1;
            this.columnList = new List<Column>();

            while (currentColumnID <= numberOfColumns)
            {   // instantiation of the 1st column for the basements excluding the top floor of the bulding from -6 to -1 
                if(currentColumnID == 1){
                    Column column = new Column(1, -numberOfBasements, -1, numberOfElevators);
                    this.columnList.Add(column);
                }
                // instantiation of the other columns of the top bulding excluding the basements
                else if (currentColumnID < numberOfColumns){
                    Column column = new Column (currentColumnID, floorCall, requestedFloor, numberOfElevators);
                    this.columnList.Add(column);
                    floorCall = requestedFloor + 1;
                    requestedFloor = requestedFloor + floorPerColumn;
                }
                else{
                    Column column = new Column(currentColumnID, floorCall, (int)BuldingFloor, numberOfElevators);
                    this.columnList.Add(column);
                }
                currentColumnID ++;
            }
            Console.WriteLine("1- Showing the parameters of constructor Battery");
            foreach (var column in columnList)
            {   
                Console.WriteLine("column " + column.ID + ", number of elevator: " + column.elevatorList.Count + ", floorLists: from " + column.floorList[0] + " to " + column.floorList.Last());
            }
              Console.WriteLine();
        }

        public void UpdateList (Elevator elevator, List<int> List, int currentFloor)
        {
            bool check = true;
            foreach(int stop in List)
            {
                if (stop == currentFloor)
                {
                    check = false;
                }
            }
            if (check)
            {
                List.Add(currentFloor);
                List.Sort();
            }
        }


        /*  Method 1: RequestElevator (FloorNumber)
        *   This method represents an elevator request on a floor or basement.
        *   Example: someone is at 54th floor and requests the 1st floor, so an
        *   elevator should be expected to pick the user up at his currentFloor
        *   and bring him back to the 1st floor.
        */
        public void RequestElevator(int FloorNumber){
            Column currentColumn = this.columnList[0];

            // Here is Finding the best column in columnList
            foreach (var column in this.columnList){
                if (FloorNumber >= column.floorList[1] && FloorNumber <= column.floorList.Last()){
                   currentColumn = column;
                   break; 
                }
            }
            
            // The if clause determine the USER Direction  
            string direction;

            if (FloorNumber < 1){
                direction = "UP";
            }
            else{
                direction = "DOWN";
            }
            Console.WriteLine("Current columnID " + currentColumn.ID + " which is from floor " + currentColumn.floorList[1] + " is goind to floor " + currentColumn.floorList.Last());
            // finding the nearest elevator by comparing destination AND elevator down: 
            //      1- the moving elevator which is arriving to the user
            //      2- the IDLE elevator
            //      3- other elevators
            int destination = 1000;
            int distance;
            /*  Define the best elevator which down and ready to serve in the 
            *   Elevator List and get the best one in the column of elevatorList
            */
            List<Elevator> FirstElevatorDown = new List<Elevator>();
            List<Elevator> SecondElevatorDown = new List<Elevator>();
            List<Elevator> ThirdElevatorDown = new List<Elevator>();
            Elevator BestElevator = currentColumn.elevatorList[0];
            foreach(Elevator elevator in currentColumn.elevatorList){
                int currentDestination;
                if (elevator.direction == "IDLE"){
                    currentDestination = 0;
                }
                else{
                    currentDestination = elevator.requestList.Last();
                }
                Console.WriteLine("ElevatorID : " + elevator.ID + ", Elevator Position = " + elevator.currentFloor + ", Elevator direction = " + elevator.direction + " and current destination is " + currentDestination);
                if (elevator.direction == direction){
                    if ((direction == "UP" && elevator.currentFloor <= FloorNumber)|(direction == "DOWN" && elevator.currentFloor >= FloorNumber)){
                        FirstElevatorDown.Add(elevator);
                    }
                }
                else if(elevator.direction =="IDLE"){
                    SecondElevatorDown.Add(elevator);
                }
                else{
                    ThirdElevatorDown.Add(elevator);
                }
            }
            if(FirstElevatorDown.Count > 0){
                foreach (Elevator elevator in FirstElevatorDown)
                {
                     distance = 1;
                           
                    if ( distance <= destination)
                    {
                        destination = distance;
                        BestElevator = elevator;
                    }
                }
            }
            else if (SecondElevatorDown.Count > 0){
                foreach (Elevator elevator in SecondElevatorDown){
                    distance = 1;
                         
                    if ( distance <= destination)
                    {
                        destination = distance;
                        BestElevator = elevator;
                    }
                }
            }
            else{
                foreach (Elevator elevator in ThirdElevatorDown){
                    distance = 1;
       
                    if ( distance <= destination)
                    {
                        destination = distance;
                        BestElevator = elevator;
                    }
                }
            }   
                Console.WriteLine("The best ElevatorID is " + BestElevator.ID);
            //  Updating the RequestList of the selected elevator
            if (BestElevator.direction == direction && BestElevator.direction == "IDLE"){
                if (BestElevator.direction == "DOWN" && BestElevator.currentFloor >= FloorNumber){
                    Console.WriteLine("Take the column " +  currentColumn.ID + " and ElevatorID: " + BestElevator.ID +  " which is currently at floor " + BestElevator.currentFloor);
                    UpdateList(BestElevator, BestElevator.requestList, FloorNumber);
                    UpdateList(BestElevator, BestElevator.requestList, 1);
                    Console.WriteLine("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);
                    moveElevator(BestElevator);
                }
                else if (BestElevator.direction == "UP" && BestElevator.currentFloor <= FloorNumber){
                    Console.WriteLine("Take the column " +  currentColumn.ID + " and ElevatorID: " + BestElevator.ID +  " which is currently at floor " + BestElevator.currentFloor);
                    UpdateList(BestElevator, BestElevator.requestList, FloorNumber);
                    UpdateList(BestElevator, BestElevator.requestList, 1);
                    Console.WriteLine("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);
                    moveElevator(BestElevator);
                }
                else if (BestElevator.direction == "IDLE"){
                    Console.WriteLine("Take the column " +  currentColumn.ID + " and ElevatorID: " + BestElevator.ID +  " which is currently at floor" + BestElevator.currentFloor);
                    UpdateList(BestElevator, BestElevator.requestList, FloorNumber);
                    UpdateList(BestElevator, BestElevator.requestList, 1);
                    Console.WriteLine("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);
                    moveElevator(BestElevator);
                }
                else{
                    Console.WriteLine("Take the column " +  currentColumn.ID + " and ElevatorID: " + BestElevator.ID +  " which is currently at floor" + BestElevator.currentFloor);
                    UpdateList(BestElevator, BestElevator.BufferList, FloorNumber);
                    UpdateList(BestElevator, BestElevator.BufferList, 1);
                    Console.WriteLine("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);
                 
                    if (FloorNumber > 1){
                        BestElevator.BufferDirection = "DOWN";
                        moveElevator(BestElevator);
                    }
                    else{
                        BestElevator.BufferDirection = "UP";
                        moveElevator(BestElevator);
                    }
                }
            } 
            else{
                Console.WriteLine("Take the column " +  currentColumn.ID + " and ElevatorID: " + BestElevator.ID +  " which is currently at floor" + BestElevator.currentFloor);
                UpdateList(BestElevator, BestElevator.BufferList, FloorNumber);
                UpdateList(BestElevator, BestElevator.requestList, 1);
                Console.WriteLine("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);

                // For Basements
                if (FloorNumber > 1){
                    BestElevator.BufferDirection = "DOWN";
                    moveElevator(BestElevator);
                }
                // For other floors
                else{
                    BestElevator.BufferDirection = "UP";
                    moveElevator(BestElevator);
                }
            }    
        }

        /*  Method 2: AssignElevator (RequestedFloor)
        *   This method will be used for the requests made on the first floor. 
        *   Example: someone is at 1st floor and requests the 20th floor, so an
        *   elevator should be expected to be sent to the user position 
        *   and get him up to the 20th floor.
        */
        public void AssignElevator(int RequestedFloor)
        {
            // Here is Finding the best column in columnList
            Column currentColumn = this.columnList[0];
            Console.WriteLine("2- Finding the best columnID and choose the best on which is goind to the same direction as the User");
            foreach(var column in this.columnList)
            {   
                Console.WriteLine("Current columnID " + column.ID + " which is from floor " + column.floorList[1] + " is goind to floor " + column.floorList.Last());
                if ( RequestedFloor >= column.floorList[1] && RequestedFloor <= column.floorList.Last()) 
                {
                    currentColumn = column;
                    break;
                }
            }
            Console.WriteLine();
            Console.WriteLine("3- Here the column has been choosen. The first one which match to the same direction as the User");
            Console.WriteLine("the columnID for floor " + RequestedFloor + " is " + currentColumn.ID + "\n"); 
            int destination = 1000;
            Elevator BestElevator = currentColumn.elevatorList[0];

            // Finding the best elevator 
            Console.WriteLine("4- Here the method findBestElevator is deployed, findind de best Elevator in column's elevatorList ");
            foreach(Elevator elevator in currentColumn.elevatorList)
            {    int distance = 1;
                    int currentDestination;
                    if (elevator.direction == "IDLE")
                    {
                        currentDestination = 0;
                    }
                    else
                    {
                        currentDestination = elevator.requestList.Last();
                    }
                    Console.WriteLine("ElevatorID : " + elevator.ID + ", Elevator Position = " + elevator.currentFloor + ", Elevator direction = " + elevator.direction + " and current destination is " + currentDestination);

                            
                    if ( distance <= destination)
                    {
                        destination = distance;
                        BestElevator = elevator;
                    }
            }
            
            Console.WriteLine();
            Console.WriteLine("5- The best Column && the nearest Elevator has been found");
            Console.WriteLine("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);

            if (BestElevator.currentFloor == 1)
            {                   
                    Console.WriteLine("The nearest ElevatorID is " + BestElevator.ID);
                    UpdateList(BestElevator, BestElevator.requestList, RequestedFloor);    
            }
            else
            {   
                    Console.WriteLine("The nearest ElevatorID is " + BestElevator.ID);
                    UpdateList(BestElevator, BestElevator.BufferList, RequestedFloor);
                    // Setting Buffer direction For Basements
                    if (RequestedFloor >= 1)
                    {
                        BestElevator.BufferDirection = "DOWN";
                    }
                    // For other floors
                    else
                    {
                        BestElevator.BufferDirection = "UP";
                    }
            }
            moveElevator(BestElevator);
        }
        
        public void moveElevator(Elevator elevator)
        {
            while (elevator.requestList.Count > 0)
            {
                if (elevator.requestList[0] > elevator.currentFloor)
                {
                    elevator.direction = "UP";
                    Console.WriteLine();
                    Console.WriteLine("6- The Elevator is moving to pick up the User and go to his destination");
                    while (elevator.currentFloor < elevator.requestList[0])
                    {
                        elevator.currentFloor += 1;
                        if (elevator.currentFloor != 0)
                        {
                            Console.WriteLine("Elevator " + elevator.ID + " is at floor " + elevator.currentFloor);
                        }
                        if (elevator.currentFloor == elevator.floorList.Last())
                        {
                            elevator.direction = "IDLE";
                        }
                    }
                    elevator.door = "OPEN";
                    Console.WriteLine("Door is opened");
                    elevator.requestList.RemoveAt(0);
                }
                else 
                {
                    elevator.direction = "DOWN";
                    Console.WriteLine();
                    Console.WriteLine("7- The Elevator is moving to pick up the User and go to his destination");
                    while (elevator.currentFloor > elevator.requestList.Last())
                    {
                        elevator.currentFloor -= 1;
                        if (elevator.currentFloor != 0)
                        {
                            Console.WriteLine("Elevator " + elevator.ID + " is at floor " + elevator.currentFloor);
                        }
                        if (elevator.currentFloor == elevator.floorList.First())
                        {
                            elevator.direction = "IDLE";
                        }
                    }
                    elevator.door = "OPEN";
                    Console.WriteLine("Door is opened");
                    elevator.requestList.RemoveAt(elevator.requestList.Count -1);
                }                
                elevator.door = "CLOSED";
                Console.WriteLine("Door is closed");
                elevator.direction = "IDLE";
            }
            if (elevator.BufferList.Count > 0)
            {
                elevator.requestList = elevator.BufferList;
                elevator.direction = elevator.BufferDirection;
                moveElevator(elevator);
            }
            else 
            {
                elevator.direction = "IDLE";
            }
        }
    }

}