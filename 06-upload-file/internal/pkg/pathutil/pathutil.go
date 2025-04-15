package pathutil

import (
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

// GenerateFilePath 根据原始文件名生成一个唯一的存储路径
// 生成的路径格式为：年/月/日/UUID.扩展名
// 例如：2023/04/15/550e8400-e29b-41d4-a716-446655440000.png
//
// 参数:
//   - filename: 原始文件名，用于提取文件扩展名
//
// 返回值:
//   - 生成的文件路径字符串
func GenerateFilePath(filename string) string {
	// 提取文件扩展名（例如 .jpg, .png 等）
	ext := filepath.Ext(filename)

	// 生成基于当前日期的目录结构（年/月/日/）
	// 使用 2006/01/02 是 Go 的时间格式化特定写法，表示年/月/日
	dateDir := time.Now().Format("2006/01/02/")

	// 生成 UUID 作为文件名，确保文件名唯一，防止同名文件覆盖
	// 最后拼接原始文件的扩展名
	return dateDir + uuid.New().String() + ext
}
