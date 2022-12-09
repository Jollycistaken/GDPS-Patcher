package main

import (
	"encoding/base64"
	"fmt"
	"github.com/fatih/color"
	"github.com/sqweek/dialog"
	"os"
	"strings"
)

func makeError(err string) {
	color.Red(err)
	fmt.Println("Press enter to exit...")
	_, err2 := fmt.Scanln()
	if err2 != nil {
		color.Red(err2.Error())
		os.Exit(1)
	}
	os.Exit(1)
}

func main() {
	filename, err := dialog.File().Filter("Executables (*.exe)", "exe").Title("Select the GD exe that you want to patch").Load()
	if err != nil {
		makeError(err.Error())
	}

	var url string
	fmt.Printf("Enter your GDPS url: ")
	_, err = fmt.Scanln(&url)
	if err != nil {
		makeError(err.Error())
	}

	url = strings.Replace(url, "http://", "", -1)
	url = strings.Replace(url, "https://", "", -1)
	if last := len(url) - 1; last >= 0 && url[last] == '/' {
		url = url[:last]
	}
	b64Url := base64.StdEncoding.EncodeToString([]byte("http://" + url))
	file, err := os.Open(filename)
	if err != nil {
		makeError(err.Error())
	}
	stat, err := file.Stat()
	if err != nil {
		err2 := file.Close()
		if err2 != nil {
			makeError(err2.Error())
		}
		makeError(err.Error())
	}
	buffer := make([]byte, stat.Size())
	_, err = file.Read(buffer)
	if err != nil {
		err2 := file.Close()
		if err2 != nil {
			makeError(err2.Error())
		}
		makeError(err.Error())
	}
	exeBody := string(buffer)
	if strings.Contains(url, "/database") {
		if len(url) != len("www.boomlings.com/database") {
			makeError("Hex editing would corrupt the exe because your url is either too short or long!")
		}
		exeBody = strings.Replace(exeBody, "www.boomlings.com/database", url, -1)
		exeBody = strings.Replace(exeBody, "aHR0cDovL3d3dy5ib29tbGluZ3MuY29tL2RhdGFiYXNl", b64Url, -1)
	} else {
		if len(url) != len("www.boomlings.com") {
			makeError("Hex editing would corrupt the exe because your url is either too short or long!")
		}
		exeBody = strings.Replace(exeBody, "www.boomlings.com", url, -1)
		exeBody = strings.Replace(exeBody, "aHR0cDovL3d3dy5ib29tbGluZ3MuY29t", b64Url, -1)
	}
	err = file.Close()
	if err != nil {
		makeError(err.Error())
	}

	file, err = os.Create(strings.Replace(filename, ".exe", "_Patched.exe", -1))
	if err != nil {
		makeError(err.Error())
	}
	_, err = file.WriteString(exeBody)
	if err != nil {
		makeError(err.Error())
	}
	err = file.Close()
	if err != nil {
		makeError(err.Error())
	}

	color.Green("Successfully patched!")
	fmt.Println("Press enter to exit...")
	_, err = fmt.Scanln()
	if err != nil {
		makeError(err.Error())
	}
	os.Exit(0)
}
