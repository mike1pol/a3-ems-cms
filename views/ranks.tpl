{{ define "content" }}
<div class="container">
    <div class="row">
        <h1>Ранги</h1>
    </div>
    <div class="row" style="margin-bottom: 20px">
        <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#editRankModal">
            Добавить ранг
        </button>
    </div>
    <div class="row">
        <table class="table table-hover">
            <thead>
            <tr>
                <th style="width: 50px;" class="col">ID</th>
                <th class="col">Ранг</th>
                <th class="col">Дней</th>
                <th class="col">Сортировка</th>
                <th class="col" style="min-width:100px;text-align:center;">Admin</th>
            </tr>
            </thead>
            <tbody>
            {{ range .List }}
            <tr>
                <th style="vertical-align: middle" scope="row">{{ .Id }}</th>
                <td style="vertical-align: middle">{{ .Name }}</td>
                <td style="text-align: center;vertical-align: middle">{{ .Next }}</td>
                <td style="text-align: center;vertical-align: middle">{{ .Sort }}</td>
                <td style="text-align: center;vertical-align: middle;">
                    <a href="#" data-id="{{ .Id }}"
                       data-name="{{ .Name }}"
                       data-next="{{ .Next }}"
                       data-sort="{{ .Sort }}"
                       data-toggle="modal"
                       data-target="#editRankModal">✏️
                    </a>
                    &nbsp
                    <a href="#" class="btnRemoveRank" data-id="{{ .Id }}" data-name="{{ .Name }}">❌</a>
                </td>
            </tr>
            {{ end }}
            </tbody>
        </table>
    </div>
</div>
<div class="modal fade" id="editRankModal" tabindex="-1" role="dialog" aria-labelledby="editRankModal"
     aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="editRankModalLabel">Изменить ранг</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <form method="POST" id="editRankForm">
                <div class="modal-body">
                    <div class="form-group">
                        <label for="rName">Название</label>
                        <input name="name" type="text" class="form-control" id="rName">
                    </div>
                    <div class="form-group">
                        <label for="rNext">До следующего ранга дней</label>
                        <input name="next" type="number" class="form-control" id="rNext">
                    </div>
                    <div class="form-group">
                        <label for="rSort">Сортировка</label>
                        <input name="sort" type="number" class="form-control" id="rSort">
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