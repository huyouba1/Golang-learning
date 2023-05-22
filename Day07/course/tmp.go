package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Person struct {
	Name string
	Age  int
}

func RandString(n int) string {
	chars := "qwertyuiopasdfghjklzxcvbnm1234567890QWERTYUIOPASDFGHJKLZXCVBNM"
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteByte(chars[rand.Intn(len(chars))])
	}
	return builder.String()
}

func (p Person) GetName() string {
	if p.Name == "" {
		return RandString(5)
	}
	return p.Name
}

func (p Person) GetAge() int {
	if p.Age == 0 {
		return rand.Intn(100)
	}
	return p.Age
}

func main() {
	person := &Person{Name: "aaa", Age: 0}
	fmt.Println(person.GetName())
	fmt.Println(person.GetAge())
}
