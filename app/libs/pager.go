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
	buf.WriteString("<div class=\"card\">")
	buf.WriteString("<div class=\"body\">")
	buf.WriteString("<ul class=\"pagination pagination-primary m-b-0\" >")
	if this.Page > 1 {
		buf.WriteString(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link" +
			"\" href=\"%s\"><i class=\"zmdi zmdi-arrow-left\"></i></a></li>",
			this.url(this.Page-1)))
	}

	if this.Page > linknum {
		buf.WriteString(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\" href=\"%s\">1...</a></li>",
			this.url(1)))
	}

	for i := from; i <= to; i++ {
		if i == this.Page {
			buf.WriteString(fmt.Sprintf("<li class=\"page-item active\"><a class=\"page-link\" href=\"#\">%d</a></li>", i))
		} else {
			buf.WriteString(fmt.Sprintf("<li  class=\"page-item\" ><a class=\"page-link\" href=\"%s\">%d</a></li>", this.url(i), i))
		}
	}

	if totalpage > to {
		buf.WriteString(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\"  href=\"%s\">...%d</a></li>", this.url(totalpage), totalpage))
	}

	if this.Page < totalpage {
		buf.WriteString(fmt.Sprintf("<li class=\"page-item\"><a class=\"page-link\"  href=\"%s\"><i" +
			" class=\"zmdi zmdi-arrow-right\"></i></a></li>",
			this.url(this.Page+1)))
	}


	buf.WriteString("</ul>")

	buf.WriteString("</div>")
	buf.WriteString("</div>")

	return buf.String()
}
