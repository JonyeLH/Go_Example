package main

import "fmt"

//闭包，相当于就是匿名函数
//形式：func function_name() func() type{···}
//运行方法：
//1、在匿名函数最后添加()，调用时将自动运行
//func function_name() func() type{···}()
//2、在匿名函数赋值后，在赋值对象后跟个()
//Function := function_name()
//Function()

//Anonymous functions are useful when you want to define a function inline without having to name it.

/*
This function intSeq returns another function, which we define anonymously in the body of intSeq.
The returned function closes over the variable i to form a closure.
 */
func intSeq() func() int{
	i := 0
	return func() int {
		i++
		return i
	}
}
func main(){
	nextInt := intSeq()

	/*
	We call intSeq, assigning the result (a function) to nextInt.
	This function value captures its own i value, which will be updated each time we call nextInt.
	See the effect of the closure by calling nextInt a few times.
	 */
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	/*
	To confirm that the state is unique to that particular function, create and test a new one.
	 */
	newInts := intSeq()
	fmt.Println(newInts())
}