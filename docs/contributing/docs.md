# Documentation

## Tooling

| Tool              | Documentation                       | Sources                           |
|-------------------|-------------------------------------|-----------------------------------|
| mkdocs            | [documentation][mkdocs]             | [Sources][mkdocs-src]             |

[mkdocs]: https://www.mkdocs.org "Mkdocs"
[mkdocs-src]: https://github.com/mkdocs/mkdocs "Mkdocs - Sources"

## Build locally

```sh
# Pre-requisite: python3, pip and virtualenv
DOCS="/tmp/extdns-docs"
mkdir "$DOCS"
virtualenv "$DOCS"
source "$DOCS/bin/activate"
pip install -r docs/scripts/requirements.txt
mkdocs serve # or mkdocs build
```
