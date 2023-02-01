package report_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/derailed/popeye/internal/client"
	"github.com/derailed/popeye/internal/issues"
	"github.com/derailed/popeye/internal/report"
	"github.com/derailed/popeye/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestBuilderHtml(t *testing.T) {
	b, ta := report.NewBuilder(), report.NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(client.NewGVR("fred"), issues.Root, config.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection(client.NewGVR("fred"), "fred", o, ta)
	b.AddError(errors.New("boom"))
	s, err := b.ToHTML()

	assert.Nil(t, err)
	assert.Equal(t, reportHTML, s)
}

func TestBuilderJunit(t *testing.T) {
	b, ta := report.NewBuilder(), report.NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(client.NewGVR("fred"), issues.Root, config.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection(client.NewGVR("fred"), "fred", o, ta)
	b.AddError(errors.New("boom"))
	s, err := b.ToJunit(config.OkLevel)

	assert.Nil(t, err)
	assert.Equal(t, reportJunit, s)
}

func TestBuilderYAML(t *testing.T) {
	b, ta := report.NewBuilder(), report.NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(client.NewGVR("fred"), issues.Root, config.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection(client.NewGVR("fred"), "fred", o, ta)
	b.AddError(errors.New("boom"))
	s, err := b.ToYAML()

	assert.Nil(t, err)
	assert.Equal(t, reportYAML, s)
}

func TestBuilderJSON(t *testing.T) {
	b, ta := report.NewBuilder(), report.NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(client.NewGVR("fred"), issues.Root, config.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection(client.NewGVR("fred"), "fred", o, ta)
	b.AddError(errors.New("boom"))
	s, err := b.ToJSON()

	assert.Nil(t, err)
	assert.Equal(t, reportJSON, s)
}

func TestPrintSummary(t *testing.T) {
	b, ta := report.NewBuilder(), report.NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(client.NewGVR("fred"), issues.Root, config.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection(client.NewGVR("fred"), "fred", o, ta)
	b.AddError(errors.New("boom"))

	buff := bytes.NewBuffer([]byte(""))
	san := report.NewSanitizer(buff, false)
	b.PrintSummary(san)

	assert.Equal(t, summaryExp, buff.String())
}

func TestPrintHeader(t *testing.T) {
	b, ta := report.NewBuilder(), report.NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(client.NewGVR("fred"), issues.Root, config.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection(client.NewGVR("fred"), "fred", o, ta)
	b.AddError(errors.New("boom"))

	buff := bytes.NewBuffer([]byte(""))
	san := report.NewSanitizer(buff, false)
	b.PrintHeader(san)

	assert.Equal(t, headerExp, buff.String())
}

func TestPrintReport(t *testing.T) {
	b, ta := report.NewBuilder(), report.NewTally()
	o := issues.Outcome{
		"blee": issues.Issues{
			issues.New(client.NewGVR("fred"), issues.Root, config.OkLevel, "Blah"),
		},
	}

	ta.Rollup(o)
	b.AddSection(client.NewGVR("fred"), "fred", o, ta)
	b.AddError(errors.New("boom"))

	buff := bytes.NewBuffer([]byte(""))
	san := report.NewSanitizer(buff, false)
	b.PrintReport(config.OkLevel, san)

	assert.Equal(t, reportExp, buff.String())
}

func TestTitleize(t *testing.T) {
	uu := map[string]struct {
		count    int
		title, e string
	}{
		"none": {count: -1, title: "FRED", e: "FRED"},
		"one":  {count: 1, title: "FRED", e: "FRED (1 SCANNED)"},
		"many": {count: 2, title: "FRED", e: "FRED (2 SCANNED)"},
	}

	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, report.Titleize(u.title, u.count))
		})
	}
}

// ----------------------------------------------------------------------------
// Helpers...

var (
	reportHTML  = "\n<html>\n<head>\n  <title>Popeye Sanitizer Report</title>\n  <script src=\"https://kit.fontawesome.com/b45e86135f.js\" crossorigin=\"anonymous\"></script>\n</head>\n<style>\n  body {\n    background-color: #111;\n    color: white;\n    font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;\n  }\n\n  .sanitizer {\n    padding: 10px 30px;\n  }\n\n  ul.outcome {\n    list-style-type: disc;\n  }\n\n  div.clear {\n    display: block;\n  }\n\n  .outcome-score {\n    float: right;\n  }\n\n  div.outcome {\n    display: inline-block;\n  }\n\n  .issue {\n    text-align: right;\n  }\n\n  ul.issues {\n    display: block;\n    list-style-type: none;\n    padding-left: 2px;\n  }\n\n  ul.sub-issues {\n    list-style-type: none;\n    padding-left: 10px;\n  }\n\n  .section {\n    padding-top: 30px;\n  }\n\n  .section-title {\n    text-transform: uppercase;\n    float: left;\n  }\n\n  .scores {\n    text-align: right;\n  }\n\n  .msg {\n    display: block;\n  }\n\n  .section-score {\n    color: purple;\n  }\n\n  .scorer {\n    padding-right: 3px;\n  }\n\n  .level-0 {\n    color: rgb(65, 255, 65);\n  }\n\n  .level-1 {\n    color: rgb(2, 156, 207);\n  }\n\n  .level-2 {\n    color: rgb(255, 193, 77);\n  }\n\n  .level-3 {\n    color: rgb(199, 39, 39);\n  }\n\n  .grade-A {\n    color: rgb(65, 255, 65);\n  }\n\n  .grade-B {\n    color: rgb(2, 156, 207);\n  }\n\n  .grade-C {\n    color: rgb(255, 193, 77);\n  }\n\n  .grade-D {\n    color: rgb(199, 39, 39);\n  }\n\n  .grade-E {\n    color: rgb(199, 39, 39);\n  }\n\n  .grade-F {\n    color: rgb(199, 39, 39);\n  }\n\n  .grade {\n    font-size: 5em;\n  }\n\n  .container {\n    color: #38ABCC;\n  }\n\n\n  span.cluster {\n    font-style: italic;\n    text-transform: uppercase;\n    color: greenyellow;\n  }\n\n  h3 {\n    border-bottom: 1px dashed black;\n    width: 50%;\n  }\n\n  span.cluster-score {\n    font-size: 3em;\n  }\n\n  div.score-summary {\n    font-size: 2em;\n    text-align: center;\n    float: left;\n  }\n\n  div.title {\n    font-size: 3em;\n    text-align: center;\n  }\n\n  a.popeye-logo {\n    display: inline-block;\n  }\n\n  div.summary {\n    display: flex;\n    align-items: center;\n    font-weight: 2em;\n  }\n\n  img.logo {\n    max-width: 175px;\n    border-radius: 10px;\n    -webkit-filter: drop-shadow(8px 8px 10px #373831);\n    filter: drop-shadow(8px 8px 10px #373831);\n  }\n\n  div.a {\n    color: blue;\n    float: left;\n    display: block;\n  }\n\n  div.scorer {\n    width: 90%;\n    text-align: right;\n  }\n</style>\n\n<body>\n  <div class=\"sanitizer\">\n    <div class=\"title\">Popeye K8s Sanitizer Report</div>\n    <div class=\"summary\">\n      <a class=\"popeye-logo\" href=\"https://github.com/derailed/popeye\">\n        <img class=\"logo\" src=\"https://github.com/derailed/popeye/raw/master/assets/popeye_logo.png\" />\n      </a>\n      <div class=\"score-summary\">\n        Scanned\n        <span class=\"cluster\"></span>\n      </div>\n      <div class=\"scorer\">\n        <span class=\"grade grade-A\">A</span>\n        <span class=\"section-score cluster-score\"> 100 </span>\n      </div>\n    </div>\n\n    \n    <div class=\"section\">\n      <hr />\n      <div class=\"section-title\">\n        \n        FRED (1 SCANNED)\n      </div>\n      <div class=\"scores\">\n        <span class=\"scorer level-3\"> <i class=\"fas fa-bomb\"></i> 0 </span>\n        <span class=\"scorer level-2\"> <i class=\"fas fa-radiation-alt\"></i> 0 </span>\n        <span class=\"scorer level-1\"> <i class=\"fas fa-info-circle\"></i> 0 </span>\n        <span class=\"scorer level-0\"> <i class=\"far fa-check-circle\"></i> 1 </span>\n        <span class=\"section-score\">100%</span>\n      </div>\n      <ul class=\"outcome\">\n        \n        <li>\n          <div class=\"outcome level-0\">\n            blee\n          </div>\n          <div class=\"outcome-score level-0\">\n            <i class=\"far fa-check-circle\"></i>\n          </div>\n          <div class=\"clear\"></div>\n          <ul class=\"issues\">\n            \n            \n            \n            <li>\n              <span class=\" msg level-0\"><i class=\"far fa-check-circle\"></i>\n                Blah\n              </span>\n            </li>\n            \n          \n          \n      </ul>\n      </li>\n      \n      </ul>\n    </div>\n    \n  </div>\n</body>\n\n</html>\n"
	reportJunit = "<testsuites name=\"Popeye\" tests=\"1\" failures=\"0\" errors=\"1\">\n\t<testsuite name=\"fred\" tests=\"1\" failures=\"0\" errors=\"0\">\n\t\t<properties>\n\t\t\t<property name=\"OK\" value=\"1\"></property>\n\t\t\t<property name=\"Info\" value=\"0\"></property>\n\t\t\t<property name=\"Warn\" value=\"0\"></property>\n\t\t\t<property name=\"Error\" value=\"0\"></property>\n\t\t\t<property name=\"Score\" value=\"100%\"></property>\n\t\t</properties>\n\t\t<testcase classname=\"\" name=\"blee\"></testcase>\n\t</testsuite>\n</testsuites>"
	reportJSON  = "{\"popeye\":{\"score\":100,\"grade\":\"A\",\"sanitizers\":[{\"sanitizer\":\"fred\",\"gvr\":\"fred\",\"tally\":{\"ok\":1,\"info\":0,\"warning\":0,\"error\":0,\"score\":100},\"issues\":{\"blee\":[{\"group\":\"__root__\",\"gvr\":\"fred\",\"level\":0,\"message\":\"Blah\"}]}}],\"errors\":[{}]}}"

	reportYAML = `popeye:
  score: 100
  grade: A
  sanitizers:
  - sanitizer: fred
    gvr: fred
    tally:
      ok: 1
      info: 0
      warning: 0
      error: 0
      score: 100
    issues:
      blee:
      - group: __root__
        gvr: fred
        level: 0
        message: Blah
  errors:
  - {}
`

	summaryExp = "\n\x1b[38;5;75mSUMMARY\x1b[0m\n\x1b[38;5;75m" + strings.Repeat("┅", 101) + "\x1b[0m\nYour cluster score: 100 -- A\n                                                                                \x1b[38;5;82mo          .-'-.     \x1b[0m\n                                                                                \x1b[38;5;82m o     __| A    `\\  \x1b[0m\n                                                                                \x1b[38;5;82m  o   `-,-`--._   `\\\x1b[0m\n                                                                                \x1b[38;5;82m []  .->'  a     `|-'\x1b[0m\n                                                                                \x1b[38;5;82m  `=/ (__/_       /  \x1b[0m\n                                                                                \x1b[38;5;82m    \\_,    `    _)  \x1b[0m\n                                                                                \x1b[38;5;82m       `----;  |     \x1b[0m\n\n"
	headerExp  = "\n\x1b[38;5;122m ___     ___ _____   _____ \x1b[0m                                                     \x1b[38;5;75mK          .-'-.     \x1b[0m\n\x1b[38;5;122m| _ \\___| _ \\ __\\ \\ / / __|\x1b[0m                                                     \x1b[38;5;75m 8     __|      `\\  \x1b[0m\n\x1b[38;5;122m|  _/ _ \\  _/ _| \\ V /| _| \x1b[0m                                                     \x1b[38;5;75m  s   `-,-`--._   `\\\x1b[0m\n\x1b[38;5;122m|_| \\___/_| |___| |_| |___|\x1b[0m                                                     \x1b[38;5;75m []  .->'  a     `|-'\x1b[0m\n\x1b[38;5;75m  Biffs`em and Buffs`em!\x1b[0m                                                        \x1b[38;5;75m  `=/ (__/_       /  \x1b[0m\n                                                                                \x1b[38;5;75m    \\_,    `    _)  \x1b[0m\n                                                                                \x1b[38;5;75m       `----;  |     \x1b[0m\n\n"
	reportExp  = "\n\x1b[38;5;75mFRED (1 SCANNED)\x1b[0m" + strings.Repeat(" ", 61) + "💥 0 😱 0 🔊 0 ✅ 1 \x1b[38;5;122m100\x1b[0m٪\n\x1b[38;5;75m" + strings.Repeat("┅", 101) + "\x1b[0m\n  · \x1b[38;5;155mblee\x1b[0m\x1b[38;5;250m" + strings.Repeat(".", 91) + "\x1b[0m✅\n    ✅ \x1b[38;5;155mBlah.\x1b[0m\n\n"
)
