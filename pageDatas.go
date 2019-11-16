package main

type IndexPageData struct {
	PageTitle string
	Context   string
}
type TagPageData struct {
	PageTitle string
	Tags      []Tag
}
type Tag struct {
	Title string
	Freq  int
	Size  int
}
