{{ define "setVacationModal" }}
<div class="modal fade" id="setVacationModal" tabindex="-1" role="dialog" aria-labelledby="setVacationModal"
     aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="setVacationModalLabel">В отпуск {{ .Person.Name }}</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <form method="POST" id="setVacationForm">
                <div class="modal-body">
                    <div class="form-group">
                        <label for="vacationStartDate">Дата начала</label>
                        <input name="Start" type="date" data-toggle="datepicker" class="form-control"
                               id="vacationStartDate">
                    </div>
                    <div class="form-group">
                        <label for="vacationEndDate">Дата окончания</label>
                        <input name="End" type="date" data-toggle="datepicker" class="form-control"
                               id="vacationEndDate">
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