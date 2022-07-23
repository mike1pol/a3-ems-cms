{{ define "content" }}
<div class="container">
  <div class="row">
    <h1>База онлайна (всего: {{ len .List }})</h1>
  </div>
  <div class="row">
    <div class="col">
      <form class="form-inline" method="GET">
        <div class="form-group mb-2">
          <label class="sr-only" for="sName">Имя</label>
          <input type="text" name="n" value="{{ .Filter.Name }}" placeholder="Имя" class="form-control"
                 id="sName">
        </div>
        <div class="form-group mb-2">
          <label class="sr-only" for="sSD">с</label>
          <input name="s" value="{{ .Filter.SDate }}" type="date" placeholder="Дата с"
                 data-toggle="datepicker" class="form-control" id="sSD">
        </div>
        <div class="form-group mb-2">
          <label class="sr-only" for="sSD">по</label>
          <input name="e" type="date" value="{{ .Filter.EDate }}" placeholder="Дата с"
                 data-toggle="datepicker" class="form-control" id="sSD">
        </div>
        <div class="form-group mb-2" style="margin-left:10px;margin-right:10px;">
          <label for="sG" style="margin-right: 5px">Разложить по дням</label>
          <input name="g" type="checkbox" value="true"{{ if .Filter.UnGroup }}
                 checked {{ end }}class="form-control" id="sG">
        </div>
        <button type="submit" class="btn btn-primary mb-2">Применить</button>
      </form>
    </div>
  </div>
  <div class="row">
    <ul class="list-group" style="min-width: 500px;">
    {{range .List}}
      <li class="list-group-item d-flex justify-content-between align-items-center">
        <div>
        {{ if ne .Personal.Id 0 }}
          <img src="/static/img/med.png"/>

        <a href="/personal/{{ .Personal.Id }}">
          [{{ .Personal.GetCurrentRank.Rank.Name }}]
        {{ end }}
        {{ .Odb.Name }}
        {{ if ne .Personal.Id 0 }}
        </a>
        {{ end }}
          ({{ if ne .Odb.FirstSeen ""}}{{ .Odb.FirstSeen }} / {{ end }}{{ .Odb.LastSeen }})
        </div>
        <span class="badge badge-primary badge-pill">{{.Odb.GetTime}}</span>
      </li>
    {{end}}
    </ul>
  </div>
</div>
{{ end }}
