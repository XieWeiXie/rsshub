package rsshub

type Rules interface {
	ToRule()
	Describe() string
}

type Rule struct {
	Host            string `json:"host"`
	HostTitle       string `json:"hostTitle"`
	HostDescription string `json:"hostDescription"`

	TargetURl string `json:"targetURl"`

	ListContainers string `json:"listContainers"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	URL            string `json:"url"`
	Date           string `json:"date"`
	Description    string `json:"description"`
	Contents       string `json:"contents"`
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
