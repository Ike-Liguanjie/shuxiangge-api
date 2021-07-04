package model

// Book 书籍model
type Book struct {
	Name         string `json:"name"`          //小说名称
	Author       string `json:"author"`        //小说作者
	Introduction string `json:"introduction"`  //小说简介
	CategoryId   string `json:"category_id"`   //所属分类
	ChapterId    int    `json:"chapter_id"`    //最新章节
	FavoritesNum int    `json:"favorites_num"` //收藏数
}
