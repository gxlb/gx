package logi

//manage a large block
type Page struct {
	//[]block
}

type Block struct {
}

//manage fix-size blocks
type FixPool struct {
	//page-set
	//page-map
	//free-list
}
type LogiPool struct {
	//[]FixPool
	//max_free_list_len
	//max_bytes
}
