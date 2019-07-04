package ximg

import (
	"bufio"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/util/gconv"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// 裁剪图像
func CutImage(w *ghttp.Response, path string, width, height int) {
	// 没有宽高，就是在加载原图像
	if width == 0 && height == 0 {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			NoImage(w)
			log.Println("file, err = os.Open(path)", err)
			return
		}
		w.Write(b)
		return
	}
	// 裁剪图像 --------------------------------------
	// 裁剪图像的组合路径
	var str = strings.Builder{}
	str.WriteString(path)
	str.WriteString("_")
	str.WriteString(strconv.Itoa(width))
	str.WriteString("_")
	str.WriteString(strconv.Itoa(height))
	CutPath := str.String()
	// 判断是否存在裁剪图像
	b, err := ioutil.ReadFile(CutPath)
	if err == nil {
		w.Write(b)
		return
	}
	// 原图像
	file, err := os.Open(path)
	if err != nil {
		NoImage(w)
		log.Println("file, err = os.Open(path)", err)
		return
	}
	defer file.Close()
	// 图片解码 --------------------------------------
	bufFile := bufio.NewReader(file)
	img, imgType, err := image.Decode(bufFile)
	if err != nil {
		NoImage(w)
		log.Println("img, imgType, err := image.Decode(bufFile)", err)
		return
	}
	// 要裁剪的宽高不能大于自身的宽高
	RWidth := img.Bounds().Max.X
	if width > RWidth {
		width = RWidth
	}
	RHeight := img.Bounds().Max.Y
	if height > RHeight {
		height = RHeight
	}
	// gif 图就不处理了
	if imgType == GIF || (width == RWidth && height == RHeight) {
		b, err := ioutil.ReadFile(CutPath)
		if err != nil {
			NoImage(w)
			log.Println("gif, err = os.Open(path)", err)
			return
		}
		w.Write(b)
		// 设置文件的偏移量 - 因为文件被 image.Decode 后文件的偏移量到尾部
		//_, _ = file.Seek(0, 0)
		// 向浏览器输出
		//_, _ = io.Copy(w.ResponseWriter.ResponseWriter, file)
		return
	}
	// 进行裁剪
	reImg := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	// 裁剪的存储
	out, err := os.Create(CutPath)
	if err != nil {
		NoImage(w)
		log.Println("out, err := os.Create(CutPath)", err)
		return
	}
	defer out.Close()
	if imgType == JPEG || imgType == JPG {
		// 保存裁剪的图片
		_ = jpeg.Encode(out, reImg, nil)
		// 向浏览器输出
		//_ = jpeg.Encode(w.ResponseWriter.ResponseWriter, reImg, nil)
		b, err := ioutil.ReadFile(CutPath)
		if err != nil {
			NoImage(w)
			log.Println("gif, err = os.Open(path)", err)
			return
		}
		w.Write(b)
	} else if imgType == PNG {
		// 保存裁剪的图片
		_ = png.Encode(out, reImg)
		b, err := ioutil.ReadFile(CutPath)
		if err != nil {
			NoImage(w)
			log.Println("gif, err = os.Open(path)", err)
			return
		}
		w.Write(b)
		// 向浏览器输出
		//_ = png.Encode(w.ResponseWriter.ResponseWriter, reImg)
	}
}

// 用于找不到图片时用
func NoImage(r *ghttp.Response) {
	// 图片流方式输出
	r.Header().Set("Content-Type", "image/png")
	// 进行图片的编码
	//_ = png.Encode(r.ResponseWriter.ResponseWriter, noImg)
	r.Write(gconv.Bytes(noImg))
}
