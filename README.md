logs a random hash every 5 secs and appends a timestamp to it

### Run with docker

```
docker build -t log-output .
docker run log-output
```

### Run in k3s

```
kubectl create deployment log-output --image=pingust/log-output
kubectl logs -f ...
```
