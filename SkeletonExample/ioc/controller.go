package ioc

var (
	ctrIoc    = make(map[string]interface{}, 100)
	CtrIocObj *controllerIoc
)

type controllerIoc struct{}

func newControllerIoc() *controllerIoc {
	return &controllerIoc{}
}

func (c *controllerIoc) Registry(svcName string, svc interface{}) {
	_, isOk := ctrIoc[svcName]
	if !isOk {
		ctrIoc[svcName] = svc
	}
}

func (c *controllerIoc) GetSvc(svcName string) interface{} {
	svc, isOk := ctrIoc[svcName]
	if isOk {
		return svc
	}
	return nil
}

func init() {
	CtrIocObj = newControllerIoc()
}
