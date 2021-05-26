# Locust

## Installation

```bash
❯ pip3 install locust

❯ locust --version
locust 1.5.3
```

## Demo

### Example of `locustfile.py`

```Python
import time
from locust import HttpUser, task, between

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def hello_world(self):
        self.client.get("/hello")
        self.client.get("/world")

    @task(3)
    def view_items(self):
        for item_id in range(10):
            self.client.get(f"/item?id={item_id}", name="/item")
            time.sleep(1)

    def on_start(self):
        self.client.post("/login", json={"username":"foo", "password":"bar"})
```

### Start Locust

```bash
❯ locust -f locustfile.py
[2021-05-24 17:51:25,740] ip-192-168-3-3.ap-northeast-1.compute.internal/INFO/locust.main: Starting web interface at http://0.0.0.0:8089 (accepting connections from all network interfaces)
[2021-05-24 17:51:25,759] ip-192-168-3-3.ap-northeast-1.compute.internal/INFO/locust.main: Starting Locust 1.5.3
```

### Locust’s web interface

Once you’ve started Locust using one of the above command lines, you should open up a browser and point it to http://127.0.0.1:8089.
