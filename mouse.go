package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var TICK_TIMEOUT = 5 * 60 * time.Second
var VERSION = "NONE"

// const TIMEOUT = 1 * 10 * time.Second

type ticker struct {
	// Ticker : 미래에 한 시점에서 반복
	period time.Duration
	ticker time.Ticker
}

func createTicker(period time.Duration) *ticker {
	return &ticker{period, *time.NewTicker(period)}
}
func (t *ticker) resetTicker() {
	t.ticker = *time.NewTicker(t.period)
}

func main() {
	defer fmt.Println("FIN")

	if runtime.GOOS == "windows" {
		// 중복 실행 삭제
		var fireArray [2]int
		fireIndex := 0

		lsCmd := exec.Command("c:/windows/system32/tasklist.exe")
		grepCmd := exec.Command("c:/windows/system32/findstr.exe", "go.automover.exe")
		taskList, _ := lsCmd.Output()
		grepIn, _ := grepCmd.StdinPipe()
		grepOut, _ := grepCmd.StdoutPipe()

		grepCmd.Start()

		grepIn.Write(taskList)
		grepIn.Close()

		grepBytes, _ := ioutil.ReadAll(grepOut)

		grepCmd.Wait()

		stringResult := string(grepBytes)
		r, _ := regexp.Compile("automover")
		stringArray := r.FindAllString(stringResult, -1)

		if len(stringArray) > 1 {
			slice := strings.Split(stringResult, " ")

			for _, str := range slice {
				pid, _ := strconv.Atoi(str)

				if pid > 1 {
					fireArray[fireIndex] = pid
					fireIndex += 1
				}
			}

			cmd := exec.Command("c:/windows/system32/taskkill.exe", "/PID", strconv.Itoa(fireArray[0]))
			cmd.Start()
			cmd.Wait()
		}
	}

	defer low()
	defer add()
}

func add() {
	data, _ := ioutil.ReadFile("go.automover.config")
	info := strings.Split(string(data), "|")

	for index, value := range info {
		if index == 0 {
			VERSION = value
		} else if index == 1 {
			timeout, _ := strconv.Atoi(value)
			TICK_TIMEOUT = time.Duration(timeout) * 60 * time.Second
		}

	}

	fmt.Println(" #######################################")
	fmt.Println(`	 ⣠⣤⡾⠿⠿⠿⢷⣤⣀⠀⠀⠀⠀⠀⣀⣤⡾⠿⠿⠿⢷⣤⣄
	⣤⡿⠋⠀⠀⠀⠀⠀⠙⣿⣄⠀⠀⠀⣠⣿⠋⠁⠀⠀⠀⠀⠙⢿⣤
	⣿⡇⠀⠀⠀⠀⠀⠀⠀⠻⣿⣿⣿⣿⣿⠿⠀⠀⠀⠀⠀⠀⠀⢸⣿
	⠻⣷⣄⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⢠⣼⠟
	⠀⠙⢿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⡿⠋
	⠀⠀⠀⢹⣿⠀⢀⣴⣿⠀⠀⠀⠀⠀⠀⠀⣿⣦⡄⠀⣿⡏
	⠀⠀⠀⢸⣿⠀⠘⢿⣿⠀⠀⠀⠀⠀⠀⠀⣿⡿⠃⠀⣿⡇
	⠀⠀⠀⠘⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣿⠃
	⠀⠀⠀⢠⣿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣄
	⠀⠀⢸⣿⣿⣿⣿⣿⣷⣄⠀⠀⠀⠀⠀⣀⣾⣿⣿⣿⣿⣿⡇
	⠀⠀⠀⢹⣿⣿⣿⡟⠛⢿⣷⣄⠀⣠⣶⡿⠛⢻⣿⣿⣿⡏⠁
	⠀⠀⠀⠀⠉⠉⠁⠀⠀⠀⠙⠿⠿⠿⠋⠀⠀⠀⠈⠉⠉`)
	fmt.Println(" ############## Automover ##############")
	fmt.Println(" Author : kingsae1004@gmail.com")
	fmt.Println(" Ticker Timeout :", TICK_TIMEOUT)
	fmt.Println(" Version : v" + string(VERSION))
	fmt.Println(" #######################################")
}

var index = 0

func moveMouseCount() {
	index += 1

	// log.Println(index)

	if index > 10 {
		index = 0
		robotgo.ScrollMouse(1, "up")
		robotgo.ScrollMouse(1, "down")
	}
}

func low() {
	EvChan := hook.Start()
	defer hook.End()

	ticker := createTicker(10 * time.Second)
	ticker.resetTicker()

	s := spinner.New(spinner.CharSets[35], 700*time.Millisecond) // Build our new spinner
	s.Restart()
	s.Prefix = " [Detecting] Event Hook : "
	s.Color("cyan")

	n := 1

	for n > 0 {
		// log.Println("Event Length : " + strconv.Itoa(len(EvChan)))
		s.Reverse() // Reverse the direction the spinner is spinning
		// s.Restart()

		s.Suffix = " (" + strconv.Itoa(len(EvChan)) + ")"

		if len(EvChan) > 30 {
			// 이벤트가 발생한 경우
			s.Prefix = " [Detected] Event Hook : "
			s.Color("magenta")

			// 이벤트 초기화
			hook.End()
			EvChan = hook.Start()
			index = 0
			time.Sleep(5 * time.Second)
		} else {
			// 이벤트가 없는 경우 이벤트 강제 발생
			s.Prefix = " [Detecting] Event Hook : "
			s.Color("cyan")
			moveMouseCount()
			time.Sleep(TICK_TIMEOUT)
		}

		ticker.resetTicker()
	}
}
