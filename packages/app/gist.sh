---

apiVersion: pkg.crossplane.io/v1
kind: Configuration
metadata:
  name: crossplane-sql
spec:
  package: vfarcic/crossplane-sql:v0.2.19

---

apiVersion: pkg.crossplane.io/v1
kind: Configuration
metadata:
  name: crossplane-app
spec:
  package: vfarcic/crossplane-app:v0.2.6

---

apiVersion: pkg.crossplane.io/v1
kind: Configuration
metadata:
  name: crossplane-k8s
spec:
  package: vfarcic/crossplane-k8s:v0.3.3

---

apiVersion: v1
kind: ServiceAccount
metadata:
  name: crossplane-provider-kubernetes
  namespace: crossplane-system

---

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: crossplane-provider-kubernetes
subjects:
- kind: ServiceAccount
  name: crossplane-provider-kubernetes
  namespace: crossplane-system
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io

---

apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: crossplane-provider-kubernetes
spec:
  serviceAccountName: crossplane-provider-kubernetes

---

apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: crossplane-provider-kubernetes
spec:
  package: crossplane/provider-kubernetes:v0.3.0
  controllerConfigRef:
    name: crossplane-provider-kubernetes

---

apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: crossplane-provider-aws
spec:
  package: crossplane/provider-aws:v0.24.1