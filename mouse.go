package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/briandowns/spinner"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

const TIMEOUT = 5 * time.Second

type ticker struct {
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
	add()
	low()
	// event()
}

func add() {
	data, error := ioutil.ReadFile("version.txt")

	if error != nil {
		fmt.Println(error)
		data = []byte("Please check version.txt")
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
	fmt.Println(" ########### Mouse Automover ###########")
	fmt.Println(" Author : kingsae1004@gmail.com")
	fmt.Println(" Timeout :", TIMEOUT)
	fmt.Println(" Version : " + string(data))
	fmt.Println(" #######################################")
	// fmt.Println(" Please press 'q' to start event hook !")
	robotgo.EventHook(hook.KeyDown, []string{"q"}, func(e hook.Event) {
		// fmt.Println(" Detect Keyboard and Mouse event hook !")
		// robotgo.EventEnd()
		s := robotgo.EventStart()
		<-robotgo.EventProcess(s)
	})

	// robotgo.EventHook(hook.MouseUp, []string{"mleft"}, func(e hook.Event) {
	// fmt.Println("mouse-left")
	// robotgo.EventEnd()
	// })

}

func low() {
	// EvChan := hook.Start()
	defer hook.End()

	ticker := createTicker(TIMEOUT)
	ticker.resetTicker()

	color := []string{"magenta", "cyan", "white", "yellow"}
	colorIndex := 0

	s := spinner.New(spinner.CharSets[35], 700*time.Millisecond) // Build our new spinner
	s.Prefix = " Detecting Event Hook : "
	s.Color(color[colorIndex])

	n := 1
	go func() {
		for n > 0 {
			select {
			//case <-done:
			//	return
			case <-ticker.ticker.C:
				// fmt.Println("Move Scroll mouse")
				colorIndex += 1

				if colorIndex >= len(color) {
					colorIndex = 0
				}
				s.Color(color[colorIndex])

				robotgo.ScrollMouse(10, "up")
			}
		}
	}()
	// for ev := range EvChan {
	for n > 0 {
		// if ev.Kind != 0 {
		// for n > 0 {
		s.Reverse() // Reverse the direction the spinner is spinning
		s.Restart()
		time.Sleep(10 * time.Second)
		s.Stop()
		ticker.resetTicker()
		// }
	}
}

// func event() {
// 	ok := robotgo.AddEvents("q")
// 	if ok {
// 		fmt.Println("add events...")
// 	}

// 	mleft := robotgo.AddEvent("mleft")
// 	if mleft {
// 		fmt.Println("you press... ", "mouse left button")
// 	}

// }
