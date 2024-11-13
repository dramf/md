## Target `/etc/ssl/private/ssl-cert-snakeoil.key`
### Vulnerabilities
| Library | Status | Vulnerability ID | Severity | Installed Version | Fixed Version | Title | References |
|---------|--------|------------------|----------|-------------------|---------------|-------|------------|
| `libcrypto3` | fixed | CVE-2024-9143 | LOW | 3.3.2-r0 | 3.3.2-r1 | openssl: Low-level invalid GF(2^m) parameters lead to OOB memory access | https://access.redhat.com/security/cve/CVE-2024-9143, <br>https://github.com/openssl/openssl/commit/72ae83ad214d2eef262461365a1975707f862712, <br>https://github.com/openssl/openssl/commit/bc7e04d7c8d509fb78fc0e285aa948fb0da04700 |
| `libssl3` | fixed | CVE-2024-9143 | LOW | 3.3.2-r0 | 3.3.2-r1 | openssl: Low-level invalid GF(2^m) parameters lead to OOB memory access | https://access.redhat.com/security/cve/CVE-2024-9143, <br>https://github.com/openssl/openssl/commit/72ae83ad214d2eef262461365a1975707f862712, <br>https://github.com/openssl/openssl/commit/bc7e04d7c8d509fb78fc0e285aa948fb0da04700 |
### Misconfigurations
| ID | Title | Description | Severity | Message | Code |
|----|-------|-------------|----------|---------|------|
| DS001 | ':latest' tag used | When using a 'FROM' statement you should use a specific tag to avoid uncontrolled behavior when the image is updated. | MEDIUM | Specify a tag in the 'FROM' statement for image 'ansibleplaybookbundle/apb-base' <br> See https://avd.aquasec.com/misconfig/ds001 | <pre>1 FROM ansibleplaybookbundle/apb-base <br></pre> |
| DS011 | COPY with more than two arguments not ending with slash | When a COPY command has more than two arguments, the last one should end with a slash. | CRITICAL | Slash is expected at the end of COPY command argument '}}' <br> See https://avd.aquasec.com/misconfig/ds011 | <pre>7 COPY . /opt/ansible/roles/{{ role_name }} <br></pre> |
| DS026 | No HEALTHCHECK defined | You should add HEALTHCHECK instruction in your docker container images to perform the health check on running containers. | LOW | Add HEALTHCHECK instruction in your Dockerfile <br> See https://avd.aquasec.com/misconfig/ds026 |  |
### Secrets
| RuleID | Category | Severity | Title |
|--------|----------|----------|-------|
| private-key | AsymmetricPrivateKey | HIGH | Asymmetric Private Key |
