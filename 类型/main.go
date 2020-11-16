package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		j int = 0
		k int = 1
	)
	fmt.Println(j, k)
	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	fmt.Println("f32 is", f32, ",f64 is", f64)

	// 零值
	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi, zf, zb, zs)

	// 变量简短声明
	i := 10

	// 指针
	pi := &i // &取地址 *取值
	fmt.Println(*pi)

	// 常量
	// const name = "111"

	// iota 常量生成器
	const (
		one = iota + 1
		two
		three
		four
	)
	fmt.Println(one, two, three, four)

	// 强制类型转换
	i2s := strconv.Itoa(i)
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)
	i2f := float64(i)
	f2i := int(f64)
	fmt.Println(i2f, f2i)

	const s1 = "H2o"
	//判断s1的前缀是否是H
	fmt.Println(strings.HasPrefix(s1, "H"))
	//在s1中查找字符串o
	fmt.Println(strings.Index(s1, "o"))
	//把s1全部转为大写
	fmt.Println(strings.ToUpper(s1))

	ii := strings.Index("飞雪无情", "飞雪")
	fmt.Println("ii", ii)

	// 集合
	array := [5]string{"a", "b", "c", "d", "e"}
	array := [...]string{"a", "b", "c", "d", "e"}

}
