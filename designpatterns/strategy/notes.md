# Strategy Design Pattern

## Example of where strategy design pattern come to rescue.
<br/>


Lets consider a example where we have a base class ans we have many sub-classes which inherts behaviour from the base class.
<br/>

Everything works fine as we want untill we have a change in the requirments.
Change is bound to happen in software development.
<br/>
<br/>

### One in the new requirment we want few class behave differently from others.
let's go through few approches to deal with.

### 1. Behaviour is handled in the sub-classes.
1. This will work just fine.
2. But it suffers with few problems.
    1. Lot of duplicated code. Many sub-classes will have the same code.
    2. when we get a new requirment which we will for sure. we have to make changes in all tens even hunderdes of classes 

### 2. Default behaviour in the base class and over-ride behavior in subclasses which have different requirment.
1. This also will work just fine. 
<br>
2. But, if we change the behaviour in the base class this will effect all the sub-classes. 
we have copuling between the classes.

### 3. Strategy Pattern
It's all about encapsulating out the behaviour into a class. And Composition over Inheriatance.
<br>

### Composition Over Inheriatance
when you inherit the behaviour is set statically at compile time. which is highly coupled and not flexable. Where as when you extend the behaviour using composition at you are not coupled to anything and we can change the behaviour dynamically at runtime. 

Which also follows open/close principle we can later add new behaviours without even touching the exisiting code. which leads to few bugs.

### lets see a example implementation of strategy design pattern.


FlyBehavior interface
```
package fly

type FlyBehavior interface {
	Fly()
}
```

FlyWithWings FlyBehavior
```
package fly

import "fmt"

type FlyWithWings struct{}

func (flyBehavior FlyWithWings) Fly() {
	fmt.Println("Flying with wings")
}

```

FlyNoWay FlyBehavior
```
package fly

import "fmt"

type FlyNoWay struct{}

func (flyBehavior FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}
```

FlyRocketPowred FlyBehavior

```
package fly

import "fmt"

type FlyRocketPowred struct{}

func (flyBehavior FlyRocketPowred) Fly() {
	fmt.Println("Flying With Rocket Powered Wings")
}
```