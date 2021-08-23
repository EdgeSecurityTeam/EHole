package finger

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func outjson(filename string, data []byte) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		defer f.Close()
		_, err = f.Write(data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func outxlsx(filename string, msg []Outrestul) {
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "url")
	xlsx.SetCellValue("Sheet1", "B1", "cms")
	xlsx.SetCellValue("Sheet1", "C1", "server")
	xlsx.SetCellValue("Sheet1", "D1", "statuscode")
	xlsx.SetCellValue("Sheet1", "E1", "length")
	xlsx.SetCellValue("Sheet1", "F1", "title")
	for k, v := range msg {
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(k+2), v.Url)
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(k+2), v.Cms)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(k+2), v.Server)
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(k+2), v.Statuscode)
		xlsx.SetCellValue("Sheet1", "E"+strconv.Itoa(k+2), v.Length)
		xlsx.SetCellValue("Sheet1", "F"+strconv.Itoa(k+2), v.Title)
	}
	err := xlsx.SaveAs(filename)
	if err != nil {
		fmt.Println(err)
	}
}

func outfile(filename string, allresult chan Outrestul) {
	file := strings.Split(filename, ".")
	var allresults []Outrestul
	for result := range allresult {
		allresults = append(allresults, result)
		if len(allresult) == 0 {
			close(allresult)
		}
	}
	if len(file) == 2 {
		if file[1] == "json" {
			buf, err := json.MarshalIndent(allresults, "", " ")
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			outjson(filename, buf)
		}
		if file[1] == "xlsx" {
			outxlsx(filename, allresults)
		}
	}

}
