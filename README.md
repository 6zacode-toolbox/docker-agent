# docker-agent
An image with drivers to engage with docker host


# How this project was started

```bash 
go mod init github.com/6zacode-toolbox/docker-agent
cobra-cli init
go mod tidy

#Updates Library
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
  resources: ["dockerhosts", "dockerhosts/status", "dockercomposerunners", "dockercomposerunners/status"]
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
  containers:
  - name: docker-agent
    image: 6zar/docker-agent:latest
    imagePullPolicy: Always
    env:
    - name: DOCKER_CERT_PATH
      value: "/certs"
    - name: DOCKER_HOST
      value: "tcp://192.168.2.162:2376"
    - name: DOCKER_TLS_VERIFY
      value: "1"
    command: ['/home/app/docker-agent', 'agent', '--crd-api-version', 'tool.6zacode-toolbox.github.io/v1', '--crd-namespace', 'default', '--crd-instance', 'dockerhost-sample', '--crd-resource', 'dockerhosts']
    volumeMounts:
    - mountPath: "/certs"
      name: docker-certs
      readOnly: true
  serviceAccountName: docker-agent-sa
  volumes:
  - name: docker-certs
    secret:
      secretName: docker-secret



```