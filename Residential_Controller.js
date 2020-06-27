/*Method: RequestElevator(RequestedFloor, Direction)
The method 1 must return choosen elevator and make sure to move the elevators in the traitement.

Method
: RequestFloor(Elevator, RequestedFloor)
The method 2 must move the elevators in his traitement*/

//Button Class with THE FloorRequestButton ID
class FloorRequestButton {
    constructor(ID){
        this.ID = ID;
        this.Pressed = false;
    }
}

//Button Class with constructor direction & floor
class Button {
    constructor(direction, floor){
        this.direction = direction;
        this.floor = floor;
        this.light = 'OFF'
    }
}

class Elevator {
    constructor(elevatorID, nbFloor){
        this.elevatorID = elevatorID;
        this.currentFloor = 1;
        this.Direction = 'UP';
        this.RequestList = [];
        this.FloorCallButton = [];
        this.Door = 'CLOSED';
        for (var i = 1; i <= nbFloor; i++){
            this.FloorCallButton.push(i);
        }
    }
}

// Column class with constructor nbEelevator & nbFloor
class Column {
    constructor(ID, nbFloor, nbElevator){
        this.ID = ID;
        this.nbElevator = nbElevator;
        this.nbFloor = nbFloor;
        this.elevatorList = [];
        this.floorList = [];
        this.buttonList = [];
        
        //Add all elevator on the list of Elevator
        for (var i = 1; i <= nbElevator; i++) {
            this.elevatorList.push(new Elevator(i , nbFloor))
        }

         //nomberOfFloors
        for (var i = 1; i <= nbFloor; i++) {
            this.floorList.push(i)
        }
        //callButton
        for (var i = 1; i <= nbFloor; i++) {
            if (i != this.nbFloor - 1 ){
                var callbutton = new Button('DOWN', i)
                this.buttonList.push(callbutton)
                //console.log(this.buttonList);
            }
            if (i != nbFloor ){
                var callbutton = new Button('UP', i)
                this.buttonList.push(callbutton)
                //console.log(this.buttonList);
            }
        }    
    }


    UpdateList(List, currentFloor){
        var List = [];
        List.push(currentFloor);
        List.sort(function(a, b){
            return a-b
        });
    }
    // finding the elevator which is closer to the user position
    nearestElevator (elevatorList, userCurrentFloor){
        var distance = Column.floorList;
        var bestNearestElevator;
        for (let i = 0; i < elevatorList; i++) {
            const elevator = elevatorList[i];
            if (Math.abs(elevator.currentFloor - userCurrentFloor) < distance){                
                bestNearestElevator = elevator;
            }
        }
            return bestNearestElevator;
    }

