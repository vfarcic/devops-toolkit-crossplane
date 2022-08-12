kubernetes: {
	annotations: {}
	attributes: workload: definition: {
		apiVersion: "devopstoolkitseries.com/v1alpha1"
		kind:       "ClusterClaim"
	}
	description: "Kubernetes cluster anywhere (AWS, Azure, Google Cloud, Civo, DigitalOcean, etc.)"
	labels: {}
	type: "component"
}

template: {
	output: {
		apiVersion: "devopstoolkitseries.com/v1alpha1"
		kind:       "ClusterClaim"
		metadata: name: context.name
		spec: {
			compositionRef: {
				name:  parameter.cluster
			}
			id: context.name
			parameters: {
				minNodeCount: parameter.minNodeCount
				nodeSize:     parameter.nodeSize
				version: 	  parameter.version
			}
			writeConnectionSecretToRef: name: context.name
		}
	}
	outputs: {}
	parameter: {
		cluster:      *"cluster-aws" | string
		minNodeCount: *3 | int
		nodeSize:     *"medium" | string
		version:      *"1.22" | string
	}
}

