package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/thetinshusasi/practice_go_grpc/proto"
	"google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
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

func writeToFile(fname string, pb proto.Message) {
	// Write to file
	//...

	out, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't marshel or serialize to bytes", err)
		return
	}

	if err := os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)

		return
	}

	fmt.Println("Wrote to file", fname)

}

func readFromFile(fname string, pb proto.Message) {
	// Read from file
	//...
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can't read from file", err)
		return
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Can't unmarshal from bytes", err)
		return
	}

	fmt.Println("Read from file", fname)

}

func DoFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)

}

// Function to convert protobuf message to JSON
func protoToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		Multiline: true,
	}
	jsonBytes, err := marshaler.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// Function to convert JSON to protobuf message
func jsonToProto(jsonData string) (*pb.Simple, error) {
	var message pb.Simple
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	err := unmarshaler.Unmarshal([]byte(jsonData), &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func main() {

	// fmt.Printf("%v \n", doSimple())
	// fmt.Printf("%v \n", doComplex())
	// fmt.Println(doEyecColor())
	// fmt.Println(doShape())

	// circle := &pb.Circle{Radius: 5}
	// shape := &pb.Shape{Shape: &pb.Shape_Circle{Circle: circle}}
	// area := CalculateArea(shape)
	// fmt.Println("Area of circle:", area)

	// rectangle := &pb.Rectangle{Width: 4, Height: 6}
	// shape = &pb.Shape{Shape: &pb.Shape_Rectangle{Rectangle: rectangle}}
	// area = CalculateArea(shape)
	// fmt.Println("Area of rectangle:", area)

	// triangle := &pb.Triangle{Base: 3, Height: 4}
	// shape = &pb.Shape{Shape: &pb.Shape_Triangle{Triangle: triangle}}
	// area = CalculateArea(shape)
	// fmt.Println("Area of triangle:", area)
	// fmt.Println(doMaps())
	// data := doMaps()
	// fmt.Println(data.Scores)

	message := &pb.Simple{
		Id:         1,
		IsSimple:   true,
		Name:       "Simple",
		SampleList: []int32{1, 2, 3, 4, 5},
	}

	// DoFile(message)

	jsonData, err := protoToJSON(message)
	if err != nil {
		log.Fatalf("Error converting to JSON: %v", err)
	}

	// Print JSON data
	fmt.Println("Protobuf message converted to JSON:")
	fmt.Println(jsonData)

	// Convert JSON data back to protobuf message
	decodedMessage, err := jsonToProto(jsonData)
	if err != nil {
		log.Fatalf("Error converting to protobuf: %v", err)
	}

	// Print decoded protobuf message
	fmt.Println("JSON data converted back to Protobuf message:")
	fmt.Println(decodedMessage)

}
