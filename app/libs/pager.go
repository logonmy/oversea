package libs

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Pager struct {
	Page     int
	Totalnum int
	Pagesize int
	urlpath  string
	urlquery string
	nopath   bool
}

func NewPager(page, totalnum, pagesize int, url string, nopath ...bool) *Pager {
	p := new(Pager)
	p.Page = page
	p.Totalnum = totalnum
	p.Pagesize = pagesize

	arr := strings.Split(url, "?")
	p.urlpath = arr[0]
	if len(arr) > 1 {
		p.urlquery = "?" + arr[1]
	} else {
		p.urlquery = ""
	}

	if len(nopath) > 0 {
		p.nopath = nopath[0]
	} else {
		p.nopath = false
	}

	return p
}

func (this *Pager) url(page int) string {
	if this.nopath { //不使用目录形式
		if this.urlquery != "" {
			return fmt.Sprintf("%s%s&page=%d", this.urlpath, this.urlquery, page)
		} else {
			return fmt.Sprintf("%s?page=%d", this.urlpath, page)
		}
	} else {
		return fmt.Sprintf("%s/page/%d%s", this.urlpath, page, this.urlquery)
	}
}

func (this *Pager) ToString() string {
	if this.Totalnum <= this.Pagesize {
		return ""
	}

	var buf bytes.Buffer
	var from, to, linknum, offset, totalpage int

	offset = 5
	linknum = 10

	totalpage = int(math.Ceil(float64(this.Totalnum) / float64(this.Pagesize)))

	if totalpage < linknum {
		from = 1
		to = totalpage
	} else {
		from = this.Page - offset
		to = from + linknum
		if from < 1 {
			from = 1
			to = from + linknum - 1
		} else if to > totalpage {
			to = totalpage
			from = totalpage - linknum + 1
		}
	}

	buf.WriteString("<ul class=\"oversea-pagination oversea-table-pagination\" style=\"margin-right:10px;\">")
	if this.Page > 1 {
		buf.WriteString(fmt.Sprintf("<li class=\"oversea-pagination-prev\"><a class=\"oversea-pagination-item-link" +
			"\" href=\"%s\"></a></li>",
			this.url(this.Page-1)))
	} else {
		buf.WriteString("<li class=\"oversea-pagination-disabled oversea-pagination-prev\"><a href=\"#\" class" +
			"=\"oversea-pagination-item-link\"></a></li>")
	}

	if this.Page > linknum {
		buf.WriteString(fmt.Sprintf("<li class=\"oversea-pagination-item\"><a class=\"oversea-pagination-item-link\"  href=\"%s\">1...</a></li>",
			this.url(1)))
	}

	for i := from; i <= to; i++ {
		if i == this.Page {
			buf.WriteString(fmt.Sprintf("<li class=\"oversea-pagination-item oversea-pagination-item-1 oversea-pagination-item-active\"><a href=\"#\">%d</a></li>", i))
		} else {
			buf.WriteString(fmt.Sprintf("<li  class=\"oversea-pagination-item\" ><a href=\"%s\">%d</a></li>", this.url(i), i))
		}
	}

	if totalpage > to {
		buf.WriteString(fmt.Sprintf("<li class=\"oversea-pagination-item\"><a class=\"oversea-pagination-item-link\"  href=\"%s\">...%d</a></li>", this.url(totalpage), totalpage))
	}

	if this.Page < totalpage {
		buf.WriteString(fmt.Sprintf("<li class=\"oversea-pagination-next\"><a class=\"oversea-pagination-item-link\"  href=\"%s\"></a></li>",
			this.url(this.Page+1)))
	} else {
		buf.WriteString(fmt.Sprintf("<li class=\"oversea-pagination-disabled oversea-pagination-next\"><a href" +
			"=\"#\" class=\"oversea-pagination-item-link\" ></a></li>"))
	}

	buf.WriteString(fmt.Sprintf("     <li class=\"oversea-pagination-options\">"+
	"<div class=\"oversea-pagination-options-quick-jumper\">"+
		"跳至<input type=\"text\" href=\"%s\" value=\"%d\">页"+
	"</div></li>",this.url(this.Page), this.Page))

	buf.WriteString("</ul>")

	return buf.String()
}
