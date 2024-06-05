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

func doComplex() *pb.Complex {
	fmt.Println("Hello, World!")
	return &pb.Complex{

		OneDummy: &pb.Dummy{
			Id:   1,
			Name: "Dummy",
		},
		MultipleDummies: []*pb.Dummy{
			&pb.Dummy{
				Id:   1,
				Name: "Dummy",
			},
			&pb.Dummy{
				Id:   2,
				Name: "Dummy",
			},
		},
	}
}

func doEyecColor() *pb.Enumeration {

	return &pb.Enumeration{
		EyeColor: pb.Color_RED,
	}
}

func doShape() *pb.Shape {
	return &pb.Shape{
		Shape: &pb.Shape_Circle{
			Circle: &pb.Circle{
				Radius: 10,
			},
		},
	}
}

// Function to calculate the area of a shape
func CalculateArea(s *pb.Shape) float64 {
	switch s.Shape.(type) {
	case *pb.Shape_Circle:
		circle := s.GetCircle()
		return float64(3.14 * circle.Radius * circle.Radius)
	case *pb.Shape_Rectangle:
		rectangle := s.GetRectangle()
		return float64(rectangle.Width * rectangle.Height)
	case *pb.Shape_Triangle:
		triangle := s.GetTriangle()
		return float64(0.5 * triangle.Base * triangle.Height)
	default:
		return 0
	}
}

func doMaps() *pb.Student {
	return &pb.Student{
		Scores: map[string]int32{
			"Maths":   90,
			"Science": 80,
			"English": 70,
		},
	}
}

func main() {

	// fmt.Printf("%v \n", doSimple())
	// fmt.Printf("%v \n", doComplex())
	fmt.Println(doEyecColor())
	fmt.Println(doShape())

	circle := &pb.Circle{Radius: 5}
	shape := &pb.Shape{Shape: &pb.Shape_Circle{Circle: circle}}
	area := CalculateArea(shape)
	fmt.Println("Area of circle:", area)

	rectangle := &pb.Rectangle{Width: 4, Height: 6}
	shape = &pb.Shape{Shape: &pb.Shape_Rectangle{Rectangle: rectangle}}
	area = CalculateArea(shape)
	fmt.Println("Area of rectangle:", area)

	triangle := &pb.Triangle{Base: 3, Height: 4}
	shape = &pb.Shape{Shape: &pb.Shape_Triangle{Triangle: triangle}}
	area = CalculateArea(shape)
	fmt.Println("Area of triangle:", area)
	fmt.Println(doMaps())
	data := doMaps()
	fmt.Println(data.Scores)

}
