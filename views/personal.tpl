{{ define "content" }}
<div class="container">
  <div class="row">
    <h1>Личный состав</h1>
  </div>
{{ if .Header.User.IsAdmin }}
  <div class="row" style="margin-bottom: 20px">
    <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#addPersonModal">
      Добавить нового сотрудника
    </button>
  </div>
{{ end }}
  <div class="row" style="margin-bottom: 20px">
    <form class="form-inline" method="GET">
      <div class="form-group mb-2">
        <label class="sr-only" for="sName">Имя</label>
        <input type="text" name="n" value="{{ .Filter.Name }}" placeholder="Имя" class="form-control"
               id="sName">
      </div>
      <div class="form-group mb-2">
        <label class="sr-only" for="sRank">Ранг</label>
        <select name="r" class="form-control" id="sRank">
          <option value="">Выберите ранг</option>
        {{ range .RanksR }}
          <option {{if .IsEqualName $.Filter.Rank }} selected{{ end }}>{{ .Name }}</option>
        {{ end }}
        </select>
      </div>
      <div class="form-group mb-2">
        <label class="sr-only" for="sStatus">Статус</label>
        <select name="s" class="form-control" id="sStatus">
          <option value="active" {{if eq .Filter.Status "active"}} selected{{ end }}>Активен</option>
          <option value="delete" {{if eq .Filter.Status "delete"}} selected{{ end }}>Уволен</option>
        </select>
      </div>
      <button type="submit" class="btn btn-primary mb-2">Применить</button>
    </form>
  </div>
{{ range .List }}
{{ block "list" . }} {{ end }}
{{ end }}

</div>

{{ template "addPersonModal" . }}
{{ end }}


{{ define "list" }}
{{ $mLen := len .List }}
{{ if gt $mLen 0 }}
<div class="row">
  <h2>{{ .Name }}{{ if ne .Name "Министр" }} ({{ $mLen }}){{ end }}</h2>
</div>
<div class="row d-flex">
{{ range .List }}
  <div class="card" style="width: 22rem; margin: 10px;">
    <div class="card-body">
      <div class="card-title"><a href="/personal/{{ .Id }}" class="card-link">{{ .Name }}</a></div>
      <h6 class="card-subtitle mb-2 text-muted">{{ .GetCurrentRankName }} ({{ .GetStatus }})</h6>
      <div class="card-text">
        <div>SteamId: <a href="https://steamcommunity.com/profiles/{{ .SteamId }}/" target="_blank">{{ .SteamId }}</a></div>
        <div>Дата назначения: {{ .GetCurrentRankDate }}</div>
      {{ if .GetNextRankDate }}
      {{ if ne .GetCurrentRank.Rank.Id 4 }}
        <div>Дата переквалификации: <span class="next-rank">{{ .GetNextRankDate }}</span></div>
      {{ end }}
      {{ end }}
        <div>В сети сегодня: {{ .GetOnline 0 }}</div>
        <div>В сети вчера: {{ .GetOnline 2 }}</div>
        <div>В сети за 7 дней: {{ .GetOnline 1 }}</div>
        <div>Был в онлайне: {{ .GetLastOnline false }}</div>
      </div>
    </div>
  </div>
{{ end }}
</div>
{{ end }}
{{ end }}
