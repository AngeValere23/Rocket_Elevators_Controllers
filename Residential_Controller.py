"""Method: RequestElevator(RequestedFloor, Direction)
The method 1 must return choosen elevator and make sure to move the elevators in the traitement.

Method
: RequestFloor(Elevator, RequestedFloor)
The method 2 must move the elevators in his traitement"""

#Button Class with THE FloorRequestButton ID
class FloorRequestButton():
    def __init__(self, ID):
        self.ID = ID
        self.Pressed = False

#Button Class with constructor direction & floor
class Button() :
    def __init__ (self, direction, floor):
        self.direction = direction
        self.floor = floor
        self.light = 'OFF'

#Elevator Class with the elevator & floor
class Elevator():
    def __init__(self, ID, nbFloor):
        self.ID = ID
        self.currentFloor = 1
        self.direction = 'IDLE'
        self.requestList = [] 
        self.floorCallButton = []
        self.door = 'CLOSED'
        for i in range (nbFloor):
            self.floorCallButton.append(i)


# Column class with constructor nbEelevator & nbFloor
class Column(): 
    def __init__ (self, ID, nbFloor, nbElevator):
        self.ID = ID
        self.nbElevator = nbElevator
        self.nbFloor = nbFloor
        self.elevatorList = []
        self.floorList = []
        self.buttonList = []

       
        #Add all elevator on the list of Elevator
        for i in range (nbElevator):
            elevator = Elevator(i + 1, nbFloor)
            self.elevatorList.append(elevator)
            #print(self.elevatorList)

        #nomberOfFloors
        for i in range (nbFloor):
            self.floorList.append(i)

        #callButton
        for i in range (nbFloor):
            if i != self.nbFloor - 1 :
                callbutton = Button('DOWN', i)
                self.buttonList.append(callbutton)
               

            if i != nbFloor :
                callbutton = Button('UP', i)
                self.buttonList.append(callbutton)

    def UpdateList(self, elList, currentFloor):
        elList = []
        elList.append(currentFloor)
        elList.sort()

    #finding the elevator which is closer to the user current floor
    def nearestElevator (self, elevatorList, userCurrentFloor):
        distance = len(self.floorList)
        bestNearestElevator = elevatorList[0]
        for elevator in elevatorList:
            if abs(elevator.currentFloor - userCurrentFloor) < distance:
                bestNearestElevator = elevator
        return bestNearestElevator

    def RequestElevator(self, requestedFloor, direction):
        self.findBestElevator(requestedFloor, direction)
        return requestedFloor

    def findBestElevator(self, requestedFloor, direction):
        print("The user is at floor ", requestedFloor , " and is going ", direction)
        BestElevator = []
        for elevator in self.elevatorList:
            print("The Elevator #", elevator.ID , " is now at floor ", elevator.currentFloor , " and its direction " , elevator.direction)
            if(elevator.currentFloor == requestedFloor and elevator.direction == direction):
                if(elevator.door == 'OPEN'):
                    self.UpdateList(elevator.requestList, requestedFloor)
                    print("The best elevator No" , elevator.ID , " is there!")
                    self.moveElevator(elevator)
                    return elevator
            elif (elevator.currentFloor == direction):
                if (elevator.currentFloor > requestedFloor and elevator.direction == 'DOWN'):
                    BestElevator.append(elevator)
                elif (elevator.currentFloor < requestedFloor and elevator.direction == 'UP'):
                    BestElevator.append(elevator)
                else:
                    BestElevator.append(elevator)
            elif (elevator.direction == 'IDLE'):
                if (elevator.currentFloor > requestedFloor and elevator.direction == 'IDLE'):
                    BestElevator.append(elevator)        
                elif (elevator.currentFloor < requestedFloor and elevator.direction == 'IDLE'):
                    BestElevator.append(elevator)
                else:
                    BestElevator.append(elevator)   
            else:
                bestElevator = self.nearestElevator(BestElevator , requestedFloor)
                self.UpdateList(elevator.requestList, requestedFloor)
                print("The elevator No" , elevator.ID , " is comming")
                self.moveElevator(bestElevator)
                return bestElevator
    
    def moveElevator(self, elevator):
        while (len(elevator.requestList) > 0):
            if(elevator.requestList[0] > elevator.currentFloor):
                elevator.direction = 'UP'
                while (elevator.currentFloor < elevator.requestList[0]):
                    elevator.currentFloor += 1
                    print("The elevator #" , elevator.ID , " is now at floor ", elevator.currentFloor)

                    if (elevator.currentFloor == len(elevator.requestList[0])):
                        elevator.direction = 'IDLE'
            else:
                elevator.direction = 'DOWN'
                while (elevator.currentFloor > elevator.requestList[0]):
                    elevator.currentFloor -= 1
                    print("The elevator #" , elevator.ID , " is at floor " , elevator.currentFloor)
                    if (elevator.direction == 1):
                        elevator.direction = 'IDLE' 
                elevator.door = 'OPEN'
                print("Door opened")
            
            elevator.door = 'CLOSED'
            print("Door is closed")   
    
    def RequestFloor(self, elevator, requestedFloor):
            self.UpdateList(elevator.requestList, requestedFloor)
            self.moveElevator(elevator)
            return elevator



# --- Scenarios ---
def Scenario1 ():
    print('******************* ******************* *******************')
    print('*******************      Scenario 1     *******************')
    print('******************* ******************* *******************')
    column = Column(1, 10, 2)

    column.elevatorList[0].CurrentFloor = 2
    column.elevatorList[0].Direction = 'IDLE'
    column.elevatorList[1].CurrentFloor = 6
    column.elevatorList[1].Direction = 'IDLE'
    print('******************* USER-1 goes from floor 3 to floor 7  *******************')
    RequestedFloor = 3
    direction = 'UP'
    destination = 7
    elevator = column.RequestElevator(RequestedFloor, direction)
    column.RequestFloor(elevator, destination)
    
# --- /Scenarios---
Scenario1() 


