{% extends "layouts/main.tpl" %}

{% block content %}
        <div class="container mt-5">
            <div class="border-bottom d-flex justify-content-between align-items-center pb-2">
                <h2 class="">Error: something going wrong</h2>
                <div>
                    <a href="/" class="btn btn-primary me-2">Home</a>
                    <a href="/settings" class="btn btn-primary me-2">Settings</a>
                </div>
            </div>
            <div class="g-4 py-5">
                <a href="javascript:void(0);" onclick="history.back();">&#8592; Back to last page</a>
            </div>
        </div>
{% endblock %}