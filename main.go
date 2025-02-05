package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

// 定义可执行的方法
type Methods struct{}

func (m Methods) Method1(args ...string) {
	fmt.Printf("Executing Method1: Hello from Method1! %v\n", args)
}

func (m Methods) Method2(args ...string) {
	fmt.Printf("Executing Method2: Start Sleep.")
	time.Sleep(20 * time.Second)
	fmt.Printf("Executing Method2: Hello from Method2! %v\n", args)
}

func (m Methods) Method3(args ...string) {
	fmt.Println("Executing Method3: Start Sleep.")
	time.Sleep(5 * time.Second)
	fmt.Println("Executing Method3: Start Panic.")
	panic("test panic")
}

func main() {
	// 检查是否提供了至少一个参数
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <method_name>")
		os.Exit(1)
	}

	// 获取方法名
	methodName := os.Args[1]
	// 入参
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	// 使用反射查找方法
	m := Methods{}
	method := reflect.ValueOf(m).MethodByName(methodName)

	if method.IsValid() {
		var values []reflect.Value
		for _, arg := range args {
			values = append(values, reflect.ValueOf(arg))
		}
		// 调用方法
		method.Call(values)
	} else {
		fmt.Printf("Error: Method '%s' not found.\n", methodName)
		os.Exit(1)
	}
}
