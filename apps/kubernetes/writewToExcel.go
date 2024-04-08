package kubernetes

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func WriteClusterResourcesToExcel(clusterResources ClusterResources, timestamp string) error {
	f := excelize.NewFile()

	sheetName := "华为云CCE集群"
	if strings.EqualFold(f.GetSheetName(0), sheetName) {
	} else {
		f.SetSheetName("Sheet1", sheetName)

	}

	// 用于跟踪当前写入的行数
	var currentRow int
	// FLAG变量，用于判断是否已写入表头
	headerWritten := false

	// 遍历环境并写入 Excel
	for env, envResources := range clusterResources {
		if !headerWritten {
			// 设置表头（仅第一次迭代时执行）
			headers := []string{"环境", "命名空间", "资源", "名称", "副本数", "REQUESTS.CPU(core)", "REQUESTS.MEM总量(MiB)", "LIMITS.CPU(core)", "LIMITS.MEM总量(MiB)"}
			for i, header := range headers {
				cell, _ := excelize.CoordinatesToCellName(i+1, 1) // 列从1开始计数
				f.SetCellValue(sheetName, cell, header)
			}
			headerWritten = true // 设置表头已写入的标志
			currentRow = 2       // 设置起始行为2（表头占用了第1行）
		}

		// 遍历命名空间并写入信息
		for ns, nsResources := range envResources {
			for resourceKey, resourceInfo := range nsResources {
				f.SetCellValue(sheetName, fmt.Sprintf("A%d", currentRow), env)
				f.SetCellValue(sheetName, fmt.Sprintf("B%d", currentRow), ns)
				f.SetCellValue(sheetName, fmt.Sprintf("C%d", currentRow), resourceKey)
				f.SetCellValue(sheetName, fmt.Sprintf("D%d", currentRow), resourceInfo.Name)
				f.SetCellValue(sheetName, fmt.Sprintf("E%d", currentRow), fmt.Sprintf("%d", resourceInfo.Replicas))
				f.SetCellValue(sheetName, fmt.Sprintf("F%d", currentRow), fmt.Sprintf("%.2f", resourceInfo.TotalRequestsCpu))
				f.SetCellValue(sheetName, fmt.Sprintf("G%d", currentRow), fmt.Sprintf("%.2f", resourceInfo.TotalRequestsMemory))
				f.SetCellValue(sheetName, fmt.Sprintf("H%d", currentRow), fmt.Sprintf("%.2f", resourceInfo.TotalLimitsCpu))
				f.SetCellValue(sheetName, fmt.Sprintf("I%d", currentRow), fmt.Sprintf("%.2f", resourceInfo.TotalLimitsMemory))
				currentRow++ // 移动到下一行以准备写入下一个资源信息
			}
		}
	}

	// 设置列的宽度（可根据需要调整）
	f.SetColWidth(sheetName, "A", "A", 20)
	f.SetColWidth(sheetName, "B", "B", 30)
	f.SetColWidth(sheetName, "C", "C", 60)
	f.SetColWidth(sheetName, "D", "D", 60)
	f.SetColWidth(sheetName, "E", "I", 15)

	// 自动筛选
	f.AutoFilter(sheetName, "A1:I1", []excelize.AutoFilterOptions{{Column: "A", Expression: "x != blanks"}})

	// 使用格式化后的时间戳来构造文件名
	dst := fmt.Sprintf("./output/华为云CCE集群_资源成本精细化分账报表_%s.xlsx", timestamp)

	// 保存 Excel 文件
	if err := f.SaveAs(dst); err != nil {
		return err
	}
	return nil
}
