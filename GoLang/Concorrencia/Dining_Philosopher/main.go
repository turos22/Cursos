package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Philosofer is a struct para armazenar infos acerca dos filosofos
type Philosopher struct {
	name      string
	rigthFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", rigthFork: 0, leftFork: 4},
	{name: "Socrates", rigthFork: 1, leftFork: 0},
	{name: "Aristotle", rigthFork: 2, leftFork: 1},
	{name: "Pascal", rigthFork: 3, leftFork: 2},
	{name: "Locke", rigthFork: 4, leftFork: 3},
}

var hunger = 3 //quantas vezes tem de comer
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex sync.Mutex
var OrderFinished []string

func main() {
	fmt.Println("Welcome to the Dining Philosopher`s problem")
	fmt.Println("-----------------------------------------")
	fmt.Println("The table is empty.")

	dine()

	fmt.Println("The table is empty.")

}



func dine() {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second
	sleepTime = 0 * time.Second
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}
	//start the meal
	for i := 0; i < len(philosophers); i++ {
		go diningProblema(philosophers[i], wg, forks, seated)
	}
	wg.Wait()

	fmt.Println("The table is empty.")
	fmt.Println("Order that they left: " + strings.Join(OrderFinished, ", "))

}

func diningProblema(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {
		//Trava a execucao caso o garfo esteja em lock

		if philosopher.leftFork > philosopher.rigthFork {
			forks[philosopher.rigthFork].Lock()
			fmt.Printf("\t %s takes the right fork\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t %s takes the left fork\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t %s takes the left fork\n", philosopher.name)
			forks[philosopher.rigthFork].Lock()
			fmt.Printf("\t %s takes the right fork\n", philosopher.name)			
		}	

		fmt.Printf("\t %s has both fork and is eating \n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t %s is thinking \n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.rigthFork].Unlock()
		forks[philosopher.leftFork].Unlock()
	
		fmt.Printf("\t %s put down the forks \n", philosopher.name)
	}

	fmt.Println(philosopher.name, "is done satisfied")
    fmt.Println(philosopher.name, "left the table")

	orderMutex.Lock()
	OrderFinished = append(OrderFinished, philosopher.name)
	orderMutex.Unlock()
}
