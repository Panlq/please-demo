package(default_visibility = ["PUBLIC"])


pip_library(
    name = "flask",
    version = "2.0.2",
    # zip_safe = False, # This is because NumPy has shared object files which can't be linked to them when zipped up
    deps = [
        ":markupsafe",
        ":werkzeug",
        ":Jinja2",
        ":click",
        ":itsdangerous"
    ]
)

pip_library(
    name = "werkzeug",
    version = "2.0.3"
)

pip_library(
    name = "markupsafe",
    version = "2.1.1",
)

pip_library(
    name = "Jinja2",
    version = "3.0.3",
)

pip_library(
    name = "click",
    version = "8.0.3",
)

pip_library(
    name = "itsdangerous",
    version = "2.1.1",
)