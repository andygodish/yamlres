# yamlres

YAML + Resume. Profound, I know. 

Yamlres is a simple web app that presents a resume in a clean, readable format from a YAML file that follows the [JSON Resume Schema](https://jsonresume.org/schema).

## Quick Start

Pull the dev image that gets created on every push to the main branch that doesn't include a release-please release. 

Mount a local `resume.yaml` file that follows the JSON Resume Schema to `/app/config/resume.yaml` ([example](https://gist.githubusercontent.com/andygodish/1eebda44b8278be98035f787b5ee3107/raw/59e579b77d178d906314b036c60cc2f23d06458d/resume.yaml)) in the container and expose port 8080 in order to reach the UI at localhost:8080. The backend server deploys inside the same container and serves the resume file from the file system as indicated by the `-v` flag in the following docker run command:

```
docker run -it -p 8080:8080 -v $(pwd)/resume.yaml:/app/config/resume.yaml ghcr.io/andygodish/yamlres:dev
```
