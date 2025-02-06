package main

type (
	Option struct {
		Title                     string
		ScreenWidth, ScreenHeight int
		LayoutWidth, LayoutHeight int
	}

	OptionFunc func(*Option)
)

func DefaultOption() Option {
	return Option{
		Title:        "Gomoku",
		ScreenWidth:  640,
		ScreenHeight: 480,
		LayoutWidth:  640,
		LayoutHeight: 480,
	}
}
