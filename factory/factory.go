//对象工厂 每个类型都包含一个Creator 对应的对象由相应的Creater生产 出来的对象都用interface{}表示
//有了这个对象工厂 就可以通过配置预定义好的字符串 生成对应的对象了
//像协议buff转协议对象这种活就都由这个factory承包了

//实现工厂模式
package factory

type Creater interface{ //Factory
	Create(arg... interface{}) interface{}
}

var(
	g_factory = NewFactoryS()
)


func AddCreater(typename string, creater Creater) error{
	return g_factory.AddCreater(typename, creater)
}
func RemoveCreater(typename string) error{
	return g_factory.RemoveCreater(typename)
}

func Create(typename string, arg... interface{}) interface{}{
	return g_factory.Create(typename, arg...)
}
