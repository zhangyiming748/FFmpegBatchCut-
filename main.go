package main

import (
	"FFmpegBatchCut/ffmpeg"
	"FFmpegBatchCut/util"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	root := "F:\\原始视频\\分割完成"
	folders, _ := util.GetFoldersWithTimestamps(root)
	for _, folder := range folders {
		fmt.Printf("out:=%v\n", folders)
		mp4, _ := util.GetUniqueMP4File(folder)
		if mp4 == "" {
			mp4, _ = util.GetUniqueMKVFile(folder)
			if mp4 == "" {
				log.Printf("在 %v 目录下 mkv 和 mp4 均未找到\n", folder)
				continue
			}
		}
		timestampsFile := strings.Join([]string{folder, "timestamps.txt"}, string(os.PathSeparator))
		timestamps := util.ReadByLine(timestampsFile)
		timestamps = removeEmptyStrings(timestamps)
		log.Printf("目录%v\n文件%v\n时间戳%v\n", folder, mp4, timestamps)
		err := ffmpeg.CutOne(mp4, timestamps)
		if err != nil {
			log.Fatal(err)
		} else {
			if err := os.Remove(timestampsFile); err != nil {
				log.Printf("删除%v失败\n", timestamps)
			} else {
				if err := os.Remove(mp4); err != nil {
					log.Printf("删除%v失败\n", mp4)
				}
				log.Printf("分割文件结束,删除%v和%v失败\n", timestamps, mp4)
			}
		}
	}
}
func removeEmptyStrings(input []string) []string {
	var result []string
	for _, str := range input {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
