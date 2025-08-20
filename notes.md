# Intermediate GO

## Magesh Kuppan
-tkmagesh77@gmail.com

## Schedule
| What | When |
|-------- | ---- |
| Commence | 9:30 AM |
| Tea Break | 11:00 AM (15 mins)|
| Lunch Break | 1:00 PM (1 hr) |
| Wind up   | 4:00 PM |

## Methodology
- NO powerpoint
- Discussion & Code

## Repository
- https://github.com/tkmagesh/Nutanix-IntermediateGo-Aug-2025

## Software Prerequisites
- go tools (https://go.dev/dl)
```shell
go version
```

## Go Prerequisites
- Data Types
- Variables , Contants, Iota
- Program Structure
- How to compile?
- Modules & Packages
- Programming constructs
    - if else, for, switch case
- Functions
    - Named Results
    - Higher Order Functions
    - Anonymous Functions
    - Variadic Functions
    - Deferred Functions
- Error Handling
- Panic & Recovery
- Structs
    - Struct composition
- Methods
- Type Assertion
- Interfaces
    - Interface Composition

# Agenda
## Concurrency In Go
- What is Concurrency?
### WaitGroup
- Semaphore based counter
- Has the ability to block the execution of a function until the counter becomes 0
### Concurrent Safe State
#### Detect Data Race
```shell
go run -race [app]
```

```shell
go build -race [app]
```

```shell
go test -race [app]
```

### Share memory by communicating
#### Channel Datatype
**Declaration**
```go
var ch chan int
```
**Intitilization**
```go
ch = make(chan int)
```
**Declaration & Initialization**
```go
var ch chan int = make(chan int)
// OR
var ch = make(chan int)
// OR
ch := make(chan int)
```
##### Channel Operations (<-  Operator)
**Send Operation**
```go
ch <- 100
```
**Receive Operation**
```go
data := <-ch
```

## Context
- Cancel Propagation
- context instances implement "context.Context" interface
    - Done()
- Creating Contexts
    - context.Background()
        - meant to act as the root (top most)
    - context.WithCancel(parentCtx)
        - Programmatic cancellation
    - context.WithTimeout() & context.WithDeadline()
        - Time based cancellation
        - Also support programmatic cancellation
    - context.WithValue()
        - Meant for sharing data across context hierarchies
        - Non cancellable

## Database Programming
## Building REST APIs


