FROM golang:alpine
WORKDIR /new_gate_server
EXPOSE 80
CMD ["sh","/new_gate_server/build.sh"]

