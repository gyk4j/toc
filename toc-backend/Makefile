SRC=api/swagger.yaml
SRCDIR=cmd/toc-server

all: clean validate generate build

validate: ${SRC}
	swagger validate ${SRC}
	
generate: ${SRC}
#	test -f restapi/configure_toc.go && \
#	mv restapi/configure_toc.go \
#	restapi/configure_toc.$(shell date +"%Y%m%d_%H%M%S").go
	swagger generate server -f ${SRC} -t .
	
build: ${SRCDIR}/main.go
	go build ./${SRCDIR}

clean:
	rm -f toc-server*
	rm -rf cmd
	rm -rf models
	find restapi/* -not -name "configure_toc.*" -exec /usr/bin/rm -rf {} \;
	