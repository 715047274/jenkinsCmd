package model

import _ "time"

// Test represents a test in the results.
type Test struct {
	Title      string      `json:"title"`
	FullTitle  string      `json:"fullTitle"`
	TimedOut   interface{} `json:"timedOut"`
	Duration   int         `json:"duration"`
	State      string      `json:"state"`
	Speed      string      `json:"speed"`
	Pass       bool        `json:"pass"`
	Fail       bool        `json:"fail"`
	Pending    bool        `json:"pending"`
	Context    interface{} `json:"context"`
	Code       string      `json:"code"`
	Err        struct{}    `json:"err"`
	UUID       string      `json:"uuid"`
	ParentUUID string      `json:"parentUUID"`
	IsHook     bool        `json:"isHook"`
	Skipped    bool        `json:"skipped"`
}

// Suite represents a test suite.
type Suite struct {
	UUID        string   `json:"uuid"`
	Title       string   `json:"title"`
	FullFile    string   `json:"fullFile"`
	File        string   `json:"file"`
	BeforeHooks []string `json:"beforeHooks"`
	AfterHooks  []string `json:"afterHooks"`
	Tests       []Test   `json:"tests"`
	Suites      []Suite  `json:"suites"`
	Passes      []string `json:"passes"`
	Failures    []string `json:"failures"`
	Pending     []string `json:"pending"`
	Skipped     []string `json:"skipped"`
	Duration    int      `json:"duration"`
	Root        bool     `json:"root"`
	RootEmpty   bool     `json:"rootEmpty"`
	Timeout     int      `json:"_timeout"`
}

// Result represents a test result.
type Result struct {
	UUID        string   `json:"uuid"`
	Title       string   `json:"title"`
	FullFile    string   `json:"fullFile"`
	File        string   `json:"file"`
	BeforeHooks []string `json:"beforeHooks"`
	AfterHooks  []string `json:"afterHooks"`
	Tests       []Test   `json:"tests"`
	Suites      []Suite  `json:"suites"`
	Passes      []string `json:"passes"`
	Failures    []string `json:"failures"`
	Pending     []string `json:"pending"`
	Skipped     []string `json:"skipped"`
	Duration    int      `json:"duration"`
	Root        bool     `json:"root"`
	RootEmpty   bool     `json:"rootEmpty"`
	Timeout     int      `json:"_timeout"`
}

// Stats represents the statistics.
type Stats struct {
	Suites          int    `json:"suites"`
	Tests           int    `json:"tests"`
	Passes          int    `json:"passes"`
	Pending         int    `json:"pending"`
	Failures        int    `json:"failures"`
	TestsRegistered int    `json:"testsRegistered"`
	PassPercent     int    `json:"passPercent"`
	PendingPercent  int    `json:"pendingPercent"`
	Other           int    `json:"other"`
	HasOther        bool   `json:"hasOther"`
	Skipped         int    `json:"skipped"`
	HasSkipped      bool   `json:"hasSkipped"`
	Start           string `json:"start"`
	End             string `json:"end"`
	Duration        int    `json:"duration"`
}

// Meta represents the metadata.
type Meta struct {
	Mocha struct {
		Version string `json:"version"`
	} `json:"mocha"`
	Mochawesome struct {
		Options struct {
			Quiet           bool   `json:"quiet"`
			ReportFilename  string `json:"reportFilename"`
			SaveHTML        bool   `json:"saveHtml"`
			SaveJSON        bool   `json:"saveJson"`
			ConsoleReporter string `json:"consoleReporter"`
			UseInlineDiffs  bool   `json:"useInlineDiffs"`
			Code            bool   `json:"code"`
		} `json:"options"`
		Version string `json:"version"`
	} `json:"mochawesome"`
	Marge struct {
		Options struct {
			ID                  string `json:"id"`
			ReportDir           string `json:"reportDir"`
			Overwrite           bool   `json:"overwrite"`
			Charts              bool   `json:"charts"`
			ReportPageTitle     string `json:"reportPageTitle"`
			EmbeddedScreenshots bool   `json:"embeddedScreenshots"`
			InlineAssets        bool   `json:"inlineAssets"`
			SaveAllAttempts     bool   `json:"saveAllAttempts"`
			Debug               bool   `json:"debug"`
			SaveJSON            bool   `json:"saveJson"`
			HTML                bool   `json:"html"`
			JSON                bool   `json:"json"`
		} `json:"options"`
		Version string `json:"version"`
	} `json:"marge"`
}

// TestReport represents the entire test report.
type TestReport struct {
	Stats   Stats    `json:"stats"`
	Results []Result `json:"results"`
	Meta    Meta     `json:"meta"`
}
