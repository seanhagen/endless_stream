
PROTO_DIR="../proto"
PROTO_FILES_IN=$(shell find "$(PROTO_DIR)" -name '*.proto')

PROTOCMD=protoc
PROTOC=$(PROTOCMD) $(IMPORTS)
IMPORTS=-I${PROTO_DIR} -I"../vendor"

PROTOC_GEN_TS_PATH=./node_modules/.bin/protoc-gen-ts
JS_OUT=proto
JS_FILES=$(subst ../proto,./src/grpc,$(PROTO_FILES_IN:.proto=_pb.js))

JS_OUT_DIR=src/grpc


JS_GRPC_FILES=$(addsuffix /service_grpc_web_pb.js,$(addprefix js/,$(SERVICES)))
JS_IMPORT_FIX=M
JS_GEN_OUT=--plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" --js_out=import_style=commonjs:$(JS_OUT_DIR) --ts_out=$(JS_OUT_DIR)
JS_GEN_WEB_OUT=--plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" --grpc-web_out=import_style=commonjs,mode=grpcweb:$(JS_OUT_DIR) --ts_out=service=true:$(JS_OUT_DIR)
JS_PROTOC=protoc $(IMPORTS) $(JS_GEN_OUT) $(JS_GEN_WEB_OUT)
NOD_MODE=./node_modules

$(NODE_MOD): package.json
	npm install

$(JS_OUT_DIR):
	mkdir -p $(JS_OUT_DIR)

./src/grpc/%_pb.js: ../proto/%.proto
	@# @echo "? => $?"
	@# @echo "^ => $^"
	@# @echo "@ => $@"
	@# @echo "< => $<"

	$(eval Y=$(subst js/,,$@))
	$(eval Y=$(patsubst %/,%,$(dir $Y)))
	$(eval P=$(subst js/,,$(@:_pb.js=.proto)))
	@echo "Generating $@ from $(PROTO_DIR)/$P"
	$(PROTOCMD) $(IMPORTS) $(JS_GEN_OUT) $?
	$(PROTOCMD) $(IMPORTS) $(JS_GEN_WEB_OUT) $?


grpc-js: $(JS_OUT_DIR) $(JS_FILES) $(NODE_MOD)
	$(PROTOCMD) $(IMPORTS) $(JS_GEN_OUT) ../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/annotations.proto
	$(PROTOCMD) $(IMPORTS) $(JS_GEN_OUT) ../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/http.proto
	$(PROTOCMD) $(IMPORTS) $(JS_GEN_WEB_OUT) ../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/annotations.proto
	$(PROTOCMD) $(IMPORTS) $(JS_GEN_WEB_OUT) ../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api/http.proto
	@echo "Fixing import paths in JS"
	@find $(JS_OUT_DIR) -name '*.js' -exec sed -i '/mwitkow/d' {} \;
	@find $(JS_OUT_DIR) -name '*.js' -exec sed -i 's~\.\./\.\./\.\./\.\./\.\./~~' {} \;
	@rm -rf $(JS_OUT_DIR)/google
	@mv $(JS_OUT_DIR)/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google $(JS_OUT_DIR)/google
	@rm -rf $(JS_OUT_DIR)/github.com/
