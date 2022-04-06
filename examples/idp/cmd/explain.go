package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain a parameter of an API",
	Long: `Explain a parameter of an API.
	
  # Show spec of the compositecluster resource
  kubectl idp compositecluster --parameter parameters

  # Show spec.parameters of the compositecluster resource
  kubectl idp compositecluster --parameter parameters`,
	Run: func(cmd *cobra.Command, args []string) {
		api := args[0]
		parameter, _ := cmd.Flags().GetString("parameter")
		explain(api, parameter)
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
	explainCmd.Flags().StringP("parameter", "p", "", "Specific spec parameter")
}

func explain(api, parameter string) {
	resource := api + ".spec"
	if len(parameter) > 0 {
		resource += "." + parameter
	}
	output, err := exec.Command("kubectl", "explain", resource).Output()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	outputString := string(output)
	words := []string{
		"claimRef",
		"compositionRevisionRef",
		"Alpha",
		"compositionUpdatePolicy",
		"publishConnectionDetailsTo",
		"resourceRefs",
	}
	for _, word := range words {
		re := regexp.MustCompile("(?m)[\r\n]+^.*" + word + ".*$")
		outputString = re.ReplaceAllString(outputString, "")
	}
	if parameter == "compositionRef" {
		expectedApiVersion, expectedKind := getApiVersionKind(api)
		outputString += getCompositionsOutput(expectedApiVersion, expectedKind, false)
	}
	if parameter == "compositionSelector.matchLabels" {
		expectedApiVersion, expectedKind := getApiVersionKind(api)
		outputString += getCompositionsOutput(expectedApiVersion, expectedKind, true)
	}
	fmt.Println(outputString)
}

func getCompositionsOutput(expectedApiVersion, expectedKind string, hasLabels bool) string {
	output := "\nAVAILABLE COMPOSITIONS:"
	if hasLabels {
		output = "\nAVAILABLE LABELS:"
	}
	compositions := getCompositions(expectedApiVersion, expectedKind)
	println(len(compositions.Items))
	for _, composition := range compositions.Items {
		output += "\n\n   " + composition.Metadata.Name
		if hasLabels {
			output += ":"
		}
		for key, value := range composition.Metadata.Labels {
			output += "\n      " + key + ": " + value
		}
	}
	return output
}

func getApiVersionKind(api string) (apiVersion, kind string) {
	apiVersion = api[strings.Index(api, ".")+1:]
	kindBytes, err := exec.Command("kubectl", "get", "crd", api, "--output", "jsonpath={.spec.names.kind}").Output()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	kind = string(kindBytes)
	return apiVersion, kind
}
