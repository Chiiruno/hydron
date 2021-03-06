// This file is automatically generated by qtc from "browser.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line browser.qtpl:1
package templates

//line browser.qtpl:1
import "github.com/bakape/hydron/common"

//line browser.qtpl:2
import "strconv"

//line browser.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line browser.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line browser.qtpl:4
func StreamBrowser(qw422016 *qt422016.Writer, page common.Page, imgs []common.CompactImage) {
	//line browser.qtpl:5
	filters := page.Filters.String()

	//line browser.qtpl:6
	title := filters

	//line browser.qtpl:7
	if title == "" {
		//line browser.qtpl:8
		title = "hydron"

		//line browser.qtpl:9
	}
	//line browser.qtpl:10
	streamhead(qw422016, title)
	//line browser.qtpl:10
	qw422016.N().S(`<body><nav id="top-banner"><div style="display: flex;"><form method="get"><input type="search" id="search" placeholder="Search" value="`)
	//line browser.qtpl:15
	qw422016.E().S(filters)
	//line browser.qtpl:15
	qw422016.N().S(`" name="q" autofocus autocomplete="off" list="search-suggestions"><script>var el = document.getElementById("search");el.selectionStart = el.selectionEnd = el.value.length;</script><datalist id="search-suggestions"></datalist><select name="order" tabindex="-1" title="Order by">`)
	//line browser.qtpl:22
	for i := common.None; i <= common.Random; i++ {
		//line browser.qtpl:22
		qw422016.N().S(`<option value="`)
		//line browser.qtpl:23
		qw422016.N().D(int(i))
		//line browser.qtpl:23
		qw422016.N().S(`"`)
		//line browser.qtpl:23
		if i == page.Order.Type {
			//line browser.qtpl:23
			qw422016.N().S(` `)
			//line browser.qtpl:23
			qw422016.N().S(`selected`)
			//line browser.qtpl:23
		}
		//line browser.qtpl:23
		qw422016.N().S(`>`)
		//line browser.qtpl:24
		qw422016.N().S(orderLabels[int(i)])
		//line browser.qtpl:24
		qw422016.N().S(`</option>`)
		//line browser.qtpl:26
	}
	//line browser.qtpl:26
	qw422016.N().S(`</select><input type="checkbox" name="reverse" tabindex="-1" title="Reverse order"`)
	//line browser.qtpl:28
	if page.Order.Reverse {
		//line browser.qtpl:28
		qw422016.N().S(` `)
		//line browser.qtpl:28
		qw422016.N().S(`checked`)
		//line browser.qtpl:28
	}
	//line browser.qtpl:28
	qw422016.N().S(`></form>`)
	//line browser.qtpl:30
	streampagination(qw422016, page)
	//line browser.qtpl:30
	qw422016.N().S(`</div><div style="width: 100%; height: 0.3em;"><div id="progress-bar"></div></div></nav><section id="browser" tabindex="1">`)
	//line browser.qtpl:37
	for i, img := range imgs {
		//line browser.qtpl:38
		StreamThumbnail(qw422016, img, page, i == 0)
		//line browser.qtpl:39
	}
	//line browser.qtpl:39
	qw422016.N().S(`</section><script src="/assets/main.js" async></script></body>`)
//line browser.qtpl:43
}

//line browser.qtpl:43
func WriteBrowser(qq422016 qtio422016.Writer, page common.Page, imgs []common.CompactImage) {
	//line browser.qtpl:43
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line browser.qtpl:43
	StreamBrowser(qw422016, page, imgs)
	//line browser.qtpl:43
	qt422016.ReleaseWriter(qw422016)
//line browser.qtpl:43
}

//line browser.qtpl:43
func Browser(page common.Page, imgs []common.CompactImage) string {
	//line browser.qtpl:43
	qb422016 := qt422016.AcquireByteBuffer()
	//line browser.qtpl:43
	WriteBrowser(qb422016, page, imgs)
	//line browser.qtpl:43
	qs422016 := string(qb422016.B)
	//line browser.qtpl:43
	qt422016.ReleaseByteBuffer(qb422016)
	//line browser.qtpl:43
	return qs422016
//line browser.qtpl:43
}

// Links to different pages on a search page

