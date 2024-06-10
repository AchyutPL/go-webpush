docker build --tag webpush-go --platform linux/x86_64 .

docker run -d --name webpush-go -p 8080:8080 webpush-go