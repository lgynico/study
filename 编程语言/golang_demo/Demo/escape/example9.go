package main

type Student struct {
	name string
}

func NewStudent(name string) Student {
	return Student{
		name: name,
	}
}

func (p *Student) Name() string {
	return p.name
}

func main() {
	stu := NewStudent("Nico")
	pStu := &stu
	println(pStu)
}
