package source

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/gookit/color"
)

func LocalFile(filename string) (urls []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Local file read error:", err)
		color.RGBStyleFromString("237,64,35").Println("[error] the input file is wrong!!!")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "http") {
			urls = append(urls, scanner.Text())
		} else {
			urls = append(urls, "https://"+scanner.Text())
		}
	}
	return
}
