package spec3

/*
processes are an instance of a running program
	- it contains its own memory:
		- virtual address spaces.
		- code, stack, heap,
	- shared libraries <across the process>
	- registers
		- program counters, etc...


   Threads: 
    threads share some context.
	many threads can exists in one process

	- they have shared context like, virtual memory, file descriptor, etc..
	- they also have unique context like, code, data registers, stack, etc..


   GoRoutines:
	this is basically a thread but in go.


  Go Runtime scheduler
  
*/

func ConcurrencyDemo() {

}
