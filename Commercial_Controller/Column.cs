using System;
using System.Collections.Generic;


namespace Commercial_Controller
{

    class Column {
        public int ID;
        public int numberOfElevator;
        public int floorCall; // User position
        public int requestedFloor; // User destination
        public List<Elevator> elevatorList;
        public List<int> floorList;
        public List<Button> buttonList;
        
        public Column(int ID, int floorCall, int requestedFloor, int numberOfElevator){
            this.ID = ID;
            this.floorCall = floorCall;
            this.requestedFloor = requestedFloor;
            this.numberOfElevator = numberOfElevator;
            this.floorList = new List<int>();

            // The for loop is Creating Floor List for each Column
            for (int i = floorCall; i <= requestedFloor; i++){
                this.floorList.Add(i);
            }

            // Creating Elevator List for each Column in floors's range(floorCall to requestedFloor)
            this.elevatorList = new List<Elevator>();
            for (int i = 1; i <= this.numberOfElevator; i = i + 1)
            {
                this.elevatorList.Add(new Elevator(i, floorCall, requestedFloor));
            }

            // The IF clause is Creating CallButtons for Basements as columnID #1 = the  basement
            // and the ELSE clause is creating for the others floors in floors's range(floorCall to requestedFloor)
            this.buttonList = new List<Button>();
            if(ID == 1){
                for(int i = requestedFloor; i <= floorCall; i++){
                    Button callButton = new Button ("UP", i);
                    this.buttonList.Add(callButton);
                }
            }
            else{
                for(int i = floorCall; i <= requestedFloor; i++){
                    Button callButton = new Button("DOWN", i);
                    this.buttonList.Add(callButton);
                }
            }
            //Console.WriteLine("ColumnID: " + this.ID.ToString() + " is at floor " + floorCall.ToString() +
            //" and going to floor " + requestedFloor.ToString() + " with the elevator number " + numberOfElevator.ToString() );
        }   
    }
}
    