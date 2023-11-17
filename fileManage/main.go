package main

import (
	_ "Img/config"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	//"gorm.io/driver/sqlite"
	//"gorm.io/gorm"
	"io"
	"os"
	"path"
	"strings"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("系统捕获错误：%+v", err)
		}
		fmt.Printf("结束")
	}()
	//pathName := "D:\\个人\\所有照片\\"
	//descPath := "D:\\个人\\data\\"
	//db, err := gorm.Open(sqlite.Open(descPath+"test.db"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//err = db.AutoMigrate(&fileImage{})
	//if err != nil {
	//	return
	//}
	//path := "D:\\个人\\手机图片处理\\全部照片"
	//fileNames, err := getFilesAndDirs(pathName)
	//if err != nil {
	//	return
	//}
	//for _, fileName := range fileNames {
	//	moveFile(fileName, descPath)
	//}
	//fmt.Printf("done")
}

func moveFile(fileImage fileImage, descDir string) {
	//panic(descDir + fileImage.ModTime + fileImage.Md5File + path.Ext(fileImage.FileName))
	_, err := copyFile(fileImage.FileName, descDir+fileImage.ModTime+path.Ext(fileImage.FileName))
	if err != nil {
		return
	}
}

func getFilesAndDirs(dirPath string) (files []fileImage, err error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	fileImages := make([]fileImage, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		//panic(info.ModTime().Format("2006-1-02 15:04:05"))
		if entry.IsDir() {
			childDir := dirPath + "\\" + info.Name()
			tempInfo, _ := getFilesAndDirs(childDir)
			fileImages = append(fileImages, tempInfo...)
		}
		if err != nil {
			return nil, err
		}
		fileName := dirPath + "\\" + info.Name()
		suffixes := []string{".png", ".PNG", "JPG", "jpg"}
		if hasSuffixes(fileName, suffixes) {
			//md5FileString, _ := getFileMd5(fileName)
			//modTime := info.ModTime().Format("2006102150405")
			modTime := info.ModTime().Format("2006-1-02-15-04-05")
			tempImage := fileImage{
				//Md5File:  md5FileString,
				ModTime:  modTime,
				Path:     dirPath,
				FileName: fileName,
			}
			fileImages = append(fileImages, tempImage)
		}
	}
	return fileImages, nil
}

func hasSuffixes(s string, suffixes []string) bool {
	for _, v := range suffixes {
		if strings.HasSuffix(s, v) {
			return true
		}
	}
	return false
}

func getFileMd5(fileName string) (s string, err error) {
	osFile, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer func(osFile *os.File) {
		err := osFile.Close()
		if err != nil {
			fmt.Printf("关闭文件失败：fileName=%v ,err= %+v ", fileName, err)
		}
	}(osFile)
	md5String := md5.New()
	_, err = io.Copy(md5String, osFile)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(md5String.Sum(nil)), nil
}

func copyFile(sourceFileName string, descFileName string) (bool, error) {
	file1, err := os.Open(sourceFileName)
	if err != nil {
		return false, err
	}
	file2, err := os.OpenFile(descFileName, os.O_WRONLY|os.O_CREATE, 0777)
	defer func(file1 *os.File, file2 *os.File) {
		_ = file1.Close()
		_ = file2.Close()
	}(file1, file2)

	_, err = io.Copy(file2, file1)
	if err != nil {
		return false, err
	}
	return true, nil
}

type fileImage struct {
	//gorm.Model
	Md5File  string
	ModTime  string
	Path     string
	FileName string
}
