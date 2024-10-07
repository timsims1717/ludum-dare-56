package loading

import (
	"github.com/gopxl/pixel"
	"github.com/timsims1717/ludum-dare-56/internal/data"
	"github.com/timsims1717/ludum-dare-56/internal/ui"
)

var (
	PauseConstructor *ui.DialogConstructor
)

func InitConstructors() {
	PauseConstructor = &ui.DialogConstructor{
		Key:    data.DialogPause,
		Width:  7,
		Height: 9,
		Elements: []ui.ElementConstructor{
			{
				Key:         "pause_resume_ct",
				Position:    pixel.V(0, 56),
				ElementType: ui.ContainerElement,
				Width:       7*16. - 8,
				Height:      16. * 1.5,
				SubElements: []ui.ElementConstructor{
					{
						Key:         "pause_resume_text",
						Text:        "Resume",
						Color:       pixel.ToRGBA(data.ChalkboardWhite),
						Position:    pixel.V(-16, 0),
						ElementType: ui.TextElement,
					},
				},
			},
			{
				Key:         "pause_restart_ct",
				Position:    pixel.V(0, 28),
				ElementType: ui.ContainerElement,
				Width:       7*16. - 8,
				Height:      16. * 1.5,
				SubElements: []ui.ElementConstructor{
					{
						Key:         "pause_restart_text",
						Text:        "Restart",
						Color:       pixel.ToRGBA(data.ChalkboardWhite),
						Position:    pixel.V(-18, 0),
						ElementType: ui.TextElement,
					},
				},
			},
			{
				Key:         "pause_options_ct",
				Position:    pixel.V(0, 0),
				ElementType: ui.ContainerElement,
				Width:       7*16. - 8,
				Height:      16. * 1.5,
				SubElements: []ui.ElementConstructor{
					{
						Key:         "pause_restart_text",
						Text:        "Options",
						Color:       pixel.ToRGBA(data.ChalkboardWhite),
						Position:    pixel.V(-17, 0),
						ElementType: ui.TextElement,
					},
				},
			},
			{
				Key:         "pause_quit_mm_ct",
				Position:    pixel.V(0, -28),
				ElementType: ui.ContainerElement,
				Width:       7*16. - 8,
				Height:      16. * 1.5,
				SubElements: []ui.ElementConstructor{
					{
						Key:         "pause_quit_mm_text",
						Text:        "Quit to Menu",
						Color:       pixel.ToRGBA(data.ChalkboardWhite),
						Position:    pixel.V(-36, 0),
						ElementType: ui.TextElement,
					},
				},
			},
			{
				Key:         "pause_quit_full_ct",
				Position:    pixel.V(0, -56),
				ElementType: ui.ContainerElement,
				Width:       7*16. - 8,
				Height:      16. * 1.5,
				SubElements: []ui.ElementConstructor{
					{
						Key:         "pause_quit_full_text",
						Text:        "Quit to Desktop",
						Color:       pixel.ToRGBA(data.ChalkboardWhite),
						Position:    pixel.V(-42, 0),
						ElementType: ui.TextElement,
					},
				},
			},
		},
	}
}
