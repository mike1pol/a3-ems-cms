{{ define "content" }}
<div class="container">
    <div class="row">
        <h1>Пользователи</h1>
    </div>
    <div class="row">
        <table class="table table-hover">
            <thead>
            <tr>
                <th style="width: 50px;" class="col">ID</th>
                <th class="col">Имя</th>
                <th class="col">Steam Id</th>
                <th class="col">МЧС</th>
                <th class="col">Admin</th>
            </tr>
            </thead>
            <tbody>
            {{ range .List }}
            <tr>
                <th style="vertical-align: middle" scope="row">{{ .Id }}</th>
                <td style="vertical-align: middle"><img src="{{ .Avatar }}"> {{ .Name }}</td>
                <td style="vertical-align: middle"><a target="_blank" href="{{ .ProfileUrl }}">{{ .SteamId }}</a></td>
                <td style="text-align: center;vertical-align: middle;">
                    <a style="text-decoration:none;" href="/user/{{ .Id }}/action?action=user">{{ if .IsUser }}
                        ✅{{ else }}❌{{ end }}</a>
                </td>
                <td style="text-align: center;vertical-align: middle;">
                    <a style="text-decoration:none;" href="/user/{{ .Id }}/action?action=admin">{{ if .IsAdmin }}
                        ✅{{ else }}❌{{ end }}</a>
                </td>
            </tr>
            {{ end }}
            </tbody>
        </table>
    </div>
</div>
{{ end }}
