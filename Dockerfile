FROM debian:9.5-slim

LABEL "com.github.actions.name"="Advent of Code"
LABEL "com.github.actions.description"="Submit Answers to AoC"
LABEL "com.github.actions.icon"="mic"
LABEL "com.github.actions.color"="purple"

LABEL "repository"="http://github.com/mattdsteele-advent-of-code"
LABEL "homepage"="https://steele.blue"
LABEL "maintainer"="Matt Steele <orphum@gmail.com>"

RUN apt update && apt install ruby-full
RUN cd $GITHUB_WORKSPACE
ENTRYPOINT ["ruby" "./cmd/sync/sync.rb"]
