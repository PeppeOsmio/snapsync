package utils

import (
	"fmt"
	"path"
	"snapsync/structs"
	"strconv"
	"strings"
)

func HumanReadableSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func GetInfoFromSnapshotBasePath(snapshotBasePath string) (snapshotName string, number int, err error) {
	items := strings.Split(snapshotBasePath, ".")
	if len(items) != 2 {
		return snapshotName, number, fmt.Errorf("snapshot name %s must be in format <name>.<interval>", snapshotBasePath)
	}
	snapshotName = items[0]
	number, err = strconv.Atoi(items[1])
	if err != nil {
		return snapshotName, number, fmt.Errorf("can't parse snapshot number: %s", snapshotBasePath)
	}
	return snapshotName, number, nil
}

func GetInfoFromSnapshotPath(snapshotAbspath string) (snapshotInfo *structs.SnapshotInfo, err error) {
	snapshotDirName := strings.TrimSuffix(path.Base(snapshotAbspath), "/")
	name, number, err := GetInfoFromSnapshotBasePath(snapshotDirName)
	if err != nil {
		return nil, err
	}
	return &structs.SnapshotInfo{
		Abspath:      snapshotAbspath,
		SnapshotName: name,
		Number:       number,
	}, nil
}
