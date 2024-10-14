FROM toc-toc-frontend AS frontend
FROM toc-toc-backend AS backend
FROM alpine

RUN apk update && apk upgrade

WORKDIR /opt/toc/

COPY --from=backend /toc-backend/toc-backend.tar.gz /tmp/toc-backend.tar.gz
RUN tar -xzvf /tmp/toc-backend.tar.gz
COPY --from=frontend /app/toc-frontend.tar.gz /tmp/toc-frontend.tar.gz
RUN tar -C web -xzvf /tmp/toc-frontend.tar.gz

ENV HOST=0.0.0.0
ENV PORT=80

ENTRYPOINT ["/opt/toc/toc"]
CMD ["--host", "0.0.0.0", "--port", "80"]
