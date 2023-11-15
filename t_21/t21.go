package main

import "fmt"

type Person struct{}

func (person *Person) checkData(p Phone) int {
	return p.checkData()
}

type Phone interface {
	checkData() int
}

type SamsungPhone struct {
	ProccChar int
}

func (s *SamsungPhone) checkData() int {
	s.ProccChar = 10
	fmt.Println("default param samsung ", s.ProccChar)
	return s.ProccChar
}

type Iphone struct {
	ProccChar int
}

func (i *Iphone) checkData() int {
	i.ProccChar = 15
	fmt.Println("default param iphone ", i.ProccChar)
	return i.ProccChar
}

type iphoneAdapter struct {
	iphone *Iphone
}

func (i *iphoneAdapter) checkData() int {
	i.iphone.ProccChar = i.iphone.checkData()
	if i.iphone.ProccChar > 10 {
		i.iphone.ProccChar = 9
		return 9
	} else {
		return i.iphone.ProccChar
	}
}

func main() {
	person := &Person{}
	samsung := &SamsungPhone{}
	person.checkData(samsung)

	iphone := &Iphone{}
	iphoneA := &iphoneAdapter{
		iphone: iphone,
	}

	v := person.checkData(iphoneA)
	fmt.Println("После использования адаптера iphone = ", v)
}
