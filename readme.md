GOOS=linux go build -o runFile
==> Tạo file test-docker build với môi trường là linux
==> không để GOOS=linux  là build với môi trường bình thường
==> chạy project : ./runFile
==> runFile : tên file build

docker build -t test-docker .
==> chạy dockerfile pull về

docker run test-docker
==> run project