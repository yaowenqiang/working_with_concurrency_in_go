package main
import (
	"fmt"
	"sync"
	"time"
)

type Philosopher  struct {
	name string
	rightFork int
	leftFork int
}

var philosophers = []Philosopher {
	{
		name: "Plato",
		leftFork: 4,
		rightFork: 0,
	},
	{
		name: "Socrates",
		leftFork: 0,
		rightFork: 1,
	},
	{
		name: "Aristotle",
		leftFork: 1,
		rightFork: 2,
	},
	{
		name: "Pascal",
		leftFork: 2,
		rightFork: 3,
	},
	{
		name: "Locke",
		leftFork: 3,
		rightFork: 4,
	},
}

var hunger = 3 // how much times does a person eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	// start the meal
	fmt.Println("Dining philosophers problems")
	fmt.Println("----------------------------")
	fmt.Println("The table is empty!")
	dine()

	// print finished message
	fmt.Println("The table is empty!")
}

func dine() {
	wg := sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks

	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(philosophers);i++ {
		fork := sync.Mutex{}
		forks[i] = &fork
	}
	// start the meal

	for i := 0; i< len(philosophers); i++ {
		// fire off a gorouting for the current philosopher
		go diningProblem(i, &wg, forks, &seated)
	}
	wg.Wait()
}
func diningProblem(philosopherIdx int, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// 获取当前哲学家
	philosopher := philosophers[philosopherIdx]

	// 每个哲学家需要进餐hunger次
	for i := 1; i <= hunger; i++ {
		// 思考一段时间
		fmt.Printf("%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		// 使用资源分级解决方案来避免死锁
		// 总是先拿起编号较小的叉子，再拿起编号较大的叉子
		var firstFork, secondFork *sync.Mutex
		var firstForkNum, secondForkNum int

		if philosopher.leftFork < philosopher.rightFork {
			firstFork = forks[philosopher.leftFork]
			secondFork = forks[philosopher.rightFork]
			firstForkNum = philosopher.leftFork
			secondForkNum = philosopher.rightFork
		} else {
			firstFork = forks[philosopher.rightFork]
			secondFork = forks[philosopher.leftFork]
			firstForkNum = philosopher.rightFork
			secondForkNum = philosopher.leftFork
		}

		// 先获取编号较小的叉子
		firstFork.Lock()
		fmt.Printf("%s picked up fork %d (smaller number).\n", philosopher.name, firstForkNum)

		// 再获取编号较大的叉子
		secondFork.Lock()
		fmt.Printf("%s picked up fork %d (larger number).\n", philosopher.name, secondForkNum)

		// 进餐
		fmt.Printf("%s is eating (meal %d/%d).\n", philosopher.name, i, hunger)
		time.Sleep(eatTime)

		// 放下编号较大的叉子
		secondFork.Unlock()
		fmt.Printf("%s put down fork %d.\n", philosopher.name, secondForkNum)

		// 放下编号较小的叉子
		firstFork.Unlock()
		fmt.Printf("%s put down fork %d.\n", philosopher.name, firstForkNum)
	}

	// 完成进餐
	fmt.Printf("%s is done eating and leaving the table.\n", philosopher.name)
	seated.Done()
}
