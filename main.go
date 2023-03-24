package main

import (
	"filespliting/file/sender"
	"fmt"
)

func main() {
	sender.InitFolders() // to initialize the core folders
	fmt.Println("Sub Project [File Spliting and Rearranging]")
	fmt.Println("Chooose:\n1.Split and Send File\n2.Rearrange File")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Enter the name of the file")
		var fileName sender.FileName
		fmt.Scan(&fileName)
		fmt.Println("The Entered file name is :" + fileName)
		fileName.HandleFile()
	}
	for {

	}
}
