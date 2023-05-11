package ioc

var (
	IocSvc Ioc
)

type Ioc interface {
	Registry(svcName string, svc interface{})
	GetSvc(svcName string) interface{}
}
