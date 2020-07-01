using System;
using System.Collections.Generic;
using System.Linq;

/*Modern Approach:
Method 1: RequestElevator (FloorNumber)
This method represents an elevator request on a floor or basement.
Method 2: AssignElevator (RequestedFloor)
This method will be used for the requests made on the first floor.
*/
namespace Commercial_Controller
{

    class Program
    {
        static void Main(string[] args)
        {
           // Scenario 1
            void Scenario1 ()
            {
                Console.WriteLine();
                Console.WriteLine("*******************      Scenario 1     *******************" + "\n");
                Console.WriteLine("******** User at floor 1. He goes UP to floor 20 **********");
                Console.WriteLine("*********** Elevator 5 from Column 2 is expected **********" + "\n");
                Battery Battery1 = new Battery(4, 66, 6, 5);
                // Initializing Elevator 1 of Column 2
                Battery1.columnList[1].elevatorList[0].currentFloor = 20;
                Battery1.columnList[1].elevatorList[0].direction = "DOWN";
                Battery1.columnList[1].elevatorList[0].requestList.Add(5);
                // Initializing Elevator 2 of Column 2
                Battery1.columnList[1].elevatorList[1].currentFloor = 2;
                Battery1.columnList[1].elevatorList[1].direction = "UP";
                Battery1.columnList[1].elevatorList[1].requestList.Add(15);
                // // Initializing Elevator 3 of Column 2
                Battery1.columnList[1].elevatorList[2].currentFloor = 13;
                Battery1.columnList[1].elevatorList[2].direction = "DOWN";
                Battery1.columnList[1].elevatorList[2].requestList.Add(1);
                // // Initializing Elevator 4 of Column 2
                Battery1.columnList[1].elevatorList[3].currentFloor = 15;
                Battery1.columnList[1].elevatorList[3].direction = "DOWN";
                Battery1.columnList[1].elevatorList[3].requestList.Add(2);
                // // Initializing Elevator 5 of Column 2
                Battery1.columnList[1].elevatorList[4].currentFloor = 6;
                Battery1.columnList[1].elevatorList[4].direction = "DOWN";
                Battery1.columnList[1].elevatorList[4].requestList.Add(1);  
                Battery1.AssignElevator(20);
                
            }

            // Scenario 2
            void Scenario2 ()
            {
                Console.WriteLine();
                Console.WriteLine("*******************      Scenario 2     *******************" + "\n");
                Console.WriteLine("******** User at floor 1. He goes UP to floor 36 **********");
                Console.WriteLine("*********** Elevator 1 from Column 3 is expected **********" + "\n");
                Battery Battery2 = new Battery(4, 66, 6, 5);
                // Initializing Elevator 1 of Column 3
                Battery2.columnList[2].elevatorList[0].currentFloor = 1;
                Battery2.columnList[2].elevatorList[0].direction = "UP";
                Battery2.columnList[2].elevatorList[0].requestList.Add(21);
                // Initializing Elevator 2 of Column 3
                Battery2.columnList[2].elevatorList[1].currentFloor = 23;
                Battery2.columnList[2].elevatorList[1].direction = "UP";
                Battery2.columnList[2].elevatorList[1].requestList.Add(28);
                // // Initializing Elevator 3 of Column 3
                Battery2.columnList[2].elevatorList[2].currentFloor = 33;
                Battery2.columnList[2].elevatorList[2].direction = "DOWN";
                Battery2.columnList[2].elevatorList[2].requestList.Add(1);
                // // Initializing Elevator 4 of Column 3
                Battery2.columnList[2].elevatorList[3].currentFloor = 40;
                Battery2.columnList[2].elevatorList[3].direction = "DOWN";
                Battery2.columnList[2].elevatorList[3].requestList.Add(24);
                // // Initializing Elevator 5 of Column 3
                Battery2.columnList[2].elevatorList[4].currentFloor = 39;
                Battery2.columnList[2].elevatorList[4].direction = "DOWN";
                Battery2.columnList[2].elevatorList[4].requestList.Add(1);
                Battery2.AssignElevator(36);
                
            }

            // Scenario 3
            void Scenario3 ()
            {
                Console.WriteLine();
                Console.WriteLine("*******************      Scenario 3     *******************" + "\n");
                Console.WriteLine("******** User at floor 54. He goes Down to floor 1 ********");
                Console.WriteLine("*********** Elevator 1 from Column 4 is expected **********" + "\n");
                Battery Battery2 = new Battery(4, 66, 6, 5);
                // Initializing Elevator 1 of Column 3
                Battery2.columnList[3].elevatorList[0].currentFloor = 58;
                Battery2.columnList[3].elevatorList[0].direction = "DOWN";
                Battery2.columnList[3].elevatorList[0].requestList.Add(1);
                // Initializing Elevator 2 of Column 3
                Battery2.columnList[3].elevatorList[1].currentFloor = 50;
                Battery2.columnList[3].elevatorList[1].direction = "UP";
                Battery2.columnList[3].elevatorList[1].requestList.Add(60);
                // // Initializing Elevator 3 of Column 3
                Battery2.columnList[3].elevatorList[2].currentFloor = 46;
                Battery2.columnList[3].elevatorList[2].direction = "DOWN";
                Battery2.columnList[3].elevatorList[2].requestList.Add(58);
                // // Initializing Elevator 4 of Column 3
                Battery2.columnList[3].elevatorList[3].currentFloor = 1;
                Battery2.columnList[3].elevatorList[3].direction = "DOWN";
                Battery2.columnList[3].elevatorList[3].requestList.Add(54);
                // // Initializing Elevator 5 of Column 3
                Battery2.columnList[3].elevatorList[4].currentFloor = 60;
                Battery2.columnList[3].elevatorList[4].direction = "DOWN";
                Battery2.columnList[3].elevatorList[4].requestList.Add(1);
                Battery2.RequestElevator(54);
                
            }

            // Scenario 4
            void Scenario4 ()
            {
                Console.WriteLine();
                Console.WriteLine("*******************      Scenario 4     *******************" + "\n");
                Console.WriteLine("******** User at floor -3. He goes UP to floor 1 **********");
                Console.WriteLine("*********** Elevator 4 from Column 1 is expected **********" + "\n");
                Battery Battery2 = new Battery(4, 66, 6, 5);
                // Initializing Elevator 1 of Column 3
                Battery2.columnList[0].elevatorList[0].currentFloor = -4;
                Battery2.columnList[0].elevatorList[0].direction = "IDLE";
                //Battery2.columnList[0].elevatorList[0].requestList.Add();
                // Initializing Elevator 2 of Column 3
                Battery2.columnList[0].elevatorList[1].currentFloor = 1;
                Battery2.columnList[0].elevatorList[1].direction = "IDLE";
                //Battery2.columnList[0].elevatorList[1].requestList.Add(28);
                // // Initializing Elevator 3 of Column 3
                Battery2.columnList[0].elevatorList[2].currentFloor = -3;
                Battery2.columnList[0].elevatorList[2].direction = "DOWN";
                Battery2.columnList[0].elevatorList[2].requestList.Add(-5);
                // // Initializing Elevator 4 of Column 3
                Battery2.columnList[0].elevatorList[3].currentFloor = -6;
                Battery2.columnList[0].elevatorList[3].direction = "UP";
                Battery2.columnList[0].elevatorList[3].requestList.Add(1);
                // // Initializing Elevator 5 of Column 3
                Battery2.columnList[0].elevatorList[4].currentFloor = -1;
                Battery2.columnList[0].elevatorList[4].direction = "DOWN";
                Battery2.columnList[0].elevatorList[4].requestList.Add(-6);
                Battery2.RequestElevator(-3);
                
            }
            Scenario1();
            Scenario2();
            //Scenario3();
            //Scenario4();
        }
    }
}
