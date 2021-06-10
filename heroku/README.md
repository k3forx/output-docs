# Heroku

## Demo

- https://devcenter.heroku.com/articles/container-registry-and-runtime

```bash
❯ heroku container:login
Login Succeeded
```

### Update submodule

```bash
git submodule update --init
```

### Create app

```bash
❯ cd alpinehelloworld

❯ heroku create
```

### Build and push

```bash
❯ heroku container:push web
```

### Make a release

```bash
❯ heroku container:release web
Releasing images web to afternoon-hamlet-44764... done
```

### Open

```bash
❯ heroku open
```

<img width="595" alt="貼り付けた画像_2021_06_10_16_50" src="https://user-images.githubusercontent.com/45956169/121486308-fcd1f000-ca0b-11eb-89c2-be3b8a4e622a.png">

### Stop the application

```bash
❯ heroku ps:scale web=0
Scaling dynos... done, now running web at 0:Free
```
