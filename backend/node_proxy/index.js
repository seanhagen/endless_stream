"use strict";
var app = require("express")();
var http = require("http").createServer(app);
var sock = require("socket.io")(http);

var inputPB = require("./grpc/input_pb");
var services = require("./grpc/endless_grpc_pb");

var uuid = require("uuid");
var grpc = require("grpc");

var protoLoader = require("@grpc/proto-loader");

var pkgDef = protoLoader.loadSync(
  [
    "./proto/endless.proto",
    "./proto/input.proto",
    "./proto/output.proto",
    "./proto/util.proto"
  ],
  {
    keepCase: true,
    enums: String,
    arrays: true,
    objects: true,
    defaults: true,
    oneofs: true
  }
);
var pd = grpc.loadPackageDefinition(pkgDef);
var client = new pd.endless.stream.v1.Game(
  "localhost:8000",
  grpc.credentials.createInsecure()
);

app.get("/", (req, res) => {
  console.log("request to server", req);
  res.send("<h1>Hello world</h1>");
});

sock.on("connection", s => {
  console.log("incoming socket connection");

  var call = client.state();
  call.on("data", msg => {
    // console.log("got data from grpc: ", msg);
    s.emit("data", msg);
  });

  call.on("end", () => {
    console.log("grpc stream done");
    s.disconnect(true);
  });

  call.on("error", e => {
    console.error("grpc stream error: ", e);
  });

  s.on("disconnect", () => {
    console.log("user disconnected");
    call.end();
  });

  s.on("input", data => {
    console.log("received data: ", data);
    // var inp = new inputPB.Input(data);
    // console.log("input: ", inp.toObject());
    call.write(data);
  });
});

http.listen(3002, () => {
  console.log("listening on localhost:3002");
});
console.log("starting server");
