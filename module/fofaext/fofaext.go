package fofaext

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Fofaext(msg [][]string, filename string) {
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "ip")
	xlsx.SetCellValue("Sheet1", "B1", "host")
	xlsx.SetCellValue("Sheet1", "C1", "title")
	xlsx.SetCellValue("Sheet1", "D1", "port")
	xlsx.SetCellValue("Sheet1", "E1", "protocol")
	for k, v := range msg {
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(k+2), v[0])
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(k+2), v[1])
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(k+2), v[2])
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(k+2), v[3])
		xlsx.SetCellValue("Sheet1", "E"+strconv.Itoa(k+2), v[4])
	}
	err := xlsx.SaveAs(filename)
	if err != nil {
		fmt.Println(err)
	}
}
