package repl

import (
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"github.com/timakin/octop/client"
)

type ResCtx struct {
	Reslines []ResLines
	selected []ResLines
	input    []rune
	heading  bool
	mutex    sync.Mutex
	loop     bool
	dirty    bool
	update   bool
	help     bool
}

var resCtx = ResCtx{
	Reslines: []ResLines{},
	input:    []rune{},
	heading:  false,
	mutex:    sync.Mutex{},
	loop:     true,
	dirty:    true,
	update:   false,
	help:     false,
}

type ResLines struct {
	line  *client.ResponseContent
	disp  string
	Owner string
	Title string
	Body  string
}

type matchedres struct {
	ResLines
	pos1     int
	pos2     int
	selected bool
}

type filteredRes []matchedres

var currentRes filteredRes

func filterResLine() {
	resCtx.mutex.Lock()
	defer resCtx.mutex.Unlock()

	defer func() {
		recover()
	}()

	if len(resCtx.input) == 0 {
		currentRes = make(filteredRes, len(resCtx.Reslines))
		for n, f := range resCtx.Reslines {
			Owner, Title, Body, _ := SplitRes(f.line)
			prev_selected := false
			for _, s := range resCtx.selected {
				if f.disp == s.disp {
					prev_selected = true
					break
				}
			}
			currentRes[n] = matchedres{
				ResLines: ResLines{
					line:  f.line,
					disp:  fmt.Sprintf("%s %s", Owner, Title),
					Owner: Owner,
					Title: Title,
					Body:  Body,
				},
				pos1:     -1,
				pos2:     -1,
				selected: prev_selected,
			}
		}
	} else {
		pat := "(?i)(?:.*)("
		for _, r := range []rune(resCtx.input) {
			pat += regexp.QuoteMeta(string(r)) + ".*?"
		}
		pat += ")"
		re := regexp.MustCompile(pat)

		currentRes = make(filteredRes, 0, len(resCtx.Reslines))
		for _, f := range resCtx.Reslines {
			Owner, Title, Body, _ := SplitRes(f.line)
			ms := re.FindAllStringSubmatchIndex(f.disp, 1)
			if len(ms) != 1 || len(ms[0]) != 4 {
				continue
			}
			prev_selected := false
			for _, s := range resCtx.selected {
				if f.disp == s.disp {
					prev_selected = true
					break
				}
			}
			currentRes = append(currentRes, matchedres{
				ResLines: ResLines{
					line:  f.line,
					disp:  fmt.Sprintf("%s %s", Owner, Title),
					Owner: Owner,
					Title: Title,
					Body:  Body,
				},
				pos1:     len([]rune(f.disp[0:ms[0][2]])),
				pos2:     len([]rune(f.disp[0:ms[0][3]])),
				selected: prev_selected,
			})
		}
	}

	if cursor_y < 0 {
		cursor_y = 0
	}
	if cursor_y >= len(currentRes) {
		cursor_y = len(currentRes) - 1
	}
}

func drawResScreen() {
	resCtx.mutex.Lock()
	defer resCtx.mutex.Unlock()

	defer func() {
		recover()
	}()

	width, height = termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	pat := ""
	for _, r := range resCtx.input {
		pat += regexp.QuoteMeta(string(r)) + ".*?"
	}
	for n := 0; n < height-3; n++ {
		if n >= len(currentRes) {
			break
		}
		x := 2
		w := 0
		//line := currentRes[n].line
		line := currentRes[n].disp

		pos1 := currentRes[n].pos1
		pos2 := currentRes[n].pos2
		selected := currentRes[n].selected
		if pos1 >= 0 {
			pwidth := runewidth.StringWidth(string([]rune(currentRes[n].disp)[0:pos1]))
			if !resCtx.heading && pwidth > width/2 {
				rline := []rune(line)
				wwidth := 0
				for i := 0; i < len(rline); i++ {
					w = runewidth.RuneWidth(rline[i])
					if wwidth+w > width/2 {
						line = "..." + string(rline[i:])
						pos1 -= i - 3
						pos2 -= i - 3
						break
					}
					wwidth += w
				}
			}
		}
		swidth := runewidth.StringWidth(line)
		if swidth+2 > width {
			rline := []rune(line)
			line = string(rline[0:width-5]) + "..."
		}
		for f, c := range []rune(line) {
			w = runewidth.RuneWidth(c)
			if x+w > width {
				break
			}
			if pos1 <= f && f < pos2 {
				if selected {
					termbox.SetCell(x, height-4-n, c, termbox.ColorDefault|termbox.AttrReverse, termbox.ColorDefault)
				} else if cursor_y == n {
					termbox.SetCell(x, height-4-n, c, termbox.ColorYellow|termbox.AttrUnderline, termbox.ColorDefault)
				} else {
					termbox.SetCell(x, height-4-n, c, termbox.ColorRed|termbox.AttrBold, termbox.ColorDefault)
				}
			} else {
				if selected {
					termbox.SetCell(x, height-4-n, c, termbox.ColorDefault|termbox.AttrReverse, termbox.ColorDefault)
				} else if cursor_y == n {
					termbox.SetCell(x, height-4-n, c, termbox.ColorYellow|termbox.AttrUnderline, termbox.ColorDefault)
				} else {
					termbox.SetCell(x, height-4-n, c, termbox.ColorDefault, termbox.ColorDefault)
				}
			}
			x += w
		}
	}
	if cursor_y >= 0 {
		printTB(0, height-4-cursor_y, termbox.ColorRed|termbox.AttrBold, termbox.ColorBlack, "> ")
	}
	if scanning >= 0 {
		printTB(0, height-3, termbox.ColorGreen|termbox.AttrBold, termbox.ColorBlack, string([]rune("-\\|/")[scanning%4]))
		scanning++
	}
	printfTB(2, height-3, termbox.ColorWhite|termbox.AttrBold, termbox.ColorBlack, "%d/%d(%d)", len(currentRes), len(resCtx.Reslines), len(resCtx.selected))
	printTB(0, height-2, termbox.ColorBlue|termbox.AttrBold, termbox.ColorBlack, "> ")
	printTB(2, height-2, termbox.ColorWhite|termbox.AttrBold, termbox.ColorBlack, string(resCtx.input))
	termbox.SetCursor(2+runewidth.StringWidth(string(resCtx.input[0:cursor_x])), height-2)
	termbox.Flush()
}

