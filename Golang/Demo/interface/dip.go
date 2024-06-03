package main

import (
	"fmt"
	"reflect"
)

type Aito struct{}

func (p *Aito) Run() {
	fmt.Println("Aito is running...")
}

type BYD struct{}

func (p *BYD) Run() {
	fmt.Println("BYD is running...")
}

type Nico struct{}

func (p *Nico) DriveAito(aito *Aito) {
	fmt.Println("Nico drive Aito")
	aito.Run()
}

func (p *Nico) DriveBYD(byd *BYD) {
	fmt.Println("Nico drive BYD")
	byd.Run()
}

func (p *Nico) Drive(car Car) {
	fmt.Println("Nico drive", reflect.TypeOf(car).Elem().Name())
	car.Run()
}

type Hime struct{}

func (p *Hime) DriveAito(aito *Aito) {
	fmt.Println("Hime drive Aito")
	aito.Run()
}

func (p *Hime) DriveBYD(byd *BYD) {
	fmt.Println("Hime drive BYD")
	byd.Run()
}

func (p *Hime) Drive(car Car) {
	fmt.Println("Hime drive", reflect.TypeOf(car).Elem().Name())
	car.Run()
}

// 添加车辆，要修改每一个司机
// 添加司机，要实现所有车辆的 Drive

type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

type Andy struct{}

func (p *Andy) Drive(car Car) {
	fmt.Println("Andy drive", reflect.TypeOf(car).Elem().Name())
	car.Run()
}

func main() {
	var (
		nico = &Nico{}
		hime = &Hime{}
		andy = &Andy{}

		aito = &Aito{}
		byd  = &BYD{}
	)

	nico.DriveAito(aito)
	nico.DriveBYD(byd)

	hime.DriveAito(aito)
	hime.DriveBYD(byd)

	fmt.Println()

	DriveCar(nico, aito)
	DriveCar(nico, byd)
	DriveCar(hime, aito)
	DriveCar(hime, byd)
	DriveCar(andy, aito)
	DriveCar(andy, byd)
}

func DriveCar(driver Driver, car Car) {
	driver.Drive(car)
}
