package structs

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	LogLevel            string   `yaml:"log_level"`
	CpPath              string   `yaml:"cp_path"`
	RSyncPath           string   `yaml:"rsync_path"`
	SnapshotsConfigsDir string   `yaml:"snapshots_configs_dir"`
	InitCommands        []string `yaml:"init_commands"`
}

type SnapshotConfig struct {
	SnapshotName                  string        `yaml:"snapshot_name"`
	Dirs                          []SnapshotDir `yaml:"dirs"`
	SnapshotsDir                  string        `yaml:"snapshots_dir"`
	Retention                     int           `yaml:"retention"`
	Cron                          string        `yaml:"cron"`
	AlwaysRunPostSnapshotCommands bool          `yaml:"always_run_post_snapshot_commands"`
	PreSnapshotCommands           []string      `yaml:"pre_snapshot_commands"`
	PostSnapshotCommands          []string      `yaml:"post_snapshot_commands"`
}

type SnapshotDir struct {
	SrcDirAbspath    string   `yaml:"src_dir_abspath"`
	DstDirInSnapshot string   `yaml:"dst_dir_in_snapshot"`
	Excludes         []string `yaml:"excludes"`
}

type SnapshotInfo struct {
	Abspath      string
	SnapshotName string
	Number       int
}

func (snapshotInfo *SnapshotInfo) CompactName() string {
	return fmt.Sprintf("%s.%d", snapshotInfo.SnapshotName, snapshotInfo.Number)
}

func (snapshotInfo *SnapshotInfo) Size() (size int64, err error) {
	err = filepath.Walk(snapshotInfo.Abspath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}
