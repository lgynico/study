package algogo

import "testing"

type (
	Foo struct{}

	Bar struct{}
)

func BenchmarkConvert(b *testing.B) {
	foos := make([]*Foo, 1000000)
	for i := range foos {
		foos[i] = &Foo{}
	}

	b.Run("Empty", func(b *testing.B) { Convert_Empty(foos) })
	b.Run("Capacity", func(b *testing.B) { Convert_Capacity(foos) })
	b.Run("Length", func(b *testing.B) { Convert_Length(foos) })
}

func Convert_Empty(foos []*Foo) []*Bar {
	bars := make([]*Bar, 0)
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}
func Convert_Capacity(foos []*Foo) []*Bar {
	bars := make([]*Bar, 0, len(foos))
	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}
	return bars
}

func Convert_Length(foos []*Foo) []*Bar {
	bars := make([]*Bar, len(foos))
	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}
	return bars
}

func fooToBar(_ *Foo) *Bar {
	return &Bar{}
}
