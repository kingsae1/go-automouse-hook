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

var HOOK_TIMEOUT = 10 * time.Second
var TICK_TIMEOUT = 5 * 60 * time.Second
var TICK_COUNT = 10
var VERSION = "NONE"
var SPINNER = spinner.New(spinner.CharSets[35], 700*time.Millisecond) // Build our new spinner
var INDEX = 0

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

	SPINNER.Prefix = " [Detecting] Event Hook : "
	SPINNER.Color("cyan")

	checkExe()
	initialize()
	eventhook()
}

func checkExe() {
	// (윈도우) 중복 실행 삭제
	if runtime.GOOS == "windows" {
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
}

func initialize() {
	data, _ := ioutil.ReadFile("go.automover.config")
	info := strings.Split(string(data), "|")

	// Config 값 로드 후, 재정의
	for index, value := range info {
		if index == 0 {
			// 버전
			VERSION = value
		} else if index == 1 {
			// 티커 타임아웃
			timeout, _ := strconv.Atoi(value)
			TICK_TIMEOUT = time.Duration(timeout) * 60 * time.Second
		} else if index == 2 {
			// 티커 횟수
			tickcount, _ := strconv.Atoi(value)
			TICK_COUNT = tickcount
		} else if index == 3 {
			// 후커 타임아웃
			hooktimeout, _ := strconv.Atoi(value)
			HOOK_TIMEOUT = time.Duration(hooktimeout) * time.Second
		}
	}

	fmt.Println("")
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
	⠀⠀⠀⠀⠉⠉⠁⠀⠀⠀⠙⠿⠿⠿⠋⠀⠀⠀⠈⠉⠉
	`)
	fmt.Println(" ############## Automover ##############")
	fmt.Println(" Author : kingsae1004@gmail.com")
	fmt.Println(" Ticker Timeout :", TICK_TIMEOUT)
	fmt.Println(" Hooker Timeout :", HOOK_TIMEOUT)
	fmt.Println(" Ticker MaxCount :", TICK_COUNT)
	fmt.Println(" Version : v" + string(VERSION))
	fmt.Println(" #######################################")
}

func moveMouseCount() {
	INDEX += 1

	// 실제 호출되는 시점 : Timeout * index
	if INDEX > TICK_COUNT {
		robotgo.ScrollMouse(1, "up")
		robotgo.ScrollMouse(1, "down")

		time.Sleep(5 * time.Second)
		SPINNER.Prefix = " [Moving] Event Hook : "
		SPINNER.Color("yellow")
		time.Sleep(5 * time.Second)
		INDEX = 0
	}
}

func eventhook() {
	INDEX = 0
	EvChan := hook.Start()
	defer hook.End()

	ticker := createTicker(10 * time.Second)
	ticker.resetTicker()

	SPINNER.Restart()
	SPINNER.Prefix = " [Detecting] Event Hook : "
	SPINNER.Color("cyan")

	n := 1

	for n > 0 {
		SPINNER.Reverse() // Reverse the direction the spinner is spinning
		SPINNER.Suffix = " (" + strconv.Itoa(INDEX) + "/" + strconv.Itoa(len(EvChan)) + ")"

		if len(EvChan) > 30 {
			// 이벤트가 발생한 경우
			INDEX = 0

			SPINNER.Prefix = " [Detected] Event Hook : "
			SPINNER.Color("magenta")

			// 이벤트 초기화
			hook.End()
			EvChan = nil
			time.Sleep(HOOK_TIMEOUT)
			EvChan = hook.Start()
		} else {
			// 이벤트가 없는 경우 이벤트 강제 발생
			SPINNER.Prefix = " [Detecting] Event Hook : "
			SPINNER.Color("cyan")
			time.Sleep(TICK_TIMEOUT)
			moveMouseCount()
		}

		ticker.resetTicker()
	}
}
