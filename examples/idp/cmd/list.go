package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/lipgloss"
	"github.com/lensesio/tableprinter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "The list of all the available components",
	Long: `The list of all the available components.
	
  # Show spec of the compositecluster resource
  kubectl idp list`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

type (
	tableRow struct {
		API       string `header:"API"`
		Name      string `header:"Name"`
		ClaimName string `header:"Claim"`
	}

	XRDs struct {
		Items []XRD
	}

	XRD struct {
		Metadata struct {
			Name string
		}
		Spec struct {
			ClaimNames KindPlural `yaml:"claimNames"`
			Names      KindPlural
		}
	}

	KindPlural struct {
		Kind   string
		Plural string
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}

func list() XRDs {
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

func getXrds() XRDs {
	output, err := exec.Command("kubectl", "get", "compositeresourcedefinitions.apiextensions.crossplane.io", "-o", "yaml").Output()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	xrds := XRDs{}
	yaml.Unmarshal([]byte(string(output)), &xrds)
	return xrds
}
