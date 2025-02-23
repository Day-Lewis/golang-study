package main

import (
	"fmt"
	"math"
	"unicode"
)

/*
	数值类型：整数、浮点数
	布尔类型：bool，值为true或false
	字符类型：byte、rune
	字符串类型：string
	其他类型：数组、指针、结构体、channel、func、slice、any、map
*/

// 无符号整型范围 0 ~ 2^n - 1
func unsignedIntegers() {
	var u uint = math.MaxUint //	uint默认是uint64
	var u8 uint8 = math.MaxUint8
	var u16 uint16 = math.MaxUint16
	var u32 uint32 = math.MaxUint32
	var u641 uint64 = 1
	var uZero uint64 = math.MaxUint64 + u641	//	上限溢出，变为0
	var u64 uint64 = math.MaxUint64
	fmt.Println(u, u8, u16, u32, uZero ,u64)
}

// 有符号整型范围 -2^(n-1) ~ 2^(n-1) - 1
func integers() {
	var i int = math.MaxInt //	int默认是int64
	var i8 int8 = math.MaxInt8
	var i16 int16 = math.MaxInt16
	var i32 int32 = math.MaxInt32
	var i64 int64 = math.MaxInt64
	fmt.Println(i, i8, i16, i32, i64)
}

// float32范围约1.4e-45 到 约3.4e38
//
//	float64范围约4.9e-324 到 约1.8e308
func floats() {
	var f32 float32 = math.MaxFloat32
	var f64 float64 = math.MaxFloat64
	fmt.Println(f32, f64)

	floatStr1 := 3.2
	fmt.Printf("%.2f\n", floatStr1)

	//	很大或很小的数，最好使用科学计数写法
	var avogadro = 6.02214129e23
	var planck = 6.62606957e-34
	fmt.Printf("%f, %.35f\n", avogadro, planck)
}

func bools() {
	var truth bool = true
	var falses bool = false
	fmt.Println(truth, falses)
}

// 字符类型：byte、rune
// byte 等价 uint8, 代表ascii码的一个字符
// rune 等价 int32, 代表一个unicode字符，处理中文、日文、其他符合字符时需要用到rune。
func char_byte() {
	//	ascii码一个字符占一个字节
	var ch byte = 'A'
	var chs byte = 65      //	ascii码中A的值时65，所以可以这么定义
	var chs2 byte = '\x41' //	16进制表示
	var chs3 byte = '\101' //	8进制表示
	fmt.Println(ch, chs, chs2, chs3)
}

func char_unicode() {
	var ni rune = 20320
	var hao rune = '\u597D'
	var hao1 = rune('好')
	//	%c表示字符，%U表示unicode码
	fmt.Printf("%c, %c, %U\n", ni, hao, hao1)
	a := unicode.IsLetter(ni)  //	判断是否是字母
	b := unicode.IsDigit(ni)   //	判断是否是数字
	c := unicode.IsSpace(hao1) //	判断是否是空白字符
	fmt.Println(a, b, c)
}

func str() {
	var str string = "Hello, World"
	fmt.Println(str)
	// str[0] = "1" //	字符串内容是不可变的，除非重新赋值
	str = "Hello, 世界"	//	注意这里用的是=，而不是:=
	fmt.Println(str)
}

func typeConversion() {
	var a int = 10
	var b float64 = 3.14
	//	go必须显示转换类型，不支持隐式转换
	fmt.Println("a + b = ", float64(a) + b)

	//	类型转换正确情况是从小类型转换大类型
	//	小类型 转换 大类型会丢失精度
	//	另外不同地层类型无法相互转换，例如bool和int
}

//	string是不可变的，但可以通过转为[]byte或[]rune来修改
func modifyString() {
	s1 := "localhost:8080"
	strByte := []byte(s1)	//	string转[]byte

	//	修改字符串
	strByte[len(s1)-1] = '8'
	fmt.Println(strByte)

	//	[]byte转string
	s2 := string(strByte)
	fmt.Println(s2)

	s3 := "你好，世界"
	runeStr := []rune(s3)
	runeStr[0] = '我'
	fmt.Println(runeStr)

	s4 := string(runeStr)
	fmt.Println(s4)
}


func main() {
	unsignedIntegers()
	integers()
	floats()
	bools()
	char_byte()    //	字符类型
	char_unicode() //	字符类型
	str()
	typeConversion()
	modifyString()
}
