# Deploying Go App in Docker

## Docker Build
```bash
$ docker build -t godocker .
```

## Docker Run
```bash
$ docker run -p5000:5000 --rm --name godocker godocker
```