# Go make it simple to build 
1. fan-in model
2. fan-out model

### whole idea is that we can balance out level of concurrency to make our app/programs effecient.
<br/>

### In go workers -> goroutines and they use channels to communicate to implement concurrency.
<br/>
<br/>





# goroutine

is a function executing concurrently with other goroutines in the same address spaces (i.e within the same program/process).
it is lightweight costing little more than the allocation of stack space.


# few more points on goroutine
1. sheduled by go runtime // scheduler
2. lighter than threads
3. go manages goroutines
4. fewer context switchs.
5. lesser threads
6. fast startups
7. comminicate via channels
<br/>
<br/>


# waitGroups

is simply a counter that have a special behaviour when they a value equal to zero.
we can increase it or decrease it

# waitGroups Api
```
func (wg waitGroup) Add(int delta) -> increments counter by delta
func (wg waitGroup) Done() -> decrements counter by one
func (wg waitGroup) Wait() -> this is the special behaviour, this simply blocks wherever its called untill counter becoms zero
```
<br/>
<br/>


# Demo:

```
func main(){

    // anonymous function
    // to make a func to goroutine is to add go key word in front of it.
    go func(){
        fmt.Println("async thing")
    }()
}
```
# Explaination.
1.here we are creating a goroutine.

2.we are making a request to go scheduler to schedule the goroutine in the future.

3.but it gets terminated as soon as the we exith the main function.
the go runtime get notified but did not get the time to schedule the goroutine.

4.for solve the same we use waitGroups.

```
func main(){

    var wg sync.WaitGroup 
    wg.Add(1)

    go func(){
        fmt.Println("async thing")
        wg.Done()
    }()

    wg.Wait()
}
```




# channels

### Help to synchronize and to communicate b/w goroutines.

```
ch := make(chan <type>)
```

Operations:
1. send message to a channel
2. read message from a channel
<br>

(similar to a message queue)
<br>
<br>


### sending message to a channel
```
ch <- "Hello!!!"
```
<br>


### reading message to a channel
```
msg := <- ch
```
<br>


## Note: Block until complementory operation is ready.
