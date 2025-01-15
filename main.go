package main

import (
	"fmt"
	"os"
	"reflect"
)

// 定义可执行的方法
type Methods struct{}

func (m Methods) Method1() {
	fmt.Println("Executing Method1: Hello from Method1!")
}

func (m Methods) Method2() {
	fmt.Println("Executing Method2: Hello from Method2!")
}

func (m Methods) Method3() {
	fmt.Println("Executing Method3: Hello from Method3!")
}

func main() {
	// 检查是否提供了至少一个参数
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <method_name>")
		os.Exit(1)
	}

	// 获取方法名
	methodName := os.Args[1]

	// 使用反射查找方法
	m := Methods{}
	method := reflect.ValueOf(m).MethodByName(methodName)

	if method.IsValid() {
		// 调用方法
		method.Call(nil)
	} else {
		fmt.Printf("Error: Method '%s' not found.\n", methodName)
		os.Exit(1)
	}
}
