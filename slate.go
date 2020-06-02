package slate

type SlateItem struct {
	Children []SlateItem `json:"children,omitempty"`
	Text string `json:"text,omitempty"`
	Type string `json:"type,omitempty"`
	Bold bool `json:"bold,omitempty"`
	Code bool `json:"code,omitempty"`
	Italic bool `json:"italic,omitempty"`
	Underlined bool `json:"underlined,omitempty"`
	URL  string `json:"url,omitempty"`
}

type SlateDocument []SlateItem

func SlateNodeHTML (node SlateItem) string {

	html:=""

	if node.Type == "heading-one" {
		html = html + "<h1>" + node.Text
	} else if node.Type == "heading-two" {
		html = html + "<h2>" + node.Text
	} else if node.Type == "paragraph" {
		html = html + "<p>" + node.Text
	} else if node.Type == "list-item" {
		html = html + "<li>" + node.Text
	} else if node.Type == "bulleted-list" {
		html = html + "<ul>" + node.Text
	} else if node.Type == "numbered-list" {
		html = html + "<ol>" + node.Text
	} else if node.Type == "image" {
		html = html + "<image src='"+ node.URL+"'"
	} else if node.Type == "link" {
		html = html + "<a href='"+ node.URL+"' target='_blank'>"+ node.Text
	} else if node.Type == "" && node.Text!=""  {
		if node.Bold {
			html = html + "<strong>" + node.Text
		} else if node.Italic {
			html = html + "<em>" + node.Text
		} else if node.Code {
			html = html + "<code>" + node.Text
		} else if node.Underlined {
			html = html + "<u>" + node.Text
		} else {
			html = html + "<span>" + node.Text
		}

	}

	for _, n := range node.Children {
		html = html + SlateNodeHTML(n)
	}

	if node.Type == "heading-one" {
		html = html + "</h1>"
	} else if node.Type == "heading-two" {
		html = html + "</h2>"
	} else if node.Type == "paragraph" {
		html = html + "</p>"
	} else if node.Type == "list-item" {
		html = html + "</li>"
	} else if node.Type == "bulleted-list" {
		html = html + "</ul>"
	} else if node.Type == "numbered-list" {
		html = html + "</ol>"
	} else if node.Type == "image" {
		html = html + "</image>"
	} else if node.Type == "link" {
		html = html + "</a>"
	} else if node.Type == "" && node.Text!="" {
		if node.Bold {
			html = html + "</strong>"
		} else if node.Italic {
			html = html + "</em>"
		} else if node.Code {
			html = html + "</code>"
		} else if node.Underlined {
			html = html + "</u>"
		} else {
			html = html + "</span>"
		}
	}

	return html
}

func SlateHTML (items []SlateItem) string {

	html:=""

	if len(items)==0 {
		return "<div>Empty Document</div>"
	}

	for _, node := range items {
		html = html + "<p>" + SlateNodeHTML(node)+ "</p>"
	}

	return html
}

