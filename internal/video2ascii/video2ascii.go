package video2ascii

import (
	"fmt"
	"image"
	"strings"

	"gocv.io/x/gocv"
)

type CharImg []string

// 用于生成字符画的像素，越往后视觉上越明显。。这是我自己按感觉排的，你可以随意调整。写函数里效率太低，所以只好放全局了
var pixels = []rune(" .,-'`:!1+*abcdefghijklmnopqrstuvwxyz<>()\\/{}[]?234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ%&@#$")

func Img2Chars(img *gocv.Mat) (result CharImg) {
	dims := img.Size()
	height, width := dims[0], dims[1]
	for row := 0; row < height; row++ {
		var lineBuilder strings.Builder
		for col := 0; col < width; col++ {
			gray := int(img.GetUCharAt(row, col))
			// 灰度是用8位表示的，最大值为255。这里将灰度转换到 0-1 之间
			// 将灰度值进一步转换到 0 到 (len(pixels) - 1) 之间，这样就和 pixels 里的字符对应起来了
			index := float32(gray) / 255 * float32(len(pixels)-1)
			// 添加字符像素（最后面加一个空格，是因为命令行有行距却没几乎有字符间距，用空格当间距）
			lineBuilder.WriteRune(pixels[int(index)])
			lineBuilder.WriteRune(' ')
		}
		result = append(result, lineBuilder.String())
	}
	return
}

func Video2Chars(videoPath string, size image.Point, seconds float64) (charImgList []CharImg, fps float64, err error) {
	// open webcam
	webcam, err := gocv.OpenVideoCapture(videoPath)
	if err != nil {
		err = fmt.Errorf("error opening capture device %v: %w", videoPath, err)
		return
	}
	defer webcam.Close()

	fps = webcam.Get(gocv.VideoCaptureFPS)
	frameCount := int(fps * seconds)

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	count := 0
	for webcam.IsOpened() && count < frameCount {
		if ok := webcam.Read(&img); !ok {
			err = fmt.Errorf("cannot read device: %v", videoPath)
			return
		}
		if img.Empty() {
			continue
		}

		gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
		gocv.Resize(img, &img, size, 0, 0, gocv.InterpolationArea)

		img := Img2Chars(&img)
		charImgList = append(charImgList, img)

		count++
	}

	return charImgList, fps, nil
}
