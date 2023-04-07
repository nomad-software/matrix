package term

import (
	"github.com/gdamore/tcell"
	"github.com/nomad-software/matrix/char"
	"github.com/nomad-software/matrix/cli"
	"github.com/nomad-software/matrix/output"
	"github.com/nomad-software/screensaver/screen/saver/digital_rain/matrix"
)

// Term is the main terminal.
type Term struct {
	tcell  tcell.Screen
	ASCII  bool
	Width  int
	Height int
}

// New contructs a new terminal.
func New(opt *cli.Options) Term {
	tc, err := tcell.NewScreen()
	output.OnError(err, "failed to create terminal")

	err = tc.Init()
	output.OnError(err, "failed to initialise terminal")

	width, height := tc.Size()

	s := Term{
		tcell:  tc,
		ASCII:  opt.ASCII,
		Width:  width,
		Height: height,
	}

	return s
}

// Draw displays the passed matrix onto the terminal.
func (t *Term) Draw(m *matrix.Matrix) {
	t.tcell.Clear()

	for x := 0; x < t.Width; x++ {
		for y := 0; y < t.Height; y++ {
			glyph := m.ColumnAtIndex(x).GlyphAtIndex(y)
			if !glyph.IsEmpty() {
				if glyph.IsHighlighted() {
					t.tcell.SetContent(x, y, char.Get(glyph.Index(), t.ASCII), nil, tcell.StyleDefault.Foreground(tcell.ColorWhite))
				} else {
					t.tcell.SetContent(x, y, char.Get(glyph.Index(), t.ASCII), nil, tcell.StyleDefault.Foreground(tcell.ColorGreen))
				}
			}
		}
	}
}

// Update commits all changes to the terminal display.
func (t *Term) Update() {
	t.tcell.Show()
}

// Destroy closes the terminal display and shows the original display.
func (t *Term) Destroy() {
	t.tcell.Fini()
}

// HandleInput waits for key presses and takes the appropriate action.
func (t *Term) HandleInput(signal chan bool) {
	go func() {
		for {
			ev := t.tcell.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(signal)
					return
				case tcell.KeyCtrlL:
					t.tcell.Sync()
				}
			case *tcell.EventResize:
				t.tcell.Sync()
			}
		}
	}()
}
