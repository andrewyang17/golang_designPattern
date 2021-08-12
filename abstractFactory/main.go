package main

import (
	"errors"
	"fmt"
	"strings"
)

type UIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

func CreateUIFactory(os string) (UIFactory, error) {
	os = strings.ToLower(os)
	if os == "macos" {
		return &MacOS{}, nil
	}
	if os == "windows" {
		return &Windows{}, nil
	}
	if os == "deepin" {
		return &Deepin{}, nil
	}
	return nil, errors.New("incorrect os passed")
}

type Button interface {
	RenderButton() string
}

type Checkbox interface {
	RenderCheckbox() string
}

type MacOS struct {}

func (*MacOS) CreateButton() Button {
	return &MacOSButton{}
}

func (*MacOS) CreateCheckbox() Checkbox {
	return &MacOSCheckbox{}
}

type Windows struct {}

func (*Windows) CreateButton() Button {
	return &WindowsButton{}
}

func (*Windows) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

type MacOSButton struct {}

func (*MacOSButton) RenderButton() string {
	return "Rendering MacOS button..."
}

type MacOSCheckbox struct {}

func (*MacOSCheckbox) RenderCheckbox() string {
	return "Rendering MacOS button..."
}

type WindowsCheckbox struct {}
func (*WindowsCheckbox) RenderCheckbox() string {
	return "Rendering Windows checkbox..."
}

type WindowsButton struct {}

func (*WindowsButton) RenderButton() string {
	return "Rendering Windows button..."
}

type Deepin struct {}

func (*Deepin) CreateButton() Button {
	return &DeepinButton{}
}

func (*Deepin) CreateCheckbox() Checkbox {
	return &DeepinCheckbox{}
}

type DeepinButton struct {}

func (*DeepinButton) RenderButton() string {
	return "Rendering Deepin button..."
}

type DeepinCheckbox struct {}

func (*DeepinCheckbox) RenderCheckbox() string {
	return "Rendering Deepin checkbox..."
}



func main() {
	macFactory, err := CreateUIFactory("macos")
	if err != nil {
		panic(err)
	}
	button := macFactory.CreateButton()
	checkbox := macFactory.CreateCheckbox()

	fmt.Println(button.RenderButton())
	fmt.Println(checkbox.RenderCheckbox())

	{
		windowsFactory, err := CreateUIFactory("windows")
		if err != nil {
			panic(err)
		}
		button := windowsFactory.CreateButton()
		checkbox := windowsFactory.CreateCheckbox()

		fmt.Println(button.RenderButton())
		fmt.Println(checkbox.RenderCheckbox())
	}

	{
		deepinFactory, err := CreateUIFactory("deepin")
		if err != nil {
			panic(err)
		}
		button := deepinFactory.CreateButton()
		checkbox := deepinFactory.CreateCheckbox()

		fmt.Println(button.RenderButton())
		fmt.Println(checkbox.RenderCheckbox())
	}
}
