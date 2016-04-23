package repl

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"github.com/timakin/octop/client"
)

var (
	duration           = 10 * time.Millisecond
	scanning           = 0
	cursor_x, cursor_y int
	width, height      int
	timer              *time.Timer
)

func printTB(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range []rune(msg) {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func printfTB(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	printTB(x, y, fg, bg, s)
}

type RepoCtx struct {
	Repolines []RepoLines
	selected  []RepoLines
	input     []rune
	heading   bool
	mutex     sync.Mutex
	loop      bool
	dirty     bool
	update    bool
	help      bool
}

var repoCtx = RepoCtx{
	Repolines: []RepoLines{},
	input:     []rune{},
	heading:   false,
	mutex:     sync.Mutex{},
	loop:      true,
	dirty:     true,
	update:    false,
	help:      false,
}

type RepoLines struct {
	line        *client.RepoNotificationCounter
	disp        string
	UnreadCount string
	Owner       string
	Repo        string
}

type matchedrepo struct {
	RepoLines
	pos1     int
	pos2     int
	selected bool
}

type filteredRepo []matchedrepo

var currentRepo filteredRepo

func filterRepoLine() {
	repoCtx.mutex.Lock()
	defer repoCtx.mutex.Unlock()

	defer func() {
		recover()
	}()

	if len(repoCtx.input) == 0 {
		currentRepo = make(filteredRepo, len(repoCtx.Repolines))
		for n, f := range repoCtx.Repolines {
			UnreadCount, Owner, Repo, _ := SplitRepo(f.line)
			prev_selected := false
			for _, s := range repoCtx.selected {
				if f.disp == s.disp {
					prev_selected = true
					break
				}
			}
			currentRepo[n] = matchedrepo{
				RepoLines: RepoLines{
					line:        f.line,
					disp:        fmt.Sprintf("%s %s %s", UnreadCount, Owner, Repo),
					UnreadCount: UnreadCount,
					Owner:       Owner,
					Repo:        Repo,
				},
				pos1:     -1,
				pos2:     -1,
				selected: prev_selected,
			}
		}
	} else {
		pat := "(?i)(?:.*)("
		for _, r := range []rune(repoCtx.input) {
			pat += regexp.QuoteMeta(string(r)) + ".*?"
		}
		pat += ")"
		re := regexp.MustCompile(pat)

		currentRepo = make(filteredRepo, 0, len(repoCtx.Repolines))
		for _, f := range repoCtx.Repolines {
			UnreadCount, Owner, Repo, _ := SplitRepo(f.line)
			ms := re.FindAllStringSubmatchIndex(f.disp, 1)
			if len(ms) != 1 || len(ms[0]) != 4 {
				continue
			}
			prev_selected := false
			for _, s := range repoCtx.selected {
				if f.disp == s.disp {
					prev_selected = true
					break
				}
			}
			currentRepo = append(currentRepo, matchedrepo{
				RepoLines: RepoLines{
					line:        f.line,
					disp:        fmt.Sprintf("%s %s", Owner, Repo),
					UnreadCount: UnreadCount,
					Owner:       Owner,
					Repo:        Repo,
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
	if cursor_y >= len(currentRepo) {
		cursor_y = len(currentRepo) - 1
	}
}

func drawRepoScreen() {
	repoCtx.mutex.Lock()
	defer repoCtx.mutex.Unlock()

	defer func() {
		recover()
	}()

	width, height = termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	pat := ""
	for _, r := range repoCtx.input {
		pat += regexp.QuoteMeta(string(r)) + ".*?"
	}
	for n := 0; n < height-3; n++ {
		if n >= len(currentRepo) {
			break
		}
		x := 2
		w := 0
		//line := currentRepo[n].line
		line := currentRepo[n].disp

		pos1 := currentRepo[n].pos1
		pos2 := currentRepo[n].pos2
		selected := currentRepo[n].selected
		if pos1 >= 0 {
			pwidth := runewidth.StringWidth(string([]rune(currentRepo[n].disp)[0:pos1]))
			if !repoCtx.heading && pwidth > width/2 {
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
	printfTB(2, height-3, termbox.ColorWhite|termbox.AttrBold, termbox.ColorBlack, "%d/%d(%d)", len(currentRepo), len(repoCtx.Repolines), len(repoCtx.selected))
	printTB(0, height-2, termbox.ColorBlue|termbox.AttrBold, termbox.ColorBlack, "> ")
	printTB(2, height-2, termbox.ColorWhite|termbox.AttrBold, termbox.ColorBlack, string(repoCtx.input))
	termbox.SetCursor(2+runewidth.StringWidth(string(repoCtx.input[0:cursor_x])), height-2)
	termbox.Flush()
}

func NewRepoLines(line *client.RepoNotificationCounter) RepoLines {
	UnreadCount, Owner, Repo, _ := SplitRepo(line)
	Repolines := RepoLines{
		line:        line,
		disp:        fmt.Sprintf("%s %s", Owner, Repo),
		UnreadCount: UnreadCount,
		Owner:       Owner,
		Repo:        Repo,
	}
	return Repolines
}

func RepoSelectInterface(repoNotificationCounters client.RepoNotificationCounters) (selected []RepoLines, err error) {
	data := repoNotificationCounters
	repoCtx.Repolines = make([]RepoLines, 0)
	for _, line := range data {
		repoCtx.Repolines = append(repoCtx.Repolines, NewRepoLines(line))
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
	refreshRepoScreen(0)
	repoMainLoop()

	selected = repoCtx.selected
	if len(selected) == 0 {
		err = fmt.Errorf("no selected")
		return
	}

	return
}

func handleRepoKeyEvent(ev termbox.Event) {
	defer func() {
		recover()
	}()

	switch ev.Key {
	case termbox.KeyTab:
		if repoCtx.help {
			repoCtx.help = false
		} else {
			repoCtx.help = true
		}
	case termbox.KeyEsc, termbox.KeyCtrlC:
		if repoCtx.help {
			repoCtx.help = false
		} else {
			termbox.Close()
			os.Exit(1)
		}
	case termbox.KeyHome, termbox.KeyCtrlA:
		cursor_x = 0
	case termbox.KeyEnd, termbox.KeyCtrlE:
		cursor_x = len(repoCtx.input)
	case termbox.KeyEnter:
		if cursor_y >= 0 && cursor_y < len(currentRepo) {
			if len(repoCtx.selected) == 0 {
				repoCtx.selected = append(repoCtx.selected, currentRepo[cursor_y].RepoLines)
			}
			repoCtx.loop = false
		}
	case termbox.KeyArrowLeft, termbox.KeyCtrlB:
		if cursor_x > 0 {
			cursor_x--
		}
	case termbox.KeyArrowRight, termbox.KeyCtrlF:
		if cursor_x < len([]rune(repoCtx.input)) {
			cursor_x++
		}
	case termbox.KeyArrowUp, termbox.KeyCtrlK, termbox.KeyCtrlP:
		if cursor_y < len(currentRepo)-1 {
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
			repoCtx.input = append(repoCtx.input[0:cursor_x-1], repoCtx.input[cursor_x:len(repoCtx.input)]...)
			cursor_x--
			repoCtx.update = true
		}
	case termbox.KeyDelete:
		if cursor_x < len([]rune(repoCtx.input)) {
			repoCtx.input = append(repoCtx.input[0:cursor_x], repoCtx.input[cursor_x+1:len(repoCtx.input)]...)
			repoCtx.update = true
		}
	default:
		if ev.Key == termbox.KeySpace {
			ev.Ch = ' '
		}
		if ev.Ch > 0 {
			out := []rune{}
			out = append(out, repoCtx.input[0:cursor_x]...)
			out = append(out, ev.Ch)
			repoCtx.input = append(out, repoCtx.input[cursor_x:len(repoCtx.input)]...)
			cursor_x++
			repoCtx.update = true
		}
	}

	// If need to update, start timer
	if scanning != -1 {
		if repoCtx.update {
			repoCtx.dirty = true
			timer.Reset(duration)
		} else {
			timer.Reset(1)
		}
	} else {
		if repoCtx.update {
			filterRepoLine()
		} else {
		}
		drawRepoScreen()
	}
}

func refreshRepoScreen(delay time.Duration) {
	if timer == nil {
		timer = time.AfterFunc(delay, func() {
			if repoCtx.dirty {
				filterRepoLine()
			}
			if repoCtx.help {
				repoCtx.input = []rune{}
				termbox.HideCursor()
			} else {
				drawRepoScreen()
			}
			repoCtx.dirty = false
		})
	} else {
		timer.Reset(delay)
	}
}

func repoMainLoop() {
	for repoCtx.loop {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventError {
			repoCtx.update = false
		} else if ev.Type == termbox.EventKey {
			handleRepoKeyEvent(ev)
		}
	}
}

func SplitRepo(counter *client.RepoNotificationCounter) (UnreadCount, Owner, repo string, err error) {
	UnreadCount = strconv.Itoa(counter.UnreadNotificationCount)
	Owner = *counter.Repository.Owner.Login
	repo = *counter.Repository.Name
	err = nil
	return
}
