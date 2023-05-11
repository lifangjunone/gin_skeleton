package ioc

var (
	dIoc      = make(map[string]interface{}, 100)
	DaoIocObj *daoIoc
)

type daoIoc struct{}

func newDaoIoc() *daoIoc {
	return &daoIoc{}
}

func (c *daoIoc) Registry(svcName string, svc interface{}) {
	_, isOk := dIoc[svcName]
	if !isOk {
		dIoc[svcName] = svc
	}
}

func (c *daoIoc) GetSvc(svcName string) interface{} {
	svc, isOk := dIoc[svcName]
	if isOk {
		return svc
	}
	return nil
}

func init() {
	DaoIocObj = newDaoIoc()
}
