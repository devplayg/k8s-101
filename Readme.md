# k8s  
  
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




  kubectl delete pods webpod
  kubectl delete webpod
  kubectl describe
  kubectl describe pod webpod
  kubectl describe pod webpod > kube.txt
  kubectl describe pod/nginx-pod -n default
  kubectl describe pods nginx-pod
  kubectl describe pods nginx-pod2
