using System;
using System.Collections.Generic;

namespace Commercial_Controller
{
    class FloorRequestButton {
        public int ID;
        public bool Pressed = false;

        public FloorRequestButton(int ID){
            this.ID = ID;
        }
        
        public int getID() {
            return this.ID;
        }
        public void setID(int newID){
            this.ID = newID;
        }
    }
}