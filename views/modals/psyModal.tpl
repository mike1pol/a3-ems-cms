{{ define "aePsy" }}
<div class="modal fade" id="aePsyModal" tabindex="-1" role="dialog" aria-labelledby="aePsyModal"
     aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="aePsyModalLabel">Добавить</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <form method="POST" id="aePsyModalForm">
                <div class="modal-body">
                    <div class="form-group">
                        <label for="pName">Имя</label>
                        <div class="typeahead__container">
                            <div class="typeahead__field">
                                <span class="typeahead__query">
                                    <input class="typeahead-search-person"
                                           name="name"
                                           id="pName"
                                           type="search"
                                           autocomplete="off">
                                </span>
                            </div>
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="pDate">Дата</label>
                        <input name="date" type="date" data-toggle="datepicker" class="form-control" id="pDate">
                    </div>
                    <div class="form-group">
                        <label for="pConclusion">Заключение
                        {{/*<button class="btn btn-sm btn-info" id="rLoadTemplate" type="button">Загрузить из шаблона*/}}
                        {{/*</button>*/}}
                        </label>
                        <textarea name="conclusion" rows="7" class="form-control" id="pConclusion"></textarea>
                    </div>
                    <div class="form-group">
                        <label for="pPractitioner">Врач</label>
                        <input name="practitioner" type="text" class="form-control" id="pPractitioner">
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