# How to reproduce

- Step 1: prerequisites
make sure you have <kind, openssl> installed locally

```
kind create cluster
```

- Step 2: generate TLS certificate
Generate the certificates for TLS, a sample script is included. Please feel free to modify it and test.

# Add head Pod as IP.2

```
sh gencert.sh
```

- Step 3: replace the secret.yaml (Do not do this. We should not update the format)
copy the content from the generated certs, use <https://lingojam.com/TexttoOneLine> to convert to one line, then paste to the secret.yaml file

- Step 4: build Ray cluster locally using Kind

```
kubectl apply -f secret.yaml
kubectl apply -f head.yaml
kubectl apply -f worker.yaml
```

- Step 5: check logs after the ray cluster finishes building

```
kubectl logs <head pod>
```
