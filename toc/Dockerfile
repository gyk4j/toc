FROM toc-frontend AS frontend
FROM toc-backend AS backend
FROM alpine

RUN apk update && apk upgrade

WORKDIR /opt/toc/

COPY --from=backend /toc-backend/toc-backend.tar.gz /tmp/toc-backend.tar.gz
RUN tar -xzvf /tmp/toc-backend.tar.gz
COPY --from=frontend /app/toc-frontend.tar.gz /tmp/toc-frontend.tar.gz
RUN tar -C web -xzvf /tmp/toc-frontend.tar.gz

ENTRYPOINT ["/opt/toc/toc"]
