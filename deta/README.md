# Deta

- https://docs.deta.sh/docs/home/

## Configuring & Installing the Deta CLI

```bash
❯ curl -fsSL https://get.deta.dev/cli.sh | sh

❯ source ~/.zshrc

❯ deta version
deta v1.1.2-beta x86_64-darwin
```

## Logging in to Deta via the CLI

```bash
❯ deta login
```

## Creating Your First Micro

```bash
❯ deta new --python first_micro
Successfully created a new micro
{
        "name": "first_micro",
        "runtime": "python3.7",
        "endpoint": "https://<path>.deta.dev",
        "visor": "enabled",
        "http_auth": "disabled"
}
```

## Updating your Micro: Dependencies and Code

Create `requirements.txt` in `first_micro` directory.

```txt
flask==1.1.2
```

Edit `main.py`.

```python
from flask import Flask

app = Flask(__name__)


@app.route('/', methods=["GET"])
def hello_world():
    return "Hello World"
```

## Deploy

```bash
❯ deta deploy
Deploying...
Successfully deployed changes
Updating dependencies...
  Downloading Flask-1.1.2-py2.py3-none-any.whl (94 kB)
Collecting click>=5.1
  Downloading click-7.1.2-py2.py3-none-any.whl (82 kB)
Collecting itsdangerous>=0.24
  Downloading itsdangerous-1.1.0-py2.py3-none-any.whl (16 kB)
Collecting Jinja2>=2.10.1
  Downloading Jinja2-2.11.3-py2.py3-none-any.whl (125 kB)
Collecting MarkupSafe>=0.23
  Downloading MarkupSafe-1.1.1-cp37-cp37m-manylinux2010_x86_64.whl (33 kB)
Collecting Werkzeug>=0.15
  Downloading Werkzeug-1.0.1-py2.py3-none-any.whl (298 kB)
Installing collected packages: MarkupSafe, Werkzeug, Jinja2, itsdangerous, click, flask
Successfully installed Jinja2-2.11.3 MarkupSafe-1.1.1 Werkzeug-1.0.1 click-7.1.2 flask-1.1.2 itsdangerous-1.1.0
```

## Check

```bash
❯ deta details
{
        "name": "first_micro",
        "runtime": "python3.7",
        "endpoint": "https://a9vexz.deta.dev",
        "dependencies": [
                "flask==1.1.2"
        ],
        "visor": "enabled",
        "http_auth": "disabled"
}

❯ curl https://a9vexz.deta.dev
Hello World
```

On UI.

![image](https://user-images.githubusercontent.com/45956169/117736105-d3b61800-b231-11eb-864a-32ee51f0104d.png)
