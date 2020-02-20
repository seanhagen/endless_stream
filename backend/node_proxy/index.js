"use strict";
const app = require("express")();
const http = require("http").createServer(app);
const sock = require("socket.io")(http);
const cors = require('cors');

app.use(cors())

const uuid = require("uuid");
const grpc = require("grpc");

const protoLoader = require("@grpc/proto-loader");
const grpcPromise = require('grpc-promise');

var pkgDef = protoLoader.loadSync(
  [
    "./proto/endless.proto",
    "./proto/input.proto",
    "./proto/output.proto",
    "./proto/util.proto"
  ],
  {
    keepCase: true,
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
grpcPromise.promisify(client, ['create']);

async function createGame() {
  return new Promise(function(resolve, reject) {
    client.create()
      .sendMessage({})
    .then(msg => { resolve(msg) })
    .catch(err => { reject(err)})
  })
}

app.post("/create", async (req, res, next) => {
  // client.create({}, (err, msg) => {
  //   if (err) {
  //     console.log("error creating game: ", err)
  //     res.send({error: err});
  //   } else {
  //     console.log("game created: ", msg)
  //     res.send({result: msg});
  //   }
  // });
  // client.create()
  //   .sendMessage({})
  //   .then(msg => {
  //     console.log("got response: ", msg)
  //     res.send(msg)
  //   })
  //   .catch(err => {console.error("error creating game: ", err)})
  console.log("creating game")
  console.log("next: ", next)
  var result = await createGame()
  console.log("result: ", result);
  //res.end(JSON.stringify(result));
  res.json(result);
  console.log("json sent");
});

sock.on("connection", s => {
  console.log("incoming socket connection");

  var call = client.state();
  call.on("data", msg => {
    s.emit("data", msg);
    if (msg.tick === undefined){
      console.log("msg: ", msg)
    }
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
    call.write(data);
  });
});

http.listen(3002, () => {
  console.log("listening on localhost:3002");
});
console.log("starting server");
