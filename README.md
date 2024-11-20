# md

## Installation

```sh
$ trivy plugin install github.com/dramf/md
```

## Usage
### Using CLI
Formatted report will be saved to the file `trivy-report.md`
```sh
$ trivy image alpine --output plugin=md --format json
```
### Using GitHub Actions
```sh
jobs:
  scan:
    name: Trivy-scan
    runs-on: ubuntu-22.04
    steps:

      - name: Install Trivy
        uses: aquasecurity/setup-trivy@v0.2.2

      - name: Install plugin
        run: trivy plugin install github.com/dramf/md

      - name: Run Trivy vulnerability scanner in repo mode
        uses: aquasecurity/trivy-action@v0.29.0
        with:
          format: 'json'
          output: 'plugin=md'
          skip-setup-trivy: true

      # Show report in a job summary
      - name: Get Summary
        run: cat trivy-report.md >> $GITHUB_STEP_SUMMARY
```

How report looks like when added to a job summary:

![summary](https://github.com/user-attachments/assets/a7d52e53-0e97-4c61-b6c8-8884fbabd010)
