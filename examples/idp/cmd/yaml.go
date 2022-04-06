package cmd

import (
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

type (
	Compositions struct {
		Items []Composition
	}

	Composition struct {
		Metadata struct {
			Name   string
			Labels map[string]string
		}
		Spec struct {
			CompositeTypeRef CompositeTypeRef `yaml:"compositeTypeRef"`
		}
	}

	XRDs struct {
		Items []XRD
	}

	XRD struct {
		Metadata struct {
			Name string
		}
		Spec struct {
			Group      string
			ClaimNames KindPlural `yaml:"claimNames"`
			Names      KindPlural
			Versions   []Version `yaml:"versions"`
		}
	}

	XR struct {
		ApiVersion string `yaml:"apiVersion"`
		Kind       string
		Metadata   struct {
			Name string
		}
		// Spec struct {
		// 	CompositionSelector        CompositionSelector        `yaml:"compositionSelector"`
		// 	WriteConnectionSecretToRef WriteConnectionSecretToRef `yaml:"writeConnectionSecretToRef"`
		// 	// CustomFields               interface{}
		// }
		Spec interface{}
	}

	Version struct {
		Name string
	}

	CompositeTypeRef struct {
		ApiVersion string `yaml:"apiVersion"`
		Kind       string
	}

	tableRow struct {
		API       string `header:"API"`
		Name      string `header:"Name"`
		ClaimName string `header:"Claim"`
	}

	KindPlural struct {
		Kind   string
		Plural string
	}

	CRD struct {
		ApiVersion string  `yaml:"apiVersion"`
		Spec       CrdSpec `yaml:"spec"`
	}

	CrdSpec struct {
		Group    string
		Versions []struct {
			Name   string
			Schema struct {
				OpenAPIV3Schema OpenAPIV3Schema `yaml:"openAPIV3Schema"`
			}
		}
		Names struct {
			Kind string
		}
	}

	OpenAPIV3Schema struct {
		Properties struct {
			Spec struct {
				Properties Properties
			}
		}
	}

	CompositionSelector struct {
		MatchLabels MatchLabels `yaml:"matchLabels"`
	}

	MatchLabels interface{}

	WriteConnectionSecretToRef struct {
		Name      string
		Namespace string
	}
)

var allCompositions = Compositions{}

func getName(fullName string) string {
	return fullName[strings.LastIndex(fullName, "/")+1:]
}

func getAllCompositions() Compositions {
	if len(allCompositions.Items) == 0 {
		allCompositions = Compositions{}
		yamlOutput, err := exec.Command("kubectl", "get", "compositions.apiextensions.crossplane.io", "-o", "yaml").Output()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		yaml.Unmarshal([]byte(string(yamlOutput)), &allCompositions)
	}
	return allCompositions
}

func getCompositions(expectedApiVersion, expectedKind string) Compositions {
	allCompositions := getAllCompositions()
	compositions := Compositions{}
	for _, composition := range allCompositions.Items {
		if strings.HasPrefix(string(composition.Spec.CompositeTypeRef.ApiVersion), expectedApiVersion) && composition.Spec.CompositeTypeRef.Kind == expectedKind {
			compositions.Items = append(compositions.Items, composition)
		}
	}
	return compositions
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
