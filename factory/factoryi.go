package factory

func NewFactoryI() *FactoryI{
	return &FactoryI{cmap: make(map[int] Creater)}
}
//int作为类型标识的工厂
type FactoryI struct{ 
	cmap map[int] Creater
}
func (me *FactoryI)AddCreater(nType int, creater Creater) error{
	return nil
}

func (me *FactoryI)RemoveCreater(nType int) error{
	return nil
}

func (me *FactoryI)Create(nType int, arg... interface{}) interface{}{
	if p,ok:=me.cmap[nType];ok{
		return p.Create(arg...)
	}
	return nil
}
