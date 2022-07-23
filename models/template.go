package models

type TemplateConfig struct {
	TemplateLayoutPath  string
	TemplateModalsPath  string
	TemplateIncludePath string
}

type Header struct {
	Page  string
	Title string
	User
}

type PageNotFound struct {
	Header
	Page string
}
