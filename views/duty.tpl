{{ define "content" }}
<div class="container">
  <div class="row">
    <h1>Зоны дежурства</h1>
  </div>
{{ if gt .Z0C 0}}
  <div class="row">
    <div class="col">
      <h2>Не распределены - {{ .Z0C }}</h2>
    {{ block "person" buildDuty .Z0 .Header.User }} {{ end }}
    </div>
  </div>
{{ end }}
  <div class="row">
  {{ if gt .Z1C 0}}
    <div class="col">
      <h2>Атира - {{ .Z1C }}</h2>
    {{ block "person" buildDuty .Z1 .Header.User }} {{ end }}
    </div>
  {{ end }}
  {{ if gt .Z2C 0}}
    <div class="col">
      <h2>Пиргос - {{ .Z2C }}</h2>
    {{ block "person" buildDuty .Z2 .Header.User }} {{ end }}
    </div>
  {{ end }}
  {{ if gt .Z3C 0}}
    <div class="col">
      <h2>ДП11 - {{ .Z3C }}</h2>
    {{ block "person" buildDuty .Z3 .Header.User }} {{ end }}
    </div>
  {{ end }}
  {{ if gt .Z4C 0}}
    <div class="col">
      <h2>Кавала - {{ .Z4C }}</h2>
    {{ block "person" buildDuty .Z4 .Header.User }} {{ end }}
    </div>
  {{ end }}
  {{ if gt .Z5C 0}}
    <div class="col">
      <h2>София - {{ .Z5C }}</h2>
    {{ block "person" buildDuty .Z5 .Header.User }} {{ end }}
    </div>
  {{ end }}
  </div>
</div>
{{ end }}

{{ define "person" }}
<ul class="list-unstyled">
{{ range .Duty }}
  <li>
    [{{ .Personal.GetCurrentRank.Rank.Name }}] {{ .Personal.Name }}
  {{ if $.User.IsUser }}
    <select data-id="{{ .Id }}" class="change_zone">
      <option>Выберите зону</option>
      <option value="1">Атира</option>
      <option value="2">Пиргос</option>
      <option value="3">ДП11</option>
      <option value="4">Кавала</option>
      <option value="5">София</option>
    </select>
  {{ end }}
  </li>
{{ end }}
</ul>
{{ end }}
