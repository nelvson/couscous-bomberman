package main

import (
  "fmt"
  "log"
  tcell "github.com/gdamore/tcell/v2"
)

type GameState struct {
  start bool
  quit bool
  lives int
  points int
}

type GameInterface interface {
  Start()
}

func (game GameState) Start() {
  game.start = true
  game.lives = 3
}

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func main() {
  state := GameState{false, false, 0, 0}

  boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

  s, err := tcell.NewScreen()
  if err != nil {
    fmt.Println("er");
  }
  if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

  s.Clear()


	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

  for {
    s.Show()
    ev := s.PollEvent()

    drawText(s, 1, 1, 10, 10, boxStyle, "Press S to start the game, Press Ctrl+C to quit!")
    switch ev := ev.(type) {
    case *tcell.EventKey:
      if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
        return
      } else  if ev.Rune() == 's' {
        state.Start()
        drawText(s, 1, 1, 10, 10, boxStyle, "game is starting!")

      }
    }
  }
  //
}
