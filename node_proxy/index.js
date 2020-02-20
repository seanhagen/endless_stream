"use strict";
const app = require("express")();
const http = require("http").createServer(app);
const sock = require("socket.io")(http);
const cors = require("cors");

const uuid = require("uuid");
const grpc = require("grpc");

const protoLoader = require("@grpc/proto-loader");
const grpcPromise = require("grpc-promise");

const mpipe = require("multipipe");

console.log("setting up endless proxy");

var pkgDef = protoLoader.loadSync(
  [
    "./proto/endless.proto",
    "./proto/input.proto",
    "./proto/output.proto",
    "./proto/util.proto"
  ],
  {
    keepCase: false,
    arrays: true,
    objects: true,
    defaults: true,
    oneofs: true
  }
);
var pd = grpc.loadPackageDefinition(pkgDef);
var client = new pd.endless.stream.v1.Game(
  "game_server:8000",
  grpc.credentials.createInsecure()
);
grpcPromise.promisify(client, ["create"]);

async function createGame() {
  return new Promise(function(resolve, reject) {
    client
      .create()
      .sendMessage({})
      .then(msg => {
        resolve(msg);
      })
      .catch(err => {
        reject(err);
      });
  });
}

app.use(cors());

app.post("/create", async (req, res, next) => {
  var result = await createGame();
  console.log("created game: ", result.code);
  res.json(result);
});

sock.on("connection", s => {
  var call = client.state();
  call = mpipe(call);

  call.on("data", msg => {
    s.emit("data", msg);
  });

  call.on("end", () => {
    console.log("grpc stream done");
    s.disconnect(true);
  });

  call.on("error", err => {
    console.error("grpc stream error: ", err.details, err.code);
    s.emit("error", { details: err.details, code: err.code });
    s.disconnect(true);
  });

  s.on("disconnect", () => {
    console.log("user disconnected");
    call.end();
  });

  s.on("input", data => {
    data = fixInput(data);
    call.write(data);
  });

  s.on("error", e => {
    console.log("socket error: ", e);
  });
});

http.listen(3000, () => {
  console.log("listening on localhost:3000");
});
console.log("starting server");

const fixInput = data => {
  if (data.charSelect !== undefined) {
    data = fixCharSelect(data);
  }

  return data;
};

const fixCharSelect = data => {
  data.charSelect.choice.class = data.charSelect.choice.pb_class;
  delete data.charSelect.choice.pb_class;
  data.char_select = data.charSelect;
  delete data.CharSelect;
  return data;
};
