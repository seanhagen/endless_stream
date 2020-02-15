#!make
include .env

NOW=$(shell date +"%s")

ifeq ($(VERSION),)
export VERSION=$(shell cat VERSION)
export BUILD=$(shell git rev-parse HEAD)
export LDFLAGSBASE=-X main.Version=${VERSION} -X main.Build=${BUILD}
endif

ifeq ($(INCPATH),)
export INCPATH=$(GOPATH)
endif

GO_BUILD_ENV:=CGO_ENABLED=0 GOOS=linux GOARCH=amd64

LDFLAGS=$(LDFLAGSBASE) -X main.Repo=${REPO}
LDFLAGBUILD=-ldflags "$(LDFLAGS) -s -w"
LDFLAGDEBUG=-ldflags "$(LDFLAGS)"

PROTO_DIR=proto

PROTOCMD=protoc
IMPORTS=-I"./${PROTO_DIR}/" -I"./vendor"
PROTOC=$(PROTOCMD) $(IMPORTS)

SERVER_SRC=backend
SERVER_TARGET=$(SERVER_SRC)/server
GO_PROTO_TARGET_DIR=backend/endless

PROTO_IN=$(shell find "$(PROTO_DIR)" -name '*.proto')
VALID_FILES_IN=$(shell find "$(PROTO_DIR)" -name '*.proto' | grep -vE "endless")
SERVICE_PROTOS=$(shell find "$(PROTO_DIR)" -name '*endless.proto')

GW_FILES=$(subst $(PROTO_DIR),$(GO_PROTO_TARGET_DIR),$(SERVICE_PROTOS:.proto=.pb.gw.go))
PB_FILES=$(subst $(PROTO_DIR),$(GO_PROTO_TARGET_DIR),$(VALID_FILES_IN:.proto=.pb.go))
SRV_PB_FILES=$(subst $(PROTO_DIR),$(GO_PROTO_TARGET_DIR),$(SERVICE_PROTOS:.proto=.pb.go))

GRPC_DESCRIPTOR=$(GO_PROTO_TARGET_DIR)/grpc_descriptor.pb

GO_PROTO_OUT=$(GW_FILES) $(PB_FILES)

test:
	@echo $(GW_FILES) $(PB_FILES) $(SRV_PB_FILES) $(GRPC_DESCRIPTOR)
	@echo $(PROTO_IN)
	@echo $(VALID_FILES_IN)
	@echo $(SERVICE_PROTOS)

$(GO_PROTO_TARGET_DIR):
	mkdir -p $@

$(GRPC_DESCRIPTOR): proto/endless.proto $(GO_PROTO_TARGET_DIR)
	@echo "Generating GRPC descriptor $@ from $<"
	@$(PROTOC) --include_imports --descriptor_set_out=$@ $<

$(SRV_PB_FILES): $(GO_PROTO_TARGET_DIR)/%.pb.go: $(PROTO_DIR)/%.proto
	@echo "Generating GRPC service $@ from $<"
	@$(PROTOC) --go_out=plugins=grpc,paths=source_relative,logtostderr=true:$(GO_PROTO_TARGET_DIR) $<

$(PB_FILES): $(GO_PROTO_TARGET_DIR)/%.pb.go: $(PROTO_DIR)/%.proto
	@echo "Generating GRPC protobuf $@ from $<"
	@$(PROTOC) --go_out=plugins=grpc,paths=source_relative,logtostderr=true:$(GO_PROTO_TARGET_DIR) $<

$(GW_FILES): $(GO_PROTO_TARGET_DIR)/%.pb.gw.go: $(PROTO_DIR)/%.proto
	@echo "Generating gateway $@ from $<"
	@$(PROTOC) --grpc-gateway_out=import_path=$(PROTO_DIR),paths=source_relative,logtostderr=true:$(GO_PROTO_TARGET_DIR) $<

pb: $(GO_PROTO_TARGET_DIR) $(PB_FILES)
gw: $(GO_PROTO_TARGET_DIR) $(GW_FILES)
srv: $(GO_PROTO_TARGET_DIR) $(SRV_PB_FILES)
desc: $(GRPC_DESCRIPTOR)

proto: pb gw srv desc
	@echo "Fixing imports"
	@find $(GO_PROTO_TARGET_DIR) -name '*.go' -type f -exec sed -i 's/org\/\/gen/org\/gen/g' {} \;

$(SERVER_TARGET):
	go build -a ${LDFLAGBUILD} -o ${SERVER_TARGET} -installsuffix cgo ./$(SERVER_SRC)

server: $(SERVER_TARGET)

all: clean proto server

clean:
	rm -rf $(SERVER_TARGET)
	rm -rf $(GO_PROTO_TARGET_DIR)
