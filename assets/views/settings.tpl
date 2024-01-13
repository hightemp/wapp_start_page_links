{% extends "layouts/main.tpl" %}

{% block content %}
    <div class="container mt-5">
        <div class="border-bottom d-flex justify-content-between align-items-center pb-2">
            <h2 class="">Settings page</h2>
            <div>
                <a href="/" class="btn btn-primary me-2">Home</a>
                <a href="javascript:void(0);" class="btn btn-primary me-2" onclick="history.back();">Back</a>
            </div>
        </div>
        <div class="g-4 py-5">
            <form method="post" action="/settings/update" enctype="multipart/form-data">
                <div class="mb-3">
                    <label for="name" class="form-label">Theme</label>
                    <select class="form-control" id="theme" name="theme">
                        {% for theme in config.Data.Settings.Themes %}
                            <option {% if theme.Name==config.Data.Settings.Theme %}selected{% endif %} value="{{theme.Name}}">{{theme.Name}}</option>
                        {% endfor %}
                    </select>
                </div>
                <div class="mb-3">
                    <input
                        class="form-check-input"
                        type="checkbox"
                        name="open_links_in_new_window"
                        value="1"
                        id="open_new_window"
                        {%if config.Data.Settings.OpenLinksInNewWindow %}checked{%endif%}
                    >
                    <label class="form-check-label" for="open_new_window">
                        Open links in new window
                    </label>
                </div>

                <button type="submit" name="save" value="1" class="btn btn-success">Save</button>
            </form>
        </div>
    </div>
{% endblock %}