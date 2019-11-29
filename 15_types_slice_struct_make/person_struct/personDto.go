package person_struct

type Person struct {
	Name     string
	lastName string
	age      string
	height   string
}

func NewPerson(name string, lastName string, age string, height string) *Person {
	return &Person{Name: name, lastName: lastName, age: age, height: height}
}




func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) GetHeight() string {
	return p.height
}

func (p *Person) SetHeight(height string) {
	p.height = height
}

func (p *Person) GetLastName() string {
	return p.lastName
}

func (p *Person) SetLastName(lastName string) {
	p.lastName = lastName
}

func (p *Person) GetAge() string {
	return p.age
}

func (p *Person) SetAge(age string) {
	p.age = age
}




