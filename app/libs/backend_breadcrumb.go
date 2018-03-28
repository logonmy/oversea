package libs

import (
	"bytes"
	"fmt"
)

type Breadcrumb struct {
	PageTitle     string
	PageDesc string
	MenuList []MenuList
}

type MenuList struct {
	MenuName string
	MenuUrl string
	IsActive bool
}


func NewBreadcrumb(pageTitle, pageDesc string) *Breadcrumb {
	p := new(Breadcrumb)
	p.PageTitle = pageTitle
	p.PageDesc = pageDesc

	return p
}

func (this *Breadcrumb) SetMenu(name, url string) *Breadcrumb {

	if name != ``  {

         var menuList MenuList

         menuList.MenuName = name
         menuList.MenuUrl = url
		 menuList.IsActive = false

         if url == `` {
         	menuList.IsActive = true
		 }

		 this.MenuList = append(this.MenuList, menuList)
	}

	return this
}

func (this *Breadcrumb) ToString() string {
	var buf bytes.Buffer
	buf.WriteString(`<div class="block-header">`)
	buf.WriteString(`<div class="row">`)
	buf.WriteString(`<div class="col-lg-7 col-md-6 col-sm-12">`)

	buf.WriteString(fmt.Sprintf("<h2>%s", this.PageTitle))

	if this.PageDesc != `` {
		buf.WriteString(fmt.Sprintf("<small>%s</small>", this.PageDesc))
	}

	buf.WriteString("</h2>")

	buf.WriteString(`</div>`)


	if len(this.MenuList) > 0 {
		buf.WriteString(`<div class="col-lg-5 col-md-6 col-sm-12">`)
		buf.WriteString(`<ul class="breadcrumb float-md-right">`)


		for i,v := range this.MenuList{

			if i == 0 {
				buf.WriteString(fmt.Sprintf("<li class=\"breadcrumb-item\"><a href=\"%s\"><i class=\"zmdi zmdi-home" +
					"\"></i> %s</a></li>", v.MenuUrl, v.MenuName))
			} else {
				if v.IsActive {
					buf.WriteString(fmt.Sprintf("<li class=\"breadcrumb-item\"><a href=\"%s\"><i " +
						"\"></i> %s</a></li>", v.MenuUrl, v.MenuName))
				} else {
					buf.WriteString(fmt.Sprintf("<li class=\"breadcrumb-item active\"><a href=\"%s\"><i" +
						"\"></i> %s</a></li>", v.MenuUrl, v.MenuName))
				}
			}
		}

		buf.WriteString(`</ul>`)
		buf.WriteString(`</div>`)
	}


	buf.WriteString(`</div>`)
	buf.WriteString(`</div>`)

	return buf.String()
}