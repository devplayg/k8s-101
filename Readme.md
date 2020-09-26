# k8s  

https://kubernetes.io/ko/docs/reference/kubectl/cheatsheet/

Create

    kubectl apply -f nginx-pod.yaml

Delete

    kubectl delete -f webpod.yaml
    kubectl delete pods webpod

Connect

    kubectl exec -it webpod -c go bash
    kubectl exec -it nginx-pod bash

List

    kubectl get pods

Log

    kubectl logs webpod go
    kubectl logs webpod srv
    kubectl logs webpod web


nginx-pod.yaml

    apiVersion: v1
    kind: Pod
    metadata:
      name: webpod
    spec:
      containers:
      - name: web
        image: nginx:latest
        ports:
          - containerPort: 80
            protocol: TCP
      - name: srv
        image: ubuntu:18.04
        command: ["tail"]
        args: ["-f", "/dev/null"]
      - name: go
        image: golang:latest
        command: ["tail"]
        args: ["-f", "/dev/null"]
