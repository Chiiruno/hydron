// This file is automatically generated by qtc from "browser.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line browser.qtpl:1
package templates

//line browser.qtpl:1
import "github.com/bakape/hydron/common"

//line browser.qtpl:2
import "strconv"

//line browser.qtpl:3
import "net/url"

//line browser.qtpl:4
import "strings"

//line browser.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line browser.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line browser.qtpl:6
func StreamBrowser(qw422016 *qt422016.Writer, params string, page, totalPages int, imgs []common.CompactImage) {
	//line browser.qtpl:7
	title := params

	//line browser.qtpl:8
	if title == "" {
		//line browser.qtpl:9
		title = "hydron"

		//line browser.qtpl:10
	}
	//line browser.qtpl:11
	streamhead(qw422016, title)
	//line browser.qtpl:11
	qw422016.N().S(`<body><nav id="top-banner"><div style="display: flex;"><form method="get"><input type="search" id="search" placeholder="Search" value="`)
	//line browser.qtpl:16
	qw422016.E().S(params)
	//line browser.qtpl:16
	qw422016.N().S(`" name="q" autofocus autocomplete="off" list="search-suggestions"><script>var el = document.getElementById("search");el.selectionStart = el.selectionEnd = el.value.length;</script><datalist id="search-suggestions"></datalist></form>`)
	//line browser.qtpl:23
	streampagination(qw422016, page, totalPages, params)
	//line browser.qtpl:23
	qw422016.N().S(`</div><div style="width: 100%; height: 0.3em;"><div id="progress-bar"></div></div></nav><div id="image-view"></div><section id="browser">`)
	//line browser.qtpl:31
	for _, img := range imgs {
		//line browser.qtpl:32
		StreamThumbnail(qw422016, img)
		//line browser.qtpl:33
	}
	//line browser.qtpl:33
	qw422016.N().S(`</section><script src="/assets/main.js" async></script></body>`)
//line browser.qtpl:37
}

//line browser.qtpl:37
func WriteBrowser(qq422016 qtio422016.Writer, params string, page, totalPages int, imgs []common.CompactImage) {
	//line browser.qtpl:37
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line browser.qtpl:37
	StreamBrowser(qw422016, params, page, totalPages, imgs)
	//line browser.qtpl:37
	qt422016.ReleaseWriter(qw422016)
//line browser.qtpl:37
}

//line browser.qtpl:37
func Browser(params string, page, totalPages int, imgs []common.CompactImage) string {
	//line browser.qtpl:37
	qb422016 := qt422016.AcquireByteBuffer()
	//line browser.qtpl:37
	WriteBrowser(qb422016, params, page, totalPages, imgs)
	//line browser.qtpl:37
	qs422016 := string(qb422016.B)
	//line browser.qtpl:37
	qt422016.ReleaseByteBuffer(qb422016)
	//line browser.qtpl:37
	return qs422016
//line browser.qtpl:37
}

// Links to different pages on a search page

//line browser.qtpl:40
func streampagination(qw422016 *qt422016.Writer, page, total int, params string) {
	//line browser.qtpl:41
	val := url.Values{
		"q": strings.Split(params, " "),
	}

	//line browser.qtpl:43
	qw422016.N().S(`<span class="spaced">`)
	//line browser.qtpl:45
	if page != 0 {
		//line browser.qtpl:46
		if page-1 != 0 {
			//line browser.qtpl:47
			streampageLink(qw422016, val, 0, "<<")
			//line browser.qtpl:48
		}
		//line browser.qtpl:49
		streampageLink(qw422016, val, page-1, "<")
		//line browser.qtpl:50
	}
	//line browser.qtpl:51
	count := 0

	//line browser.qtpl:52
	for i := page - 5; i < total && count < 10; i++ {
		//line browser.qtpl:53
		if i < 0 {
			//line browser.qtpl:54
			continue
			//line browser.qtpl:55
		}
		//line browser.qtpl:56
		count++

		//line browser.qtpl:57
		if i != page {
			//line browser.qtpl:58
			streampageLink(qw422016, val, i, strconv.Itoa(i+1))
			//line browser.qtpl:59
		} else {
			//line browser.qtpl:59
			qw422016.N().S(`<b>`)
			//line browser.qtpl:60
			qw422016.N().D(i + 1)
			//line browser.qtpl:60
			qw422016.N().S(`</b>`)
			//line browser.qtpl:61
		}
		//line browser.qtpl:62
	}
	//line browser.qtpl:63
	if page != total-1 {
		//line browser.qtpl:64
		streampageLink(qw422016, val, page+1, ">")
		//line browser.qtpl:65
		if page+1 != total-1 {
			//line browser.qtpl:66
			streampageLink(qw422016, val, total-1, ">>")
			//line browser.qtpl:67
		}
		//line browser.qtpl:68
	}
	//line browser.qtpl:68
	qw422016.N().S(`</span>`)
//line browser.qtpl:70
}

//line browser.qtpl:70
func writepagination(qq422016 qtio422016.Writer, page, total int, params string) {
	//line browser.qtpl:70
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line browser.qtpl:70
	streampagination(qw422016, page, total, params)
	//line browser.qtpl:70
	qt422016.ReleaseWriter(qw422016)
//line browser.qtpl:70
}

//line browser.qtpl:70
func pagination(page, total int, params string) string {
	//line browser.qtpl:70
	qb422016 := qt422016.AcquireByteBuffer()
	//line browser.qtpl:70
	writepagination(qb422016, page, total, params)
	//line browser.qtpl:70
	qs422016 := string(qb422016.B)
	//line browser.qtpl:70
	qt422016.ReleaseByteBuffer(qb422016)
	//line browser.qtpl:70
	return qs422016
//line browser.qtpl:70
}

// Link to a different paginated search page

//line browser.qtpl:73
func streampageLink(qw422016 *qt422016.Writer, values url.Values, page int, text string) {
	//line browser.qtpl:74
	values.Set("page", strconv.Itoa(page))

	//line browser.qtpl:74
	qw422016.N().S(`<a href="?`)
	//line browser.qtpl:75
	qw422016.N().S(values.Encode())
	//line browser.qtpl:75
	qw422016.N().S(`">`)
	//line browser.qtpl:76
	qw422016.N().S(text)
	//line browser.qtpl:76
	qw422016.N().S(`</a>`)
//line browser.qtpl:78
}

//line browser.qtpl:78
func writepageLink(qq422016 qtio422016.Writer, values url.Values, page int, text string) {
	//line browser.qtpl:78
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line browser.qtpl:78
	streampageLink(qw422016, values, page, text)
	//line browser.qtpl:78
	qt422016.ReleaseWriter(qw422016)
//line browser.qtpl:78
}

//line browser.qtpl:78
func pageLink(values url.Values, page int, text string) string {
	//line browser.qtpl:78
	qb422016 := qt422016.AcquireByteBuffer()
	//line browser.qtpl:78
	writepageLink(qb422016, values, page, text)
	//line browser.qtpl:78
	qs422016 := string(qb422016.B)
	//line browser.qtpl:78
	qt422016.ReleaseByteBuffer(qb422016)
	//line browser.qtpl:78
	return qs422016
//line browser.qtpl:78
}
