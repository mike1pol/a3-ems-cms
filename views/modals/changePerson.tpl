{{ define "changePersonModal" }}
<div class="modal fade" id="changePersonModal" tabindex="-1" role="dialog" aria-labelledby="changePersonModal"
     aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="changePersonModalLabel">Изменить {{ .Person.Name }}</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <form method="POST">
                <div class="modal-body">
                    <div class="form-group">
                        <label for="editName">Имя Фамилия</label>
                        <input name="Name" type="text" class="form-control" id="editName" value="{{ .Person.Name }}">
                    </div>
                    <div class="form-group">
                        <label for="edotSteamId">SteamID</label>
                        <input name="steamId" type="text" class="form-control" id="editSteamId" value="{{ .Person.SteamId }}">
                    </div>
                    <div class="form-group">
                        <label for="editStatus">Статус</label>
                        <select name="Status" class="form-control" id="editStatus">
                            <option value="active" {{if eq .Person.Status "active"}} selected{{ end }}>Активен</option>
                            <option value="delete" {{if eq .Person.Status "delete"}} selected{{ end }}>Уволен</option>
                        </select>
                    </div>
                    <div class="form-group{{if eq .Person.Status "active"}} hidden{{ end }}"
                         id="editDismissalDateBlock">
                        <label for="editDismissalDate">Дата увольнения</label>
                        <input name="DismissalDate" type="date" data-toggle="datepicker" class="form-control"
                               id="editDismissalDate">
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Отменить</button>
                    <button type="submit" class="btn btn-primary">Изменить</button>
                </div>
            </form>
        </div>
    </div>
</div>
{{ end }}
