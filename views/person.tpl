{{ define "content" }}
<div class="container">
  <div class="row">
    <div class="col">
      <h1>{{ .Person.Name }}
        <span class="card-subtitle text-muted"
              style="font-size: 18px;vertical-align:top;">{{ .Person.GetCurrentRankName }}
          ({{ .Person.GetStatus }}{{if eq .Person.Status "delete"}} {{ .Person.GetDismissalDate }}{{ end }})
                </span>
      </h1>
    </div>
  </div>
{{ if .Header.User.IsAdmin }}
  <div class="row" style="margin-bottom: 20px">
    <div class="col">
      <div class="btn-group" role="group" aria-label="Basic example">
        <button type="button" class="btn btn-info" data-toggle="modal" data-target="#changePersonModal">
          Изменить
        </button>
        <button type="button" class="btn btn-warning" data-toggle="modal" data-target="#changeRankModal">
          Должность
        </button>
        <button type="button" class="btn btn-danger" data-person="{{ .Person.Id }}" data-toggle="modal"
                data-target="#newRebukeModal">
          Выговор
        </button>
        <button type="button" data-person="{{ .Person.Id }}" class="btn btn-secondary" data-toggle="modal"
                data-target="#setVacationModal">
          В отпуск
        </button>
      </div>
    </div>
  </div>
{{ end }}
  <div class="row">
    <div class="col">
      <h2>Информация:</h2>
      <p>
        SteamId: <a href="https://steamcommunity.com/profiles/{{ .Person.SteamId }}/" target="_blank">{{ .Person.SteamId }}</a><br/>
      {{if eq .Person.Status "delete"}}
        Уволен: {{ .Person.GetDismissalDate }}<br/>
      {{ end }}
        В сети сегодня: {{ .Person.GetOnline 0 }}<br/>
        В сети вчера: {{ .Person.GetOnline 2 }}<br/>
        В сети за 7 дней: {{ .Person.GetOnline 1 }}<br/>
        Был в онлайне: {{ .Person.GetLastOnline true }}
      </p>
    </div>
  </div>
  <div class="row">
    <div class="col">
      <h2>Должности:</h2>
      <ul class="list-unstyled">
      {{ range .Person.Ranks }}
      {{ $nextDate := .GetNextRankDate $.Person.Vacations }}
        <li>
          <strong>{{ .GetDate }}</strong> {{ .Rank.Name }}
        {{ if $nextDate }}({{ $nextDate }}){{ end }}
        {{ if $.Header.User.IsAdmin }}
          <a href="#" data-person="{{ $.Person.Id }}" data-id="{{ .Id }}"
             data-rank="{{ .Rank.Id }}" data-date="{{ .Date }}" data-toggle="modal"
             data-target="#editPersonRankModal">✏️</a>
          &nbsp;
          <a href="#" class="btnRemovePersonalRank" data-id="{{ .Id }}"
             data-person="{{ $.Person.Id }}"
             data-name="{{ .Rank.Name }}">❌
          </a>
        {{ end }}
        </li>
      {{ end }}
      </ul>
    </div>
    <div class="col">
      <h2>Отпуска:</h2>
      <ul class="list-unstyled">
      {{ range .Person.Vacations }}
        <li>
        {{ .FormatedDate }}
        {{ if $.Header.User.IsAdmin }}
          <a href="#" data-person="{{ $.Person.Id }}" data-id="{{ .Id }}"
             data-start="{{ .Start }}" data-end="{{ .End }}" data-toggle="modal"
             data-target="#setVacationModal">✏️</a>
          &nbsp;
          <a href="#" class="btnRemovePersonalVac" data-id="{{ .Id }}"
             data-person="{{ $.Person.Id }}"
             data-date="{{ .FormatedDate }}">❌
          </a>
        {{ end }}
        </li>
      {{ end }}
      </ul>
    </div>
  </div>
{{ $lrb := len .Person.Rebukes}}
{{ if gt $lrb 0 }}
  <div class="row">
    <div class="col">
      <h2>Выговоры:</h2>
      <ul class="list-unstyled">
      {{ range .Person.Rebukes }}
        <li>
        {{ .FormatedDate }}
          <strong>{{ .Reason }}</strong> - {{ .Description }}
        {{ if $.Header.User.IsAdmin }}
          <a href="#" data-person="{{ $.Person.Id }}" data-id="{{ .Id }}"
             data-date="{{ .Date }}" data-reason="{{ .Reason }}" data-description="{{ .Description }}"
             data-toggle="modal"
             data-target="#newRebukeModal">✏️</a>
          &nbsp;
          <a href="#" class="btnRemovePersonalRebuke" data-id="{{ .Id }}"
             data-person="{{ $.Person.Id }}"
             data-reason="{{ .Reason }}">❌
          </a>
        {{ end }}
        </li>
      {{ end }}
      </ul>
    </div>
  </div>
{{ end }}
  <div class="row" style="margin-bottom: 40px">
    <div class="col">
      <div id="chartContainer" style="height: 370px; width: 100%;"></div>
      <script>
        const data = [
        {{ range .Person.Odbs }}
          {y: {{ .Time }}, toolTipContent: "{{ .GetTime }}", label: "{{ .GetDate }}"},
        {{ end }}
        ];
        window.onload = function () {
          const chart = new CanvasJS.Chart("chartContainer", {
            animationEnabled: true,
            theme: "light2",
            title: {
              text: "Онлайн за неделю"
            },
            axisY: {
              labelFormatter(e) {
                const t = e.value;
                const h = Math.round(t / 60);
                const m = Math.round(t - (h * 60));
                return `${h}h ${m}m`;
              }
            },
            data: [{
              type: "column",
              dataPoints: data.reverse()
            }]
          });
          chart.render();
        }
      </script>
    </div>
  </div>
  <div class="row">
    <div class="col">
      <ul class="list-group" style="min-width: 500px;">
      {{ range .Person.Odbs }}
        <li class="list-group-item d-flex justify-content-between align-items-center">
        {{ .GetDate }} ({{ .Last }})
          <span class="badge badge-primary badge-pill">{{.GetTime}}</span>
        </li>
      {{ end }}
      </ul>
    </div>
  </div>
</div>
{{ template "changePersonModal" . }}
{{ template "changeRankModal" . }}
{{ template "setVacationModal" . }}
{{ template "editPersonRankModal" . }}
{{ template "newRebukeModal" . }}
{{ end }}
