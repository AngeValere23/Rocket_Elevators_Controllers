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
        self.Direction = 'UP'
        self.RequestList = [] 
        self.FloorCallButton = []
        self.Door = 'CLOSED'
        for i in range (nbFloor):
            self.FloorCallButton.append(i)


# Column class with constructor nbEelevator & nbFloor
class Column(): 
    def __init__ (self, columnID, nbFloor, nbElevator):
        self.columnID = columnID
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

    #finding the elevator which is closer to the user position
    def nearestElevator(self, elevatorsList, userPosition):
            distance = len(self.floorList)
            bestNearestElevator = elevatorsList[0]       
            for elevator in elevatorsList:
                    if abs(elevator.Position - userPosition) < distance:                
                            bestNearestElevator = elevator                    
            return bestNearestElevator

    def RequestElevator(self, RequestedFloor, direction):
        print("The user is at floor " , RequestedFloor , " and is going "  , direction) 
        self.findBestElevator(RequestedFloor, direction)    
    
    def findBestElevator(self, RequestedFloor, direction):
        BestElevator = None
        for elevator in self.elevatorList:
            print("The Elevator #" , elevator.ID , " is now at floor "  , elevator.currentFloor , ' floor and its direction is ' , elevator.Direction)
            if (elevator.currentFloor == RequestedFloor and elevator.Direction == direction):
                if (elevator.Door == 'OPEN'):
                    self.UpdateList(elevator.RequestList, RequestedFloor)
                    print("the best elevator No: " , elevator.elevatorID , ' is there!')
                    self.moveElevator(elevator)
                    return elevator
                            
            elif(elevator.currentFloor > RequestedFloor and elevator.Direction == 'IDLE'):           
                BestElevator = elevator
            
            elif(direction == 'UP' and elevator.currentFloor < RequestedFloor):
                BestElevator = elevator
            
            elif(direction == 'DOWN' and elevator.currentFloor > RequestedFloor):
                BestElevator = elevator
            
            elif(direction == 'DOWN' and elevator.Direction == 'DOWN'):
                BestElevator = elevator
            
            elif(direction == 'UP' and elevator.Direction == 'UP'):
                BestElevator = elevator
            
            else:
                BestElevator = self.nearestElevator( RequestedFloor, direction)
                self.UpdateList(elevator.RequestList, RequestedFloor)
                print("The elevator No : " , elevator.elevatorID , ' is comming!')
                self.moveElevator(BestElevator)
                return BestElevator       

    def moveElevator(self, elevator):
        while (len(elevator.RequestList) > 0) :
            if (elevator.RequestList[0] > elevator.currentFloor) :
                elevator.Direction = 'UP'
                while (elevator.currentFloor < elevator.RequestList[0]):
                    elevator.currentFloor += 1
                    print('The Elevator # ', elevator.elevatorID , ' is now at floor ', elevator.currentFloor)
                    
                    if (elevator.currentFloor == len(elevator.RequestList)):
                        elevator.Direction == 'IDLE'
                elevator.Door = 'OPEN'
                print('Door is open')
                
            else:
                elevator.Direction = 'DOWN'
                while (elevator.currentFloor > elevator.RequestList[0]):
                    elevator.currentFloor -= 1
                    print('Elevator ', elevator.ID, ' is at Floor ', elevator.currentFloor)
                    if (elevator.currentFloor == 1):
                        elevator.Direction = 'IDLE'
                elevator.Door = 'OPEN'
                print('Door is open')
               
            elevator.Door = 'CLOSED'
            print('Door is closed')
            elevator.Direction = 'IDLE'      

    def RequestFloor(self, elevator, RequestedFloor):
            self.UpdateList(elevator.RequestList, RequestedFloor)
            self.moveElevator(elevator)
            return       

 
# --- Scenarios ---
def Scenario1 ():
    print('******************* ******************* *******************')
    print('*******************      Scenario 1     *******************')
    print('******************* ******************* *******************')
    column = Column(1, 10, 2)

    column.elevatorList[0].Position = 2
    column.elevatorList[0].Direction = 'IDLE'
    column.elevatorList[1].Position = 6
    column.elevatorList[1].Direction = 'IDLE'
    print('******************* USER-1 goes from floor 3 to floor 7  *******************')
    RequestedFloor = 3
    Direction = 'UP'
    elevatorList = [1, 2]
    column.RequestElevator(RequestedFloor, Direction)
    column.nearestElevator(elevatorList, 3)
    column.RequestFloor(1, 7)
    



#  --- /Scenarios---

Scenario1()

