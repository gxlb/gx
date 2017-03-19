package factory

type Creater interface { //Factory
	Create(arg ...interface{}) interface{}
}

var (
	g_factory = NewFactoryS()
)

func AddCreater(typename string, creater Creater) error {
	return g_factory.AddCreater(typename, creater)
}
func RemoveCreater(typename string) error {
	return g_factory.RemoveCreater(typename)
}

func Create(typename string, arg ...interface{}) interface{} {
	return g_factory.Create(typename, arg...)
}
