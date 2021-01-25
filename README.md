![HYBRID LOGO](/images/hybrid.png)
# Interval ⏱️

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Build Passing](https://github.com/Hybrid-Performance-Method/interval/workflows/build/badge.svg)

Interval makes scheduling Jupyter notebook runs using Github Actions easy. It's a [Go](https://golang.org/) utility that runs [Papermill](https://github.com/nteract/papermill) behind the scenes to paramaterize notebooks for scripting jobs.

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
        output: s3://notebooks/output.ipynb
        hasDate: "true"
        parameters: parameters.yml
        requirements: "package1==1.0.0 package2==1.1.0"
        secret: ${{ secret.MY_SECRETS }}
```

## Parameters

- Interval takes a parameters input as a `parameters.yml` file (not `.yaml`) or a yaml string. See [papermill](https://github.com/nteract/papermill) for more details on notebook parameterization.

## Requirements

- It's often desirable to package notebooks with all of their requirements and install them in the first code cell.

- If additional requirements need to be stated outside of the noteboook they can be provided as a single space delimited string. 

- It's always a good idea to include versions in requirements specifications.

## Output

- Optionally specifies the file path for the artifact. the `hasdate` parameter adds a "YYYY-MM-DD" timestamp between the filename and the suffix.  

- To persist the output of the job, specificy a file storage filepath such as `s3://notebooks/output.ipynb` along with the appropriate credentials. 

## Secrets

- Secrets must be passed in a `with` field and will be available in the `INPUT_SECRET` environment variable. 

- Secrets are passed into the actions container at run time using the workflow expression syntax.

- To pass multiple secrets into a notebook environment create a new secret composed of one or more secrets separated by a github approved character. 
Remember that it's best practice to create a unique secret with least privileges for each job.

Here's a simple way handle multiple secrets in one secret variable in a notebook cell.

```python

secrets_string = os.environ["INPUT_SECRET"]

def handle_interval_secret(secrets_string: str, delimiter: str) -> List[str]:
  return secret_sring.split(delimiter)
  
secrets = handle_interval_secret(secrets_string, "_")
```

# Contributing

Contributions are Welcome!

The action is uses a simple Go tool to manipulate the Python environment. Check out the [code of conduct](CONDUCT) before contributing.

## Steps

1. Fork the interval repo

2.  Clone the repo

``` bash
$ git clone https://github.com/Hybrid-Performance-Method/interval.git`
```

3. Git checkout a branch for local development 

```bash
$ git checkout -b name-of-branch
```

4. Create a python virtual environment, install requirements and fetch any go dependencies

```bash
# create python dev environment
$ python -m venv venv
$ source venv/bin/activate
$ pip install -r requirements.txt

# get go deps
$ go mod download
```

5. `make docker` is useful for local runs and debugging

6. Open a PR and explain the feature or bug fix

# References
[Versioning Guide](https://github.com/actions/toolkit/blob/master/docs/action-versioning.md)
