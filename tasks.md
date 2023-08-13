1. Can you make a Golang program that reads links from a text file and provides them as an argument to an executable called lux ? lux should run with those arguments even if it returns an error in any of the executions.

2. revision 873a5ca - works

3. latest revision - fix io package - stalls when executed with go run .
   We need to fix it. It's purpose was to display output from lux command on
   every new execution of the lux command. 

4. If 3 is fixed display the timer - how much time is left from lux to the console
   
