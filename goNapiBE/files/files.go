package files

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type File struct {
	Path      string `json:"path"`
	Name      string `json:"fileName"`
	Extension string `json:"fileExtension"`
	Size      int64  `json:"sizeInBytes"`
	ModTime   string `json:"lastModification"`
}

type Directory struct {
	Name    string `json:"folderName"`
	Size    int64  `json:"sizeInBytes"`
	ModTime string `json:"lastModification"`
}

func ContainsString(slice []string, value string) bool {
	for _, sliceEntry := range slice {
		if sliceEntry == value {
			return true
		}
	}
	return false
}

func GetFullDirectoryList(path string) []Directory {
	ls, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	var directories []Directory

	for _, d := range ls {
		if d.IsDir() && d.Name()[0:1] != "." {
			var info, _ = d.Info()
			var dir = Directory{
				Name:    d.Name(),
				Size:    info.Size(),
				ModTime: info.ModTime().Local().Format("2006-01-02-15:04:05"),
			}
			directories = append(directories, dir)
		}
	}

	return directories
}

func ObtainDirectoryList(path string) []string {
	ls, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	var directories []string

	for _, d := range ls {
		if d.IsDir() && d.Name()[0:1] != "." {
			directories = append(directories, d.Name())
		}
	}

	return directories
}

func ObtainFileList(rootPath string, extensions []string) []File {

	var fileList []File

	filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if d.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		if ContainsString(extensions, ext) {
			var info, _ = d.Info()
			var newFile = File{
				Path:      path,
				Name:      d.Name(),
				Extension: ext,
				Size:      info.Size(),
				ModTime:   info.ModTime().Local().Format("2006-01-02-15:04:05"),
			}
			fileList = append(fileList, newFile)
		}

		return nil
	})

	return fileList
}
