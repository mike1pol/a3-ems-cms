{{ define "content" }}
<div class="container">
  <div class="row">
    <h1>
      Серверы
      <a style="text-decoration:none;" href="/server/refresh">🔄</a>
    </h1>
  </div>
  <div class="row" style="margin-bottom: 20px">
    <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#addServerModal">
      Добавить сервер
    </button>
  </div>
  <div class="row">
    <table class="table table-hover">
      <thead>
      <tr>
        <th style="width: 50px;" class="col">ID</th>
        <th class="col">Название</th>
        <th class="col">IP:Port</th>
        <th class="col">Статус</th>
        <th class="col">Онлайн</th>
        <th class="col">Последнее обновление</th>
        <th class="col"></th>
      </tr>
      </thead>
      <tbody>
      {{ range .List }}
      <tr>
        <th style="vertical-align: middle" scope="row">{{ .Id }}</th>
        <td style="vertical-align: middle">{{ .Name }}</td>
        <td style="vertical-align: middle">{{ .Ip }}:{{ .Port }}</td>
        <td style="vertical-align: middle">{{ if .Status }} Online {{ else }} Offline {{ end }}</td>
        <td style="vertical-align: middle">{{ .Online }}/{{ .MaxPlayers }}</td>
        <td style="vertical-align: middle">{{ .GetLastUpdate }}</td>
        <td style="text-align: center;vertical-align: middle;min-width: 80px;">
          <a style="text-decoration:none;" href="/server/{{ .Id }}/action?type=refresh">🔄</a>
          <a style="text-decoration:none;" href="/server/{{ .Id }}/action?type=delete">❌</a>
        </td>
      </tr>
      {{ end }}
      </tbody>
    </table>
  </div>
</div>
<div class="modal fade" id="addServerModal" tabindex="-1" role="dialog" aria-labelledby="addServerModal"
     aria-hidden="true">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="addServerModalLabel">Добавить</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <form method="POST">
        <div class="modal-body">
          <div class="form-group">
            <label for="addName">Name</label>
            <input name="Name" type="string" class="form-control" id="addName">
          </div>
          <div class="form-group">
            <label for="addIp">IP</label>
            <input name="Ip" type="string" class="form-control" id="addIp">
          </div>
          <div class="form-group">
            <label for="addPort">Port</label>
            <input name="Port" type="string" class="form-control" id="addPort">
          </div>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-dismiss="modal">Отменить</button>
          <button type="submit" class="btn btn-primary">Добавить</button>
        </div>
      </form>
    </div>
  </div>
</div>
{{ end }}
