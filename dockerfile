# final stage
FROM alpine
WORKDIR /app
COPY ./urlShortener /app/
ENTRYPOINT ./urlShortener