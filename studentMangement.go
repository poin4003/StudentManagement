package main

import (
	"fmt"
)

type Student struct {
	name   string
	age    int
	class  string
	gender string
}

func findStudentByName(name string, listStudent map[string]Student) Student {
	for _, element := range listStudent {
		if name == element.name {
			return element
		}
	}

	return Student{}
}

func addStudent(listStudent map[string]Student, id string, name string, age int, class string, gender string) {
	student := Student{name, age, class, gender}
	listStudent[id] = student
}

func updateStudent(listStudent map[string]Student, id string, student Student) {
	listStudent[id] = student
}

func removeStudent(listStudent map[string]Student, id string) {
	delete(listStudent, id)
}

func filterByGender(listStudent map[string]Student, gender string) map[string]Student {
	var replaceMap = make(map[string]Student)

	for key, element := range listStudent {
		if element.gender == gender {
			replaceMap[key] = element
		}
	}
	return replaceMap
}

func main() {
	var a = map[string]Student{
		"id1": {"Pchuy", 20, "Pfp-01", "Male"},
		"id2": {"QKhang", 20, "Pfp-01", "Male"},
		"id3": {"Phuong", 20, "Pfp-01", "Female"},
		"id4": {"My", 19, "Pfp-02", "Female"},
		"id5": {"Khanh", 22, "Pfp-02", "Female"},
	}

	fmt.Print(findStudentByName("Pchuy", a))

	fmt.Println(a)
	fmt.Println("----------")

	addStudent(a, "id6", "Thao", 20, "Pfp-01", "Male")

	studentForUpdate := Student{"Pham Minh Thao", 21, "Pfp-02", "Male"}

	updateStudent(a, "id6", studentForUpdate)

	removeStudent(a, "id6")
	// fmt.Println(a)

	fmt.Print(filterByGender(a, "Female"))

}
