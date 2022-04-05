package cmd

import "strings"

// type (
// 	Resources struct {
// 		Items []Resource
// 	}

// 	Resource struct {
// 		Metadata struct {
// 			Name string
// 		}
// 		Spec struct {
// 			ClaimNames KindPlural `yaml:"claimNames"`
// 			Names      KindPlural
// 		}
// 	}

// 	KindPlural struct {
// 		Kind   string
// 		Plural string
// 	}
// )

func getName(fullName string) string {
	return fullName[strings.LastIndex(fullName, "/")+1:]
}
