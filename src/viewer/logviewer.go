package logviewer

import (
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Run() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	content, err := os.ReadFile("log.txt")
	if err != nil {
		log.Fatal(err)
	}

	p := widgets.NewParagraph()
	p.Text = string(content)
	p.SetRect(0, 0, 252, 65)
	p.Border = true

	ui.Render(p)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
