package cmd

import (
	"bytes"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/lensesio/tableprinter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "The list of all the available components",
	Long: `The list of all the available components.
	
  # Show spec of the compositecluster resource
  kubectl idp list`,
	Run: func(cmd *cobra.Command, args []string) {
		listXrds()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listXrds() XRDs {
	xrds := getXrds()
	table := []tableRow{}
	for _, xrd := range xrds.Items {
		table = append(
			table,
			tableRow{
				API:       xrd.Metadata.Name,
				Name:      xrd.Spec.Names.Kind,
				ClaimName: xrd.Spec.ClaimNames.Kind,
			},
		)
	}
	var buf bytes.Buffer
	tableprinter.Print(&buf, table)
	style := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder())
	fmt.Println(style.Render(buf.String()))
	return xrds
}
