module github.com/ttsubo2000/goext_example

go 1.18

require (
	github.com/cloudwan/gohan v1.0.1-0.20220809095246-14a769bf7a4b
	github.com/onsi/ginkgo v1.6.0
	github.com/onsi/gomega v1.4.2
)

require (
	github.com/golang/mock v1.3.0 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859 // indirect
	golang.org/x/sys v0.0.0-20190801041406-cbf593c0f2f3 // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace github.com/onsi/ginkgo => github.com/cloudwan/ginkgo v1.6.1-0.20190213151947-95174e8d10cd

replace github.com/cloudwan/gohan => ../gohan
