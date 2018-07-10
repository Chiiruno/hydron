// This file is automatically generated by qtc from "browser.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line browser.qtpl:1
package templates

//line browser.qtpl:1
import "github.com/bakape/hydron/common"

//line browser.qtpl:2
import "github.com/bakape/hydron/files"

//line browser.qtpl:3
import "strconv"

//line browser.qtpl:4
import "net/url"

//line browser.qtpl:5
import "strings"

//line browser.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line browser.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line browser.qtpl:7
func StreamBrowser(qw422016 *qt422016.Writer, params string, page, totalPages int, imgs []common.CompactImage) {
	//line browser.qtpl:8
	title := params

	//line browser.qtpl:9
	if title == "" {
		//line browser.qtpl:10
		title = "hydron"

		//line browser.qtpl:11
	}
	//line browser.qtpl:12
	streamhead(qw422016, title)
	//line browser.qtpl:12
	qw422016.N().S(`<body><nav id="top-banner"><form method="get"><input type="search" id="search" placeholder="Search" value="`)
	//line browser.qtpl:16
	qw422016.E().S(params)
	//line browser.qtpl:16
	qw422016.N().S(`" name="q" autofocus autocomplete="off" list="search-suggestions"><script>var el = document.getElementById("search");el.selectionStart = el.selectionEnd = el.value.length;</script><datalist id="search-suggestions"></datalist></form>`)
	//line browser.qtpl:23
	streampagination(qw422016, page, totalPages, params)
	//line browser.qtpl:23
	qw422016.N().S(`</nav><div id="overlay"></div><section id="browser">`)
	//line browser.qtpl:27
	for _, img := range imgs {
		//line browser.qtpl:27
		qw422016.N().S(`<label data-type="`)
		//line browser.qtpl:28
		qw422016.N().D(int(img.Type))
		//line browser.qtpl:28
		qw422016.N().S(`" data-sha1="`)
		//line browser.qtpl:28
		qw422016.N().S(img.SHA1)
		//line browser.qtpl:28
		qw422016.N().S(`"><input type="checkbox" name="img:`)
		//line browser.qtpl:29
		qw422016.N().S(img.SHA1)
		//line browser.qtpl:29
		qw422016.N().S(`"><div class="background"></div><img width="`)
		//line browser.qtpl:31
		qw422016.N().S(strconv.FormatUint(img.Thumb.Width, 10))
		//line browser.qtpl:31
		qw422016.N().S(`"height="`)
		//line browser.qtpl:32
		qw422016.N().S(strconv.FormatUint(img.Thumb.Height, 10))
		//line browser.qtpl:32
		qw422016.N().S(`"src="`)
		//line browser.qtpl:33
		qw422016.N().S(files.NetThumbPath(img.SHA1, img.Thumb.IsPNG))
		//line browser.qtpl:33
		qw422016.N().S(`"></label>`)
		//line browser.qtpl:36
	}
	//line browser.qtpl:36
	qw422016.N().S(`</section><script src="/assets/main.js" async></script></body>`)
//line browser.qtpl:40
}

//line browser.qtpl:40
func WriteBrowser(qq422016 qtio422016.Writer, params string, page, totalPages int, imgs []common.CompactImage) {
	//line browser.qtpl:40
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line browser.qtpl:40
	StreamBrowser(qw422016, params, page, totalPages, imgs)
	//line browser.qtpl:40
	qt422016.ReleaseWriter(qw422016)
//line browser.qtpl:40
}

//line browser.qtpl:40
func Browser(params string, page, totalPages int, imgs []common.CompactImage) string {
	//line browser.qtpl:40
	qb422016 := qt422016.AcquireByteBuffer()
	//line browser.qtpl:40
	WriteBrowser(qb422016, params, page, totalPages, imgs)
	//line browser.qtpl:40
	qs422016 := string(qb422016.B)
	//line browser.qtpl:40
	qt422016.ReleaseByteBuffer(qb422016)
	//line browser.qtpl:40
	return qs422016
//line browser.qtpl:40
}

// Links to different pages on a search page

