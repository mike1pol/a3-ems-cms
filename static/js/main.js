$(document).ready(function () {
  if ($('.typeahead-search-person').length > 0) {
    $.typeahead({
      input: '.typeahead-search-person',
      minLength: 3,
      maxItem: 15,
      group: false,
      order: 'asc',
      hint: false,
      dynamic: true,
      filter: false,
      delay: 300,
      source: {
        odb: {
          data: function (query) {
            var deferred = $.Deferred()
            query = this.query;
            fetch(`/api/v1/odb/search?name=${query}`)
              .then((data) => data.json())
              .then((data) => deferred.resolve(data.Data))
            return deferred
          }
        }
      }
    });
  }
  $('[data-toggle="datepicker"]')
    .datepicker({
      language: 'ru-RU',
      format: 'yyyy-mm-dd',
      zIndex: 1050,
      autoHide: true
    });
  $("#editStatus").change(function (e) {
    if (e.target.value == "delete") {
      $('#editDismissalDateBlock').removeClass('hidden');
    } else {
      $('#editDismissalDateBlock').addClass('hidden');
    }
  });
  $('.next-rank')
    .each(function () {
      const diff = moment($(this).text(), 'DD.MM.YYYY').diff(moment().format('YYYY-MM-DD'), 'days');
      if (diff <= 0) {
        $(this).parent().addClass('rank-red');
      } else if (diff < 2) {
        $(this).parent().addClass('rank-yellow');
      }
    });
  if (document.querySelector('#report')) {
    html2canvas(document.querySelector('#report'))
      .then(canvas => {
        const dwnl = $('#dwn-canvas');
        dwnl.attr('href', canvas.toDataURL().replace('image/png', 'image/octet-stream'));
        dwnl.attr('download', 'report.png');
        dwnl.show();
      });
  }
  $('#editPersonRankModal')
    .on('show.bs.modal', function (e) {
      const button = $(e.relatedTarget);
      const id = button.data('id');
      const rank = button.data('rank');
      const person = button.data('person');
      const date = button.data('date');
      const modal = $(this);
      $('#editPersonalRankForm').attr('action', `/personal/${person}/rank/${id}`);
      modal.find('select[name="Rank"]').val(rank).change();
      modal.find('input[name="Date"]').val(date);
    });
  $('.btnRemovePersonalRank')
    .click(function (e) {
      e.preventDefault();
      const person = $(this).data('person');
      const id = $(this).data('id');
      const name = $(this).data('name');
      const cf = confirm(`Вы действительно хотите удалить ранг ${name} у пользователя`);
      if (cf) {
        fetch(`/personal/${person}/rank/${id}`, {
          method: 'DELETE',
          credentials: 'same-origin',
        })
          .then(() => location.reload())
          .catch(console.error);
      }
    })
  $('#editRankModal')
    .on('show.bs.modal', function (e) {
      const button = $(e.relatedTarget);
      const id = button.data('id');
      const name = button.data('name');
      const next = button.data('next');
      const sort = button.data('sort');
      const modal = $(this);
      $('#editRankForm').attr('action', id ? `/rank/${id}` : '/rank');
      if (id) {
        modal.find('input[name="name"]').val(name);
        modal.find('input[name="next"]').val(next);
        modal.find('input[name="sort"]').val(sort);
        modal.find('button[type="submit"]').text('Изменить');
      } else {
        modal.find('button[type="submit"]').text('Создать');
      }
    });
  $('.btnRemoveRank')
    .click(function (e) {
      e.preventDefault();
      const id = $(this).data('id');
      const name = $(this).data('name');
      const cf = confirm(`Вы действительно хотите удалить ранг ${name} (ранг будет удален у всех пользователей)`);
      if (cf) {
        fetch(`/rank/${id}`, {
          method: 'DELETE',
          credentials: 'same-origin',
        })
          .then(() => location.reload())
          .catch(console.error);
      }
    });
  $('#setVacationModal')
    .on('show.bs.modal', function (e) {
      const button = $(e.relatedTarget);
      const id = button.data('id');
      const person = button.data('person');
      const start = button.data('start');
      const end = button.data('end');
      const modal = $(this);
      $('#setVacationForm').attr('action', `/personal/${person}/vacation${id ? `/${id}` : ''}`);
      if (id) {
        modal.find('input[name="Start"]').val(start);
        modal.find('input[name="End"]').val(end);
        modal.find('button[type="submit"]').text('Изменить');
      } else {
        modal.find('button[type="submit"]').text('Создать');
      }
    });
  $('.btnRemovePersonalVac')
    .click(function (e) {
      e.preventDefault();
      const id = $(this).data('id');
      const person = $(this).data('person');
      const date = $(this).data('date');
      const cf = confirm(`Вы действительно хотите удалить отпуск ${date}`);
      if (cf) {
        fetch(`/personal/${person}/vacation/${id}`, {
          method: 'DELETE',
          credentials: 'same-origin',
        })
          .then(() => location.reload())
          .catch(console.error);
      }
    });
  $('#newRebukeModal')
    .on('show.bs.modal', function (e) {
      const button = $(e.relatedTarget);
      const id = button.data('id');
      const person = button.data('person');
      const date = button.data('date');
      const reason = button.data('reason');
      const description = button.data('description');
      const modal = $(this);
      $('#newRebukeForm').attr('action', `/personal/${person}/rebuke${id ? `/${id}` : ''}`);
      if (id) {
        modal.find('input[name="date"]').val(date);
        modal.find('input[name="reason"]').val(reason);
        modal.find('input[name="description"]').val(description);
        modal.find('button[type="submit"]').text('Изменить');
      } else {
        modal.find('button[type="submit"]').text('Создать');
      }
    });
  $('.btnRemovePersonalRebuke')
    .click(function (e) {
      e.preventDefault();
      const id = $(this).data('id');
      const person = $(this).data('person');
      const reason = $(this).data('reason');
      const cf = confirm(`Вы действительно хотите удалить выговор ${reason}`);
      if (cf) {
        fetch(`/personal/${person}/rebuke/${id}`, {
          method: 'DELETE',
          credentials: 'same-origin',
        })
          .then(() => location.reload())
          .catch(console.error);
      }
    })
  $('#aeBlacklistModal')
    .on('show.bs.modal', function (e) {
      const button = $(e.relatedTarget);
      const id = button.data('id');
      const name = button.data('name');
      const date = button.data('date');
      const reason = button.data('reason');
      const isActive = button.data('isactive');
      const modal = $(this);
      $('#aeBlacklistModalForm').attr('action', `/blacklist${id ? `/${id}` : ''}`);
      if (id) {
        modal.find('input[name="name"]').val(name);
        modal.find('input[name="date"]').val(date);
        modal.find('input[name="reason"]').val(reason);
        if (isActive) {
          modal.find('input[name="isActive"]').attr('checked', 'checked');
        } else {
          modal.find('input[name="isActive"]').removeAttr('checked');
        }
        modal.find('button[type="submit"]').text('Изменить');
        modal.find('#aeBlacklistModalLabel').text('Изменить');
      } else {
        modal.find('#aeBlacklistModalLabel').text('Добавить');
        modal.find('button[type="submit"]').text('Добавить');
      }
    });
  $('.btnRemoveBlacklist')
    .click(function (e) {
      e.preventDefault();
      const id = $(this).data('id');
      const name = $(this).data('name');
      const reason = $(this).data('reason');
      const cf = confirm(`Вы действительно хотите удалить из черного списка ${name}.\n причина: ${reason}`);
      if (cf) {
        fetch(`/blacklist/${id}`, {
          method: 'DELETE',
          credentials: 'same-origin',
        })
          .then(() => location.reload())
          .catch(console.error);
      }
    });
  var refDBTemplates = {
    work: `хирург - здоров
невропатолог - здоров
психиатр - здоров
нарколог - здоров
оториноларинголог - здоров
офтальмолог - здоров
терапевт - здоров`,
    weapon: `психиатр - здоров
офтальмолог - здоров
нарколог - здоров
терапевт - здоров`
  }
  $('#aeRefDBModal select[name="type"]').change(function (e) {
    const form = $(this).parent().parent().parent()
    const id = form.data('id')
    if (id == 0) {
      if (e.target.value == 1) {
        form.find('textarea[name="conclusion"]').val(refDBTemplates.work);
      } else if (e.target.value == 2) {
        form.find('textarea[name="conclusion"]').val(refDBTemplates.weapon);
      }
    }
  });
  $('#rLoadTemplate').click(function (e) {
    e.preventDefault();
    const form = $(this).parent().parent().parent()
    const val = form.find('select[name="type"]').val();
    if (val == 1) {
      form.find('textarea[name="conclusion"]').val(refDBTemplates.work);
    } else if (val == 2) {
      form.find('textarea[name="conclusion"]').val(refDBTemplates.weapon);
    }
  });
  $('#aeRefDBModal')
    .on('show.bs.modal', function (e) {
      const button = $(e.relatedTarget);
      const id = button.data('id');
      const name = button.data('name');
      const type = button.data('type');
      const date = button.data('date');
      const conclustion = button.data('conclusion');
      const practitioner = button.data('practitioner');
      const modal = $(this);
      $('#aeRefDBModalForm').data('id', id || 0);
      $('#aeRefDBModalForm').attr('action', `/refdb${id ? `/${id}` : ''}`);
      if (id) {
        modal.find('input[name="name"]').val(name);
        modal.find('select[name="type"]').val(type).change();
        modal.find('input[name="date"]').val(date);
        modal.find('textarea[name="conclusion"]').val(conclustion);
        modal.find('input[name="practitioner"]').val(practitioner);
        modal.find('button[type="submit"]').text('Изменить');
        modal.find('#aeRefDBModalLabel').text('Изменить');
      } else {
        modal.find('input[name="name"]').val('');
        modal.find('select[name="type"]').val('0').change();
        modal.find('input[name="date"]').val('');
        modal.find('textarea[name="conclusion"]').val('');
        modal.find('input[name="practitioner"]').val('');
        modal.find('#aeRefDBModalLabel').text('Добавить');
        modal.find('button[type="submit"]').text('Добавить');
      }
    });
  $('.btnRemoveRefDB')
    .click(function (e) {
      e.preventDefault();
      const id = $(this).data('id');
      const name = $(this).data('name');
      const type = $(this).data('type');
      const cf = confirm(`Вы действительно хотите удалить справку ${type} для ${name}`);
      if (cf) {
        fetch(`/refdb/${id}`, {
          method: 'DELETE',
          credentials: 'same-origin',
        })
          .then(() => location.reload())
          .catch(console.error);
      }
    });
  const clipboardRefMsg = new ClipboardJS('.rCopyMessageForForum')
  clipboardRefMsg.on('success', function (e) {
    const body = `
    <div class="alert alert-success alert-dismissible fade show" role="alert">
      Справка скопированна в буфер обмена
      <button type="button" class="close" data-dismiss="alert" aria-label="Close">
          <span aria-hidden="true">&times;</span>
      </button>
    </div>
    `
    $('#alertZone').append(body);
    e.clearSelection();
  });
  $('#aePsyModal')
    .on('show.bs.modal', function (e) {
      const button = $(e.relatedTarget);
      const id = button.data('id');
      const name = button.data('name');
      const date = button.data('date');
      const conclustion = button.data('conclusion');
      const practitioner = button.data('practitioner');
      const modal = $(this);
      $('#aePsyModalForm').data('id', id || 0);
      $('#aePsyModalForm').attr('action', `/psy${id ? `/${id}` : ''}`);
      if (id) {
        modal.find('input[name="name"]').val(name);
        modal.find('input[name="date"]').val(date);
        modal.find('textarea[name="conclusion"]').val(conclustion);
        modal.find('input[name="practitioner"]').val(practitioner);
        modal.find('button[type="submit"]').text('Изменить');
        modal.find('#aePsyModalLabel').text('Изменить');
      } else {
        modal.find('input[name="name"]').val('');
        modal.find('input[name="date"]').val('');
        modal.find('textarea[name="conclusion"]').val('');
        modal.find('input[name="practitioner"]').val('');
        modal.find('#aePsyModalLabel').text('Добавить');
        modal.find('button[type="submit"]').text('Добавить');
      }
    });
  $('.btnRemovePsy')
    .click(function (e) {
      e.preventDefault();
      const id = $(this).data('id');
      const name = $(this).data('name');
      const cf = confirm(`Вы действительно хотите удалить из базы психически больных ${name}`);
      if (cf) {
        fetch(`/psy/${id}`, {
          method: 'DELETE',
          credentials: 'same-origin',
        })
          .then(() => location.reload())
          .catch(console.error);
      }
    });
  const clipboardPsy = new ClipboardJS('.pCopyMessageForForum')
  clipboardPsy.on('success', function (e) {
    const body = `
    <div class="alert alert-success alert-dismissible fade show" role="alert">
      Справка скопированна в буфер обмена
      <button type="button" class="close" data-dismiss="alert" aria-label="Close">
          <span aria-hidden="true">&times;</span>
      </button>
    </div>
    `
    $('#alertZone').append(body);
    e.clearSelection();
  });

  $('.change_zone').change(function () {
    const id = $(this).data('id');
    const val = $(this).val();
    location.href = `/duty/action?id=${id}&to=${val}`;
  });
  if (document.querySelector('#changeStatus')) {
    $('#changeStatus').change(function (e) {
      const status = $(this).val();
      location.href = `/raport/${raportID}/status?status=${status}`;
    });
  }
});
