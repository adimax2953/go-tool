package gencode

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"math/rand"
	"time"

	LogTool "github.com/adimax2953/log-tool"
	"github.com/fogleman/gg"
)

type ColorData struct {
	R float64
	G float64
	B float64
	A float64
}

func GenCode(width, height, length int) string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 创建一个 4 位数字验证码
	code := generateVerificationCode(length)

	// 创建一个 widthxheight 像素的图像
	dc := gg.NewContext(width, height)

	// 设置图像背景颜色
	dc.SetRGBA(1, 1, 1, 1) // 白色背景
	dc.Clear()

	// 设置验证码文本颜色
	dc.SetRGB(0, 0, 0) // 黑色文本

	// 设置字体大小
	fontSize := 48
	f, _ := gg.LoadFontFace("ALGER.TTF", float64(fontSize))
	dc.SetFontFace(f)

	// 将验证码绘制到图像上
	dc.DrawStringAnchored(code, float64(width)/2, float64(height)/2, 0.5, 0.5)

	var buf bytes.Buffer
	if err := png.Encode(&buf, dc.Image()); err != nil {
		LogTool.LogErrorf("GenCode", "png encode error: %v", err)
	}
	base64Image := base64.StdEncoding.EncodeToString(buf.Bytes())

	return "data:image/png;base64," + base64Image
}

func generateVerificationCode(length int) string {
	const charset = "0123456789"
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
