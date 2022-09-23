package main

//import package
import (
	"fmt"
	"os"
	"strconv"
)

// struct person
type Person struct {
	//inisialisasi dalam struct
	absen     string
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

// function main
func main() {

	/*
		package os berguna untuk menyimpan data dalam bentuk array dengan pemisah adalah tanda spasi.
		contoh :  go run biodata.go 1 maka akan menghasilkan data array dalam bentuk string
		=> []string{"C:\\Users\\elang\\AppData\\Local\\Temp\\go-build3150221260\\b001\\exe\\biodata.exe", "1"} (tampilan %#v Go-syntax representation of the value)
		=> [C:\Users\elang\AppData\Local\Temp\go-build921977684\b001\exe\biodata.exe 1] (tampilan %s atau string)
		array ke 0 adalah "C:\\Users\\elang\\AppData\\Local\\Temp\\go-build3150221260\\b001\\exe\\biodata.exe"
		array ke 1 adalah "1"
	*/

	var argsRaw = os.Args

	//memasukkan data argsRaw di array 1 ke args
	var args = argsRaw[1]

	//merubah string ke integer
	num, err := strconv.Atoi(args)

	//err digunakan untuk menampung isian bernilai nil
	_ = err

	//memanggil function biodata dengan parameter num
	biodata(num)
}

// function biodata dengan parameter arr bertipe data integer
func biodata(arr int) {

	//inisialisasi variabel people beserta isi datanya
	var people = []Person{ //memanggil struct Person
		{absen: "1", nama: "Elang", alamat: "Sidoarjo", pekerjaan: "Backend", alasan: "Ingin menambah pengetahuan"},
		{absen: "2", nama: "Eka", alamat: "Surabaya", pekerjaan: "Frontend", alasan: "Ingin bisa bahasa go"},
		{absen: "3", nama: "Marga", alamat: "Madiun", pekerjaan: "Fullstack", alasan: "Menambah ilmu"},
		{absen: "4", nama: "Putra", alamat: "Jembrana, Bali", pekerjaan: "DevOps", alasan: "Mendapat ilmu bahasa go"},
		{absen: "5", nama: "Artup", alamat: "Situbondo", pekerjaan: "IoT", alasan: "Mendapatkan ilmu tentang bahasa go"},
	}

	fmt.Println("=============================================")
	//print/menampilkan data pada variabel people
	fmt.Printf(
		"=> Absen\t: %s\n=> Nama\t\t: %s\n=> Alamat\t: %s\n=> Pekerjaan\t: %s\n=> Alasan\t: %s\n",
		people[arr-1].absen, people[arr-1].nama, people[arr-1].alamat, people[arr-1].pekerjaan, people[arr-1].alasan) // arr-1 karena dalam inputan dimulai dari 1, sedangkan array pertama merupakan [0]
	fmt.Println("=============================================")
}
