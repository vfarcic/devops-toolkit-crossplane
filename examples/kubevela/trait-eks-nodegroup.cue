eksnodegroup: {
	annotations: {}
	attributes: {
		appliesToWorkloads: ["kubernetes"]
		conflictsWith:      []
		podDisruptive:      false
		workloadRefPath:    ""
	}
	description: "EKS nodegroup"
	labels:      {}
	type:        "trait"
}

template: {
	outputs: eksnodegroup: {
		apiVersion: "eks.aws.crossplane.io/v1alpha1"
		kind:       "NodeGroup"
		metadata: {
			name: context.name + "-" + parameter.name
		}
		spec: forProvider: {
			region:      "us-east-1"
			clusterName: context.name
			nodeRoleRef: name: context.name + "-nodegroup"
			subnetRefs:  [
				{name: context.name + "-1a"},
				{name: context.name + "-1b"},
				{name: context.name + "-1c"}
			]
			scalingConfig: {
				minSize:     context.output.spec.parameters.minNodeCount
				maxSize:     10
				desiredSize: context.output.spec.parameters.minNodeCount
			}
			instanceTypes: [parameter.instanceType]
		}
	}
	parameter: {
		name:         string
		instanceType: *"t3.medium" | string
	}
}

