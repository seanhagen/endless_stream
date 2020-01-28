PROTO_DIR=proto

PROTOCMD=protoc
IMPORTS=-I"./${PROTO_DIR}/" -I"./vendor"
PROTOC=$(PROTOCMD) $(IMPORTS)

GO_TARGET_DIR=backend/grpc

PROTO_IN=proto/endless.proto
GW_FILES=$(subst proto,backend/grpc,$(PROTO_IN:.proto=.pb.gw.go))
PB_FILES=$(subst proto,backend/grpc,$(PROTO_IN:.proto=.pb.go))
GRPC_DESCRIPTOR=backend/grpc/grpc_descriptor.pb

GO_PROTO_OUT=$(GW_FILES) $(PB_FILES)

test:
	echo $(GW_FILES) $(PB_FILES) $(GRPC_DESCRIPTOR)

$(GO_TARGET_DIR):
	mkdir -p $@

$(GRPC_DESCRIPTOR): proto/endless.proto $(GO_TARGET_DIR)
	@echo "Generating GRPC descriptor $@ from $<"
	$(PROTOC) --include_imports --descriptor_set_out=$@ $<
	ls backend/grpc

$(PB_FILES): proto/endless.proto $(GO_TARGET_DIR)
	@echo "Generating GRPC protobuf $@ from $<"
	$(PROTOC) --go_out=plugins=grpc,paths=source_relative,logtostderr=true:$(GO_TARGET_DIR) $<
	ls backend/grpc

$(GW_FILES): proto/endless.proto $(GO_TARGET_DIR)
	@echo "Generating gateway $@ from $<"
	$(PROTOC) --grpc-gateway_out=import_path=$(PROTO_DIR),paths=source_relative,logtostderr=true:$(GO_TARGET_DIR) $<
	ls backend/grpc


pb: $(PB_FILES)
gw: $(GW_FILES)
desc: $(GRPC_DESCRIPTOR)

proto: pb gw desc

clean:
	rm -rf backend/grpc
