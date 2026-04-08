package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"strconv"

	"gioui.org/op/clip"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Calculator struct {
	display  string
	stored   float64
	operator string
	newInput bool
}

func NewCalculator() *Calculator {
	return &Calculator{
		display: "0",
	}
}

func (c *Calculator) inputDigit(d string) {
	if c.newInput {
		c.display = d
		c.newInput = false
	} else {
		if c.display == "0" {
			c.display = d
		} else {
			c.display += d
		}
	}
}

func (c *Calculator) applyOp() {
	val, _ := strconv.ParseFloat(c.display, 64)

	switch c.operator {
	case "+":
		c.stored += val
	case "-":
		c.stored -= val
	case "*":
		c.stored *= val
	case "/":
		if val != 0 {
			c.stored /= val
		}
	default:
		c.stored = val
	}

	c.display = fmt.Sprintf("%g", c.stored)
}

func (c *Calculator) setOperator(op string) {
	c.applyOp()
	c.operator = op
	c.newInput = true
}

func (c *Calculator) equals() {
	c.applyOp()
	c.operator = ""
	c.newInput = true
}

func (c *Calculator) clear() {
	c.display = "0"
	c.stored = 0
	c.operator = ""
	c.newInput = false
}

func main() {
	go func() {
		window := new(app.Window)
		if err := run(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops

	calc := NewCalculator()

	buttons := map[string]*widget.Clickable{}
	keys := []string{
		"7", "8", "9", "/",
		"4", "5", "6", "*",
		"1", "2", "3", "-",
		"0", "C", "=", "+",
	}

	for _, k := range keys {
		buttons[k] = new(widget.Clickable)
	}

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    0,
						Bottom: 0,
						Left:   0,
						Right:  0,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						cs := gtx.Constraints
						defer clip.Rect{Max: image.Pt(cs.Max.X, cs.Max.Y)}.Push(gtx.Ops).Pop()
						paint.ColorOp{Color: color.NRGBA{R: 32, G: 33, B: 36, A: 255}}.Add(gtx.Ops)
						paint.PaintOp{}.Add(gtx.Ops)
						return layout.Inset{
							Top:    0,
							Bottom: 0,
							Left:   24,
							Right:  24,
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							label := material.H2(theme, calc.display)
							label.Alignment = text.End
							label.Color = color.NRGBA{R: 245, G: 245, B: 245, A: 255}
							return layout.Center.Layout(gtx, label.Layout)
						})
					})
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:    8,
						Bottom: 8,
						Left:   8,
						Right:  8,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							buildRow(gtx, theme, calc, buttons, "7", "8", "9", "/"),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: 8}.Layout(gtx) }),
							buildRow(gtx, theme, calc, buttons, "4", "5", "6", "*"),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: 8}.Layout(gtx) }),
							buildRow(gtx, theme, calc, buttons, "1", "2", "3", "-"),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions { return layout.Spacer{Height: 8}.Layout(gtx) }),
							buildRow(gtx, theme, calc, buttons, "0", "C", "=", "+"),
						)
					})
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}

func buildRow(gtx layout.Context, theme *material.Theme, calc *Calculator, buttons map[string]*widget.Clickable, keys ...string) layout.FlexChild {
	return layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			func() []layout.FlexChild {
				var children []layout.FlexChild
				for i, k := range keys {
					key := k
					btn := buttons[key]

					children = append(children, layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						for btn.Clicked(gtx) {
							handleInput(calc, key)
						}
						var btnStyle material.ButtonStyle
						btnStyle = material.Button(theme, btn, key)
						btnStyle.CornerRadius = 12
						btnStyle.Inset = layout.UniformInset(8)
						switch key {
						case "+", "-", "*", "/":
							btnStyle.Background = color.NRGBA{R: 60, G: 60, B: 70, A: 255}
							btnStyle.Color = color.NRGBA{R: 220, G: 220, B: 255, A: 255}
						case "=":
							btnStyle.Background = color.NRGBA{R: 80, G: 120, B: 80, A: 255}
							btnStyle.Color = color.NRGBA{R: 240, G: 240, B: 240, A: 255}
						case "C":
							btnStyle.Background = color.NRGBA{R: 90, G: 60, B: 60, A: 255}
							btnStyle.Color = color.NRGBA{R: 255, G: 200, B: 200, A: 255}
						default:
							btnStyle.Background = color.NRGBA{R: 48, G: 49, B: 54, A: 255}
							btnStyle.Color = color.NRGBA{R: 220, G: 220, B: 220, A: 255}
						}
						return layout.Inset{
							Left: 4, Right: 4,
						}.Layout(gtx, btnStyle.Layout)
					}))
					// Add spacing between buttons except after the last one
					if i < len(keys)-1 {
						children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Width: 0}.Layout(gtx)
						}))
					}
				}
				return children
			}()...,
		)
	})
}

func handleInput(calc *Calculator, key string) {
	switch key {
	case "+", "-", "*", "/":
		calc.setOperator(key)
	case "=":
		calc.equals()
	case "C":
		calc.clear()
	default:
		calc.inputDigit(key)
	}
}

func paintRect(gtx layout.Context, width, height int, col color.NRGBA) {
	defer clip.Rect{Max: image.Pt(width, height)}.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
}
