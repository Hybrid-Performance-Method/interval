![HYBRID LOGO](/images/hybrid.png)
# Interval ⏱️

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Build Passing](https://github.com/Hybrid-Performance-Method/interval/workflows/build/badge.svg)

Interval makes scheduling Jupyter notebooks in github actions schedules easy. It uses [Papermill](https://github.com/nteract/papermill) behind the scenes to paramaterize notebooks for scripting jobs.

Interval is built with the idea that simple is effective, just like a workout.

# Usage
See [action.yml](action.yml)

Basic:  

```yaml
steps:
- uses: Hybrid-Performance-Method/hybrid-interval@v1
  with:
    notebook: notebook.ipynb
```

Full Workflow:

```yaml
name: nightly-job
on:
  schedule:
    - cron: '0 0 * * *'

jobs:
  run:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Run interval
      uses: Hybrid-Performance-Method/hybrid-interval@v1
      with:
        notebook: notebook.ipynb
        parameters: parameters.yml
        requirements: "package1==1.0.0 package2==1.1.0"
        secret: ${{ secret.MY_SECRETS }}
```

## Parameters

- Interval takes a parameters input as a `parameters.yml` file (not `.yaml`) or a yaml string. See [papermill](https://github.com/nteract/papermill) for more details on notebook parameterization.

## Requirements

- It's often desirable to package notebooks with all of their requirements and install them in the first code cell.

- If additional requirements need to be stated outside of the noteboook it can be provided as a space delimited string. Please state versions. 

## Secrets

- Secrets must be passed in a `with` field and will be available in the `INPUT_SECRET` environment variable. 


- Secrets are passed into the actions container at run time using the workflow expression syntax.

- To pass multiple secrets into a notebook environment create a new secret composed of one or more secrets separated by a github approved character. 
Remember that it's best practice to create a unique secret with least privileges for each job.

Here's an example of how to handle multiple secrets in one secret variable in a notebook cell.

```python
delimiter = "_"
secrets = os.ENVIRON["INPUT_SECRET"]

def handle_interval_secret(secret_string: str, delimiter: str) -> List[str]:
  return secret_sring.split(delimiter)
```

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

5. Create a python virtual environment, install requirements and fetch any go dependencies

```bash
# create python dev environment
$ python -m venv venv
$ source venv/bin/activate
$ pip install -r requirements.txt

# get go deps
$ go mod download
```

6. `make docker` is useful for local debugging

# References
[Versioning Guide](https://github.com/actions/toolkit/blob/master/docs/action-versioning.md)
