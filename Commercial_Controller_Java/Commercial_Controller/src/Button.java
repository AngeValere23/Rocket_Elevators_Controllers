/**
 * Button
 */
public class Button {

    
    String direction;
    int floor;
    String light;

    public Button(String direction, int floor){
        this.direction = direction;
        this.floor = floor;
        this.light = "OFF";
    }

    public String getDirection(){
        return this.direction;
    }

    public void setDirection(String newDirection){
        this.direction = newDirection;
    }

    public int getFloor(){
        return this.floor;
    }

    public void getFloor(int newFloor){
        this.floor = newFloor;
    }
}