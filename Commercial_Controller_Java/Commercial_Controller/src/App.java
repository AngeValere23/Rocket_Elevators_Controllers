/*Modern Approach:
Method 1: RequestElevator (FloorNumber)
This method represents an elevator request on a floor or basement.
Method 2: AssignElevator (RequestedFloor)
This method will be used for the requests made on the first floor.
*/

public class App {
   

        static void Scenario1 ()
        {
                System.out.println();
                System.out.println("*******************      Scenario 1     *******************" + "\n");
                System.out.println("******** User at floor 1. He goes UP to floor 20 **********");
                System.out.println("*********** Elevator 5 from Column 2 is expected **********" + "\n");
                Battery Battery1 = new Battery(4, 66, 6, 5);
                // Initializing Elevator 1 of Column 2
                Battery1.columnList.get(1).elevatorList.get(0).currentFloor = 20;
                Battery1.columnList.get(1).elevatorList.get(0).direction = "DOWN";
                Battery1.columnList.get(1).elevatorList.get(0).requestList.add(5);
                // Initializing Elevator 2 of Column 2
                Battery1.columnList.get(1).elevatorList.get(1).currentFloor = 2;
                Battery1.columnList.get(1).elevatorList.get(1).direction = "UP";
                Battery1.columnList.get(1).elevatorList.get(1).requestList.add(15);
                // // Initializing Elevator 3 of Column 2
                Battery1.columnList.get(1).elevatorList.get(2).currentFloor = 13;
                Battery1.columnList.get(1).elevatorList.get(2).direction = "DOWN";
                Battery1.columnList.get(1).elevatorList.get(2).requestList.add(1);
                // // Initializing Elevator 4 of Column 2
                Battery1.columnList.get(1).elevatorList.get(3).currentFloor = 15;
                Battery1.columnList.get(1).elevatorList.get(3).direction = "DOWN";
                Battery1.columnList.get(1).elevatorList.get(3).requestList.add(2);
                // // Initializing Elevator 5 of Column 2
                Battery1.columnList.get(1).elevatorList.get(4).currentFloor = 6;
                Battery1.columnList.get(1).elevatorList.get(4).direction = "DOWN";
                Battery1.columnList.get(1).elevatorList.get(4).requestList.add(1);  
                Battery1.AssignElevator(20);
                
        }

        // Scenario 2
        static void Scenario2 ()
        {
            System.out.println();
            System.out.println("*******************      Scenario 2     *******************" + "\n");
            System.out.println("******** User at floor 1. He goes UP to floor 36 **********");
            System.out.println("*********** Elevator 1 from Column 3 is expected **********" + "\n");
            Battery Battery2 = new Battery(4, 66, 6, 5);
            // Initializing Elevator 1 of Column 3
            Battery2.columnList.get(2).elevatorList.get(0).currentFloor = 1;
            Battery2.columnList.get(2).elevatorList.get(0).direction = "UP";
            Battery2.columnList.get(2).elevatorList.get(0).requestList.add(21);
            // Initializing Elevator 2 of Column 3
            Battery2.columnList.get(2).elevatorList.get(1).currentFloor = 23;
            Battery2.columnList.get(2).elevatorList.get(1).direction = "UP";
            Battery2.columnList.get(2).elevatorList.get(1).requestList.add(28);
            // // Initializing Elevator 3 of Column 3
            Battery2.columnList.get(2).elevatorList.get(2).currentFloor = 33;
            Battery2.columnList.get(2).elevatorList.get(2).direction = "DOWN";
            Battery2.columnList.get(2).elevatorList.get(2).requestList.add(1);
            // // Initializing Elevator 4 of Column 3
            Battery2.columnList.get(2).elevatorList.get(3).currentFloor = 40;
            Battery2.columnList.get(2).elevatorList.get(3).direction = "DOWN";
            Battery2.columnList.get(2).elevatorList.get(3).requestList.add(24);
            // // Initializing Elevator 5 of Column 3
            Battery2.columnList.get(2).elevatorList.get(4).currentFloor = 39;
            Battery2.columnList.get(2).elevatorList.get(4).direction = "DOWN";
            Battery2.columnList.get(2).elevatorList.get(4).requestList.add(1);
            Battery2.AssignElevator(36);
                
        }

