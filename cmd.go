package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/littlefish12345/bilinetdrive"
)

func main() {
	var command, parameter, parameter1 string
	for {
		fmt.Printf("> ")
		fmt.Scanln(&command, &parameter, &parameter1)
		if command == "setSESSDATA" {
			bilinetdrive.SetSESSDATA(parameter)
		} else if command == "initRootNode" {
			hash, err := bilinetdrive.InitializeRootNode()
			if err != nil {
				fmt.Println(err)
			} else {
				bilinetdrive.SetRootNode(hash)
			}
		} else if command == "setRootNode" {
			bilinetdrive.SetRootNode(parameter)
		} else if command == "getRootNodeHash" {
			hash, err := bilinetdrive.GetRootNodeHash()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(hash)
			}
		} else if command == "ls" {
			fileList, err := bilinetdrive.ListFile()
			if err != nil {
				fmt.Println(err)
			} else {
				for i := 0; i < len(fileList); i++ {
					fmt.Printf(fileList[i][0])
					if fileList[i][1] == "0" {
						fmt.Printf("/")
					}
					fmt.Printf("\n")
				}
			}
		} else if command == "pwd" {
			path, err := bilinetdrive.GetPwd()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(path)
			}
		} else if command == "cd" {
			err := bilinetdrive.SwitchDir(parameter)
			if err != nil {
				fmt.Println(err)
			}
		} else if command == "mkdir" {
			err := bilinetdrive.MakeFolder(parameter)
			if err != nil {
				fmt.Println(err)
			}
		} else if command == "rm" {
			err := bilinetdrive.RemoveNode(parameter)
			if err != nil {
				fmt.Println(err)
			}
		} else if command == "rn" {
			err := bilinetdrive.RenameNode(parameter, parameter1)
			if err != nil {
				fmt.Println(err)
			}
		} else if command == "upload" {
			startTime := float64(time.Now().UnixNano()) / 1e9
			f, err := os.Open(parameter)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fileStat, _ := f.Stat()
			defer f.Close()
			err = bilinetdrive.UploadFile(f, filepath.Base(parameter))
			f.Close()
			if err != nil {
				fmt.Println(err)
				continue
			}
			endTime := float64(time.Now().UnixNano()) / 1e9
			fmt.Println(float64(fileStat.Size()) / (endTime - startTime) / 1024 / 1024)
		} else if command == "download" {
			startTime := float64(time.Now().UnixNano()) / 1e9
			f, err := os.Create(parameter)
			if err != nil {
				fmt.Println(err)
				continue
			}
			defer f.Close()
			//data, err := bilinetdrive.DownloadData(parameter)
			num, err := bilinetdrive.DownloadFile(parameter, f)
			if err != nil {
				fmt.Println(err)
				continue
			}
			//_, err = f.Write(data.Bytes())
			//if err != nil {
			//	fmt.Println(err)
			///	continue
			//}
			f.Close()
			endTime := float64(time.Now().UnixNano()) / 1e9
			fmt.Println(float64(num) / (endTime - startTime) / 1024 / 1024)
		} else if command == "exit" {
			hash, err := bilinetdrive.GetRootNodeHash()
			if err == nil {
				fmt.Println(hash)
			}
			os.Exit(0)
		} else {
			fmt.Println("Unknow command")
		}
	}
}
