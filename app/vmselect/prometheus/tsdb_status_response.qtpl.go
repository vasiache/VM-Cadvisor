// Code generated by qtc from "tsdb_status_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/prometheus/tsdb_status_response.qtpl:1
package prometheus

//line app/vmselect/prometheus/tsdb_status_response.qtpl:1
import "github.com/VictoriaMetrics/VictoriaMetrics/lib/storage"

// TSDBStatusResponse generates response for /api/v1/status/tsdb .

//line app/vmselect/prometheus/tsdb_status_response.qtpl:5
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/prometheus/tsdb_status_response.qtpl:5
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/prometheus/tsdb_status_response.qtpl:5
func StreamTSDBStatusResponse(qw422016 *qt422016.Writer, status *storage.TSDBStatus) {
//line app/vmselect/prometheus/tsdb_status_response.qtpl:5
	qw422016.N().S(`{"status":"success","data":{"seriesCountByMetricName":`)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:9
	streamtsdbStatusEntries(qw422016, status.SeriesCountByMetricName)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:9
	qw422016.N().S(`,"labelValueCountByLabelName":`)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:10
	streamtsdbStatusEntries(qw422016, status.LabelValueCountByLabelName)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:10
	qw422016.N().S(`,"seriesCountByLabelValuePair":`)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:11
	streamtsdbStatusEntries(qw422016, status.SeriesCountByLabelValuePair)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:11
	qw422016.N().S(`}}`)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
}

//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
func WriteTSDBStatusResponse(qq422016 qtio422016.Writer, status *storage.TSDBStatus) {
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
	StreamTSDBStatusResponse(qw422016, status)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
}

//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
func TSDBStatusResponse(status *storage.TSDBStatus) string {
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
	WriteTSDBStatusResponse(qb422016, status)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
	return qs422016
//line app/vmselect/prometheus/tsdb_status_response.qtpl:14
}

//line app/vmselect/prometheus/tsdb_status_response.qtpl:16
func streamtsdbStatusEntries(qw422016 *qt422016.Writer, a []storage.TopHeapEntry) {
//line app/vmselect/prometheus/tsdb_status_response.qtpl:16
	qw422016.N().S(`[`)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:18
	for i, e := range a {
//line app/vmselect/prometheus/tsdb_status_response.qtpl:18
		qw422016.N().S(`{"name":`)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:20
		qw422016.N().Q(e.Name)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:20
		qw422016.N().S(`,"value":`)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:21
		qw422016.N().D(int(e.Count))
//line app/vmselect/prometheus/tsdb_status_response.qtpl:21
		qw422016.N().S(`}`)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:23
		if i+1 < len(a) {
//line app/vmselect/prometheus/tsdb_status_response.qtpl:23
			qw422016.N().S(`,`)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:23
		}
//line app/vmselect/prometheus/tsdb_status_response.qtpl:24
	}
//line app/vmselect/prometheus/tsdb_status_response.qtpl:24
	qw422016.N().S(`]`)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
}

//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
func writetsdbStatusEntries(qq422016 qtio422016.Writer, a []storage.TopHeapEntry) {
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
	streamtsdbStatusEntries(qw422016, a)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
}

//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
func tsdbStatusEntries(a []storage.TopHeapEntry) string {
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
	writetsdbStatusEntries(qb422016, a)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
	qs422016 := string(qb422016.B)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
	return qs422016
//line app/vmselect/prometheus/tsdb_status_response.qtpl:26
}
