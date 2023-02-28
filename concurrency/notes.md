## Go make it simple to build 
1. fan-in model
2. fan-out model

### whole idea is that we can balance out level of concurrency to many app effecient.

### in go workers -> goroutines and they use channels to communicate to implement concurrency.





## goroutine

is a function executing concurrently with other goroutines in the same address spaces (i.e within the same program/process).
it is lightweight costing little more than the allocation of stack space.


### few more points on goroutine
1. sheduled by go runtime // scheduler
2. lighter than threads
3. go manages goroutines
4. fewer context switchs.
5. lesser threads
6. fast startups
7. comminicate via channels


## waitGroups

is simply a counter that have a special behaviour when they a value equal to zero.
we can increase it or decrease it

waitGroups Api
func (wg waitGroup) Add(int delta) -> increments counter by delta
func (wg waitGroup) Done() -> decrements counter by one
func (wg waitGroup) Wait() -> this is the special behaviour, this simply blocks wherever its called untill counter becoms zero




## channels

