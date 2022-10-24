package main

// 测试Go的函数式编程 有点冗余
type Handler func(input string)

func Input(input string, handler Handler) {
	handler(input)
}

//func main() {
//	//此处的func即时已经被第5行的type定义过 仍然需要重写一遍 没有对应的语法糖吗？
//	Input("zzzz", func(input string) {
//		fmt.Println("zzzzz")
//	})
//}
