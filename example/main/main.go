package main

import (
	"github.com/Hidayathamir/logrusy"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "hidayat", Age: 24}
	logrusy.Info(u)
	// {"level":"info","msg":"{hidayat 24}","parsed_msg":{"Name":"hidayat","Age":24},"time":"2023-12-06T00:48:33+07:00"}

	numbers := []int{0, 1, 2, 3, 4, 5}
	logrusy.Info(numbers)
	// {"level":"info","msg":"[0 1 2 3 4 5]","parsed_msg":[0,1,2,3,4,5],"time":"2023-12-06T00:48:33+07:00"}
}
