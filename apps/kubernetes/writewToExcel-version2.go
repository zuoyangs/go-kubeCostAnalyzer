package kubernetes

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func WriteClusterResourcesToExcel(clusterResources ClusterResources, timestamp string) error {
	f := excelize.NewFile()

	// 遍历环境并写入 Excel
	for env, envResources := range clusterResources {
		// 为每个环境创建一个 sheet
		sheetName := fmt.Sprintf("集群环境%s", env)
		f.SetSheetName("Sheet1", sheetName)
		f.NewSheet(sheetName)

		// 设置表头
		headers := []string{"命名空间", "资源", "名称", "副本数", "REQUESTS.CPU(core)", "REQUESTS.MEM总量(MiB)", "LIMITS.CPU(core)", "LIMITS.MEM总量(MiB)"}
		for i, header := range headers {
			cell, _ := excelize.CoordinatesToCellName(1+i, 1)
			f.SetCellValue(sheetName, cell, header)
		}

		row := 2
		// 遍历命名空间并写入信息
		for ns, nsResources := range envResources {
			for resourceKey, resourceInfo := range nsResources {
				f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), fmt.Sprintf("%s", resourceInfo.ClusterName))
				f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), ns)
				f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), resourceKey)
				f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), resourceInfo.Name)
				f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), fmt.Sprintf("%d", resourceInfo.Replicas))
				f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), fmt.Sprintf("%.2f", resourceInfo.TotalRequestsCpu))
				f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), fmt.Sprintf("%.2f", resourceInfo.TotalRequestsMemory))
				f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), fmt.Sprintf("%.2f", resourceInfo.TotalLimitsCpu))
				f.SetCellValue(sheetName, fmt.Sprintf("I%d", row), fmt.Sprintf("%.2f", resourceInfo.TotalLimitsMemory))
				row++
			}
		}

		// 设置列的宽度
		f.SetColWidth(sheetName, "A", "A", 20)
		f.SetColWidth(sheetName, "B", "B", 30)
		f.SetColWidth(sheetName, "C", "H", 15)

		// 自动筛选
		f.AutoFilter(sheetName, "A1:H1", []excelize.AutoFilterOptions{{Column: "A", Expression: "x != blanks"}})
	}

	// 使用格式化后的时间戳来构造文件名
	dst := fmt.Sprintf("./output/华为云CCE集群_资源成本精细化分账报表_%s.xlsx", timestamp)

	// 保存 Excel 文件
	if err := f.SaveAs(dst); err != nil {
		return err
	}
	return nil
}
