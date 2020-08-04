package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Angka struct {
	angka1, angka2, angka3 float64
}

func Perkalian(v1, v2 int) int {
	return v1 * v2
}

func Pembagian(v1, v2 int) int {
	return v1 / 2
}

func Pertambahan(v1, v2 int) int {
	return v1 + v2
}

func Pengurangan(v1, v2 int) int {
	return v1 - v2
}

func Pengakaran(v1 float64) float64 {
	return math.Sqrt(v1)
}

func Pangkat(v1, v2 float64) float64 {
	pangkat := math.Pow(v1, v2)
	return pangkat
}

func LuasPersegi(v1, v2 int) int {
	// 2 * (panjang * lebar)
	luas := 2 * (v1 * v2)
	return luas
}

func luasLingkaran(v1 float64) float64 {
	//  phi * ( jari jari * diameter )
	phi := 3.14
	luas := phi * (v1 * v1)
	return luas
}

func VolumeTabung(v1, v2 float64) float64 {
	// r = luas alas
	// t = tinggi
	// V = π x r2 x t
	phi := 3.14
	volume := phi * (v1 * v1) * v2
	return volume
}

func VolumeBalok(v1, v2, v3 float64) float64 {
	// panjang * luas * tinggi
	// V = p x l x t
	volume := (v1 * v2 * v3)
	return volume
}

func VolumePrisma(v1, v2, v3 float64) float64 {
	// L = Luas Alas + Luas Tutup + Luas Selimut
	v := v1 * v2 * v3
	return v
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var input string = ""
	for input != "99" {
		fmt.Println("CALCULATOR SAMPLE")
		fmt.Println("1. Perkalian")
		fmt.Println("2. Pembagian")
		fmt.Println("3. Pertambahan")
		fmt.Println("4. Pengurangan")
		fmt.Println("5. AkarKuadrat")
		fmt.Println("6. Perpangkatan")
		fmt.Println("7. Hitung Luas Persegi")
		fmt.Println("8. Hitung Luas Lingkaran")
		fmt.Println("9. Hitung Volume Tabung")
		fmt.Println("10. Hitung volume Balok")
		fmt.Println("11. Hitung volume prisma")
		fmt.Println("99. exit")

		fmt.Print("input pilihan :")
		scanner.Scan()
		input = scanner.Text()
		switch input {
		case "1":
			fmt.Println("1. Perkalian")
			fmt.Println("Masukan Angka 1 : ")
			scanner.Scan()
			v1, _ := strconv.Atoi(scanner.Text())
			fmt.Println("Masukan Angka 2 : ")
			scanner.Scan()
			v2, _ := strconv.Atoi(scanner.Text())

			kali := Perkalian(v1, v2)
			fmt.Println(kali)

			break
		case "2":
			fmt.Println("2. Pembagian")
			fmt.Println("Masukan Angka 1 : ")
			scanner.Scan()
			v1, _ := strconv.Atoi(scanner.Text())
			fmt.Println("Masukan Angka 2 : ")
			scanner.Scan()
			v2, _ := strconv.Atoi(scanner.Text())

			bagi := Pembagian(v1, v2)
			fmt.Println(bagi)

			break
		case "3":
			fmt.Println("3. Pertambahan")
			fmt.Println("Masukan Angka 1 : ")
			scanner.Scan()
			v1, _ := strconv.Atoi(scanner.Text())
			fmt.Println("Masukan Angka 2 : ")
			scanner.Scan()
			v2, _ := strconv.Atoi(scanner.Text())

			tambah := Pertambahan(v1, v2)
			fmt.Println(tambah)

			break
		case "4":
			fmt.Println("4. Pengurangan")
			fmt.Println("Masukan Angka 1 : ")
			scanner.Scan()
			v1, _ := strconv.Atoi(scanner.Text())
			fmt.Println("Masukan Angka 2 : ")
			scanner.Scan()
			v2, _ := strconv.Atoi(scanner.Text())

			kurang := Pengurangan(v1, v2)
			fmt.Println(kurang)

			break
		case "5":
			fmt.Println("5. AkarKuadrat ")
			fmt.Println("Masukan Angka : ")
			scanner.Scan()
			f, _ := strconv.ParseFloat(scanner.Text(), 64)
			akar := Pengakaran(f)
			fmt.Println(akar)

			break
		case "6":
			fmt.Println("6. Perpangkatan ")
			fmt.Println("Masukan Angka : ")
			scanner.Scan()
			v1, _ := strconv.ParseFloat(scanner.Text(), 64)
			fmt.Println("Masukan Angka 2 : ")
			scanner.Scan()
			v2, _ := strconv.ParseFloat(scanner.Text(), 64)
			kurang := Pangkat(v1, v2)
			fmt.Println(kurang)

			break
		case "7":
			fmt.Println("7. Luas Persegi ")
			fmt.Println("Masukan Panjang : ")
			scanner.Scan()
			v1, _ := strconv.Atoi(scanner.Text())
			fmt.Println("Masukan Lebar : ")
			scanner.Scan()
			v2, _ := strconv.Atoi(scanner.Text())
			persegi := LuasPersegi(v1, v2)
			fmt.Println(persegi)

			break
		case "8":
			fmt.Println("8. Luas Lingkaran ")
			fmt.Println("Masukan Jari Jari : ")
			scanner.Scan()
			v1, _ := strconv.ParseFloat(scanner.Text(), 64)

			lingkaran := luasLingkaran(v1)
			fmt.Println(lingkaran)

			break
		case "9":
			fmt.Println("9. Volume Tabung ")
			fmt.Println("Masukan Luas Alas : ")
			scanner.Scan()
			v1, _ := strconv.ParseFloat(scanner.Text(), 64)
			fmt.Println("Masukan Tinggi : ")
			scanner.Scan()
			v2, _ := strconv.ParseFloat(scanner.Text(), 64)

			tabung := VolumeTabung(v1, v2)
			fmt.Println(tabung)

			break
		case "10":
			fmt.Println("10. Volume Balok ")
			fmt.Println("Masukan Panjang : ")
			scanner.Scan()
			v1, _ := strconv.ParseFloat(scanner.Text(), 64)
			fmt.Println("Masukan Luas : ")
			scanner.Scan()
			v2, _ := strconv.ParseFloat(scanner.Text(), 64)
			fmt.Println("Masukan Tinggi : ")
			scanner.Scan()
			v3, _ := strconv.ParseFloat(scanner.Text(), 64)
			balok := VolumeBalok(v1, v2, v3)
			fmt.Println(balok)

			break
		case "11":
			fmt.Println("11. Volume Prisma ")
			fmt.Println("Masukan Luas Alas : ")
			scanner.Scan()
			v1, _ := strconv.ParseFloat(scanner.Text(), 64)
			fmt.Println("Masukan Luas Tutup : ")
			scanner.Scan()
			v2, _ := strconv.ParseFloat(scanner.Text(), 64)
			fmt.Println("Masukan Luas Prisma : ")
			scanner.Scan()
			v3, _ := strconv.ParseFloat(scanner.Text(), 64)
			prisma := VolumePrisma(v1, v2, v3)
			fmt.Println(prisma)

			break
		case "99":
			fmt.Println("good by")
			break
		default:
			fmt.Println("wrong input")
			break
		}
	}

}
