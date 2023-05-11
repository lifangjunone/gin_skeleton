package ioc

var (
	svcIoc    = make(map[string]interface{}, 100)
	SvcIocObj *serviceIoc
)

type serviceIoc struct{}

func newServiceIoc() *serviceIoc {
	return &serviceIoc{}
}

func (s *serviceIoc) Registry(svcName string, svc interface{}) {
	_, isOk := svcIoc[svcName]
	if !isOk {
		svcIoc[svcName] = svc
	}
}

func (s *serviceIoc) GetSvc(svcName string) interface{} {
	svc, isOk := svcIoc[svcName]
	if isOk {
		return svc
	}
	return nil
}

func init() {
	SvcIocObj = newServiceIoc()
}
