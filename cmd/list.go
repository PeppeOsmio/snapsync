/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"snapsync/configs"
	"snapsync/core"
	"snapsync/utils"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the snapshots",
	Long:  `List the snapshots`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configFilePath, err := cmd.Flags().GetString("config-file")
		if err != nil {
			slog.Error("can 't get configs-dir flag")
			return
		}
		expandVars, err := cmd.Flags().GetBool("expand-vars")
		if err != nil {
			slog.Error("can 't get expand-vars flag")
			return
		}
		config, err := configs.LoadConfig(configFilePath, expandVars)
		if err != nil {
			slog.Error("can't get " + configFilePath + ": " + err.Error())
			return
		}
		err = core.RunInitCommands(config)
		if err != nil {
			slog.Error(err.Error())
			return
		}

		snapshotToList := args[0]
		snapshotsInfo, err := core.GetSnapshotsInfo(config.SnapshotsConfigsDir, expandVars, snapshotToList)
		if err != nil {
			slog.Error("Can't get snapshots of snapshot " + snapshotToList + ": " + err.Error())
			return
		}
		for _, snapshotInfo := range snapshotsInfo {
			size, err := snapshotInfo.Size()
			sizeStr := ""
			if err != nil {
				sizeStr = fmt.Sprintf("can't evaluate snapshot size: %s", err.Error())
			} else {
				sizeStr = utils.HumanReadableSize(size)
			}
			fmt.Printf("%s, size: %s\n", snapshotInfo.CompactName(), sizeStr)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
