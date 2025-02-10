# Word Count ([wc](https://en.wikipedia.org/wiki/Wc_(Unix))) using Golang:

Implemented the unix version of wc

Reference: https://codingchallenges.fyi/challenges/challenge-wc

#### Prerequisites:
* GoLang
* BSD/Linux Environment

#### Usage:
* `make build`
* `./gwc <options> <filename>`

#### Tasks
- [ ] Implement option --libxo
- [x] Add support for multiple files
- [x] Add test cases
- [x] Add default flags when none are passed
- [ ] Fix to handle typed input data
- [ ] Increase test coverage
- [ ] Add proper exit codes
- [ ] Handle signals: signal.Notify from system SIGTERM etc

#### Upgrades or Options to consider
* Use [pflag](github.com/spf13/pflag) instead of [argp](github.com/tdewolff/argp)?
* Use io.Read for better performance over bufio?