//line browser.qtpl:43
func streampagination(qw422016 *qt422016.Writer, page, total int, params string) {
	//line browser.qtpl:44
	val := url.Values{
		"q": strings.Split(params, " "),
	}

	//line browser.qtpl:46
	qw422016.N().S(`<span class="spaced">`)
	//line browser.qtpl:48
	if page != 0 {
		//line browser.qtpl:49
		if page-1 != 0 {
			//line browser.qtpl:50
			streampageLink(qw422016, val, 0, "<<")
			//line browser.qtpl:51
		}
		//line browser.qtpl:52
		streampageLink(qw422016, val, page-1, "<")
		//line browser.qtpl:53
	}
	//line browser.qtpl:54
	count := 0

	//line browser.qtpl:55
	for i := page - 5; i < total && count < 10; i++ {
		//line browser.qtpl:56
		if i < 0 {
			//line browser.qtpl:57
			continue
			//line browser.qtpl:58
		}
		//line browser.qtpl:59
		count++

		//line browser.qtpl:60
		if i != page {
			//line browser.qtpl:61
			streampageLink(qw422016, val, i, strconv.Itoa(i+1))
			//line browser.qtpl:62
		} else {
			//line browser.qtpl:62
			qw422016.N().S(`<b>`)
			//line browser.qtpl:63
			qw422016.N().D(i + 1)
			//line browser.qtpl:63
			qw422016.N().S(`</b>`)
			//line browser.qtpl:64
		}
		//line browser.qtpl:65
	}
	//line browser.qtpl:66
	if page != total-1 {
		//line browser.qtpl:67
		streampageLink(qw422016, val, page+1, ">")
		//line browser.qtpl:68
		if page+1 != total-1 {
			//line browser.qtpl:69
			streampageLink(qw422016, val, total-1, ">>")
			//line browser.qtpl:70
		}
		//line browser.qtpl:71
	}
	//line browser.qtpl:71
	qw422016.N().S(`</span>`)
//line browser.qtpl:73
}

//line browser.qtpl:73
func writepagination(qq422016 qtio422016.Writer, page, total int, params string) {
	//line browser.qtpl:73
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line browser.qtpl:73
	streampagination(qw422016, page, total, params)
	//line browser.qtpl:73
	qt422016.ReleaseWriter(qw422016)
//line browser.qtpl:73
}

//line browser.qtpl:73
func pagination(page, total int, params string) string {
	//line browser.qtpl:73
	qb422016 := qt422016.AcquireByteBuffer()
	//line browser.qtpl:73
	writepagination(qb422016, page, total, params)
	//line browser.qtpl:73
	qs422016 := string(qb422016.B)
	//line browser.qtpl:73
	qt422016.ReleaseByteBuffer(qb422016)
	//line browser.qtpl:73
	return qs422016
//line browser.qtpl:73
}

// Link to a different paginated search page

//line browser.qtpl:76
func streampageLink(qw422016 *qt422016.Writer, values url.Values, page int, text string) {
	//line browser.qtpl:77
	values.Set("page", strconv.Itoa(page))

	//line browser.qtpl:77
	qw422016.N().S(`<a href="?`)
	//line browser.qtpl:78
	qw422016.N().S(values.Encode())
	//line browser.qtpl:78
	qw422016.N().S(`">`)
	//line browser.qtpl:79
	qw422016.N().S(text)
	//line browser.qtpl:79
	qw422016.N().S(`</a>`)
//line browser.qtpl:81
}

//line browser.qtpl:81
func writepageLink(qq422016 qtio422016.Writer, values url.Values, page int, text string) {
	//line browser.qtpl:81
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line browser.qtpl:81
	streampageLink(qw422016, values, page, text)
	//line browser.qtpl:81
	qt422016.ReleaseWriter(qw422016)
//line browser.qtpl:81
}

//line browser.qtpl:81
func pageLink(values url.Values, page int, text string) string {
	//line browser.qtpl:81
	qb422016 := qt422016.AcquireByteBuffer()
	//line browser.qtpl:81
	writepageLink(qb422016, values, page, text)
	//line browser.qtpl:81
	qs422016 := string(qb422016.B)
	//line browser.qtpl:81
	qt422016.ReleaseByteBuffer(qb422016)
	//line browser.qtpl:81
	return qs422016
//line browser.qtpl:81
}
