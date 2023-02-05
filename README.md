# docker-agent
An image with drivers to engage with docker host


# How this project was started

```bash 
go mod init github.com/6zacode-toolbox/docker-agent
cobra-cli init
go mod tidy
go get github.com/6zacode-toolbox/docker-operator/operator/api/v1
cobra-cli add test
go run . test 
```


# Sample Deploy

```yaml

apiVersion: v1
kind: ServiceAccount
metadata:
  name: docker-agent-sa
  namespace: default
--- 
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: docker-agent-cr
rules:
  - apiGroups: ["tool.6zacode-toolbox.github.io"]
    resources: ["dockerhosts","dockerhosts/status"]
    verbs: ["*"]
--- 
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: docker-agent-crb
subjects:
- kind: ServiceAccount
  name: docker-agent-sa
  namespace: default
roleRef:
  kind: ClusterRole
  name: docker-agent-cr
  apiGroup: rbac.authorization.k8s.io

---
apiVersion: v1
kind: Pod
metadata:
  name: docker-agent
spec:
  serviceAccountName: docker-agent-sa
  containers:
  - name: docker-agent
    image: 6zar/docker-agent:latest
    imagePullPolicy: Always
    command: ['/home/app/docker-agent', 'agent', '--crd-api-version', 'tool.6zacode-toolbox.github.io/v1', '--crd-namespace', 'default', '--crd-instance', 'dockerhost-sample', '--crd-resource', 'dockerhosts']

```