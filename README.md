# HYBRID Interval ⏱️
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

"A lean Jupyter notebook runner for simple and fast Github Workflows"

Interval makes running Jupyter notebooks on cron schedules easy. It uses [Papermill](https://github.com/nteract/papermill)
to paramaterize notebooks for scripting jobs.

Interval is built with the idea that simple is effective, just like a workout.

# Usage
see [action.yml](action.yml)

# Development Notes
    - paramterization
    - secrets support
    - usage docs
    - add tests

# Contributing
Contributions are Welcome!

The action is uses a simple Go tool to manipulate the Python environment. Check out the [code of conduct](CONDUCT) before contributing.

## Steps
1. Fork the interval repo
2. Clone locally git clone https://github.com/Hybrid-Performance-Method/interval.git
3. Run `make interval` to run the program against a sample notebook
4. Git checkout a branch for local development 
```bash
$ git checkout -b name-of-branch
```
5. create a python virtual environment, install requirements and fetch any go dependencies
```bash
# create python dev environment
$ python -m venv venv
$ source venv/bin/activate
$ pip install -r requirements.txt

# get go deps
$ go mod download
```
5. Use `make run` the docker image locally

# References
[Versioning Guide](https://github.com/actions/toolkit/blob/master/docs/action-versioning.md)
