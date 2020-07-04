/**
 * FloorRequestFloor
 */
public class FloorRequestFloor {
    
    public int ID;
    public Boolean Pressed = false;

    public void FloorRequestButton(final int ID) {
        this.ID = ID;
    }

    public int getID() {
        return this.ID;
    }

    public void setID(final int newID) {
        this.ID = newID;
    }
}