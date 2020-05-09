package yesterday

import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person) SetAge(age int)  {
	p.age = age
}

func (p *Person) SetName(name string)  {
	p.name = name
}

func (p *Person) sayHi() {
	fmt.Printf("My name is %s, and I'm %d years old\n", p.name, p.age)
}
