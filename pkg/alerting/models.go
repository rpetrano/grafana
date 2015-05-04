package alerting

import (
	"fmt"

	"bosun.org/graphite"
)

type CheckDef struct {
	// Freq      uint32
	// Offset    uint8 // offset on top of "even" minute/10s/.. intervals
	// Expr      string // "median(foo...) "
	// LevelWarn float64 // > 5 , < 10
	// LevelCrit float64 // > 10, < 5
	CritExpr string
	WarnExpr string
}

type Check struct {
	Id           int64
	OrgId        int64
	DataSourceId int64
	Definition   CheckDef
}

// func (check *Check) getDataSource() {
// 	dsQuery := m.GetDataSourceByIdQuery{Id: check.Id}
//
// 	if err := bus.Dispatch(&dsQuery); err != nil {
// 		return nil, err
// 	}
//
// 	return dsQuery.Result
// }

type CheckEvaluator interface {
	Eval() (*CheckEvalResult, error)
}

type GraphiteCheckEvaluator struct {
	Context graphite.Context
	Check   *CheckDef
}

type CheckEvalResult int

const (
	EvalResultOK CheckEvalResult = 1 << iota
	EvalResultWarn
	EvalResultCrit
)

func (c CheckEvalResult) String() string {
	switch c {
	case EvalResultOK:
		return "OK"
	case EvalResultWarn:
		return "Warning"
	case EvalResultCrit:
		return "Critical"
	default:
		panic(fmt.Sprintf("Invalid CheckEvalResult value %d", int(c)))
	}
}

type DefaultCheckEvaluator struct {
}

func (ce *GraphiteCheckEvaluator) Eval() (*CheckEvalResult, error) {
	// 	exp, err := expr.New(job.expr, expr.Graphite)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	// create cache
	// 	// this is so that when bosun queries the same graphite query multiple times
	// 	// like in (median(graphite("foo")> 10 || avg(graphite("foo") > 20)
	// 	// it reuses the same resultsets internally.
	// 	// cache is unbounded so that we are guaranteed internally consistent results
	// 	// TODO recreate new cache at each second because cache is pointless at the next interval
	// 	cacheObj := cache.New(0)
	//
	// 	// TODO once auth works, do it without rh. should work too because RH should only be used for bosun's own alerts and notifications
	// 	rh := &sched.RunHistory{
	// 		Cache:           cacheObj,
	// 		Start:           job.ts, // this sets an explicit "until" to match the data this alert run is meant for, even when we are delayed
	// 		Events:          make(map[expr.AlertKey]*sched.Event),
	// 		Context:         nil,
	// 		GraphiteContext: ce.Context,
	// 		Logstash:        make([]string, 0),
	// 	}
	// 	results, _, err := exp.Execute(rh.Context, rh.GraphiteContext, rh.Logstash, rh.Cache, nil, rh.Start, 0, true, nil, nil, nil)
	// 	fmt.Println(job.ts, job.expr)
	// 	for _, res := range results.Results {
	// 		fmt.Println(res.Group, res.Value)
	// 	}
	// 	spew.Dump(err)
	// }

	return nil, nil
}
