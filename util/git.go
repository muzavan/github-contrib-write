package util

import (
	"fmt"
	"os/exec"
	"time"
)

// GitAdd specific file or pattern
func GitAdd(folderPath string, filePath string) error {
	// Add
	cmd := exec.Command("git", "add", filePath)
	cmd.Dir = folderPath

	return cmd.Run()
}

// GitCommit git commit in folderPath at date
func GitCommit(folderPath string, date time.Time) error {
	dateStr := date.Format("2006-01-02")
	dateParam := fmt.Sprintf(`--date="%s"`, dateStr)

	cmd := exec.Command("git", "commit", dateParam, "-m", `"Commit a generated content"`)
	cmd.Dir = folderPath
	return cmd.Run()
}

// GitInit on specific folder
func GitInit(folderPath string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = folderPath
	return cmd.Run()
}

// GitConfig ...
func GitConfig(folderPath string, configName string, configValue string) error {
	cmd := exec.Command("git", "config", "--local", configName, fmt.Sprintf(`"%s"`, configValue))
	cmd.Dir = folderPath
	return cmd.Run()
}