    RequestElevator(RequestedFloor, direction){
        console.log("The user is at floor " + RequestedFloor + " and is going "  + direction); 
        this.findBestElevator(RequestedFloor, direction);

        //Here's the steps to move the elevator once the user request the elevator
       
            if (Elevator.RequestList > Elevator.currentFloor){
                Elevator.Direction = 'UP';
                while (Elevator.currentFloor < elevator.RequestList[0]){
                    Elevator.currentFloor += 1;
                    console.log("The Elevator # " + elevator.elevatorID + " is now at floor "  + elevator.currentFloor);
                    if (Elevator.currentFloor == Elevator.RequestList){
                        Elevator.Direction = 'IDLE';
                    }
                }
                elevator.Door = 'OPEN';
                console.log('Door is open');
                
            }
            else {
                Elevator.Direction = 'DOWN';
                while (Elevator.currentFloor > Elevator.RequestList){
                    Elevator.currentFloor -= 1;
                    console.log("The elevator No: " + elevator.elevatorID + " is at floor "   + elevator.currentFloor);
                    if (Elevator.currentFloor == 1){
                        Elevator.Direction = 'IDLE';
                    }
                }
                Elevator.Door = 'OPEN';
                console.log('Door is open');
                
            }
            
            Elevator.Door = 'CLOSED';
            console.log('Door is closed');
            Elevator.Direction = 'IDLE';
        
    }
    findBestElevator(RequestedFloor, direction){

            var BestElevator = null ;

       for (let elevator of this.elevatorList) {
            console.log("The Elevator #" + elevator.elevatorID + " is now at floor "  + RequestedFloor + ' floor and its direction is ' + direction);
            if (elevator.currentFloor == RequestedFloor && elevator.Direction == direction){
                if (elevator.Door == 'OPEN'){
                    this.UpdateList(elevator.RequestList, RequestedFloor);
                    console.log("the best elevator No: " + elevator.elevatorID + ' is there!');
                    return elevator;
                }
            }
            else if(elevator.currentFloor > RequestedFloor && elevator.Direction ==='IDLE'){           
                BestElevator = elevator;
            }
            else if(direction === 'UP' && elevator.currentFloor < RequestedFloor){
                BestElevator = elevator;
            }
            else if(direction === 'DOWN' && elevator.currentFloor > RequestedFloor){
                BestElevator = elevator;
            }
            else if(direction === 'DOWN' && elevator.Direction === 'DOWN'){
                BestElevator = elevator;
            }
            else if(direction === 'UP' && elevator.Direction === 'UP'){
                BestElevator = elevator;
            }
            else{
                BestElevator = this.nearestElevator( BestElevator, direction);
                this.UpdateList(this.RequestList, RequestedFloor);
                console.log("The elevator No : " + elevator.elevatorID + ' is comming!');
                return BestElevator;
            }


        }
        
    }
    RequestFloor(elevator, RequestedFloor){ 
        this.UpdateList(elevator, RequestedFloor);

        //Here's the steps to move the elevator once the user is in the elevator
       
            if (Elevator.RequestList > Elevator.currentFloor){
                Elevator.Direction = 'UP';
                while (elevator.currentFloor < Elevator.RequestList[0]){
                    Elevator.currentFloor += 1;
                    console.log("The Elevator # " + elevator.elevatorID + " is now at floor "  + elevator.currentFloor);
                    if (Elevator.currentFloor == Elevator.RequestList){
                        Elevator.Direction = 'IDLE';
                    }
                }
                Elevator.Door = 'OPEN';
                console.log('Door is opened');
                
            }
            else {
                Elevator.Direction = 'DOWN';
                while (Elevator.currentFloor > Elevator.RequestList){
                    Elevator.currentFloor -= 1;
                    console.log("The elevator No: " + elevator.elevatorID + " is at floor "   + elevator.currentFloor);
                    if (elevator.currentFloor == 1){
                        elevator.Direction = 'IDLE';
                    }
                }
                Elevator.Door = 'OPEN';
                console.log('Door is open');
                
            }
            
            Elevator.Door = 'CLOSED';
            console.log('Door is closed');
            Elevator.Direction = 'IDLE';
        
    }   
}


// --- /Classes ----   
// --- Scenarios ---    
function Scenario1(){
    console.log('******************* ******************* *******************');
    console.log('*******************      Scenario 2     *******************');
    console.log('******************* ******************* *******************');
    column1 = new Column(1, 10, 2);

    column1.elevatorList[0].currentFloor = 1;
    column1.elevatorList[0].Direction = 'IDLE';
    column1.elevatorList[1].currentFloor = 7;
    column1.elevatorList[1].Direction = 'IDLE';
    console.log('******************* USER-1 goes from floor 1 to floor 6  *******************');
    RequestedFloor = 2;
    Direction = 'DOWN';
    Destination = 6;
    elevatorList = [1, 2];
    column1.RequestElevator(RequestedFloor, Direction);
    column1.nearestElevator(elevatorList, 1);
    column1.RequestFloor(2, 7);
    console.log('******************* USER-3 goes from floor 9 to floor 2  *******************');
    RequestedFloor = 9;
    Direction = 'DOWN';
    Destination = 2;
    elevatorList = [1, 2];
    column1.RequestElevator(RequestedFloor, Direction);
    column1.nearestElevator(elevatorList, 9);
    column1.RequestFloor(2, 2);
}

/*column = new Column(1, 10, 2);
best = column.findBestElevator(3, 'DOWN')
console.log(best)
best2 = column.findBestElevator(6, 'UP')
console.log(best2)

elevator = column.RequestFloor(1, 7)
console.log(elevator)
console.log(column)*/
//  --- /Scenarios---

Scenario1();

