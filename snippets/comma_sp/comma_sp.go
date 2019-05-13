package main


// 加上最后可选的 ,
func f1(a int, b string,) (x bool, y int,) {
	return true, 1
}

var f2 func (a int) (x bool)
// 下面情况必须加 , 才能通过
var f3 func (a int,
	) (x bool,
	)


func main() {
	// 加上最后可选的 ,
	var _ = []int{2, 3,}
	// 最后的 , 必须 /// 不过也没人这样写吧
	var _ = []int{2, 3,
		}

	f1(1, "a")
	f1(2, "b",
		)
}