        // Scenario 3
        static void Scenario3 ()
        {
            System.out.println();
            System.out.println("*******************      Scenario 3     *******************" + "\n");
            System.out.println("******** User at floor 54. He goes Down to floor 1 ********");
            System.out.println("*********** Elevator 1 from Column 4 is expected **********" + "\n");
            Battery Battery3 = new Battery(4, 66, 6, 5);
            // Initializing Elevator 1 of Column 3
            Battery3.columnList.get(3).elevatorList.get(0).currentFloor = 58;
            Battery3.columnList.get(3).elevatorList.get(0).direction = "DOWN";
            Battery3.columnList.get(3).elevatorList.get(0).requestList.add(1);
            // Initializing Elevator 2 of Column 3
            Battery3.columnList.get(3).elevatorList.get(1).currentFloor = 50;
            Battery3.columnList.get(3).elevatorList.get(1).direction = "UP";
            Battery3.columnList.get(3).elevatorList.get(1).requestList.add(60);
            // // Initializing Elevator 3 of Column 3
            Battery3.columnList.get(3).elevatorList.get(2).currentFloor = 46;
            Battery3.columnList.get(3).elevatorList.get(2).direction = "DOWN";
            Battery3.columnList.get(3).elevatorList.get(2).requestList.add(58);
            // // Initializing Elevator 4 of Column 3
            Battery3.columnList.get(3).elevatorList.get(3).currentFloor = 1;
            Battery3.columnList.get(3).elevatorList.get(3).direction = "DOWN";
            Battery3.columnList.get(3).elevatorList.get(3).requestList.add(54);
            // // Initializing Elevator 5 of Column 3
            Battery3.columnList.get(3).elevatorList.get(4).currentFloor = 60;
            Battery3.columnList.get(3).elevatorList.get(4).direction = "DOWN";
            Battery3.columnList.get(3).elevatorList.get(4).requestList.add(1);
            Battery3.RequestElevator(54);
                
        }

        // Scenario 4
        static void Scenario4 ()
        {
            System.out.println();
            System.out.println("*******************      Scenario 4     *******************" + "\n");
            System.out.println("******** User at floor -3. He goes UP to floor 1 **********");
            System.out.println("*********** Elevator 4 from Column 1 is expected **********" + "\n");
            Battery Battery4 = new Battery(4, 66, 6, 5);
            // Initializing Elevator 1 of Column 3
            Battery4.columnList.get(0).elevatorList.get(0).currentFloor = -4;
            Battery4.columnList.get(0).elevatorList.get(0).direction = "IDLE";
            // Initializing Elevator 2 of Column 3
            Battery4.columnList.get(0).elevatorList.get(1).currentFloor = 1;
            Battery4.columnList.get(0).elevatorList.get(1).direction = "IDLE";
            // // Initializing Elevator 3 of Column 3
            Battery4.columnList.get(0).elevatorList.get(2).currentFloor = -3;
            Battery4.columnList.get(0).elevatorList.get(2).direction = "DOWN";
            Battery4.columnList.get(0).elevatorList.get(2).requestList.add(-5);
            // // Initializing Elevator 4 of Column 3
            Battery4.columnList.get(0).elevatorList.get(3).currentFloor = -6;
            Battery4.columnList.get(0).elevatorList.get(3).direction = "UP";
            Battery4.columnList.get(0).elevatorList.get(3).requestList.add(1);
            // // Initializing Elevator 5 of Column 3
            Battery4.columnList.get(0).elevatorList.get(4).currentFloor = -1;
            Battery4.columnList.get(0).elevatorList.get(4).direction = "DOWN";
            Battery4.columnList.get(0).elevatorList.get(4).requestList.add(-6);
            Battery4.RequestElevator(-3);
                
        }
    
    public static void main(String[] args) throws Exception {
            Scenario1();
            //Scenario2();
            //Scenario3();
            //Scenario4();
    }
}
