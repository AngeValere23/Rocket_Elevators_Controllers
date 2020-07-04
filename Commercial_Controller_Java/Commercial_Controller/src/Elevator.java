import java.util.ArrayList;
import java.util.List;

/**
 * Elevator
 */
public class Elevator {
    
    public int ID;
    public int currentFloor;
    public String direction = "UP";
    public List<Integer> requestList;
    public List<Integer> floorList;
    public String door = "CLOSED";
    public String BufferDirection;
    public List<Integer> BufferList;
        

    public Elevator (int ID, int floorCall, int requestedFloor){
        this.ID = ID;
        this.currentFloor = floorCall;
        this.requestList = new ArrayList<Integer>();
        this.floorList = new ArrayList<Integer>();
        this.BufferDirection = "UP";
        this.BufferList = new ArrayList<Integer>();
        for (int i = floorCall; i <= requestedFloor; i++){
            this.floorList.add(i);
        }
    }
}