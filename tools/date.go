package date

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type BitFlag int

const (
	Active  BitFlag = 1 << iota // 1 << 0 == 1
	Send                        // 1 << 1 == 2
	Receive                     // 1 << 2 == 4
)

//为 int 定义一个别名类型 Day，定义一个字符串数组它包含一周七天的名字，为类型 Day 定义 String() 方法，它输出星期几的名字。
// 使用 iota 定义一个枚举常量用于表示一周的中每天（MO、TU...）。

type Day int

const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

//func main() {
//	fmt.Println(KB)
//	fmt.Println(MB)
//	fmt.Println(GB)
//	fmt.Println(TB)
//	flag := Active | Send // == 3
//	fmt.Println(flag)
//
//	var th Day = 3
//	fmt.Printf("The 3rd day is: %s\n", th)
//	var day = SU
//	fmt.Println(day)
//	fmt.Println(0, MO, 1, TU)
//}
