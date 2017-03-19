package factory

func NewFactoryS() *FactoryS{
	return &FactoryS{cmap: make(map[string] Creater)}
}
//string作为类型标识的工厂
type FactoryS struct{ 
	cmap map[string] Creater
}
func (me *FactoryS)AddCreater(typename string, creater Creater) error{
	return nil
}

func (me *FactoryS)RemoveCreater(typename string) error{
	return nil
}

func (me *FactoryS)Create(typename string, arg... interface{}) interface{}{
	if p,ok:=me.cmap[typename];ok{
		return p.Create(arg...)
	}
	return nil
}