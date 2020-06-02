package main

import (
	"encoding/json"
	"testing"
)

func TestSlateRender(t *testing.T) {

	t.Run("Slatejs JSON convertion into HTML", func(t *testing.T) {

		var doc SlateDocument
		bSlate:=[]byte(`[{"children":[{"text":"Header 1"}],"type":"heading-one"},{"children":[{"text":"Header 2"}],"type":"heading-two"},{"children":[{"text":"This is "},{"text":"bold","bold":true},{"text":" "},{"text":"italic","italic":true},{"text":" "},{"text":"understroke","underlined":true},{"text":" "},{"text":"text","underlined":true},{"text":" "}]},{"type":"numbered-list","children":[{"children":[{"text":"item 1"}],"type":"list-item"},{"children":[{"text":"item 2"}],"type":"list-item"},{"children":[{"text":"item 3"}],"type":"list-item"}]},{"children":[{"text":"Link "},{"type":"link","url":"https://nba.com","children":[{"text":"here"}]},{"text":""}]},{"children":[{"text":""},{"type":"link","url":"https://nba.com","children":[{"text":""}]},{"text":"Image"}]},{"type":"image","url":"https://join-us-today.com/public/img/content/cb-1591112549049-703.png","children":[{"text":""}]},{"children":[{"text":""}]},{"children":[{"text":""}]}]`)

		want:=`<p><h1><span>Header 1</span></h1></p><p><h2><span>Header 2</span></h2></p><p><span>This is </span><strong>bold</strong><span> </span><em>italic</em><span> </span><u>understroke</u><span> </span><u>text</u><span> </span></p><p><ol><li><span>item 1</span></li><li><span>item 2</span></li><li><span>item 3</span></li></ol></p><p><span>Link </span><a href='https://nba.com' target='_blank'><span>here</span></a></p><p><a href='https://nba.com' target='_blank'></a><span>Image</span></p><p><image src='https://join-us-today.com/public/img/content/cb-1591112549049-703.png'</image></p><p></p><p></p>`

		if err := json.Unmarshal(bSlate, &doc); err != nil {
			t.Error(err)
		}

		got:=SlateHTML(doc)

		if got != want {
			t.Errorf("got \n%s\n want \n%s", got, want)
		}
	})

}


