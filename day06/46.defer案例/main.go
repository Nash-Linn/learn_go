package main

import "fmt"

func f1() int {
	fmt.Println("-------------------f1--------------------")
	i := 5
	defer func() {
		i++
	}()
	return i
}

func f2() *int {
	fmt.Println("-------------------f2--------------------")
	i := 5
	defer func() {
		i++
		fmt.Printf("内%p\n", &i)
	}()

	fmt.Printf("外%p\n", &i)
	return &i ///rval = i的地址 , i = 6 , ret rval
}

func f3() (result int) {
	fmt.Println("-------------------f3--------------------")
	defer func() {
		result++
	}()
	return 5 //reslut = 5;ret result(result 替换了 rval)
}

func f4() (result int) {
	fmt.Println("-------------------f4--------------------")
	defer func() {
		result++
	}()
	return result //reslut = 0;result = 1;ret result
}

func f5() (r int) {
	fmt.Println("-------------------f5--------------------")
	t := 5
	defer func() {
		t = t + 1
	}()
	return t // r = t = 5;ret r
}

func f6() (r int) {
	fmt.Println("-------------------f6--------------------")
	fmt.Println(&r)
	defer func(r int) {
		r = r + 1
		fmt.Println(&r)
	}(r)
	return 5 // r= 5;
}

func f7() (r int) {
	fmt.Println("-------------------f7--------------------")
	defer func(x int) {
		r = x + 1
	}(r)
	return 5 // r = 5; r = 0 + 1 = 1 ;ret r
}

func main() {
	fmt.Println(f1())

	fmt.Printf("f2:%d\n", *f2())

	fmt.Println(f3())

	fmt.Println(f4())

	fmt.Println(f5())

	fmt.Println(f6())

	fmt.Println(f7())
}
