Inspired from TGI Kubernetes (Episode 40) hosted by Kris Nóva (@kris-nova). The original code can be found [here](https://github.com/vmware-tanzu/tgik/tree/master/episodes/040/live).

```sh
# initialize!
$ kubebuilder init --domain bensooraj.com

# Create the k8s API boilerplate template
$ kubebuilder create api --group birds --version v1beta1 --kind Puffin
Create Resource [y/n]
y
Create Controller [y/n]
y
Writing scaffold for you to edit...
api/v1beta1/puffin_types.go
controllers/puffin_controller.go
Running make:
$ make
/Users/Bensooraj/go/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
go vet ./...
go build -o bin/manager main.go
```

Install the CRDs into the cluster:

```sh
$ make install
/Users/Bensooraj/go/bin/controller-gen "crd:trivialVersions=true" rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
kustomize build config/crd | kubectl apply -f -
customresourcedefinition.apiextensions.k8s.io/puffins.birds.bensooraj.com created
```

Run the controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

```sh
$ make run
# /Users/Bensooraj/go/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
/Users/Bensooraj/go/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
go vet ./...
/Users/Bensooraj/go/bin/controller-gen "crd:trivialVersions=true" rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
go run ./main.go

2020-12-15T13:37:27.472+0530	INFO	controller-runtime.metrics	metrics server is starting to listen	{"addr": ":8080"}
2020-12-15T13:37:27.473+0530	INFO	setup	starting manager
2020-12-15T13:37:27.473+0530	INFO	controller-runtime.manager	starting metrics server	{"path": "/metrics"}
2020-12-15T13:37:27.474+0530	INFO	controller-runtime.controller	Starting EventSource	{"controller": "puffin", "source": "kind source: /, Kind="}
2020-12-15T13:37:27.574+0530	INFO	controller-runtime.controller	Starting Controller	{"controller": "puffin"}
2020-12-15T13:37:27.574+0530	INFO	controller-runtime.controller	Starting workers	{"controller": "puffin", "worker count": 1}

2020-12-15T13:37:52.731+0530	INFO	controllers.Puffin	Setting logged to true	{"puffin": "default/tufted-puffin"}
2020-12-15T13:37:52.731+0530	INFO	controllers.Puffin	[PENDING] tufted-puffin will be assigned a color shortly!
	{"puffin": "default/tufted-puffin"}
2020-12-15T13:37:52.768+0530	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "puffin", "request": "default/tufted-puffin"}
2020-12-15T13:37:52.773+0530	INFO	controllers.Puffin	[COLORED] tufted-puffin is assigned the color red!
	{"puffin": "default/tufted-puffin"}
2020-12-15T13:37:52.773+0530	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "puffin", "request": "default/tufted-puffin"}
2020-12-15T13:37:52.821+0530	INFO	controllers.Puffin	[PENDING] horned-puffin will be assigned a color shortly!
	{"puffin": "default/horned-puffin"}
2020-12-15T13:37:52.868+0530	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "puffin", "request": "default/horned-puffin"}
2020-12-15T13:37:52.868+0530	INFO	controllers.Puffin	[COLORED] horned-puffin is assigned the color orange!
	{"puffin": "default/horned-puffin"}
2020-12-15T13:37:52.868+0530	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "puffin", "request": "default/horned-puffin"}
```

Install Instances of Custom Resources

```sh
$ kubectl apply -f config/samples/
```

Cleanup

```sh
$ make uninstall
```

https://github.com/kubernetes-sigs/kubebuilder/issues/806

```go
    // Watch ReplicaSets and enqueue ReplicaSet object key
	if err := c.Watch(&source.Kind{Type: &appsv1.ReplicaSet{}}, &handler.EnqueueRequestForObject{}); err != nil {
		entryLog.Error(err, "unable to watch ReplicaSets")
		os.Exit(1)
	}
```

Resource:

1. https://insujang.github.io/2020-02-13/programming-kubernetes-crd/
