package model

// Chapter 章节model
type Chapter struct {
	Id          string `json:"id"`           // 章节ID
	ChapterName string `json:"chapter_name"` //章节名称
	ContentUrl  string `json:"content_url"`  //章节内容路径
	Index       string `json:"index"`        //章节索引
}
