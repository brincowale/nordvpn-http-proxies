package output

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"nordvpn-proxies/proxies"
	"strconv"
)

var f = excelize.NewFile()
var pos = 2
func init() {
	f.NewSheet("Proxies")
	f.DeleteSheet("Sheet1")
	f.SetActiveSheet(0)
	f.SetCellValue("Proxies", "A1", "ID")
	f.SetCellValue("Proxies", "B1", "Name")
	f.SetCellValue("Proxies", "C1", "Hostname")
	f.SetCellValue("Proxies", "D1", "Load")
	f.SetCellValue("Proxies", "E1", "Station")
}

func Write(proxy proxies.Proxy) {
	row := strconv.Itoa(pos)
	pos++
	f.SetCellValue("Proxies", "A" + row, proxy.ID)
	f.SetCellValue("Proxies", "B" + row, proxy.Name)
	f.SetCellValue("Proxies", "C" + row, proxy.Hostname)
	f.SetCellValue("Proxies", "D" + row, proxy.Load)
	f.SetCellValue("Proxies", "E" + row, proxy.Station)
}

func Save() {
	err := f.SaveAs("./proxies.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}