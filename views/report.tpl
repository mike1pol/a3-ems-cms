{{ define "content" }}
<div class="container">
    <div class="row">
        <h1>Отчет с {{ .Filter.GetHuman 0 }} по {{ .Filter.GetHuman 1 }} <a style="display: none" id="dwn-canvas">Скачать</a>
        </h1>
    </div>
    <div class="row">
        <form class="form-inline" method="GET">
            <div class="form-group mb-2">
                <label class="sr-only" for="sSD">с</label>
                <input name="start" value="{{ .Filter.GetISO 0 }}" type="date" placeholder="Дата с"
                       data-toggle="datepicker" class="form-control" id="sSD">
            </div>
            <div class="form-group mb-2">
                <label class="sr-only" for="sSD">по</label>
                <input name="end" type="date" value="{{ .Filter.GetISO 1 }}" placeholder="Дата с"
                       data-toggle="datepicker" class="form-control" id="sSD">
            </div>
            <div class="form-group mb-2">
                <label class="sr-only" for="sName">Имя</label>
                <input type="text" name="n" value="{{ .Filter.Name }}" placeholder="Имя" class="form-control"
                       id="sName">
            </div>
            <div class="form-group mb-2">
                <label class="sr-only" for="sRank">Ранг</label>
                <select name="r" class="form-control" id="sRank">
                    <option value="">Выберите ранг</option>
                {{ range .Ranks }}
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

    <div class="row">
        <table class="table table-hover table-bordered" id="report">
            <thead>
            <tr>
                <th colspan="{{ .Ch }}" style="text-align: center">
                    Отчет с {{ .Filter.GetHuman 0 }} по {{ .Filter.GetHuman 1 }}
                </th>
            </tr>
            <tr>
                <th style="width: 50px;" class="col">#</th>
                <th class="col">Должность / Имя</th>
            {{ range .Days }}
                <th class="col">{{ . }}</th>
            {{ end }}
            </tr>
            </thead>
            <tbody>
            {{ range $index, $s := .List }}
            <tr>
                <th style="vertical-align: middle" scope="row">{{ inc $index }}</th>
                <td style="vertical-align: middle">[{{ $s.GetCurrentRank.Rank.Name }}] {{ $s.Name }}</td>
            {{ range $s.Reports }}
                <td style="vertical-align:middle;background-color:{{ .Color }};text-align:center;font-size: 10px;">
                {{ .Time }}
                </td>
            {{ end }}
            </tr>
            {{ end }}
            </tbody>
        </table>
    </div>
</div>
{{ end }}
