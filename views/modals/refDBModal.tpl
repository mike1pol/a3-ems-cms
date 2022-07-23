{{ define "aeRefDB" }}
<div class="modal fade" id="aeRefDBModal" tabindex="-1" role="dialog" aria-labelledby="aeRefDBModal"
     aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="aeRefDBModalLabel">Добавить</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <form method="POST" id="aeRefDBModalForm">
                <div class="modal-body">
                    <div class="form-group">
                        <label for="rName">Имя</label>
                        <div class="typeahead__container">
                            <div class="typeahead__field">
                                <span class="typeahead__query">
                                    <input class="typeahead-search-person"
                                           name="name"
                                           id="rName"
                                           type="search"
                                           autocomplete="off">
                                </span>
                            </div>
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="rType">Тип справки</label>
                        <select name="type" class="form-control" id="rType">
                            <option selected value="">Тип справки</option>
                            <option value="1">На трудоустройство</option>
                            <option value="2">На оружие</option>
                        </select>
                    </div>
                    <div class="form-group">
                        <label for="rDate">Дата</label>
                        <input name="date" type="date" data-toggle="datepicker" class="form-control" id="rDate">
                    </div>
                    <div class="form-group">
                        <label for="rConclusion">Заключение
                            <button class="btn btn-sm btn-info" id="rLoadTemplate" type="button">Загрузить из шаблона
                            </button>
                        </label>
                        <textarea name="conclusion" rows="7" class="form-control" id="rConclusion"></textarea>
                    </div>
                    <div class="form-group">
                        <label for="rPractitioner">Врач</label>
                        <input name="practitioner" type="text" class="form-control" id="rPractitioner">
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