using System;
using System.Collections.Generic;

namespace Commercial_Controller
{
    
    class Elevator{
        public int ID;
        public int currentFloor;
        public string direction = "IDLE";
        public List<int> requestList;
        public List<int> floorList;
        public string door = "CLOSED";
        public string BufferDirection;
        public List<int> BufferList;
        

        public Elevator (int ID, int floorCall, int requestedFloor){
            this.ID = ID;
            this.currentFloor = floorCall;
            this.requestList = new List<int>();
            this.floorList = new List<int>();
            this.BufferDirection = "UP";
            this.BufferList = new List<int>();
            for (int i = floorCall; i <= requestedFloor; i++){
                this.floorList.Add(i);
            }
        }
    }
}