package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

const InsertHere = "INSERT_HERE"

type (
	Properties interface{}
)

var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Output sample YAML manifest",
	Long: `Output sample YAML manifest
	
# TODO:`,
	Run: func(cmd *cobra.Command, args []string) {
		api := args[0]
		composition, _ := cmd.Flags().GetString("composition")
		outputManifest(api, composition)
	},
}

func init() {
	rootCmd.AddCommand(outputCmd)
	outputCmd.Flags().StringP("composition", "c", "", "Specific composition that should be used to generate XR")
}

func outputManifest(api, composition string) {
	crd := getCRD(api)
	xr := getXR(crd, composition)
	println(getXRYaml(xr))
	println("\n\n")
}

func getCRD(api string) CRD {
	output, err := exec.Command("kubectl", "get", "crd", api, "-o", "yaml").Output()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	crd := CRD{}
	yaml.Unmarshal([]byte(string(output)), &crd)
	return crd
}

func getXR(crd CRD, compositionName string) XR {
	xr := XR{}
	xr.ApiVersion = crd.Spec.Group + "/" + crd.Spec.Versions[0].Name
	xr.Kind = crd.Spec.Names.Kind
	xr.Metadata.Name = InsertHere
	spec := make(map[interface{}]interface{})
	properties := crd.Spec.Versions[0].Schema.OpenAPIV3Schema.Properties.Spec.Properties.(map[interface{}]interface{})
	for key, value := range properties {
		switch key {
		case "claimRef", "compositionUpdatePolicy", "resourceRefs", "compositionRef", "compositionRevisionRef", "publishConnectionDetailsTo":
			// Ignore
		case "compositionSelector":
			labels := make(map[string]string)
			if len(compositionName) == 0 {
				labels["SOME_KEY"] = InsertHere
				labels["SOME_OTHER_KEY"] = InsertHere
			} else {
				compositions := getCompositions(xr.ApiVersion, xr.Kind)
				for _, composition := range compositions.Items {
					if composition.Metadata.Name == compositionName {
						for key, value := range composition.Metadata.Labels {
							labels[key] = value
						}
						break
					}
				}
			}
			matchLabels := make(map[string]interface{})
			matchLabels["matchLabels"] = labels
			spec["CompositionSelector"] = matchLabels
		case "writeConnectionSecretToRef":
			secrets := make(map[string]string)
			secrets["name"] = InsertHere + " # The name of the secret with authentication (string)"
			secrets["namespace"] = InsertHere + " # The namespace for the secret (string)"
			spec["writeConnectionSecretToRef"] = secrets
		default:
			switch v := value.(type) {
			case map[interface{}]interface{}:
				subKey := fmt.Sprintf("%v", key)
				spec[subKey] = processMapInterface(v)
			}
		}
		xr.Spec = spec
	}
	return xr
}

func processMapInterface(properties interface{}) interface{} {
	var subProperties interface{}
	hasProperties := false
	description := ""
	xDefault := InsertHere
	xType := ""
	for key, value := range properties.(map[interface{}]interface{}) {
		if key == "properties" {
			hasProperties = true
			subProperties = value
		}
		keyString := fmt.Sprintf("%v", key)
		valueString := fmt.Sprintf("%v", value)
		switch keyString {
		case "description":
			description = valueString
		case "type":
			xType = valueString
		case "default":
			xDefault = valueString
		}
	}
	if hasProperties {
		newSubProperties := make(map[interface{}]interface{})
		for key, value := range subProperties.(map[interface{}]interface{}) {
			newSubProperties[key] = processMapInterface(value)
		}
		return newSubProperties
	}
	return fmt.Sprintf(
		"%s # %s (%s)",
		xDefault,
		description,
		xType,
	)
}

func getXRYaml(xr XR) string {
	yamlData, err := yaml.Marshal(xr)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	return strings.ReplaceAll(string(yamlData), "'", "")
}

// TODO: Add XRC
