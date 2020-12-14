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




Resource:
1. https://insujang.github.io/2020-02-13/programming-kubernetes-crd/