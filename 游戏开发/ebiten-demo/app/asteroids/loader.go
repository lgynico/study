package asteroids

import (
	"embed"
	"image"
	_ "image/png"

	"io/fs"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(pattern string) []*ebiten.Image {
	var images []*ebiten.Image

	// 遍历嵌入资源文件系统
	err := fs.WalkDir(assets, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 跳过目录，只处理文件
		if d.IsDir() {
			return nil
		}

		// 使用通配符匹配文件名
		matched, err := filepath.Match(pattern, path)
		if err != nil {
			return err
		}

		if matched {
			// 使用现有的 mustLoadImage 加载匹配的图片
			images = append(images, mustLoadImage(path))
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	return images
}
