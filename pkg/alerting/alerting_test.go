package alerting

import (
	"encoding/json"
	"testing"

	"bosun.org/graphite"
	. "github.com/smartystreets/goconvey/convey"
)

type fakeGraphite struct {
	resp graphite.Response
}

func (f *fakeGraphite) Query(r *graphite.Request) (graphite.Response, error) {
	return nil, nil
}

func TestAlerting(t *testing.T) {

	Convey("when evaluating graphite checks", t, func() {
		Convey("Series median above threshold should trigger alert", func() {
			checkDef := &CheckDef{
				CritExpr: `median(graphite(test), "2m", "", "") < 100`,
			}

			fg := &fakeGraphite{}
			fg.resp = graphite.Response{
				graphite.Series{
					Target: "test",
					Datapoints: []graphite.DataPoint{
						graphite.DataPoint{json.Number("150"), json.Number("1234567890")},
					},
				}}

			evaluator := &GraphiteCheckEvaluator{
				Context: fg,
				Check:   checkDef,
			}

			Convey("Should eval threshold reached", func() {
				res, err := evaluator.Eval()
				So(err, ShouldBeNil)
				So(res, ShouldEqual, EvalResultCrit)
			})
		})

		Convey("Series median above threshold should trigger alert", func() {
			checkDef := &CheckDef{
				CritExpr: `median(graphite(test), "2m", "", "") < 100`,
			}

			fg := &fakeGraphite{}
			fg.resp = graphite.Response{
				graphite.Series{
					Target: "test",
					Datapoints: []graphite.DataPoint{
						graphite.DataPoint{json.Number("50"), json.Number("1234567890")},
					},
				}}

			evaluator := &GraphiteCheckEvaluator{
				Context: fg,
				Check:   checkDef,
			}

			Convey("Should eval threshold reached", func() {
				res, err := evaluator.Eval()
				So(err, ShouldBeNil)
				So(res, ShouldEqual, EvalResultOK)
			})
		})

	})
}
