package report

var htmlTemplate = `
<html>
<head>
  <title>Popeye Sanitizer Report</title>
  <script src="https://kit.fontawesome.com/b45e86135f.js" crossorigin="anonymous"></script>
</head>
<style>
  body {
    background-color: #111;
    color: white;
    font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
  }

  .sanitizer {
    padding: 10px 30px;
  }

  ul.outcome {
    list-style-type: disc;
  }

  div.clear {
    display: block;
  }

  .outcome-score {
    float: right;
  }

  div.outcome {
    display: inline-block;
  }

  .issue {
    text-align: right;
  }

  ul.issues {
    display: block;
    list-style-type: none;
    padding-left: 2px;
  }

  ul.sub-issues {
    list-style-type: none;
    padding-left: 10px;
  }

  .section {
    padding-top: 30px;
  }

  .section-title {
    text-transform: uppercase;
    float: left;
  }

  .scores {
    text-align: right;
  }

  .msg {
    display: block;
  }

  .section-score {
    color: purple;
  }

  .scorer {
    padding-right: 3px;
  }

  .level-0 {
    color: rgb(65, 255, 65);
  }

  .level-1 {
    color: rgb(2, 156, 207);
  }

  .level-2 {
    color: rgb(255, 193, 77);
  }

  .level-3 {
    color: rgb(199, 39, 39);
  }

  .grade-A {
    color: rgb(65, 255, 65);
  }

  .grade-B {
    color: rgb(2, 156, 207);
  }

  .grade-C {
    color: rgb(255, 193, 77);
  }

  .grade-D {
    color: rgb(199, 39, 39);
  }

  .grade-E {
    color: rgb(199, 39, 39);
  }

  .grade-F {
    color: rgb(199, 39, 39);
  }

  .grade {
    font-size: 5em;
  }

  .container {
    color: #38ABCC;
  }


  span.cluster {
    font-style: italic;
    text-transform: uppercase;
    color: greenyellow;
  }

  h3 {
    border-bottom: 1px dashed black;
    width: 50%;
  }

  span.cluster-score {
    font-size: 3em;
  }

  div.score-summary {
    font-size: 2em;
    text-align: center;
    float: left;
  }

  div.title {
    font-size: 3em;
    text-align: center;
  }

  a.popeye-logo {
    display: inline-block;
  }

  div.summary {
    display: flex;
    align-items: center;
    font-weight: 2em;
  }

  img.logo {
    max-width: 175px;
    border-radius: 10px;
    -webkit-filter: drop-shadow(8px 8px 10px #373831);
    filter: drop-shadow(8px 8px 10px #373831);
  }

  div.a {
    color: blue;
    float: left;
    display: block;
  }

  div.scorer {
    width: 90%;
    text-align: right;
  }
</style>

<body>
  <div class="sanitizer">
    <div class="title">Popeye K8s Sanitizer Report</div>
    <div class="summary">
      <a class="popeye-logo" href="https://github.com/derailed/popeye">
        <img class="logo" src="https://github.com/derailed/popeye/raw/master/assets/popeye_logo.png" />
      </a>
      <div class="score-summary">
        Scanned
        <span class="cluster">{{ .ClusterName }}</span>
      </div>
      <div class="scorer">
        <span class="grade grade-{{ .Report.Grade }}">{{ .Report.Grade }}</span>
        <span class="section-score cluster-score"> {{ .Report.Score }} </span>
      </div>
    </div>

    {{ range $section := .Report.Sections }}
    <div class="section">
      <hr />
      <div class="section-title">
        {{ $count := len $section.Outcome }}
        {{ toTitle $section.Title $count }}
      </div>
      <div class="scores">
        <span class="scorer level-3"> <i class="{{ toEmoji 3 }}"></i> {{ $section.Tally.MarshalYAML.Error }} </span>
        <span class="scorer level-2"> <i class="{{ toEmoji 2 }}"></i> {{ $section.Tally.MarshalYAML.Warn }} </span>
        <span class="scorer level-1"> <i class="{{ toEmoji 1 }}"></i> {{ $section.Tally.MarshalYAML.Info }} </span>
        <span class="scorer level-0"> <i class="{{ toEmoji 0 }}"></i> {{ $section.Tally.MarshalYAML.OK }} </span>
        <span class="section-score">{{ $section.Tally.Score }}%</span>
      </div>
      <ul class="outcome">
        {{ range $issueName, $issues := $section.Outcome }}
        <li>
          <div class="outcome level-{{ $issues.MaxSeverity }}">
            {{ $issueName }}
          </div>
          <div class="outcome-score level-{{ $issues.MaxSeverity }}">
            <i class="{{ toEmoji $issues.MaxSeverity }}"></i>
          </div>
          <div class="clear"></div>
          <ul class="issues">
            {{ $group := "" }}
            {{ range $_, $issue := $issues.Sort 0 }}
            {{ if isRoot $issue.Group }}
            <li>
              <span class=" msg level-{{ $issue.Level }}"><i class="{{ toEmoji $issue.Level }}"></i>
                {{ $issue.Message }}
              </span>
            </li>
            {{ else }}
            {{ if ne $group $issue.Group }}
            {{ if ne $group ""}}
          </ul>
          {{ end }}
          {{ $group = $issue.Group }}
        <li class="container">
          <i class="fab fa-docker"></i> {{ $issue.Group }}
        </li>
        <ul class="sub-issues">
          <li>
            <span class=" msg level-{{ $issue.Level }}"><i class="{{ toEmoji $issue.Level }}"></i>
              {{ $issue.Message }}
            </span>
          </li>
          {{ else }}
          <li>
            <span class=" msg level-{{ $issue.Level }}"><i class="{{ toEmoji $issue.Level }}"></i>
              {{ $issue.Message }}
            </span>
          </li>
          {{ end }}
          {{ end }}
          {{ end }}
          {{ if ne $group ""}}
        </ul>
        {{ end }}
      </ul>
      </li>
      {{ end }}
      </ul>
    </div>
    {{ end }}
  </div>
</body>

</html>
`
