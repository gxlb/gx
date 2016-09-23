package factory

//create load store
type Creator interface {
	Create() interface{}
}

var (
	gMapCreator map[string]Creator
	gMapId      map[int]string
)

func RegCreator(name string, creator Creator) {

}

//create by name
func Create(name string) interface{} {
	if v, ok := gMapCreator[name]; ok {
		return v.Create()
	}
	return nil
}

//create by id
func CreateById(id int) interface{} {
	if v, ok := gMapId[id]; ok {
		return Create(v)
	}
	return nil
}
