{% extends "layouts/main.tpl" %}

{% block content %}
        <div class="container mt-5">
            <div class="border-bottom d-flex justify-content-between align-items-center pb-2">
                <h2 class="">Site edit record page</h2>
                <div>
                    <a href="/" class="btn btn-primary me-2">Home</a>
                    <a href="/settings" class="btn btn-primary me-2">Settings</a>
                    <a href="/edit_list" class="btn btn-primary me-2">Back</a>
                </div>
            </div>
            <div class="g-4 py-5">
                <form method="post" action="/edit_list/update/{{id}}" enctype="multipart/form-data">
                    <div class="mb-3">
                        <label for="name" class="form-label">Name</label>
                        <input type="text" class="form-control" id="name" name="name" value="{{ item.Name }}">
                    </div>

                    <div class="mb-3">
                        <label for="description" class="form-label">Description</label>
                        <textarea class="form-control" id="description" name="description">{{ item.Description }}</textarea>
                    </div>

                    <div class="mb-3">
                        <label for="image" class="form-label">Image URL</label>
                        <input type="text" class="form-control" id="image" name="image" value="{{ item.Image }}">
                    </div>

                    <div class="mb-3">
                        <label for="url" class="form-label">URL</label>
                        <input type="text" class="form-control" id="url" name="url" value="{{ item.Url }}">
                    </div>

                    <button type="submit" name="save" value="1" class="btn btn-success">Save</button>
                    <button type="submit" name="save_and_edit" value="1" class="btn btn-success">Save and edit</button>
                </form>
            </div>
        </div>
{% endblock %}