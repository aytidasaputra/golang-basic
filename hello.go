package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

func main() {

	// cetak hello
	// fmt.Println("Hello World")

	// variable
	// var a int = 10
	// var b int = 0
	// b = 19

	// fmt.Println(a + b)

	// var c string = "Hello"
	// var d string = "Aditya"
	// fmt.Printf("Yooo... %s %s!\n", c, d)
	// fmt.Println("halo", c, d+"!")

	// var firstName string = "john"
	// lastName := "wick"
	// fmt.Printf("halo %s %s!\n", firstName, lastName)

	// _ = "belajar Golang"
	// _ = "Golang itu mudah"
	// name, _ := "john", "wick"

	// fmt.Println(name)

	// name := new(string)
	// fmt.Println(name)  // 0x20818a220
	// fmt.Println(*name) // ""

	// var decimalNumber = 2.62
	// fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	// fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// var i = 0
	// for {
	// 	fmt.Println("Angka", i)
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}
	// 	fmt.Println("Angka", i)
	// }

	// 	for i := 0; i < 5; i++ {
	// 		for j := i; j < 5; j++ {
	// 			fmt.Print(j, " ")
	// 		}
	// 		fmt.Println()
	// 	}

	// outerLoop:
	// 	for i := 0; i < 5; i++ {
	// 		for j := 0; j < 5; j++ {
	// 			if i == 3 {
	// 				break outerLoop
	// 			}
	// 			fmt.Print("matriks [", i, "][", j, "]", "\n")
	// 		}
	// 	}

	// array horizontal
	// var buah [4]string

	// buah = [4]string{"apel", "melon", "rambutan", "jambu"}

	// buah1 := []string{
	// 	"apel",
	// 	"melon",
	// 	"rambutan",
	// 	"jambu",
	// }

	// buah2 := []string{
	// 	"stroberi",
	// 	"nanas",
	// }
	// fmt.Println(buah)
	// fmt.Println(buah1)

	// for i := 0; i < len(buah1); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, buah1[i])
	// }

	// for _, buah := range buah1 {
	// 	fmt.Printf("nama buah : %s\n", buah)
	// }

	// fmt.Println(buah1[0:2])

	// buahs := append(buah1, "nanas")
	// fmt.Println(buahs)

	// var chicken = map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// 	"maret":    34,
	// 	"april":    67,
	// }

	// for key, val := range chicken1 {
	// 	fmt.Println(key, " \t:", val)

	// }

	// var value, isExist = chicken["april"]

	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Print("Data is not exist")
	// }

	// var chickens = []map[string]string{
	// 	map[string]string{"name": "chicken blue", "gender": "male"},
	// 	map[string]string{"name": "chicken red", "gender": "male"},
	// 	map[string]string{"name": "chicken yellow", "gender": "female"},
	// }
	// for _, chicken := range chickens {
	// 	fmt.Println(chicken["gender"], chicken["name"])
	// }

	// var data = []map[string]string{
	// 	{"name": "chicken blue", "gender": "male", "color": "brown"},
	// 	{"address": "mangga street", "id": "k001"},
	// 	{"community": "chicken lovers"},
	// }

	// for _, data1 := range data {
	// 	fmt.Println(data1["address"], data1["name"])
	// }

	// var names = []string{"john", "saputra"}
	// printMessage("Heii", names)

	// rand.Seed(time.Now().Unix())
	// var randomValue int
	// randomValue = randoWithRange(2, 10)
	// fmt.Println("random number:", randomValue)
	// randomValue = randoWithRange(2, 10)
	// fmt.Println("random number:", randomValue)
	// randomValue = randoWithRange(2, 10)
	// fmt.Println("random number:", randomValue)

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, "-")
	fmt.Println(message, nameString)
}

func randoWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)

}

func calculated(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d/2, 2)

	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference

}
