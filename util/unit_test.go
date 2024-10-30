package util

import (
	"fmt"
	"testing"
)

/*
windows下测试通过
*/
func TestGetFolders(t *testing.T) {
	dir := "E:\\Downloads\\My Pack" // 替换为你要检查的目录
	folders, err := getFoldersWithTimestamps(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Folders containing timestamps.txt:")
	for _, folder := range folders {
		fmt.Println(folder)
	}
}
