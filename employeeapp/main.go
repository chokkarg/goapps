package main

import (
	"strings"
	"strconv"
	"os"
	"bufio"
	"fmt"
)

type person struct {
	name string
	age int
}

func (p *person) setName(name string){
	p.name = name
}

func (p *person) setAge(age int){
	p.age = age
}

func (p person) getName() string{
	return p.name
}

func (p person) getAge() int {
	return p.age
}

func main()  {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	name,_ := reader.ReadString('\n')
	
	fmt.Println("Enter your age")
	age,_ := reader.ReadString('\n')
	ageRemovedPrefix :=	strings.TrimSuffix(age, "\r\n")
	ageValue,_ :=  strconv.Atoi(ageRemovedPrefix)
	fmt.Println("age value" , ageValue)

	p := person{}

	p.setName(name)
	p.setAge(ageValue)

	fmt.Println(p.getName())
	fmt.Println(p.getAge())
	

}