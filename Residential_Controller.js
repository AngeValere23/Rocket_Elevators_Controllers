/*Method: RequestElevator(RequestedFloor, Direction)
The method 1 must return choosen elevator and make sure to move the elevators in the traitement.

Method
: RequestFloor(Elevator, RequestedFloor)
The method 2 must move the elevators in his traitement*/

// Column class with constructor nbEelevator & nbFloor
class Column{
    constructor(nbElevator, nbFloor){
        this.nbElevator = nbElevator;
        this.nbFloor = nbFloor;
        this.elevatorList = [];
        this.floorList = [];
        this.buttonList = [];

        //Add all elevator on the list of Elevator
        for(var i = 1; i <= nbElevator; i++){
            this.elevatorList.push(new Elevator(i, nbFloor));
        }

        //nomberOfFloors
        for (var i = 1; i <= nbFloor; i++) {
            this.floorList.push(i)
        }

        //callButton
        for (var i = 1; i <= nbFloor; i++) {
            if (i != this.nbFloor - 1 ){
                var callButton = new Button('DOWN', i)
                this.buttonList.push(callButton)
                console.log(this.buttonList);
            }
            if (i != nbFloor ){
                callButton = new Button('UP', i)
                this.buttonList.push(callButton)
                console.log(this.buttonList);
            }
        }    
    }
}

//Button Class with constructor direction & floor
class Button {
    constructor(direction, floor){
        this.direction = direction;
        this.floor = floor;
        this.light = "OFF";
    }
}

// Elevator class with constructor elevatorID & nbFfloor
class Elevator {
    constructor(elevatorID, nbFloor){
        this.elevatorID = elevatorID;
        this.currentFloor = 1;
        this.Direction = 'UP';
        this.elevatorRequestList = [];
        this.FloorCallButton = [];
        this.Door = 'CLOSED';
        for (var i = 1; i <= nbFloor; i++){
            this.FloorCallButton.push(i);
        }
    }
}

function findBestElevator(RequestedFloor, direction){
    var BestElevator;

    for (var i = 1; i <= this.elevatorList; i++){ 
        if(elevator.currentFloor == RequestedFloor && elevator.Direction == direction){           
            BestElevator = elevator;
        }
        else if(elevator.currentFloor > requestedFloor && elevator.Direction ==='idle'){           
            BestElevator = elevator;
        }
        else if(direction === 'UP' && elevator.currentFloor < requestedFloor){
            BestElevator = elevator;
        }
        else if(direction === 'down' && elevator.currentFloor > requestedFloor){
            BestElevator = elevator;
        }
        else if(direction === 'down' && elevator.Direction === 'down'){
            BestElevator = elevator;
        }
        else if(direction === 'up' && elevator.Direction === 'up'){
            BestElevator = elevator;
        }
        else{
             BestElevator = this.nearestElevator( RequestedFloor, direction);
            return BestElevator;
        }
    } 

}

function nearestElevator(elevatorList, RequestedFloor){

    // finding the nearest elevator which is closer to the user position or 
        var distance = this.column.floorList.length;
        var best;
        for (var elevator of elevatorList){
            if (Math.abs(elevator.Position - RequestedFloor) < distance){
                    best = elevator;
            }
        }
        return best;
}
class Controller {

     RequestElevator(RequestedFloor, direction){
        var BestElevator = this.findBestElevator(RequestedFloor, direction)
        BestElevator.elevatorRequestList.push(requestedFloor);
        console.log("The Elevator #" + elevator.elevatorID + " is now at floor "  + elevator.currentFloor + ' floor and its direction is ' + elevator.Direction);
        
    }

    RequestFloor(Elevator, RequestedFloor){
        Elevator.elevatorRequestList.push(RequestedFloor)
        Elevator.elevatorRequestList.sort();
        this.move(elevator);
        return;    
    }
}
//Here's the steps to move the elevator once the user is in the elevator
function moveElevator(elevator){
    while (elevator.elevatorRequestList[0] > 0){
        if (elevator.elevatorRequestList[0] > elevator.currentFloor){
            elevator.Direction = 'UP';
            while (elevator.currentFloor < elevator.elevatorRequestList[0]){
                elevator.currentFloor ++;
                console.log("The Elevator #" + elevator.elevatorID + " is now at floor "  + elevator.currentFloor);
                if (elevator.currentFloor == this.colum){
                    elevator.Direction === 'IDLE';
                }
            }
            elevator.Door === 'OPEN';
            console.log('Door is open');
        }
        else {
            elevator.Direction = 'DOWN';
            while (elevator.currentFloor > elevator.elevatorRequestList[0]){
                elevator.currentFloor -= 1;
                console.log('Elevator ' + elevator.ID + ' is at floor ' + elevator.currentFloor);
                if (elevator.currentFloor == 1){
                    elevator.Direction = 'IDLE';
                }
            }
            elevator.Door = 'OPEN';
            console.log('Door is open');
        }
        
        elevator.Door = 'CLOSED';
        console.log('Door is closed');
        elevator.Direction = 'IDLE';
    }
}
// --- Scenarios ---

    console.log('******************* ******************* *******************');
    console.log('*******************  Just  Litlle Test  *******************');
    console.log('******************* ******************* *******************');
    test = new Column(2, 6);
    console.log(test);
    test1 = new Controller(3, 'DOWN')
    console.log(test1);
    test3 = new findBestElevator(5,'UP')
    console.log(test3)
    //findBestElevator(6,'UP');
    var column = new Column(2, 10);
    //Elevator 0 Idle at floor 2 and elevator 1  Idle at floor 6

    column.elevatorList[0].currentDirection = 'DOWN';
   column.elevatorList[0].currentFloor = 2;
   column.elevatorList[1].currentDirection = 'idle';
   column.elevatorList[1].currentFloor = 6;
   console.log(column)

   console.log(moveElevator (2));

    