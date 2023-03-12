package hub

type Rules interface {
	ToRule()
	Describe() string
}

type Rule struct {
	Host            string `json:"host"`            // 目标网站 URL
	HostTitle       string `json:"hostTitle"`       // 目标网站 Title
	HostDescription string `json:"hostDescription"` // 目标网站 Description

	TargetURl string `json:"targetURl"` // 提取网站

	ListContainers string   `json:"listContainers"` // 内容列表
	Title          string   `json:"title"`          // 内容
	Author         string   `json:"author"`         // 内容
	URL            string   `json:"url"`            // 内容
	Date           string   `json:"date"`           // 日前
	Description    string   `json:"description"`    // 简介
	Contents       string   `json:"contents"`       // 内容
	Images         []string `json:"images"`         // 图片
}

func (r *Rule) ToHost(host string) *Rule {
	r.Host = host
	return r
}

func (r *Rule) ToTarget(target string) *Rule {
	r.TargetURl = target
	return r
}

func (r *Rule) ToContainers(containers string) *Rule {
	r.ListContainers = containers
	return r
}

func (r *Rule) ToTitle(title string) *Rule {
	r.Title = title
	return r
}

func (r *Rule) ToAuthor(author string) *Rule {
	r.Author = author
	return r
}

func (r *Rule) ToURL(url string) *Rule {
	r.URL = url
	return r
}

func (r *Rule) ToDate(date string) *Rule {
	r.Date = date
	return r
}

func (r *Rule) ToDescription(description string) *Rule {
	r.Description = description
	return r
}
func (r *Rule) ToContents(contents string) *Rule {
	r.Contents = contents
	return r
}
