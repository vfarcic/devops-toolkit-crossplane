gitops: {
	annotations: {}
	attributes: {
		appliesToWorkloads: ["kubernetes"]
		conflictsWith:      []
		podDisruptive:      false
		workloadRefPath:    ""
	}
	description: "GitOps tools (Argo CD, Flux, Rancher Fleet, etc.) installed and configured"
	labels:      {}
	type:        "trait"
}

template: {
	outputs: gitops: {
		apiVersion: "devopstoolkitseries.com/v1alpha1"
		kind:       "GitOpsClaim"
		metadata: {
			name: context.name + "-gitops"
		}
		spec: {
			compositionSelector: matchLabels: provider: parameter.provider
			id: context.name + "-gitops"
			parameters: {
				gitOpsRepo: parameter.repo
				environment: parameter.environment
				kubeConfig: {
					secretName:      context.name + "-cluster"
					secretNamespace: "crossplane-system"
				}
			}
		}
	}
	parameter: {
		provider:    *"argo" | string
		repo:        string
		environment: *"production" | string
	}
}

