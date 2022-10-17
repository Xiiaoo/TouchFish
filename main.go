package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/atotto/clipboard"
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
	return strings.Contains(s, "calendar_time")
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

func printResult() string {
	f, err := os.Open("info.txt")
	if err != nil {
		fmt.Println("err:", err)
		return err.Error()
	}
	defer f.Close()
	content, err := readTxt(f)
	if err != nil {
		fmt.Println("err:", err)
		return err.Error()
	}
	return "开始时间：" + content[0] + "\n结束时间：" + content[1]
}

func printRand() string {
	var start_time time.Time
	var end_time time.Time
	if time.Now().Weekday().String() != "Monday" {
		start_time = time.Now().AddDate(0, 0, -1)
		end_time = time.Now()
	} else {
		start_time = time.Now().AddDate(0, 0, -2)
		end_time = time.Now().AddDate(0, 0, -1)
	}
	start_year := strconv.Itoa(start_time.Year())
	start_month := convertMonth(start_time.Month().String())
	start_day := strconv.Itoa(start_time.Day())

	rand.Seed(time.Now().UnixNano())
	start_hour := rand.Intn(4) + 18
	start_minute := rand.Intn(49) + 10
	start_second := rand.Intn(49) + 10
	start_msecond := rand.Intn(889) + 100

	s_ymd_hmsms := start_year + "-" + start_month + "-" + start_day + " " + strconv.Itoa(start_hour) + ":" + strconv.Itoa(start_minute) + ":" + strconv.Itoa(start_second) + "." + strconv.Itoa(start_msecond)

	end_month := convertMonth(time.Now().Month().String())
	end_hour := start_hour - 13
	end_minute := rand.Intn(49) + 10
	end_second := rand.Intn(49) + 10
	end_msecond := rand.Intn(889) + 100

	e_ymd_hmsms := strconv.Itoa(time.Now().Year()) + "-" + end_month + "-" + strconv.Itoa(end_time.Day()) + " " + "0" + strconv.Itoa(end_hour) + ":" + strconv.Itoa(end_minute) + ":" + strconv.Itoa(end_second) + "." + strconv.Itoa(end_msecond)

	result := "开始时间：" + s_ymd_hmsms + "\n结束时间：" + e_ymd_hmsms

	return result
}

func main() {
	var output string
	if printResult() == "open info.txt: The system cannot find the file specified." {
		fmt.Println("随机生成的时间已粘贴到剪切板！")
		output = printRand()
	} else {
		fmt.Println("实际读取的时间已粘贴到剪切板！")
		output = printResult()
	}
	clipboard.WriteAll(output)
}
