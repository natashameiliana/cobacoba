package main

import (
	"fmt"
	"errors"
    
)

func GreetingsTemplate(name string, age int) string {
	return fmt.Sprintf("Hello gais my name is %s, i am %d years old", name, age)

}

func checkAge(name string, age int) (string, error){
	if age < 18 {
		return "", errors.New("terlalu young")
	} else {
		greetingText := GreetingsTemplate(name, age)
		return greetingText, nil
	}
}

func main (){
	greeting, err := checkAge("Nata",16)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(greeting)
	}
}