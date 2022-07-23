{{ define "footer" }}
<script src="/static/js/jquery.slim.min.js"></script>
<script src="/static/js/popper.min.js"></script>
<script src="/static/js/bootstrap.min.js"></script>
<script src="/static/js/canvasjs.min.js"></script>
<script src="/static/js/datepicker.min.js"></script>
<script src="/static/js/moment.js"></script>
<script src="/static/js/main.js"></script>
<script src="/static/js/html2canvas.min.js"></script>
<script src="/static/js/clipboard.min.js"></script>
<script src="/static/js/jquery.typeahead.min.js"></script>
<!-- Global site tag (gtag.js) - Google Analytics -->
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-27075540-10"></script>
<script>
    window.dataLayer = window.dataLayer || [];

    function gtag() {
        dataLayer.push(arguments);
    }

    gtag('js', new Date());

    gtag('config', 'UA-27075540-10');
</script>
{{ end }}
