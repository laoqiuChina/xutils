package ximg

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

// 输出图片
func (m ImgServer) Get(w http.ResponseWriter, r *http.Request) {
	urlParse := r.URL.String()
	// 组合文件完整路径
	filePath := UrlParse(urlParse[1:])
	if filePath == "" {
		_, _ = w.Write(showMain())
		return
	}
	// 获取要裁剪图像的宽度、高度
	width := StringToInt(r.FormValue("w"))  // 宽度
	height := StringToInt(r.FormValue("h")) // 高度
	// 加载图片
	CutImage(w, filePath, width, height)
}

// 上传图片
func (m ImgServer) Post(w http.ResponseWriter, r *http.Request) {

	// 响应返回
	res := new(UpdateResponse)

	// 上传表单 --------------------------------------

	// 缓冲的大小 - 4M
	_ = r.ParseMultipartForm(1024 << 12)
	//是上传表单域的名字fileHeader
	upfile, upFileInfo, err := r.FormFile("userfile")
	if err != nil {

		res.Code = StatusForm
		res.Msg = StatusText(StatusForm)
		_, _ = w.Write(ResponseJson(res))

		return
	}
	defer upfile.Close()

	// 图片解码 --------------------------------------

	// 读入缓存
	bufUpFile := bufio.NewReader(upfile)
	// 进行图片的解码
	img, imgType, err := image.Decode(bufUpFile)
	if err != nil {

		res.Code = StatusImgDecode
		res.Msg = StatusText(StatusImgDecode)
		_, _ = w.Write(ResponseJson(res))

		return
	}

	// 判断是否有这个图片类型
	if !IsType(imgType) {

		res.Code = StatusImgIsType
		res.Msg = StatusText(StatusImgIsType)
		_, _ = w.Write(ResponseJson(res))

		return
	}

	// 设置文件读写下标 --------------------------------

	// 设置下次读写位置（移动文件指针位置）
	_, err = upfile.Seek(0, 0)
	if err != nil {

		res.Code = StatusFileSeek
		res.Msg = StatusText(StatusFileSeek)
		_, _ = w.Write(ResponseJson(res))

		return
	}

	// 计算文件的 MD5 值 -----------------------------

	// 初始化 MD5 实例
	md5Hash := md5.New()
	// 读入缓存
	bufFile := bufio.NewReader(upfile)
	_, err = io.Copy(md5Hash, bufFile)
	if err != nil {

		res.Code = StatusFileMd5
		res.Msg = StatusText(StatusFileMd5)
		_, _ = w.Write(ResponseJson(res))

		return
	}
	// 进行 MD5 算计，返回 16进制的 byte 数组
	fileMd5FX := md5Hash.Sum(nil)
	fileMd5 := fmt.Sprintf("%x", fileMd5FX)

	// 目录计算 --------------------------------------

	// 组合文件完整路径
	dirPath := JoinPath(fileMd5) + "/" // 目录
	filePath := dirPath + fileMd5      // 文件路径

	// 获取目录信息，并创建目录
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		err = os.MkdirAll(dirPath, 0666)
		if err != nil {

			res.Code = StatusMkdir
			res.Msg = StatusText(StatusMkdir)
			_, _ = w.Write(ResponseJson(res))

			return
		}
	} else {
		if !dirInfo.IsDir() {
			err = os.MkdirAll(dirPath, 0666)
			if err != nil {

				res.Code = StatusMkdir
				res.Msg = StatusText(StatusMkdir)
				_, _ = w.Write(ResponseJson(res))

				return
			}
		}
	}

	// 存入文件 --------------------------------------

	_, err = os.Stat(filePath)
	if err != nil {

		// 打开一个文件,文件不存在就会创建
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {

			res.Code = StatusOpenFile
			res.Msg = StatusText(StatusOpenFile)
			_, _ = w.Write(ResponseJson(res))

			return
		}
		defer file.Close()

		if imgType == PNG {
			err = png.Encode(file, img)

		} else if imgType == JPG || imgType == JPEG {
			err = jpeg.Encode(file, img, nil)

		} else if imgType == GIF {

			// 重新对 gif 格式进行解码
			// image.Decode 只能读取 gif 的第一帧

			// 设置下次读写位置（移动文件指针位置）
			_, err = upfile.Seek(0, 0)
			if err != nil {

				res.Code = StatusFileSeek
				res.Msg = StatusText(StatusFileSeek)
				_, _ = w.Write(ResponseJson(res))

				return
			}

			gifimg, giferr := gif.DecodeAll(upfile)
			if giferr != nil {

				res.Code = StatusImgDecode
				res.Msg = StatusText(StatusImgDecode)
				_, _ = w.Write(ResponseJson(res))

				return
			}
			err = gif.EncodeAll(file, gifimg)

		}

		if err != nil {
			res.Code = StatusImgEncode
			res.Msg = StatusText(StatusImgEncode)
			_, _ = w.Write(ResponseJson(res))

			return
		}

	}

	res.Success = true
	res.Code = StatusOK
	res.Msg = StatusText(StatusOK)
	res.Data.ImgId = fileMd5
	res.Data.Mime = imgType
	res.Data.Size = upFileInfo.Size

	_, _ = w.Write(ResponseJson(res))

}
