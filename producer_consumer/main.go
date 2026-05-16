// 生产者-消费者模式示例
// 这个程序模拟了一个披萨店的运营，展示如何使用Go的channel实现生产者-消费者模式
// 详情参考: https://en.wikipedia.org/wiki/Producer%E2%80%93consumer_problem
package main

import (
	"fmt"
	"github.com/fatih/color" // 用于彩色终端输出
	"math/rand"
	"time"
)

const NumberOfPizzas = 10 // 要制作的披萨总数

// 全局统计变量
var pizzasMade, pizzasFailed, total int

// Producer 生产者结构体
// data: 用于传递披萨订单的数据通道
// quit: 用于优雅关闭的通道，传递error通道用于错误处理
type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

// PizzaOrder 披萨订单结构体
// pizzaNumber: 披萨编号
// message: 订单状态消息
// success: 制作是否成功
type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

// Close 关闭生产者，实现优雅关闭
// 创建一个error通道并通过quit通道发送，然后等待关闭完成
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch // 将error通道发送到quit通道，通知生产者要关闭
	return <-ch  // 等待关闭完成并返回可能的错误
}

// makePizza 制作披萨的函数（生产者逻辑）
// pizzaNumber: 当前披萨编号
// 返回一个指向PizzaOrder的指针
func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++ // 披萨编号从1开始

	// 检查是否超过要制作的披萨总数
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1    // 随机制作时间1-5秒
		fmt.Printf("received order #%d\n", pizzaNumber)
		rnd := rand.Intn(12) + 1     // 随机数1-12，用于决定是否成功
		msg := ""
		success := false

		// 根据随机数决定制作是否成功（模拟真实情况）
		if rnd < 5 {
			pizzasFailed++ // 失败计数
		} else {
			pizzasMade++   // 成功计数
		}

		total++ // 总尝试次数
		fmt.Printf("Making pizza #%d, It will take %d seconds....\n", pizzaNumber, delay)

		// 模拟制作披萨的耗时
		time.Sleep(time.Duration(delay) * time.Second)

		// 根据随机数生成不同的结果消息
		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while makeing pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready", pizzaNumber)
		}
		p := PizzaOrder{
			pizzaNumber,
			msg,
			success,
		}
		return &p
	}

	// 超过要制作的披萨总数，返回空订单用于结束程序
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

// pizzeria 披萨店生产者函数，在独立goroutine中运行
// pizzaMaker: 生产者结构体指针，包含数据通道和退出通道
// 这个函数不断制作披萨并通过通道发送，直到收到退出信号
func pizzeria(pizzaMaker *Producer) {
	// 跟踪当前正在制作的披萨编号
	var i = 0

	// 无限循环，直到收到退出通知
	// 这是生产者的主循环

	// 尝试制作披萨

	for {
		// 尝试制作一个披萨
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// 尝试将披萨发送到数据通道（消费者会从这里接收）
			case pizzaMaker.data <- *currentPizza:
				// 发送成功，继续下一个披萨
			case quitChan := <-pizzaMaker.quit:
				// 收到退出信号，关闭所有通道
				close(pizzaMaker.data)
				close(quitChan)
				return // 结束生产者goroutine
			}
		}

	}

}

// main 主函数，程序的入口点
func main() {
	// 初始化随机数生成器，使用当前时间作为种子
	rand.Seed(time.Now().UnixNano())

	// 打印欢迎消息
	color.Cyan("The pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// 创建一个生产者实例
	// 生产者包含两个通道：一个用于数据，一个用于优雅关闭

	pizzaJob := &Producer{
		data: make(chan PizzaOrder), // 创建披萨订单通道
		quit: make(chan chan error), // 创建退出通道
	}

	// 在后台goroutine中运行生产者
	// 这是并发编程的关键：生产者在独立的goroutine中运行
	go pizzeria(pizzaJob)

	// 创建并运行消费者
	// 主函数充当消费者的角色，从通道中读取披萨订单

	for i := range pizzaJob.data {
		// 检查是否仍在制作披萨数量范围内
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				// 披萨制作成功，显示绿色消息
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				// 披萨制作失败，显示红色消息
				color.Red(i.message)
				color.Red("The cusomter is really mad!")
			}
		} else {
			// 制作完成所有披萨，准备关闭
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	// 打印结束消息
	color.Cyan("-----------------")
	color.Cyan("Done for the day.")
	color.Cyan("We made %d pizzas,but failed to make %d, with %d attemptes in total.", pizzasMade, pizzasFailed, total)

	// 根据失败数量给出不同的评价
	switch {
	case pizzasFailed > 9:
		color.Red("It was a awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was a okay day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day...")
	default:
		color.Green("It was a good day...")
	}
}
