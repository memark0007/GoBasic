package main //โปรแกรมจะเริ่มต้น ทำที่ main ก่อน
// import Package มาใช้งาน
import (
	"fmt"
	// "GoBasic/goPackage" //ที่มันพังน่าจะเป็นเพราะ file มันอยู่บน could
)

func main() {
	// Go language is C++ combine Python

	// // นิยามตัวแปร
	// var name string
	// name = "Kittisak Lamnoi"
	// var age int = 21
	// score := 98.3
	// fmt.Println("My name is ", name)
	// fmt.Printf("My age is %v\n", age)
	// fmt.Println("My score is ", score)
	// fmt.Println("Hello Go Programming")

	// // รับค่าจากผู้ใช้
	// var studentCode string
	// fmt.Print("ป้อนรหัสนักศึกษา = ")
	// fmt.Scanf("%s", &studentCode)
	// fmt.Println("สวัสดี ", studentCode)

	// // Array ไม่ค่อยยืดยุ่น
	// numbers := [...]int{1, 2, 3, 4, 5}
	// fmt.Println(numbers)

	// // Slice คล้ายๆ list ใน Python
	// names := []string{"สมหมาย", "แก้วตา"}
	// names = append(names, "สมปอง")
	// fmt.Println(names[:])

	// // map [key]value
	// coin := map[string]int{"BTC": 10, "ETH": 5}
	// fmt.Println(coin)

	// // for loop
	// for count := 1; count <= 3; count++ {
	// 	fmt.Println("count = ", count)
	// }
	// numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// // if not used the index : sign _
	// for _, value := range numbers {
	// 	fmt.Println("value = ", value)
	// }

	// // function
	// showMessage("Mark", 21)
	// result, check := summation(100, 200)
	// fmt.Println("Result = ", result)
	// fmt.Println("Check = ", check)
	// result2 := summation2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// fmt.Println("result2 = ", result2)

	// // Sructure เก็บข้อมูลต่างชนิดกัน หรือเป็นการสร้างข้อมูลชนิดใหม่ขึ้นมาเอง
	// Product1 := Product{name: "ปากกา", price: 50.5, catagory: "เครื่องเขียน", discount: 10}
	// fmt.Println(Product1)

	// // Package การแยก file ออกเป็นกลุ่มต่างๆ ถ้าจะใช้ก็ไปดึงมา เพื่อ code จะได้มีน้ำหนักเบา
	// fmt.Println(calculator.Add(1, 2))
}

func showMessage(name string, age int) int {
	fmt.Println("hello", name)
	fmt.Println("your age is", age)
	return age
}
func summation(num1, num2 int) (int, string) {
	total := num1 + num2
	status := ""
	if total%2 == 0 {
		status = "เลขคู่"
	} else {
		status = "เลขคี่"
	}
	return total, status
}

// paramitor ไม่จำกัดจำนวน
func summation2(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

type Product struct {
	name     string
	price    float64
	catagory string
	discount int
}
