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
    kubectl get po

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



### Replicaset

kube.yaml
    
    apiVersion: apps/v1
    kind: ReplicaSet
    metadata:
        name: replicaset-was
    spec:
        replicas: 3
        selector:
            matchLabels:
                app: app-was
        template:
            metadata:
                name: tpl-was
                labels:
                    app: app-was
            spec:
                containers:
                - name: was
                  image: devplayg/webcounter:latest
                  ports:
                  - containerPort: 80
                  
Create

    $ kubectl apply -f kube.yaml
    
    $ kubectl get po
    NAME                   READY   STATUS    RESTARTS   AGE
    replicaset-was-6bxhn   1/1     Running   0          57s
    replicaset-was-cz5f7   1/1     Running   0          57s
    replicaset-was-nbt2j   1/1     Running   0          57s
    
    $ kubectl get po --show-labels
    NAME                   READY   STATUS    RESTARTS   AGE   LABELS
    replicaset-was-6bxhn   1/1     Running   0          60s   app=app-was
    replicaset-was-cz5f7   1/1     Running   0          60s   app=app-was
    replicaset-was-nbt2j   1/1     Running   0          60s   app=app-was
    
    $kubectl get rs
    NAME             DESIRED   CURRENT   READY   AGE
    replicaset-was   3         3         3       3m20s


Delete

     kubectl delete rs replicaset-was
