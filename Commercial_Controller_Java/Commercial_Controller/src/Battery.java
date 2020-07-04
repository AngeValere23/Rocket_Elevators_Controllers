import java.util.ArrayList;
import java.util.List;
import java.util.Collections;

/*Modern Approach:
Method 1: RequestElevator (FloorNumber)
This method represents an elevator request on a floor or basement.
Method 2: AssignElevator (RequestedFloor)
This method will be used for the requests made on the first floor.
*/
public class Battery {
    
    public int numberOfColumns;
    public int numberOfElevators;
    public int numberOfFloors;
    public int numberOfBasements;
    public List<Column> columnList;

    /** Construction of a battery with four parameters :
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
        int floorPerColumn = (int)Math.floor(BuldingFloor / (numberOfColumns - 1));
        int floorCall = 2; 
        int requestedFloor = floorPerColumn;
        int currentColumnID = 1;
        this.columnList = new ArrayList<Column>();

        while (currentColumnID <= numberOfColumns) {
            // instantiation of the 1st column for the basements excluding the top floor of the bulding from -6 to -1 
            if(currentColumnID == 1){
                Column column = new Column(1, -numberOfBasements, -1, numberOfElevators);
                this.columnList.add(column);
            }
            // instantiation of the other columns of the top bulding excluding the basements
            else if (currentColumnID < numberOfColumns){
                Column column = new Column (currentColumnID, floorCall, requestedFloor, numberOfElevators);
                this.columnList.add(column);
                floorCall = requestedFloor + 1;
                requestedFloor = requestedFloor + floorPerColumn;
            }
            else{
                Column column = new Column(currentColumnID, floorCall, (int)BuldingFloor, numberOfElevators);
                this.columnList.add(column);
            }
                currentColumnID ++;
        }
        System.out.println("1- Showing the parameters of constructor Battery");
        for (Column column : columnList) {
            System.out.println("column " + column.ID + ", number of elevator: " + column.elevatorList.stream().count() + ", floorLists: from " + column.floorList.get(0) + " to " + column.floorList.get(column.floorList.size()-1));
        }
        System.out.println();
    }

    //Update the list of the elevator in each use and after sort them
    public void UpdateList (Elevator elevator, List<Integer> List, int currentFloor)
    {
        Boolean check = true;
        for (Integer stop : List) {
            if (stop == currentFloor)
            {
                check = false;
            }
        }
        if (check)
        {
            List.add(currentFloor);
            Collections.sort(List);
        }
    }

    // Calculate the gap and find the nearest elevator and the  lowest distance 
    // between the Elevator current floor and the user current floor
    public int calculateGap (Elevator elevator, int UserCurrentFloor, String UserDirection){

        if (elevator.direction != "IDLE" | elevator.requestList.stream().count() != 0){
            if (elevator.direction == UserDirection){               
                if (elevator.direction == "UP" && elevator.currentFloor <= UserCurrentFloor){
                    return Math.abs(elevator.currentFloor - UserCurrentFloor);
                }
                else if (elevator.direction == "DOWN" && elevator.currentFloor >= UserCurrentFloor){
                    return Math.abs(elevator.currentFloor - UserCurrentFloor);
                }
                else{
                    return Math.abs(elevator.requestList.get(elevator.requestList.size()-1) - elevator.currentFloor) + Math.abs(elevator.requestList.get(elevator.requestList.size()-1) - UserCurrentFloor);    
                }
            }
            else{
                    return Math.abs(elevator.requestList.get(elevator.requestList.size()-1) - elevator.currentFloor) + Math.abs(elevator.requestList.get(elevator.requestList.size()-1) - UserCurrentFloor);  
            }
        }
        else {
                return Math.abs(elevator.currentFloor - UserCurrentFloor);
        }
    }

    /** Method 1: RequestElevator (FloorNumber)
    *   This method represents an elevator request on a floor or basement.
    *   Example: someone is at 54th floor and requests the 1st floor, so an
    *   elevator should be expected to pick the user up at his currentFloor
    *   and bring him back to the 1st floor.
    */

