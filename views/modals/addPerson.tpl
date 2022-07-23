{{ define "addPersonModal" }}
<div class="modal fade" id="addPersonModal" tabindex="-1" role="dialog" aria-labelledby="addPersonModal"
     aria-hidden="true">
    <div class="modal-dialog" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="addPersonModalLabel">Добавить</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <form method="POST">
                <div class="modal-body">
                    <div class="form-group">
                        <label for="addName">Имя Фамилия</label>
                        <div class="typeahead__container">
                            <div class="typeahead__field">
                                <span class="typeahead__query">
                                    <input class="typeahead-search-person"
                                           name="Name"
                                           id="addName"
                                           type="search"
                                           autocomplete="off">
                                </span>
                            </div>
                        </div>
                    {{/*<input name="Name" type="text" class="form-control" id="addName">*/}}
                    </div>
                    <div class="form-group">
                        <label for="steamId">SteamID</label>
                        <input name="steamId" type="text" class="form-control" id="steamId">
                    </div>
                    <div class="form-group">
                        <label for="addRank">Ранг</label>
                        <select name="Rank" class="form-control" id="addRank">
                            <option selected value="">Выберите ранг</option>
                        {{ range .Ranks }}
                            <option value="{{ .Id }}">{{ .Name }}</option>
                        {{ end }}
                        </select>
                    </div>
                    <div class="form-group">
                        <label for="addDate">Дата поступления на службу</label>
                        <input name="Date" type="date" data-toggle="datepicker" class="form-control" id="addDate">
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
