FROM debian:9.5-slim

RUN apt-get update && apt-get install -y ruby-full
CMD cd $GITHUB_WORKSPACE && ruby ./cmd/sync/sync.rb
