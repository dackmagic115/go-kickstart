package main

import (
	"fmt"
	"reflect"
)
// type save like interface
type Doctor struct {
	number int
	actorName string
	companions []string
}

type Animal struct {
	Name string
	Origin string
}

type Bird struct{
	Animal
	SpeedKPH float32
	Canfly bool
}

type BirdDescription struct {
	Name string `requiredmax:"100"`
	Origin string
}

func main(){
	aDoctor := Doctor {
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.number)
	bee()
	description()
}

func bee(){
	b := Bird{}
	b.Name = "Bee"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.Canfly = false
	fmt.Println(b)
}

func description(){
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}