
comma := ,

ifeq ($(origin REGISTRY), undefined)
REGISTRY := harbor.alauda.cn/3rdparty
endif

ifeq ($(origin OPERATORS_INDEX), undefined)
OPERATORS_INDEX := operators-index
endif

OPS := $(foreach p,$(wildcard operators/*/*/*),$(subst /,&, $(p:operators/%=%)))

FROM_INDEX_ARG := ''
ifeq ($(FROM_INDEX), undefined)
FROM_INDEX_ARG := '--from-index $(REGISTRY)/$(OPERATORS_INDEX):$(FROM_INDEX)'
endif

.PHONY: export-bundle create-bundle.% add-index all

export-bundle:
	@./export.sh $(PKG) $(VERSION)

create-bundle.%:
	$(eval PKG := $(word 1,$(subst &, ,$*)))
	$(eval CHANNEL := $(word 2,$(subst &, ,$*)))
	$(eval TAG := $(word 3,$(subst &, ,$*)))
	@rm -rf ./bundle.Dockerfile
	@operator-sdk bundle create $(REGISTRY)/bundles/$(PKG)-bundle:$(TAG) --directory operators/$(PKG)/$(CHANNEL)/$(TAG)/manifests --package $(PKG) --channels $(CHANNEL) --default-channel $(CHANNEL)
	@docker push $(REGISTRY)/bundles/$(PKG)-bundle:$(TAG)

add-index:
	@opm index add --local-bundle --bundles $(REGISTRY)/bundles/$(PKG)-bundle:$(TAG) --tag $(REGISTRY)/$(OPERATORS_INDEX):$(INDEX_TAG) $(FROM_INDEX_ARG) -u docker

all: $(addprefix create-bundle., $(OPS))
	$(eval BUNDLES := $(shell echo $(foreach p,$(OPS),$(REGISTRY)/bundles/$(word 1,$(subst &, ,$(p)))-bundle:$(word 3,$(subst &, ,$(p)))) | tr ' ' ','))
	@opm index add --local-bundle --bundles $(BUNDLES) --tag $(REGISTRY)/$(OPERATORS_INDEX):$(INDEX_TAG) -u docker
	@docker push $(REGISTRY)/$(OPERATORS_INDEX):$(INDEX_TAG)
