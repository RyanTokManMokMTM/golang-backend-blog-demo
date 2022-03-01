package upload

import (
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

//GetFileName rename the file
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext) //remove the fileName suffix which is the extension type

	//encoding the file name
	fileName = util.EnCodeMD5(fileName)
	return fileName + ext //return a hash name.ext
}
func GetFileExt(name string) string {
	return path.Ext(name) //get the file extension .css .jpg etc
}
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

//Before image is being uploaded to server

func CheckSavePath(dest string) bool {
	_, err := os.Stat(dest)
	return os.IsNotExist(err) //err is not exist return false
}

//CheckAllowedExt check allowed uploaded file extension
func CheckAllowedExt(t FileType, name string) bool {
	//get file ext
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext) //change all ext to upper case for comparison

	switch t {
	case TypeImage:
		for _, allowed := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowed) == ext {
				return true
			}
		}
	}
	return false
}

//CheckMaxSize check uploaded file size
func CheckMaxSize(t FileType, file multipart.File) bool {
	fileBytes, _ := ioutil.ReadAll(file)
	fileSize := len(fileBytes)
	switch t {
	case TypeImage:
		//current file size is max or equal to allowed size
		if fileSize >= global.AppSetting.UploadImageMaxSize*1024*1024 { //in MB
			return true
		}
	}

	return false
}

//CheckPermission check server permission
func CheckPermission(dest string) bool {
	_, err := os.Stat(dest)
	return os.IsPermission(err) //err is not having permission return false
}

//After checking

//CreateSavePath create image resource path
func CreateSavePath(dest string, perm os.FileMode) error {
	err := os.Mkdir(dest, perm)
	if err != nil {
		return err
	}

	return nil
}

//SaveFile saving uploaded file to dest
func SaveFile(file *multipart.FileHeader, dest string) error {
	src, err := file.Open()
	if err != nil {
		return err
	} //open

	defer src.Close()

	fileOut, err := os.Create(dest) //create the file
	if err != nil {
		return err
	}
	defer fileOut.Close()

	_, err = io.Copy(fileOut, src) //copy to the source
	if err != nil {
		return err
	}
	return nil
}
