{{ define "content" }}
<div class="container">
    <div class="row">
        <h1>
            База справок
        </h1>
    </div>
    <div class="row" style="margin-bottom: 20px">
        <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#aeRefDBModal">
            Добавить
        </button>
    </div>
    <div class="row" style="margin-top: 20px;margin-bottom: 20px">
        <form class="form-inline" method="GET">
            <div class="form-group mb-2">
                <label class="sr-only" for="sName">Имя</label>
                <input type="text" name="name" value="{{ .Name }}" placeholder="Имя" class="form-control"
                       id="sName">
            </div>
            <div class="form-group mb-2">
                <label class="sr-only" for="sType">Тип</label>
                <select name="type" class="form-control" id="sType">
                    <option value="0" {{if eq .Type 0}} selected{{ end }}>Все</option>
                    <option value="1" {{if eq .Type 1}} selected{{ end }}>На работу</option>
                    <option value="2" {{if eq .Type 2}} selected{{ end }}>На оружие</option>
                </select>
            </div>
            <button type="submit" class="btn btn-primary mb-2">Применить</button>
        </form>
    </div>
    <div class="row">
        <div id="alertZone"></div>
        <table class="table table-hover">
            <thead>
            <tr>
                <th class="col">Имя</th>
            {{ if eq .Type 0 }}
                <th class="col">Тип справки</th>
            {{ end }}
                <th class="col">Дата</th>
                <th class="col">Заключение</th>
                <th class="col">Врач</th>
            {{ if .Header.User.IsAdmin }}
                <th class="col"></th>
            {{ end }}
            </tr>
            </thead>
            <tbody>
            {{ range .List }}
            <tr>
                <td style="vertical-align: middle;min-width: 200px;">{{ .Name }}</td>
            {{ if eq $.Type 0 }}
                <td style="vertical-align: middle;min-width: 180px;">{{ .GetType }}</td>
            {{ end }}
                <td style="vertical-align: middle;min-width: 100px;">{{ .GetDate }}</td>
                <td style="vertical-align: middle;min-width: 500px;">{{ .GetConclusion }}</td>
                <td style="vertical-align: middle;min-width: 200px;">{{ .Practitioner }}</td>
            {{ if $.Header.User.IsAdmin }}
                <td style="text-align: center;vertical-align: middle;min-width: 120px;">
                    <a style="text-decoration:none;" class="rCopyMessageForForum"
                       data-clipboard-text="{{ .GetFormMsg }}"
                       href="#">✉️️</a>
                    <a style="text-decoration:none;" data-toggle="modal" data-target="#aeRefDBModal"
                       data-id="{{ .Id }}" data-name="{{ .Name }}" data-type="{{ .Type }}" data-date="{{ .Date }}"
                       data-conclusion="{{ .Conclusion }}"
                       data-practitioner="{{ .Practitioner }}" href="#">✏️</a>
                    <a style="text-decoration:none;" class="btnRemoveRefDB" data-id="{{ .Id }}"
                       data-type="{{ .GetType }}" data-name="{{ .Name }}" href="#">❌</a>
                </td>
            {{ end }}
            </tr>
            {{ end }}
            </tbody>
        </table>
    </div>
</div>
{{ template "aeRefDB" . }}
{{ end }}