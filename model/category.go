package model


// BooksCategory 分类model
type BooksCategory struct {
	Id    string `json:"id"`    //分类ID
	Name  string `json:"name"`  //分类名称
	Order string `json:"order"` // 分类展示排序
}
