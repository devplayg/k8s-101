Command
---

Create and build dockerfile

    docker build --no-cache --tag was:0.1 .
    docker run -itd --rm --name was -p 80:80 was:0.1
    docker stop was && docker rmi was:0.1

Loggings

    docker logs -f was
