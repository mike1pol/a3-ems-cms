{{ define "changeRankModal" }}
<div class="modal fade" id="changeRankModal" tabindex="-1" role="dialog" aria-labelledby="changeRankModal"
     aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="changeRankModalLabel">Изменить {{ .Person.Name }}</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <form method="POST" action="/personal/{{ .Person.Id }}/rank">
                <div class="modal-body">
                    <div class="form-group">
                        <label for="changeRank">Ранг</label>
                        <select name="Rank" class="form-control" id="changeRank">
                            <option selected value="">Выберите ранг</option>
                        {{ range .Ranks }}
                            <option value="{{ .Id }}">{{ .Name }}</option>
                        {{ end }}
                        </select>
                    </div>
                    <div class="form-group">
                        <label for="changeDate">Дата поступления на службу</label>
                        <input name="Date" type="date" data-toggle="datepicker" class="form-control" id="changeDate">
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