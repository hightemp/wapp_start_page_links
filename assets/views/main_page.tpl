{% extends "layouts/main.tpl" %}

{% block content %}
        <div class="main-page container mt-5">
            <div class="border-bottom d-flex justify-content-between align-items-center pb-2">
                <h2 class="">Links</h2>
                <div>
                    <a href="/edit_list" class="btn btn-primary me-2">Edit</a>
                    <a href="/settings" class="btn btn-primary me-2">Settings</a>
                </div>
            </div>
            <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-4 g-4 py-5">
                {% for site in config.Data.List %}
                <div class="col d-flex align-items-start">
                    {% if site.Image %}
                        <img src="{{site.Image}}" class="text-body-secondary flex-shrink-0 me-3" alt="" style="width:25px;height:auto" >
                    {% else %}
                        <svg xmlns="http://www.w3.org/2000/svg" width="25" height="25" fill="currentColor" class="bi bi-link text-body-secondary flex-shrink-0 me-3" viewBox="0 0 16 16">
                          <path d="M6.354 5.5H4a3 3 0 0 0 0 6h3a3 3 0 0 0 2.83-4H9q-.13 0-.25.031A2 2 0 0 1 7 10.5H4a2 2 0 1 1 0-4h1.535c.218-.376.495-.714.82-1z"/>
                          <path d="M9 5.5a3 3 0 0 0-2.83 4h1.098A2 2 0 0 1 9 6.5h3a2 2 0 1 1 0 4h-1.535a4 4 0 0 1-.82 1H12a3 3 0 1 0 0-6z"/>
                        </svg>
                    {% endif %}
                    <div>
                        <h3 class="fw-bold mb-0 fs-4 text-body-emphasis">
                            <a
                                href="{{ site.Url }}"
                                {% if config.Data.Settings.OpenLinksInNewWindow %}
                                target="_blank"
                                {% endif %}
                            >
                                {{ site.Name }}
                            </a>
                        </h3>
                        <p>{{ site.Description }}</p>
                    </div>
                </div>
                {% endfor %}
            </div>
        </div>
{% endblock %}