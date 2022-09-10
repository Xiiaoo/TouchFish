package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func convertMonth(s string) string {
	var result string
	switch s {
	case "January":
		result = "01"
	case "February":
		result = "02"
	case "March":
		result = "03"
	case "April":
		result = "04"
	case "May":
		result = "05"
	case "June":
		result = "06"
	case "July":
		result = "07"
	case "August":
		result = "08"
	case "September":
		result = "09"
	case "October":
		result = "10"
	case "November":
		result = "11"
	case "December":
		result = "12"
	}
	return result
}

//判断行中是否包含calendar_time
func haveCalendarTime(s string) bool {
	isTrue := strings.Contains(s, "calendar_time")
	return isTrue
}

//读取txt
func readTxt(r io.Reader) ([]string, error) {
	reader := bufio.NewReader(r)
	var timeSlice []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		if haveCalendarTime(string(line)) {
			timeSlice = append(timeSlice, strings.Trim(string(line[21:44]), ""))
		}

	}
	var result []string
	result = append(result, timeSlice[0])
	result = append(result, timeSlice[len(timeSlice)-1])
	return result, nil
}

func printResult() {
	f, err := os.Open("info.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer f.Close()
	content, err := readTxt(f)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("开始时间：", content[0])
	fmt.Println("结束时间：", content[1])
}

func printRand() {
	start_time := time.Now().AddDate(0, 0, -1)

	start_year := strconv.Itoa(start_time.Year())
	start_month := convertMonth(start_time.Month().String())
	start_day := strconv.Itoa(start_time.Day())

	rand.Seed(time.Now().UnixNano())
	start_hour := rand.Intn(4) + 18
	start_minute := rand.Intn(49) + 10
	start_second := rand.Intn(49) + 10
	start_msecond := rand.Intn(889) + 100

	s_ymd_hmsms := start_year + "-" + start_month + "-" + start_day + " " + strconv.Itoa(start_hour) + ":" + strconv.Itoa(start_minute) + ":" + strconv.Itoa(start_second) + "." + strconv.Itoa(start_msecond)

	fmt.Println("开始时间：", s_ymd_hmsms)

	end_month := convertMonth(time.Now().Month().String())
	end_hour := start_hour - 13
	end_minute := rand.Intn(49) + 10
	end_second := rand.Intn(49) + 10
	end_msecond := rand.Intn(889) + 100

	e_ymd_hmsms := strconv.Itoa(time.Now().Year()) + "-" + end_month + "-" + strconv.Itoa(time.Now().Day()) + " " + "0" + strconv.Itoa(end_hour) + ":" + strconv.Itoa(end_minute) + ":" + strconv.Itoa(end_second) + "." + strconv.Itoa(end_msecond)

	fmt.Println("结束时间：", e_ymd_hmsms)
}

func main() {
	fmt.Println("========================")
	fmt.Println("以下为随机生成的虚拟时间:")
	printRand()
	fmt.Println("========================")
	fmt.Println("以下为读取TXT生成的实际的时间:")
	fmt.Println("需要TXT与BAT在同一目录下!")
	printResult()
	fmt.Println("========================")
}