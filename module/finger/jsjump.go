package finger

import (
	"log"
	"regexp"
	"strings"
)

func xegexpjs(reg string, resp string) (reslut1 [][]string) {
	reg1 := regexp.MustCompile(reg)
	if reg1 == nil {
		log.Println("regexp err")
		return nil
	}
	result1 := reg1.FindAllStringSubmatch(resp, -1)
	return result1
}

func Jsjump(str string, url string) []string {
	regs := []string{`(window|top)\.location\.href = "(.*?)"`, `redirectUrl = "(.*?)"`, `<meta.*?http-equiv=.*?refresh.*?url=(.*?)>`}
	var results []string
	for _, reg := range regs {
		result1 := xegexpjs(reg, str)
		//fmt.Println(result1)
		if result1 != nil {
			if len(result1) > 0 {
				for _, m := range result1 {
					s := len(m)
					if strings.Contains(m[s-1], "http") {
						continue
					} else {
						str2 := strings.Trim(m[s-1], "/")
						str2 = strings.ReplaceAll(str2, "../", "/")
						if len(str2) != 0 {
							if str2[:1] == "/" {
								results = append(results, url+str2)
							} else {
								results = append(results, url+"/"+str2)
							}
						}
					}
				}
			} else {
				continue
			}
		}
	}
	return results
}
