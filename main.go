package main

import (
	"fmt"

	pb "github.com/thetinshusasi/practice_go_grpc/proto"
)

func doSimple() *pb.Simple {
	fmt.Println("Hello, World!")
	return &pb.Simple{
		Id:         1,
		IsSimple:   true,
		Name:       "Simple",
		SampleList: []int32{1, 2, 3, 4, 5},
	}

}
func main() {

	fmt.Printf("%v", doSimple())

}
