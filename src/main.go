package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
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

type CalculatorUI struct {
	window       *gtk.Window
	displayLabel *gtk.Label
	calc         *Calculator
	buttons      map[string]*gtk.Button
}

func NewCalculatorUI() *CalculatorUI {
	return &CalculatorUI{
		calc:    NewCalculator(),
		buttons: make(map[string]*gtk.Button),
	}
}

func (ui *CalculatorUI) buildUI() {
	ui.window = gtk.NewWindow()
	ui.window.SetTitle("Sobkulačka")
	ui.window.SetDefaultSize(400, 500)
	ui.window.SetResizable(false)

	mainBox := gtk.NewBox(gtk.OrientationVertical, 0)
	mainBox.SetMarginStart(8)
	mainBox.SetMarginEnd(8)
	mainBox.SetMarginTop(8)
	mainBox.SetMarginBottom(8)

	ui.displayLabel = gtk.NewLabel(ui.calc.display)
	ui.displayLabel.SetHExpand(true)
	ui.displayLabel.SetVExpand(false)

	displayProvider := gtk.NewCSSProvider()
	displayProvider.LoadFromData(`
label {
	background-color: #202124;
	color: #f5f5f5;
	font-size: 48px;
	padding: 24px;
	border-radius: 12px;
}
`)
	ui.displayLabel.StyleContext().AddProvider(displayProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	mainBox.Append(ui.displayLabel)

	gridBox := gtk.NewBox(gtk.OrientationVertical, 8)
	gridBox.SetVExpand(true)
	gridBox.SetMarginTop(8)

	keys := [][]string{
		{"7", "8", "9", "/"},
		{"4", "5", "6", "*"},
		{"1", "2", "3", "-"},
		{"0", "C", "=", "+"},
	}

	operatorProvider := gtk.NewCSSProvider()
	operatorProvider.LoadFromData(`
button.operator-button {
	background-color: #3c3c46;
	color: #dcdcff;
	font-size: 24px;
	padding: 16px;
	border-radius: 12px;
	border: none;
}
button.operator-button:hover {
	opacity: 0.8;
}
button.operator-button:active {
	opacity: 0.6;
}
`)

	equalsProvider := gtk.NewCSSProvider()
	equalsProvider.LoadFromData(`
button.equals-button {
	background-color: #507850;
	color: #f0f0f0;
	font-size: 24px;
	padding: 16px;
	border-radius: 12px;
	border: none;
}
button.equals-button:hover {
	opacity: 0.8;
}
button.equals-button:active {
	opacity: 0.6;
}
`)

	clearProvider := gtk.NewCSSProvider()
	clearProvider.LoadFromData(`
button.clear-button {
	background-color: #5a3c3c;
	color: #ffc8c8;
	font-size: 24px;
	padding: 16px;
	border-radius: 12px;
	border: none;
}
button.clear-button:hover {
	opacity: 0.8;
}
button.clear-button:active {
	opacity: 0.6;
}
`)

	numberProvider := gtk.NewCSSProvider()
	numberProvider.LoadFromData(`
button.number-button {
	background-color: #303136;
	color: #dcdcdc;
	font-size: 24px;
	padding: 16px;
	border-radius: 12px;
	border: none;
}
button.number-button:hover {
	opacity: 0.8;
}
button.number-button:active {
	opacity: 0.6;
}
`)

	windowProvider := gtk.NewCSSProvider()
	windowProvider.LoadFromData(`
window {
	background-color: #0d0d0d;
}
`)

	for _, row := range keys {
		rowBox := gtk.NewBox(gtk.OrientationHorizontal, 8)
		rowBox.SetHomogeneous(true)

		for _, key := range row {
			btn := gtk.NewButtonWithLabel(key)
			btn.SetVExpand(true)
			btn.SetHExpand(true)
			btn.SetCanFocus(false)

			switch key {
			case "+", "-", "*", "/":
				btn.AddCSSClass("operator-button")
				btn.StyleContext().AddProvider(operatorProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
			case "=":
				btn.AddCSSClass("equals-button")
				btn.StyleContext().AddProvider(equalsProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
			case "C":
				btn.AddCSSClass("clear-button")
				btn.StyleContext().AddProvider(clearProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
			default:
				btn.AddCSSClass("number-button")
				btn.StyleContext().AddProvider(numberProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
			}

			keyStr := key
			btn.ConnectClicked(func() {
				ui.handleInput(keyStr)
				ui.updateDisplay()
			})

			ui.buttons[key] = btn
			rowBox.Append(btn)
		}

		gridBox.Append(rowBox)
	}

	mainBox.Append(gridBox)

	eventController := gtk.NewEventControllerKey()
	eventController.ConnectKeyPressed(func(keyval uint, keycode uint, state gdk.ModifierType) bool {
		return ui.handleKeyEvent(keyval, keycode, state)
	})
	ui.window.AddController(eventController)

	ui.window.StyleContext().AddProvider(windowProvider, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	ui.window.SetChild(mainBox)
}

func (ui *CalculatorUI) handleInput(key string) {
	switch key {
	case "+", "-", "*", "/":
		ui.calc.setOperator(key)
	case "=":
		ui.calc.equals()
	case "C":
		ui.calc.clear()
	default:
		ui.calc.inputDigit(key)
	}
}

func (ui *CalculatorUI) handleKeyEvent(keyval uint, keycode uint, state gdk.ModifierType) bool {
	switch keyval {
	case gdk.KEY_0:
		ui.handleInput("0")
	case gdk.KEY_1:
		ui.handleInput("1")
	case gdk.KEY_2:
		ui.handleInput("2")
	case gdk.KEY_3:
		ui.handleInput("3")
	case gdk.KEY_4:
		ui.handleInput("4")
	case gdk.KEY_5:
		ui.handleInput("5")
	case gdk.KEY_6:
		ui.handleInput("6")
	case gdk.KEY_7:
		ui.handleInput("7")
	case gdk.KEY_8:
		ui.handleInput("8")
	case gdk.KEY_9:
		ui.handleInput("9")
	case gdk.KEY_plus, gdk.KEY_equal:
		if (state & gdk.ShiftMask) != 0 {
			ui.handleInput("+")
		} else {
			ui.handleInput("=")
		}
	case gdk.KEY_minus:
		ui.handleInput("-")
	case gdk.KEY_asterisk:
		ui.handleInput("*")
	case gdk.KEY_slash:
		ui.handleInput("/")
	case gdk.KEY_Return, gdk.KEY_KP_Enter:
		ui.handleInput("=")
	case gdk.KEY_BackSpace:
		ui.handleInput("C")
	case gdk.KEY_c, gdk.KEY_C:
		if (state & gdk.ControlMask) != 0 {
			return false
		}
		ui.handleInput("C")
	case gdk.KEY_Escape:
		ui.handleInput("C")
	default:
		return false
	}

	ui.updateDisplay()
	return true
}

func (ui *CalculatorUI) updateDisplay() {
	ui.displayLabel.SetLabel(ui.calc.display)
}

func (ui *CalculatorUI) show() {
	ui.window.Present()
}

func main() {
	app := gtk.NewApplication("me.vypal.sobkulacka", 0)
	app.ConnectActivate(func() {
		ui := NewCalculatorUI()
		ui.buildUI()
		ui.window.SetApplication(app)
		ui.show()
	})

	if code := app.Run([]string{}); code > 0 {
		os.Exit(code)
	}
}
