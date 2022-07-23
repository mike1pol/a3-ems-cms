{{ define "aeBlacklistModal" }}
<div class="modal fade" id="aeBlacklistModal" tabindex="-1" role="dialog" aria-labelledby="aeBlacklistModal"
     aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="aeBlacklistModalLabel">Добавить</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <form method="POST" id="aeBlacklistModalForm">
                <div class="modal-body">
                    <div class="form-group">
                        <label for="bName">Имя</label>
                        <div class="typeahead__container">
                            <div class="typeahead__field">
                                <span class="typeahead__query">
                                    <input class="typeahead-search-person"
                                           name="name"
                                           id="bName"
                                           type="search"
                                           autocomplete="off">
                                </span>
                            </div>
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="bDate">Дата</label>
                        <input name="date" type="date" data-toggle="datepicker" class="form-control" id="bDate">
                    </div>
                    <div class="form-group">
                        <label for="bReason">Причина</label>
                        <input name="reason" type="text" class="form-control" id="bReason">
                    </div>
                    <div class="form-group form-check">
                        <input type="checkbox" class="form-check-input" name="isActive" id="bActive" checked="checked">
                        <label class="form-check-label" for="bActive">Активен?</label>
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