//line browser.qtpl:46
func streampagination(qw422016 *qt422016.Writer, page common.Page) {
	//line browser.qtpl:46
	qw422016.N().S(`<span class="spaced">`)
	//line browser.qtpl:48
	current := int(page.Page)

	//line browser.qtpl:49
	total := int(page.PageTotal)

	//line browser.qtpl:50
	if current != 0 {
		//line browser.qtpl:51
		if current-1 != 0 {
			//line browser.qtpl:52
			streampageLink(qw422016, page, 0, "<<")
			//line browser.qtpl:53
		}
		//line browser.qtpl:54
		streampageLink(qw422016, page, current-1, "<")
		//line browser.qtpl:55
	}
	//line browser.qtpl:56
	count := 0

	//line browser.qtpl:57
	for i := current - 5; i < total && count < 10; i++ {
		//line browser.qtpl:58
		if i < 0 {
			//line browser.qtpl:59
			continue
			//line browser.qtpl:60
		}
		//line browser.qtpl:61
		count++

		//line browser.qtpl:62
		if i != current {
			//line browser.qtpl:63
			streampageLink(qw422016, page, i, strconv.Itoa(i+1))
			//line browser.qtpl:64
		} else {
			//line browser.qtpl:64
			qw422016.N().S(`<b>`)
			//line browser.qtpl:65
			qw422016.N().D(i + 1)
			//line browser.qtpl:65
			qw422016.N().S(`</b>`)
			//line browser.qtpl:66
		}
		//line browser.qtpl:67
	}
	//line browser.qtpl:68
	if current != total-1 {
		//line browser.qtpl:69
		streampageLink(qw422016, page, current+1, ">")
		//line browser.qtpl:70
		if current+1 != total-1 {
			//line browser.qtpl:71
			streampageLink(qw422016, page, total-1, ">>")
			//line browser.qtpl:72
		}
		//line browser.qtpl:73
	}
	//line browser.qtpl:73
	qw422016.N().S(`</span>`)
//line browser.qtpl:75
}

//line browser.qtpl:75
func writepagination(qq422016 qtio422016.Writer, page common.Page) {
	//line browser.qtpl:75
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line browser.qtpl:75
	streampagination(qw422016, page)
	//line browser.qtpl:75
	qt422016.ReleaseWriter(qw422016)
//line browser.qtpl:75
}

//line browser.qtpl:75
func pagination(page common.Page) string {
	//line browser.qtpl:75
	qb422016 := qt422016.AcquireByteBuffer()
	//line browser.qtpl:75
	writepagination(qb422016, page)
	//line browser.qtpl:75
	qs422016 := string(qb422016.B)
	//line browser.qtpl:75
	qt422016.ReleaseByteBuffer(qb422016)
	//line browser.qtpl:75
	return qs422016
//line browser.qtpl:75
}

// Link to a different paginated search page

//line browser.qtpl:78
func streampageLink(qw422016 *qt422016.Writer, page common.Page, i int, text string) {
	//line browser.qtpl:79
	page.Page = uint(i)

	//line browser.qtpl:79
	qw422016.N().S(`<a href="`)
	//line browser.qtpl:80
	qw422016.N().S(page.URL())
	//line browser.qtpl:80
	qw422016.N().S(`" tabindex="2">`)
	//line browser.qtpl:81
	qw422016.N().S(text)
	//line browser.qtpl:81
	qw422016.N().S(`</a>`)
//line browser.qtpl:83
}

//line browser.qtpl:83
func writepageLink(qq422016 qtio422016.Writer, page common.Page, i int, text string) {
	//line browser.qtpl:83
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line browser.qtpl:83
	streampageLink(qw422016, page, i, text)
	//line browser.qtpl:83
	qt422016.ReleaseWriter(qw422016)
//line browser.qtpl:83
}

//line browser.qtpl:83
func pageLink(page common.Page, i int, text string) string {
	//line browser.qtpl:83
	qb422016 := qt422016.AcquireByteBuffer()
	//line browser.qtpl:83
	writepageLink(qb422016, page, i, text)
	//line browser.qtpl:83
	qs422016 := string(qb422016.B)
	//line browser.qtpl:83
	qt422016.ReleaseByteBuffer(qb422016)
	//line browser.qtpl:83
	return qs422016
//line browser.qtpl:83
}
