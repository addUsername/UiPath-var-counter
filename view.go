package main

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

func text() {

	// quick use package func
	color.Redp("Simple to use color")
	color.Redln("Simple to use color")

	color.Cyanln("Simple to use color")
	color.Yellowln("Simple to use color")

	color.Println("<fg=11aa23>he</><bg=120,35,156>llo</>, <fg=167;bg=232>wel</><fg=red>come</>")

	// tips message
	color.Info.Tips("tips style message")
	color.Warn.Tips("tips style message")
}

func Pinfo(info string) {
	//color.Info.Tips(info)
	if v {
		color.Greenln(info)
	}
}

func Perror(e error) {
	color.Redln(e)
}

func Print(x string) {
	fmt.Println(x)
}

func Pdlog(x string) {
	if v {
		color.Blueln(x)
	}
}

func Pwarn(x string) {
	myStyle := color.New(color.FgLightYellow, color.BgDefault, color.OpBold)
	myStyle.Println(x)
}

func Ptitle(i string, total string, filename string) {
	// añadir barras horizontales
	myStyle := color.New(color.FgDarkGray, color.BgLightGreen, color.OpBold)

	myStyle.Printf(" [%s - %s]    %s ", i, total, filename)
	myStyle.Println()
	fmt.Println("┌" + strings.Repeat("─", 27) + "┬" + strings.Repeat("─", 29) + "┬" + strings.Repeat("─", 6) + "┐")
}
func Prow(name string, object string, count string, odd bool) {

	//https://stackoverflow.com/questions/25637440/golang-how-to-pad-a-number-with-zeros-when-printing
	name, object, count = padColumns(name, object, count)

	/*
		if odd {
			myStyle := color.New(color.FgDarkGray, color.BgBlack, color.OpBold)
			//	myStyle.Printf("%6d", "%6d", "%6d", y[0], y[1], y[2])
			myStyle.Println()
			return
		}
		myStyle := color.New(color.FgBlack, color.BgBlack, color.OpBold)
		//myStyle.Printf("%6s", "%6s", "%6s", y[0], y[1], y[2])
		myStyle.Println()

		//	myStyle := color.New(color.FgGreen, color.BgBlack, color.OpBold)
		//myStyle.Println(x)
	*/
	fmt.Printf("%s │ %s │ %s", name, object, count)
	fmt.Println()
}
func Pfooter() {
	fmt.Println("└" + strings.Repeat("─", 27) + "┴" + strings.Repeat("─", 29) + "┴" + strings.Repeat("─", 6) + "┘")
}

func padColumns(name, object, count string) (string, string, string) {

	// name: max 20ch
	// object: max 20ch
	// count 4

	var n int

	if len(name) < 25 {
		n = 25 - len(name)
		name = "  " + name + strings.Repeat(" ", n)
	}

	object = strings.Join(strings.Split(object, ":")[1:], ":")

	if len(object) < 25 {
		n = 25 - len(object)
		object = "  " + object + strings.Repeat(" ", n)
	}

	if len(count) < 4 {
		n = 4 - len(count)
		count = "  " + count + strings.Repeat(" ", n)
	}

	return name, object, count

}
