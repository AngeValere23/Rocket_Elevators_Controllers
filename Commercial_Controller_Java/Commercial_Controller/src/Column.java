import java.util.ArrayList;
import java.util.List;

/**
 * Elevator
 */
public class Column {
    
    public int ID;
        public int numberOfElevator;
        public int floorCall; // User position
        public int requestedFloor; // User destination
        public List<Elevator> elevatorList;
        public List<Integer> floorList;
        public List<Button> buttonList;
        
        public Column(int ID, int floorCall, int requestedFloor, int numberOfElevator){
            this.ID = ID;
            this.floorCall = floorCall;
            this.requestedFloor = requestedFloor;
            this.numberOfElevator = numberOfElevator;
            this.floorList = new ArrayList<Integer>();

            // The for loop is Creating Floor List for each Column
            for (int i = floorCall; i <= requestedFloor; i++){
                this.floorList.add(i);
            }

            // Creating Elevator List for each Column in floors's range(floorCall to requestedFloor)
            this.elevatorList = new ArrayList<Elevator>();
            for (int i = 1; i <= this.numberOfElevator; i = i + 1)
            {
                this.elevatorList.add(new Elevator(i, floorCall, requestedFloor));
            }

            // The IF clause is Creating CallButtons for Basements as columnID #1 = the  basement
            // and the ELSE clause is creating for the others floors in floors's range(floorCall to requestedFloor)
            this.buttonList = new ArrayList<Button>();
            if(ID == 1){
                for(int i = requestedFloor; i <= floorCall; i++){
                    Button callButton = new Button ("UP", i);
                    this.buttonList.add(callButton);
                }
            }
            else{
                for(int i = floorCall; i <= requestedFloor; i++){
                    Button callButton = new Button("DOWN", i);
                    this.buttonList.add(callButton);
                }
            }
            //Console.WriteLine("ColumnID: " + this.ID.ToString() + " is at floor " + floorCall.ToString() +
            //" and going to floor " + requestedFloor.ToString() + " with the elevator number " + numberOfElevator.ToString() );
        }   
}