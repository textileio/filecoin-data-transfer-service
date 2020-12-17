package cmd

import (
	"fmt"
	"os"

	"github.com/blang/semver"
	"github.com/caarlos0/spin"
	"github.com/logrusorgru/aurora"
	su "github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
	bi "github.com/textileio/filecoin-data-transfer-service/buildinfo"
)

func install(assetURL string) error {
	s := spin.New("%s Downloading release")
	s.Start()
	defer s.Stop()
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	if err := su.UpdateTo(assetURL, exe); err != nil {
		return err
	}
	return nil
}

func getLatestRelease() (*su.Release, error) {
	s := spin.New("%s Checking latest fts release")
	s.Start()
	defer s.Stop()
	config := su.Config{
		Filters: []string{
			"fts",
		},
	}
	updater, err := su.NewUpdater(config)
	if err != nil {
		return nil, err
	}

	latest, found, err := updater.DetectLatest(Repo)
	if err != nil {
		return nil, err
	} else if !found {
		return nil, fmt.Errorf("Release not found")
	}
	return latest, nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update fts",
	Long:  `Update the installed fts CLI version to latest release.`,
	Args:  cobra.ExactArgs(0),
	Run: func(c *cobra.Command, args []string) {
		version := bi.Version

		latest, err := getLatestRelease()
		if err != nil {
			Warning("Unable to fetch latest public release.")
			checkErr(err)
		} else {
			current, err := semver.ParseTolerant(version)
			if err == nil {
				if current.LT(latest.Version) {
					if err = install(latest.AssetURL); err != nil {
						Warning("Error: install failed.")
						checkErr(err)
					} else {
						version = latest.Version.String()
						Message("Success: fts updated.")
					}
				} else {
					Message("Already up-to-date.")
				}
			} else {
				if err = install(latest.AssetURL); err != nil {
					Warning("Error: install failed.")
					checkErr(err)
				} else {
					version = latest.Version.String()
					Message("Success: fts updated.")
				}
			}
		}
		if version == "git" {
			Message("Custom version:")
			renderTable(
				os.Stdout,
				[]string{"GitBranch", "GitState", "GitSummary"},
				[][]string{{
					bi.GitBranch,
					bi.GitState,
					bi.GitSummary,
				}},
			)
			Message("%s (%s)", aurora.Green(bi.GitCommit), bi.BuildDate)
		} else {
			Message("%s", aurora.Green(version))
		}
	},
}
