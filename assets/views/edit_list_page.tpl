{% extends "layouts/main.tpl" %}

{% block content %}
        <div class="container mt-5">
            <div class="border-bottom d-flex justify-content-between align-items-center pb-2">
                <h2 class="">Links edit page</h2>
                <div>
                    <a href="/" class="btn btn-primary me-2">Home</a>
                    <a href="/settings" class="btn btn-primary me-2">Settings</a>
                    <a href="/edit_list/add" class="btn btn-primary me-2">Add</a>
                </div>
            </div>
            <div class="g-4 py-5">

            <table class="table table-striped table-hover">
                <tr>
                    <th>#</th>
                    <th width="30">icon</th>
                    <th>name</th>
                    <th>description</th>
                    <th width="100">actions</th>
                </tr>
                {% for site in config.Data.List %}
                <tr>
                    <td>
                        <input type="checkbox" name="select[{{ forloop.Counter-1 }}]" value="1" />
                    </td>
                    <td>
                        {% if site.Image %}
                            <img src="{{site.Image}}" class="text-body-secondary flex-shrink-0" alt="" style="width:25px;height:auto" >
                        {% else %}
                            <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-link text-body-secondary flex-shrink-0" viewBox="0 0 16 16">
                            <path d="M6.354 5.5H4a3 3 0 0 0 0 6h3a3 3 0 0 0 2.83-4H9q-.13 0-.25.031A2 2 0 0 1 7 10.5H4a2 2 0 1 1 0-4h1.535c.218-.376.495-.714.82-1z"/>
                            <path d="M9 5.5a3 3 0 0 0-2.83 4h1.098A2 2 0 0 1 9 6.5h3a2 2 0 1 1 0 4h-1.535a4 4 0 0 1-.82 1H12a3 3 0 1 0 0-6z"/>
                            </svg>
                        {% endif %}
                    </td>
                    <td>
                        {{ site.Name }}
                    </td>
                    <td>
                        {{ site.Description }}
                    </td>
                    <td class="d-flex p-2 gap-2">
                        <a href="/edit_list/edit/{{ forloop.Counter-1 }}" class="btn btn-outline-primary">Edit</a>
                        <a href="/edit_list/delete/{{ forloop.Counter-1 }}" class="btn btn-outline-danger">Delete</a>
                        <a
                            href="{{ site.Url }}"
                            {% if config.Data.Settings.OpenLinksInNewWindow %}
                            target="_blank"
                            {% endif %}
                            class="btn btn-outline-primary"
                        >Visit</a>
                    </td>
                </tr>
                {% endfor %}
            </table>

            </div>
        </div>
{% endblock %}