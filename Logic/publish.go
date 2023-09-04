package Logic

import (
	"awesomeProject4/common"
	"awesomeProject4/common/model"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func PublishVideo(userName string, saveFile string, title string, finalName string) {
	db := common.GetDB()
	var userInfoTable model.UserInfoTable
	db.Take(&userInfoTable, "Name = ?", userName)

	videoPath := saveFile
	outputDir := "./public"
	generateVideoCover(videoPath, outputDir, finalName)
	playUrl := "http://192.168.37.1:8080/static/" + finalName
	coverUrl := "http://192.168.37.1:8080/static/" + finalName + ".jpg"
	userInfoTable.WorkCount = userInfoTable.WorkCount + 1
	db.Save(&userInfoTable)
	videoTable := model.VideoTable{
		UserInfoTableId: userInfoTable.ID,
		User:            userInfoTable,
		PlayUrl:         playUrl,
		CoverUrl:        coverUrl,
		Title:           title,
		PublishTime:     time.Now().Unix(),
	}
	db.Create(&videoTable)

}
func generateVideoCover(videoPath, outputDir, finalName string) {
	outputCover := filepath.Join(outputDir, finalName+".jpg")

	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "00:00:03", "-vframes", "1", outputCover)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("无法生成封面:", err)
		return
	}
	fmt.Println("封面已生成:", outputCover)
}
