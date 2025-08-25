package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Slices_func() {
	/*
		slices常用函数
	*/

	nums := []int{1, 2, 3, 4, 8, 7, 6}
	nums2 := []int{1, 2, 3, 4, 8, 7, 6}
	strs := []string{"abc", "abcd", "abcde"}
	s1 := "ABCD"
	//var s2 string = "ABCDE"
	slices.Sort(nums)

	//slices.Sort(strs)
	//slices.Min(nums)
	max_list := slices.Max(nums)
	min_list := slices.Min(nums)
	fmt.Println("获取到列表中最大的值", max_list)
	fmt.Println("获取到列表中最大的值", min_list)
	fmt.Println("比较两个列表中的值", slices.Equal(nums, nums2))
	// Reverse：针对切片操作,通过rune转化为切片
	s2 := []rune(s1)
	slices.Reverse(s2)
	// 切片执行后再返回字符串
	s3 := string(s2)
	fmt.Println("s1翻转后的字符串为", s3)

	fmt.Println(strs)
	fmt.Println(nums)
	fmt.Println("Hello World")
}
func Math_func() {
	/*
		math常用函数
	*/
	x := 2.0
	x2 := -3.0
	fmt.Println(math.Abs(x2))
	fmt.Println(math.Max(x, x2))
	fmt.Println(math.Pow(x, x2))
	fmt.Println(math.Sqrt(x))
	fmt.Println("返回不小于x的最小整数", math.Ceil(x))
	fmt.Println("返回不大于x的最大整数", math.Floor(x))
}
func String_func() {
	/*
		字符串基本操作
	*/
	var ss string = "AB_CD_EF%GHI_%KLM"
	ss2 := ss[0]
	length := len(ss)
	ss3 := &strings.Builder{} // 创建空字符串
	ss3.WriteString("NOP")
	ss3.WriteByte('R') // 追加单个字符，单引号
	ss4 := ss[1:2]
	fmt.Println(ss4)
	fmt.Println(ss2)
	fmt.Println(ss3)
	fmt.Println(length)

}

func For_func() {
	str1 := "ABCDEFG"
	// 对应ASCII表
	for i := 0; i < len(str1); i++ {
		str1_1 := str1[i]
		fmt.Println(str1_1)
	}

	for index, char := range str1 {
		fmt.Printf("索引:%d,字符%c", index, char)
	}
}

func String2_func() {
	/*
		字符串切片操作
	*/
	var str1 string = "AB_CD_EF%GHI_%KLM"
	fmt.Println(strings.Split(str1, "_"))                    // 根据指定字符分割
	fmt.Println(strings.Join(strings.Split(str1, "&"), "_")) // 使用分隔符连接成为一个新的字符串
	//[AB CD EF%GHI %KLM]
	//AB_CD_EF%GHI_%KLM
	res1 := strings.Contains(str1, "AB")   // 判断子串
	res2 := strings.HasPrefix(str1, "ABC") // 判断是否以某一字段开头
	fmt.Println("==============")
	fmt.Println(res1)
	fmt.Println(res2)
	str_to_int := "123"
	int_to_str := 234
	fmt.Println("==============")
	strconv.Atoi(str_to_int)
	strconv.Itoa(int_to_str)
	fmt.Println(strconv.Atoi(str_to_int))
	fmt.Println(strconv.Itoa(int_to_str))

}
func Slice2() {
	l := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(l))
	fmt.Println(cap(l))
	// go语言中不包含步长概念
	fmt.Println(l[:2])
}

/*
数据结构与算法
*/

func Map_func() {
	/*
		哈希表、字典操作
	*/
	// 初始化好一个字典
	m := map[string]int{"Q": 1, "W": 2}
	fmt.Println(m)

	map1 := make(map[string]int)
	map1["A"] = 1
	map1["B"] = 2
	// 获取值
	count := map1["A"]
	fmt.Printf("获取到的A值：", count)

	// 检查键值是否存在
	value, ok := map1["B"]
	if ok {
		fmt.Printf("<UNK>B<UNK>", value)
	} else {
		fmt.Println("<UNK>")
	}
	// 删除值 delete(map,key)
	delete(map1, "B")

	// 遍历字典
	m2 := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}
	for key := range m2 {
		fmt.Println(key) // 仅遍历key值
	}
	// 字典遍历是无序的，每次便利的结果可能不同
	for k, v := range m2 {
		fmt.Println("键：", k, "值：", v)
	}
	// 字典的线程是不安全的，在写入字典时需要加锁
}
