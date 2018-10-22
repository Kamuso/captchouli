// This file is automatically generated by qtc from "captcha.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line captcha.qtpl:2
package templates

//line captcha.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line captcha.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line captcha.qtpl:2
func StreamCaptcha(qw422016 *qt422016.Writer, colour, background, tag string, id [64]byte, images [9][16]byte) {
	//line captcha.qtpl:2
	qw422016.N().S(`<style>.captchouli-checkbox {display: none;}.captchouli-checkbox:checked ~ .captchouli-img {transform: scale(0.85);}.captchouli-img {margin: 2px;user-select: none;}.captchouli-width {width: 462px;}.captchouli-margin {margin: 4px 0;}</style><form class="captchouli-width" style="height:525px; background:`)
	//line captcha.qtpl:21
	qw422016.E().S(background)
	//line captcha.qtpl:21
	qw422016.N().S(`; color:`)
	//line captcha.qtpl:21
	qw422016.E().S(colour)
	//line captcha.qtpl:21
	qw422016.N().S(`; font-family:Sans-Serif;"><input type="text" name="id" hidden value="`)
	//line captcha.qtpl:22
	streamencodeID(qw422016, id)
	//line captcha.qtpl:22
	qw422016.N().S(`"><header class="captchouli-width captchouli-margin" style="text-align:center; font-size:130%;">Select all images of <b>`)
	//line captcha.qtpl:24
	qw422016.E().S(tag)
	//line captcha.qtpl:24
	qw422016.N().S(`</b></header><div class="captchouli-width" style="height:462px;">`)
	//line captcha.qtpl:27
	buf := make([]byte, 4096)

	//line captcha.qtpl:28
	for i, img := range images {
		//line captcha.qtpl:28
		qw422016.N().S(`<label><input type="checkbox" name="`)
		//line captcha.qtpl:30
		qw422016.N().D(i)
		//line captcha.qtpl:30
		qw422016.N().S(`" class="captchouli-checkbox"><img class="captchouli-img" draggable="false" src="`)
		//line captcha.qtpl:31
		streamthumbnail(qw422016, img, buf)
		//line captcha.qtpl:31
		qw422016.N().S(`"></label>`)
		//line captcha.qtpl:33
	}
	//line captcha.qtpl:33
	qw422016.N().S(`</div><input type="submit" class="captchouli-width captchouli-margin"></form>`)
//line captcha.qtpl:37
}

//line captcha.qtpl:37
func WriteCaptcha(qq422016 qtio422016.Writer, colour, background, tag string, id [64]byte, images [9][16]byte) {
	//line captcha.qtpl:37
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line captcha.qtpl:37
	StreamCaptcha(qw422016, colour, background, tag, id, images)
	//line captcha.qtpl:37
	qt422016.ReleaseWriter(qw422016)
//line captcha.qtpl:37
}

//line captcha.qtpl:37
func Captcha(colour, background, tag string, id [64]byte, images [9][16]byte) string {
	//line captcha.qtpl:37
	qb422016 := qt422016.AcquireByteBuffer()
	//line captcha.qtpl:37
	WriteCaptcha(qb422016, colour, background, tag, id, images)
	//line captcha.qtpl:37
	qs422016 := string(qb422016.B)
	//line captcha.qtpl:37
	qt422016.ReleaseByteBuffer(qb422016)
	//line captcha.qtpl:37
	return qs422016
//line captcha.qtpl:37
}