    public void RequestElevator(int FloorNumber){
        Column currentColumn = this.columnList.get(0);

        // Here is Finding the best column in columnList
        for (Column column : columnList) {
            if (FloorNumber >= column.floorList.get(1) && FloorNumber <= column.floorList.get(column.floorList.size()-1)){
                currentColumn = column;
                break; 
             }
        }
        
        // The if clause determine the USER Direction  
        String direction;

        if (FloorNumber < 1){
            direction = "UP";
        }
        else{
            direction = "DOWN";
        }
        System.out.println("Current columnID " + currentColumn.ID + " which is from floor " + currentColumn.floorList.get(1) + " is goind to floor " + currentColumn.floorList.get(currentColumn.floorList.size()-1));
        // finding the nearest elevator by comparing gap AND elevator down: 
        //      1- the moving elevator which is arriving to the user
        //      2- the IDLE elevator
        //      3- other elevators
        int gap = 1000;
        int distance;
        /*  Define the best elevator which down and ready to serve in the 
        *   Elevator List and get the best one in the column of elevatorList
        */
        List<Elevator> FirstElevatorDown = new ArrayList<Elevator>();
        List<Elevator> SecondElevatorDown = new ArrayList<Elevator>();
        List<Elevator> ThirdElevatorDown = new ArrayList<Elevator>();
        Elevator BestElevator = currentColumn.elevatorList.get(0);
        for (Elevator elevator : currentColumn.elevatorList) {
            int currentDestination;
            if (elevator.direction == "IDLE"){
                currentDestination = 0;
            }
            else{
                currentDestination = elevator.requestList.get(elevator.requestList.size()-1) ;
            }
            System.out.println("ElevatorID : " + elevator.ID + ", Elevator Position = " + elevator.currentFloor + ", Elevator direction = " + elevator.direction + " and current destination is " + currentDestination);
            distance = calculateGap(elevator, FloorNumber, direction);
            if (elevator.direction == direction){
                if ((direction == "UP" && elevator.currentFloor <= FloorNumber)|(direction == "DOWN" && elevator.currentFloor >= FloorNumber)){
                    FirstElevatorDown.add(elevator);
                }
            }
            else if(elevator.direction =="IDLE"){
                SecondElevatorDown.add(elevator);
            }
            else{
                ThirdElevatorDown.add(elevator);
            }
        }
        if(FirstElevatorDown.stream().count() > 0){
            for (Elevator elevator : FirstElevatorDown) {
                 distance = calculateGap(elevator, FloorNumber, direction);
                       
                if ( distance <= gap)
                {
                    gap = distance;
                    BestElevator = elevator;
                }
            }
        }
        else if (SecondElevatorDown.stream().count() > 0){
            for (Elevator elevator : SecondElevatorDown) {
                distance = calculateGap(elevator, FloorNumber, direction);
                     
                if ( distance <= gap)
                {
                    gap = distance;
                    BestElevator = elevator;
                }
            }
        }
        else{
            for (Elevator elevator : ThirdElevatorDown){
                 distance = calculateGap(elevator, FloorNumber, direction);
   
                if ( distance <= gap)
                {
                    gap = distance;
                    BestElevator = elevator;
                }
            }
        }   

        System.out.println("The best ElevatorID is " + BestElevator.ID);
        //  Updating the RequestList of the selected elevator
        if (BestElevator.direction == direction && BestElevator.direction == "IDLE"){
            if (BestElevator.direction == "DOWN" && BestElevator.currentFloor >= FloorNumber){
                System.out.println("Take the column " +  currentColumn.ID + " and ElevatorID: " + BestElevator.ID +  " which is currently at floor " + BestElevator.currentFloor);
                UpdateList(BestElevator, BestElevator.requestList, FloorNumber);
                UpdateList(BestElevator, BestElevator.requestList, 1);
                System.out.println("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);
                moveElevator(BestElevator);
            }
            else if (BestElevator.direction == "UP" && BestElevator.currentFloor <= FloorNumber){
                System.out.println("Take the column " +  currentColumn.ID + " and ElevatorID: " + BestElevator.ID +  " which is currently at floor " + BestElevator.currentFloor);
                UpdateList(BestElevator, BestElevator.requestList, FloorNumber);
                UpdateList(BestElevator, BestElevator.requestList, 1);
                System.out.println("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);
                moveElevator(BestElevator);
            }
            else if (BestElevator.direction == "IDLE"){
                System.out.println("Take the column " +  currentColumn.ID + " and ElevatorID: " + BestElevator.ID +  " which is currently at floor" + BestElevator.currentFloor);
                UpdateList(BestElevator, BestElevator.requestList, FloorNumber);
                UpdateList(BestElevator, BestElevator.requestList, 1);
                System.out.println("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);
                moveElevator(BestElevator);
            }
        } 
        else{
            System.out.println("Take the column " +  currentColumn.ID + " and ElevatorID: " + BestElevator.ID +  " which is currently at floor" + BestElevator.currentFloor);
            UpdateList(BestElevator, BestElevator.BufferList, FloorNumber);
            UpdateList(BestElevator, BestElevator.requestList, 1);
            System.out.println("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);

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
        Column currentColumn = this.columnList.get(0);
        System.out.println("2- Finding the best columnID and choose the best on which is goind to the same direction as the User");
        for(Column column : columnList)
        {   
            System.out.println("Current columnID " + column.ID + " which is from floor " + column.floorList.get(1) + " is goind to floor " + column.floorList.get(currentColumn.floorList.size()-1));
            if ( RequestedFloor >= column.floorList.get(1) && RequestedFloor <= column.floorList.get(currentColumn.floorList.size()-1)) 
            {
                currentColumn = column;
                break;
            }
        }
        System.out.println();
        System.out.println("3- Here the column has been choosen. The first one which match to the same direction as the User");
        System.out.println("the columnID for floor " + RequestedFloor + " is " + currentColumn.ID + "\n"); 
        int gap = 1000;
        int distance;
        Elevator BestElevator = currentColumn.elevatorList.get(0);

        String UserDirection;
       if (RequestedFloor > 1)
       {
           UserDirection = "UP";
       }
       else
       {
           UserDirection = "DOWN";
       }

        // Finding the best elevator 
        System.out.println("4- Here the method findBestElevator is deployed, findind de best Elevator in column's elevatorList ");
        for(Elevator elevator : currentColumn.elevatorList)
        {    
            
                int currentDestination;
                if (elevator.direction == "IDLE")
                {
                    currentDestination = 0;
                }
                else
                {
                    currentDestination = elevator.requestList.get(elevator.requestList.size()-1);
                }
                System.out.println("ElevatorID : " + elevator.ID + ", Elevator Position = " + elevator.currentFloor + ", Elevator direction = " + elevator.direction + " and current destination is " + currentDestination);
            
                distance = calculateGap(elevator, 1, UserDirection);      
                if ( distance <= gap)
                {
                    gap = distance;
                    BestElevator = elevator;
                }
        }
        
        System.out.println();
        System.out.println("5- The best Column && the nearest Elevator has been found");
        System.out.println("Go and Take the column " + currentColumn.ID + ", and the nearest elevator " + BestElevator.ID);

        if (BestElevator.currentFloor == 1)
        {                   
                System.out.println("The nearest ElevatorID is " + BestElevator.ID);
                UpdateList(BestElevator, BestElevator.requestList, RequestedFloor);    
        }
        else
        {   
                System.out.println("The nearest ElevatorID is " + BestElevator.ID);
                UpdateList(BestElevator, BestElevator.BufferList, RequestedFloor);
                // Setting Buffer direction For Basements
                if (RequestedFloor > 1)
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
    
    // Here's the steps to move the elevator once the user is in the elevator
    public void moveElevator(Elevator elevator)
    {
        while (elevator.requestList.stream().count() > 0)
        {
            if (elevator.requestList.get(0) > elevator.currentFloor)
            {
                elevator.direction = "UP";
                System.out.println();
                System.out.println("6- The Elevator is moving to pick up the User and go to his destination");
                while (elevator.currentFloor < elevator.requestList.get(0))
                {
                    elevator.currentFloor ++;
                    if (elevator.currentFloor != 0)
                    {
                        System.out.println("Elevator " + elevator.ID + " is at floor " + elevator.currentFloor);
                    }
                    if (elevator.currentFloor == elevator.floorList.get(elevator.floorList.size()-1))
                    {
                        elevator.direction = "IDLE";
                    }
                }
                elevator.door = "OPEN";
                System.out.println("Door is opened");
                elevator.requestList.remove(0);
            }
            else 
            {
                elevator.direction = "DOWN";
                System.out.println();
                System.out.println("7- The Elevator is moving to pick up the User and go to his destination");
                while (elevator.currentFloor > elevator.requestList.get(elevator.requestList.size()-1))
                {
                    elevator.currentFloor --;
                    if (elevator.currentFloor != 0)
                    {
                        System.out.println("Elevator " + elevator.ID + " is at floor " + elevator.currentFloor);
                    }
                    if (elevator.currentFloor == elevator.floorList.get(0))
                    {
                        elevator.direction = "IDLE";
                    }
                }
                elevator.door = "OPEN";
                System.out.println("Door is opened");
                elevator.requestList.remove(elevator.requestList.stream().count()-1);
            }                
            elevator.door = "CLOSED";
            System.out.println("Door is closed");
            elevator.direction = "IDLE";
        }
        if (elevator.BufferList.stream().count() > 0)
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
