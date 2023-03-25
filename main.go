package main

import (
	"filespliting/file"
	"filespliting/file/reciever"
	"filespliting/file/sender"
	"fmt"
)

func main() {
	sender.InitFolders() // to initialize the core folders
	fmt.Println("Sub Project [File Spliting and Rearranging]")
	fmt.Println("Chooose:\n1.Split and Send File\n2.Rearrange File\n3.List Available files")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Enter the name of the file")
		var fileName sender.FileName
		fmt.Scan(&fileName)
		fmt.Println("The Entered file name is :" + fileName)
		fileName.HandleFile()
		break
	case 2:
		fmt.Println("The following are all the file infos availabe to the node")
		fmt.Println("=========================================================")
		fileInfoList := file.ListAvailableFiles()
		if fileInfoList == nil {
			break
		}
		for i, file := range fileInfoList {
			fmt.Println(i, "| [NAME:"+file.FileName+"]"+"[TYPE:"+file.FileType+"]"+"[ID:"+file.UniqueID.String()+"]")
		}
		var ch int
		fmt.Scan(&ch)
		fileInfo := reciever.FileBasicInfo(fileInfoList[ch]).RetreiveFullInfo()
		fileInfo.Merge()
		break
	case 3:
		fmt.Println("The following are all the file infos availabe to the node")
		fmt.Println("=========================================================")
		fileInfoList := file.ListAvailableFiles()
		if fileInfoList == nil {
			break
		}
		for i, file := range fileInfoList {
			fmt.Println(i, "| [NAME:"+file.FileName+"]"+"[TYPE:"+file.FileType+"]"+"[ID:"+file.UniqueID.String()+"]")
		}
		break
	}
	for {

	}
}
