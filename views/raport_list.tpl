{{ define "content" }}
<div class="container">
    <div class="row">
        <h1>
            Список рапортов
        </h1>
    </div>
  <div class="row" style="margin-bottom: 20px">
    <form class="form-inline" method="GET">
      <div class="form-group mb-2">
        <label class="sr-only" for="sStatus">Статус</label>
        <select name="s" class="form-control" id="sStatus">
          <option value="0" {{if eq .Filter.Status 0}} selected{{ end }}>Все</option>
          <option value="1" {{if eq .Filter.Status 1}} selected{{ end }}>Новый</option>
          <option value="2" {{if eq .Filter.Status 2}} selected{{ end }}>В работе</option>
          <option value="3" {{if eq .Filter.Status 3}} selected{{ end }}>Закрыт</option>
        </select>
      </div>
      <button type="submit" class="btn btn-primary mb-2">Применить</button>
    </form>
  </div>
    <div class="row">
        <table class="table table-hover">
            <thead>
            <tr>
                <th style="width:50px;" class="col">#</th>
                <th class="col">Тема</th>
                <th style="min-width:200px;" class="col">На кого</th>
                <th style="min-width:200px;" class="col">Отправитель</th>
                <th class="col">Дата</th>
                <th class="col">Статус</th>
                <th class="col">Комментариев</th>
            </tr>
            </thead>
            <tbody>
              {{ range .List }}
              <tr>
                <th scope="row"><a href="/raport/{{ .ID }}">{{ .ID }}</a></th>
                <td>{{ .Subject }}</td>
                <td>{{ if .To.Id }} {{ .To.Name }} {{ else }} Не выбрано  {{ end }}</td>
                <td>{{ .From.Name }}</td>
                <td>{{ .GetDate }}</td>
                <td>{{ .GetStatus }}</td>
                <td>{{ .GetCountComments }}</td>
              </tr>
              {{ end }}
            </tbody>
        </table>
    </div>
</div>
{{ end }}
