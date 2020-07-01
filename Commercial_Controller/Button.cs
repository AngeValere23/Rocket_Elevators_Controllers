using System;
using System.Collections.Generic;

namespace Commercial_Controller
{

    class Button{
        public string direction;
        public int floor;
        public string light;

        public Button(string direction, int floor){
            this.direction = direction;
            this.floor = floor;
            this.light = "OFF";
        }

        public string getDirection(){
            return this.direction;
        }

        public void setDirection(string newDirection){
            this.direction = newDirection;
        }

        public int getFloor(){
            return this.floor;
        }

        public void getFloor(int newFloor){
            this.floor = newFloor;
        }
    }
}