func NewResLines(line *client.ResponseContent) ResLines {
	Owner, Title, Body, _ := SplitRes(line)
	Reslines := ResLines{
		line:  line,
		disp:  fmt.Sprintf("%s %s", Owner, Title),
		Owner: Owner,
		Title: Title,
		Body:  Body,
	}
	return Reslines
}

func ResSelectInterface(responseContents client.ResponseContents) (selected []ResLines, err error) {
	data := responseContents
	resCtx.Reslines = make([]ResLines, 0)
	for _, line := range data {
		resCtx.Reslines = append(resCtx.Reslines, NewResLines(line))
	}
	err = termbox.Init()
	if err != nil {
		return
	}
	if isTty() {
		termbox.SetInputMode(termbox.InputEsc)
	}
	defer termbox.Close()

	// Termbox init
	termbox.SetInputMode(termbox.InputEsc)
	refreshResScreen(0)
	resMainLoop()

	selected = resCtx.selected
	if len(selected) == 0 {
		err = fmt.Errorf("no selected")
		return
	}

	return
}

func handleResKeyEvent(ev termbox.Event) {
	defer func() {
		recover()
	}()

	switch ev.Key {
	case termbox.KeyTab:
		if resCtx.help {
			resCtx.help = false
		} else {
			resCtx.help = true
		}
	case termbox.KeyEsc, termbox.KeyCtrlC:
		if resCtx.help {
			resCtx.help = false
		} else {
			termbox.Close()
			os.Exit(1)
		}
	case termbox.KeyHome, termbox.KeyCtrlA:
		cursor_x = 0
	case termbox.KeyEnd, termbox.KeyCtrlE:
		cursor_x = len(resCtx.input)
	case termbox.KeyEnter:
		if cursor_y >= 0 && cursor_y < len(currentRes) {
			if len(resCtx.selected) == 0 {
				resCtx.selected = append(resCtx.selected, currentRes[cursor_y].ResLines)
			}
			resCtx.loop = false
		}
	case termbox.KeyArrowLeft, termbox.KeyCtrlB:
		if cursor_x > 0 {
			cursor_x--
		}
	case termbox.KeyArrowRight, termbox.KeyCtrlF:
		if cursor_x < len([]rune(resCtx.input)) {
			cursor_x++
		}
	case termbox.KeyArrowUp, termbox.KeyCtrlK, termbox.KeyCtrlP:
		if cursor_y < len(currentRes)-1 {
			if cursor_y < height-4 {
				cursor_y++
			}
		}
	case termbox.KeyArrowDown, termbox.KeyCtrlJ, termbox.KeyCtrlN:
		if cursor_y > 0 {
			cursor_y--
		}
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		if cursor_x > 0 {
			resCtx.input = append(resCtx.input[0:cursor_x-1], resCtx.input[cursor_x:len(resCtx.input)]...)
			cursor_x--
			resCtx.update = true
		}
	case termbox.KeyDelete:
		if cursor_x < len([]rune(resCtx.input)) {
			resCtx.input = append(resCtx.input[0:cursor_x], resCtx.input[cursor_x+1:len(resCtx.input)]...)
			resCtx.update = true
		}
	default:
		if ev.Key == termbox.KeySpace {
			ev.Ch = ' '
		}
		if ev.Ch > 0 {
			out := []rune{}
			out = append(out, resCtx.input[0:cursor_x]...)
			out = append(out, ev.Ch)
			resCtx.input = append(out, resCtx.input[cursor_x:len(resCtx.input)]...)
			cursor_x++
			resCtx.update = true
		}
	}

	// If need to update, start timer
	if scanning != -1 {
		if resCtx.update {
			resCtx.dirty = true
			timer.Reset(duration)
		} else {
			timer.Reset(1)
		}
	} else {
		if resCtx.update {
			filterResLine()
		} else {
		}
		drawResScreen()
	}
}

func refreshResScreen(delay time.Duration) {
	if timer == nil {
		timer = time.AfterFunc(delay, func() {
			if resCtx.dirty {
				filterResLine()
			}
			if resCtx.help {
				resCtx.input = []rune{}
				termbox.HideCursor()
			} else {
				drawResScreen()
			}
			resCtx.dirty = false
		})
	} else {
		timer.Reset(delay)
	}
}

func resMainLoop() {
	for resCtx.loop {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventError {
			resCtx.update = false
		} else if ev.Type == termbox.EventKey {
			handleResKeyEvent(ev)
		}
	}
}

func SplitRes(content *client.ResponseContent) (Title, Owner, Body string, err error) {
	Title = content.Title
	Owner = content.Owner
	Body = content.Body
	err = nil
	return
}
