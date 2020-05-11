package cmd

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/twpayne/chezmoi/next/internal/chezmoi"
)

var managedCmd = &cobra.Command{
	Use:     "managed",
	Args:    cobra.NoArgs,
	Short:   "List the managed entries in the destination directory",
	Long:    mustGetLongHelp("managed"),
	Example: getExample("managed"),
	RunE:    config.runManagedCmd,
}

type managedCmdConfig struct {
	include *chezmoi.IncludeBits
}

func init() {
	rootCmd.AddCommand(managedCmd)

	persistentFlags := managedCmd.PersistentFlags()
	persistentFlags.VarP(config.managed.include, "include", "i", "include entry types")
}

func (c *Config) runManagedCmd(cmd *cobra.Command, args []string) error {
	s, err := c.getSourceState()
	if err != nil {
		return err
	}

	entries := s.Entries()
	targetNames := make([]string, 0, len(entries))
	for targetName, sourceStateEntry := range entries {
		targetStateEntry, err := sourceStateEntry.TargetStateEntry()
		if err != nil {
			return err
		}
		if !c.managed.include.Include(targetStateEntry) {
			continue
		}
		targetNames = append(targetNames, targetName)
	}

	sort.Strings(targetNames)
	sb := &strings.Builder{}
	for _, targetName := range targetNames {
		sb.WriteString(filepath.FromSlash(filepath.Join(c.DestDir, targetName)) + eolStr)
	}
	return c.writeOutputString(sb.String())
}
