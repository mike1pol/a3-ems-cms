{{ define "content" }}
<div class="container">
    <div class="row">
        <h1>
            Черный список
        </h1>
    </div>
    <div class="row">
        <h2 style="margin: 20px auto;">
        {{if .IsActive }}
            Активные&nbsp;|&nbsp;<a href="?close=on">Неактивные</a>
        {{ else }}
            <a href="/blacklist">Активные</a>&nbsp;|&nbsp;Неактивные
        {{ end }}
        </h2>
    </div>
{{ if .Header.User.IsAdmin }}
    <div class="row" style="margin-bottom: 20px">
        <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#aeBlacklistModal">
            Добавить
        </button>
    </div>
{{ end }}
    <div class="row">
        <table class="table table-hover">
            <thead>
            <tr>
                <th class="col">Имя</th>
                <th class="col">Причина</th>
                <th class="col">Дата</th>
            {{ if .Header.User.IsAdmin }}
                <th class="col"></th>
            {{ end }}
            </tr>
            </thead>
            <tbody>
            {{ range .List }}
            <tr>
              <td style="vertical-align: middle">
                {{ if .IsOnline }}
                <div class="online"></div>
                {{ else }}
                <div class="offline"></div>
                {{ end }}
                {{ .Name }}
              </td>
              <td style="vertical-align: middle;min-width: 700px;">{{ .Reason }}</td>
              <td style="vertical-align: middle;min-width: 130px;">{{ .GetDate }}</td>

              {{ if $.Header.User.IsAdmin }}
              <td style="text-align: center;vertical-align: middle;min-width: 80px;">
                <a style="text-decoration:none;" data-toggle="modal" data-target="#aeBlacklistModal"
                   data-id="{{ .Id }}" data-name="{{ .Name }}" data-date="{{ .Date }}" data-reason="{{ .Reason }}"
                   data-isactive="{{ .IsActive }}" href="#">✏️</a>
                {{ if .IsActive }}
                <a style="text-decoration:none;" class="btnRemoveBlacklist" data-id="{{ .Id }}"
                   data-reason="{{ .Reason }}" data-name="{{ .Name }}" href="#">❌</a>
                {{ end }}
              </td>
              {{ end }}
            </tr>
            {{ end }}
            </tbody>
        </table>
    </div>
</div>
{{ template "aeBlacklistModal" . }}
{{ end }}
