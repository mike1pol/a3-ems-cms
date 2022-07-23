{{ define "newRebukeModal" }}
<div class="modal fade" id="newRebukeModal" tabindex="-1" role="dialog" aria-labelledby="newRebukeModal"
     aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title">Выговор для {{ .Person.Name }}</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <form method="POST" id="newRebukeForm" action="/personal/{{ .Person.Id }}/rebuke">
                <div class="modal-body">
                    <div class="form-group">
                        <label for="rebukeDate">Дата выговора</label>
                        <input name="date" type="date" data-toggle="datepicker" class="form-control" id="rebukeDate">
                    </div>
                    <div class="form-group">
                        <label for="rebukeReason">Причина</label>
                        <input name="reason" type="text" class="form-control" id="rebukeReason">
                    </div>
                    <div class="form-group">
                        <label for="rebukeDescription">Описание</label>
                        <input name="description" type="text" class="form-control" id="rebukeDescription">
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