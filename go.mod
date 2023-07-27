module github.com/mochen302/gengox

go 1.13

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/google/go-cmp v0.5.9
	github.com/google/gofuzz v1.1.0
	github.com/spf13/pflag v1.0.5
	golang.org/x/tools v0.7.0
	google.golang.org/protobuf v1.31.0
	k8s.io/api v0.27.4 // indirect
	k8s.io/apimachinery v0.27.4
	k8s.io/klog/v2 v2.90.1
	sigs.k8s.io/yaml v1.3.0
)

replace (
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
)
