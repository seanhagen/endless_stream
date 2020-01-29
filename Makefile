PROTO_DIR=proto

PROTOCMD=protoc
IMPORTS=-I"./${PROTO_DIR}/" -I"./vendor"
PROTOC=$(PROTOCMD) $(IMPORTS)

GO_TARGET_DIR=backend/grpc

PROTO_IN=$(shell find "$(PROTO_DIR)" -name '*.proto')
VALID_FILES_IN=$(shell find "$(PROTO_DIR)" -name '*.proto' | grep -vE "endless")
SERVICE_PROTOS=$(shell find "$(PROTO_DIR)" -name '*endless.proto')

GW_FILES=$(subst $(PROTO_DIR),$(GO_TARGET_DIR),$(SERVICE_PROTOS:.proto=.pb.gw.go))
PB_FILES=$(subst $(PROTO_DIR),$(GO_TARGET_DIR),$(VALID_FILES_IN:.proto=.pb.go))
SRV_PB_FILES=$(subst $(PROTO_DIR),$(GO_TARGET_DIR),$(SERVICE_PROTOS:.proto=.pb.go))

GRPC_DESCRIPTOR=$(GO_TARGET_DIR)/grpc_descriptor.pb

GO_PROTO_OUT=$(GW_FILES) $(PB_FILES)

test:
	@echo $(GW_FILES) $(PB_FILES) $(SRV_PB_FILES) $(GRPC_DESCRIPTOR)
	@echo $(PROTO_IN)
	@echo $(VALID_FILES_IN)
	@echo $(SERVICE_PROTOS)

$(GO_TARGET_DIR):
	mkdir -p $@

$(GRPC_DESCRIPTOR): proto/endless.proto $(GO_TARGET_DIR)
	@echo "Generating GRPC descriptor $@ from $<"
	@$(PROTOC) --include_imports --descriptor_set_out=$@ $<

$(SRV_PB_FILES): $(GO_TARGET_DIR)/%.pb.go: $(PROTO_DIR)/%.proto
	@echo "Generating GRPC service $@ from $<"
	@$(PROTOC) --go_out=plugins=grpc,paths=source_relative,logtostderr=true:$(GO_TARGET_DIR) $<

$(PB_FILES): $(GO_TARGET_DIR)/%.pb.go: $(PROTO_DIR)/%.proto
	@echo "Generating GRPC protobuf $@ from $<"
	@$(PROTOC) --go_out=plugins=grpc,paths=source_relative,logtostderr=true:$(GO_TARGET_DIR) $<

$(GW_FILES): $(GO_TARGET_DIR)/%.pb.gw.go: $(PROTO_DIR)/%.proto
	@echo "Generating gateway $@ from $<"
	@$(PROTOC) --grpc-gateway_out=import_path=$(PROTO_DIR),paths=source_relative,logtostderr=true:$(GO_TARGET_DIR) $<

pb: $(GO_TARGET_DIR) $(PB_FILES)
gw: $(GO_TARGET_DIR) $(GW_FILES)
srv: $(GO_TARGET_DIR) $(SRV_PB_FILES)
desc: $(GRPC_DESCRIPTOR)

proto: pb gw srv desc
	@echo "Fixing imports"
	@find backend/grpc -name '*.go' -type f -exec sed -i 's/org\/\/gen/org\/gen/g' {} \;

clean:
	rm -rf backend/grpc
