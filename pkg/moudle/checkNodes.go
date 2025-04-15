package moudle

import (
	"SafelineAPI/internal/app/safeLineApi"
	"time"
)

func CheckNodes(nodes safeLineApi.Nodes, n int) safeLineApi.Nodes {
	var need safeLineApi.Nodes
	date := time.Now()
	for _, node := range nodes {
		days := int(node.ValidBefore.Sub(date).Hours() / 24)
		if days <= n {
			need = append(need, node)
		}
	}
	return need
}
