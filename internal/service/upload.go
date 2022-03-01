package service

import (
	"errors"
	"github.com/RyanTokManMokMTM/blog-service/global"
	"github.com/RyanTokManMokMTM/blog-service/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessURL string
}

/*
UploadFile
1.Get File Info from FileHeader
2.Get Save Path
3.Check image extension
4.Check save path is existed
5.Check Image size
6.Check save path permission
7.Save
*/
func (serve *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	//Get the file name
	fileName := upload.GetFileName(fileHeader.Filename)

	//Get Saving path
	savePath := upload.GetSavePath()

	//set the destination
	dest := savePath + "/" + fileName

	//check extension
	if !upload.CheckAllowedExt(fileType, fileName) {
		//not allowed
		return nil, errors.New("")
	}

	//check save path exist
	if upload.CheckSavePath(savePath) { //not exist
		err := upload.CreateSavePath(savePath, os.ModePerm) //set to 0777
		if err != nil {
			return nil, errors.New("failed to create save path/directory")
		}
	}

	//check size
	if upload.CheckMaxSize(fileType, file) { //over the max size
		return nil, errors.New("exceeded maximum file limit")
	}
	//check permission
	if upload.CheckPermission(savePath) { //no permission to access save path
		return nil, errors.New("deny permission")
	}
	//save file
	err := upload.SaveFile(fileHeader, dest)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		Name:      fileName,
		AccessURL: global.AppSetting.UploadSavePathURL + "/" + fileName,
	}, nil
}
