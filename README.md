# Advent of Code

![shield](https://img.shields.io/endpoint?url=https%3A%2F%2Fus-central1-sandbox-254315.cloudfunctions.net%2FShields)

## CI/CD Tooling

In the `cmd/` folder are the following utilities:

`setup/` - Preps a day by downloading the input

`submit/` - Submits an answer

`sync/` - Identifies any answers not submitted successfully, and submits each in turn

### Automation

* Submits entries to AOC via https://github.com/mattdsteele/aoc-submit-action
* Stars via https://github.com/mattdsteele/aoc-stars-shield
