package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var tanggal [][]int

var hari [][]time.Weekday
var bulan []time.Month

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func setDay(value []int) {
	tanggal = append(tanggal, value)
}

func setDayName(value []time.Weekday) {
	hari = append(hari, value)
}

func setMonth(value time.Time) {
	bulan = append(bulan, value.Month())
}

func lastDay(value time.Time) int {
	lastDay := value.AddDate(0, 1, -1)
	return lastDay.Day()
}

func firstDay(value time.Time) int {
	return value.Day()
}

func getDay(value time.Time) time.Weekday {
	return value.Weekday()
}

func generateDay(value time.Time) {
	first := firstDay(value)
	days := getDay(value)
	last := lastDay(value)
	var block []int
	var blockDay []time.Weekday
	for first <= last {
		block = append(block, first)
		blockDay = append(blockDay, days)
		days++
		if days >= 7 {
			days = 0
		}
		first++
	}
	setDay(block)
	setDayName(blockDay)
}

func generateCalender(tahun int) {
	var year = 12
	for i := 1; i <= year; i++ {
		bulan := Date(tahun, i, 1)
		setMonth(bulan)
		generateDay(bulan)
	}

}

func printSlice(value []int) {
	for i := 0; i < len(value); i++ {
		fmt.Print("   ", value[i], "\t")
	}
	fmt.Println()
}
func printDay(value []time.Weekday) {
	for i := 0; i < len(value); i++ {
		fmt.Print(value[i], " ")
	}
	fmt.Println("\n")
}

func printCalender(tahun int) {
	fmt.Println("CALENDER TAHUN ", tahun)
	for i := 0; i < 12; i++ {
		fmt.Println("=====================", bulan[i], tahun, "=====================")
		fmt.Println("\n")
		slice1 := tanggal[i][:7]
		slice2 := tanggal[i][7:14]
		slice3 := tanggal[i][14:21]
		slice4 := tanggal[i][21:28]
		slice5 := tanggal[i][28:]
		day1 := hari[i][:7]
		day2 := hari[i][7:14]
		day3 := hari[i][14:21]
		day4 := hari[i][21:28]
		day5 := hari[i][28:]

		printSlice(slice1)
		printDay(day1)
		printSlice(slice2)
		printDay(day2)
		printSlice(slice3)
		printDay(day3)
		printSlice(slice4)
		printDay(day4)
		printSlice(slice5)
		printDay(day5)

		fmt.Println("========================================================")
	}
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var input string = ""
	for input != "99" {
		fmt.Println("APLIKASI CREATE CALENDER :")
		fmt.Println("author : ragilmaulana@gmail.com \n")
		fmt.Println("1. generate calender")
		fmt.Println("99. exit")

		fmt.Print("input pilihan :")
		scanner.Scan()
		input = scanner.Text()
		switch input {
		case "1":
			fmt.Println("input tahun yang akan di generate \t")
			fmt.Print("Input Tahun : ")
			scanner.Scan()
			tahun, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}

			generateCalender(tahun)
			printCalender(tahun)

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
