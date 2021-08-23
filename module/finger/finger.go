package finger

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gookit/color"
)

type Outrestul struct {
	Url        string `json:"url"`
	Cms        string `json:"cms"`
	Server     string `json:"server"`
	Statuscode int    `json:"statuscode"`
	Length     int    `json:"length"`
	Title      string `json:"title"`
}

func MapToJson(param map[string][]string) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func fingerscan(allresult chan Outrestul, ch chan string, result chan Outrestul, finpx *Packjson) {
	chsize := len(ch)
	cycle := 1
	for len(ch) != 0 {
		url := <-ch
		var data *resps
		data, err := httprequest(url, cycle, chsize)
		if err != nil {
			continue
		}
		if data.statuscode == 400 {
			data, err = httprequest(url, cycle, chsize)
			if err != nil {
				continue
			}
		}
		cycle++
		for _, jurl := range data.jsurl {
			if jurl != "" {
				ch <- jurl
			}
		}
		headers := MapToJson(data.header)
		var cms []string
		for _, finp := range finpx.Fingerprint {
			if finp.Location == "body" {
				if finp.Method == "keyword" {
					if iskeyword(data.body, finp.Keyword) {
						cms = append(cms, finp.Cms)
					}
				}
				if finp.Method == "faviconhash" {
					if data.favhash == finp.Keyword[0] {
						cms = append(cms, finp.Cms)
					}
				}
				if finp.Method == "regular" {
					if isregular(data.body, finp.Keyword) {
						cms = append(cms, finp.Cms)
					}
				}
			}
			if finp.Location == "header" {
				if finp.Method == "keyword" {
					if iskeyword(headers, finp.Keyword) {
						cms = append(cms, finp.Cms)
					}
				}
				if finp.Method == "regular" {
					if isregular(headers, finp.Keyword) {
						cms = append(cms, finp.Cms)
					}
				}
			}
			if finp.Location == "title" {
				if finp.Method == "keyword" {
					if iskeyword(data.title, finp.Keyword) {
						cms = append(cms, finp.Cms)
					}
				}
				if finp.Method == "regular" {
					if isregular(data.title, finp.Keyword) {
						cms = append(cms, finp.Cms)
					}
				}
			}
		}
		cms = RemoveDuplicatesAndEmpty(cms)
		cmss := strings.Join(cms, ",")
		out := Outrestul{data.url, cmss, data.server, data.statuscode, data.length, data.title}
		allresult <- out
		if len(out.Cms) != 0 {
			s := fmt.Sprintf("[ %s | %s | %s | %d | %d | %s ]", out.Url, out.Cms, out.Server, out.Statuscode, out.Length, out.Title)
			color.RGBStyleFromString("237,64,35").Println(s)
			result <- out
		} else {
			s := fmt.Sprintf("[ %s | %s | %s | %d | %d | %s ]", out.Url, out.Cms, out.Server, out.Statuscode, out.Length, out.Title)
			fmt.Println(s)
		}

	}
}

func Fingermain(urls []string, thread int, output string) {
	err1 := LoadWebfingerprint("./finger.json")
	if err1 != nil {
		//log.Println("fingerprint file read error:", err1)
		color.RGBStyleFromString("237,64,35").Println("[error] fingerprint file error!!!")
		os.Exit(1)
	}
	ch := make(chan string, len(urls))
	result := make(chan Outrestul, len(urls)*2)
	allresult := make(chan Outrestul, len(urls)*2)
	finpx := GetWebfingerprint()
	//fmt.Println(finpx)
	for _, url := range urls {
		ch <- url
	}
	for i := 0; i <= thread; i++ {
		go fingerscan(allresult, ch, result, finpx)
	}
	for {
		//fmt.Println(len(ch))
		if len(ch) > 0 {
			time.Sleep(10 * time.Second)
		} else {
			time.Sleep(5 * time.Second)
			if len(ch) == 0 {
				break
			} else {
				continue
			}

		}
	}
	close(ch)
	color.RGBStyleFromString("244,211,49").Println("\n重点资产：")
	if len(result) == 0 {
		close(result)
	}
	for aas := range result {
		fmt.Printf(fmt.Sprintf("[ %s | ", aas.Url))
		color.RGBStyleFromString("237,64,35").Printf(fmt.Sprintf("%s", aas.Cms))
		fmt.Printf(fmt.Sprintf(" | %s | %d | %d | %s ]\n", aas.Server, aas.Statuscode, aas.Length, aas.Title))
		if len(result) == 0 {
			close(result)
		}
	}
	if output != "" {
		outfile(output, allresult)
	}
}
