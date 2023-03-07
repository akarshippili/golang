# Strategy Design Pattern

## Example of where strategy design pattern come to rescue.
<br/>


Lets consider a example where we have a base class ans we have many sub-classes which inherts behavior from the base class.
<br/>

Everything works fine as we want untill we have a change in the requirments.
Change is bound to happen in software development.
<br/>
<br/>

### One in the new requirment we want few class behave differently from others.
let's go through few approches to deal with.

### 1. behavior is handled in the sub-classes.
1. This will work just fine.
2. But it suffers with few problems.
    1. Lot of duplicated code. Many sub-classes will have the same code.
    2. when we get a new requirment which we will for sure. we have to make changes in all tens even hunderdes of classes 

### 2. Default behavior in the base class and over-ride behavior in subclasses which have different requirment.
1. This also will work just fine. 
<br>
2. But, if we change the behavior in the base class this will effect all the sub-classes. 
we have copuling between the classes.

### 3. Strategy Pattern
It's all about encapsulating out the behavior into a class. And Composition over Inheriatance.
<br>

For example, we have different kind of ducks which have differnt flying behavior:
1. Normal duck. -> Normal Flying behavior
2. Wodden duck. -> No Flying behavior
3. Rubber duck. -> No Flying behavior

so, we encapsulate out the different flying behavior into differnet classes.

### Composition Over Inheriatance
when you inherit the behavior is set statically at compile time. which is highly coupled and not flexable. Where as when you extend the behavior using composition at you are not coupled to anything and we can change the behavior dynamically at runtime. 

Which also follows open/close principle we can later add new behaviors without even touching the exisiting code. which leads to few bugs.

### Dependency Inversion Principle
Higher-level components should not depend on lower-level components insted both higher ans lower-level components depends on abstractions. which leads to loosly-coupled systems.

Here in our example Duck is higher level-component and the flying behavior is the lower-level component. 
So, both Duck and all the different flying behaviors dependent on abstraction. (Flyingbehavior)


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

Similar to Flybehavior we have encapsulated out the Quackbehavior


QuackBehavior Interface/Abstraction on which both the lower level components (different implemntaion of Quackbehavior and Duck depends.)
```
package quack

type QuackBehavior interface {
	Quack()
}

```

QuackingBehavior

```
package quack

import "fmt"

type QuackingBehavior struct{}

func (quack QuackingBehavior) Quack() {
	fmt.Println("Quack")
}
```

MuteBehavior
```
package quack

import "fmt"

type MuteBehavior struct{}

func (quack MuteBehavior) Quack() {
	fmt.Println("I can't quack")
}
```

SqueakBehavior
```
package quack

import "fmt"

type SqueakBehavior struct{}

func (quack SqueakBehavior) Quack() {
	fmt.Println("Squeak squeak!!!")
}
```

Here is the Duck (higher-level component).
which depends on abstraction, and extends the behavior using composition and delegation.

```
package duck

import (
	"github.com/akarshippili/golang/designpatterns/strategy/fly"
	"github.com/akarshippili/golang/designpatterns/strategy/quack"
)

type Duck struct {
	FlyBehavior   fly.FlyBehavior
	QuackBehavior quack.QuackBehavior
}

func (duck *Duck) SetFlyingBehavior(fly fly.FlyBehavior) {
	duck.FlyBehavior = fly
}

func (duck *Duck) SetQuackingBehavior(quack quack.QuackBehavior) {
	duck.QuackBehavior = quack
}

func (duck *Duck) Behave() {
	duck.QuackBehavior.Quack()
	duck.QuackBehavior.Quack()
	duck.FlyBehavior.Fly()
}
```


### main.go

```
package main

import (
	"github.com/akarshippili/golang/designpatterns/strategy/duck"
	"github.com/akarshippili/golang/designpatterns/strategy/fly"
	"github.com/akarshippili/golang/designpatterns/strategy/quack"
)

func main() {

	duck := duck.Duck{
		FlyBehavior:   fly.FlyWithWings{},
		QuackBehavior: quack.QuackingBehavior{},
	}
	duck.Behave()

	// changing the flyBehavior and quackBehavior at runtime.
	duck.SetFlyingBehavior(fly.FlyNoWay{})
	duck.SetQuackingBehavior(quack.MuteBehavior{})
	duck.Behave()

	duck.SetFlyingBehavior(fly.FlyRocketPowred{})
	duck.SetQuackingBehavior(quack.SqueakBehavior{})
	duck.Behave()
}

```

### Output:

```
Quack
Quack
Flying with wings
I can't quack
I can't quack
I can't fly
Squeak squeak!!!
Squeak squeak!!!
Flying With Rocket Powered Wings
```

## References:
1. Head First Design Patterns Book
2. https://www.youtube.com/watch?v=yjZsAl13trk

### Note:
1. This is my understanding on strategy design pattern.
2. Pls, feel free to comment any suggestions.