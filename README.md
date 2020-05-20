# Deprecated

Now it's possible to use GitHub Action's own emails by adding the following workflow to your `.github/worflow/email.yml`

```yaml
on:
  check_suite:
    type: ['completed']

name: Email about Cirrus CI failures
jobs:
  continue:
    name: After Cirrus CI
    if: github.event.check_suite.app.name == 'Cirrus CI' && github.event.check_suite.conclusion != 'success'
    runs-on: ubuntu-latest
    steps:
      - uses: octokit/request-action@v2.x
        id: get_failed_check_run
        with:
          route: GET /repos/${{ github.repository }}/check-suites/${{ github.event.check_suite.id }}/check-runs?status=completed
          mediaType: '{"previews": ["antiope"]}'
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
      - run: |
          echo "Cirrus CI ${{ github.event.check_suite.conclusion }} on ${{ github.event.check_suite.head_branch }} branch!"
          echo "SHA ${{ github.event.check_suite.head_sha }}"
          echo $MESSAGE
          echo "##[error]See $CHECK_RUN_URL for details" && false
        env:
          CHECK_RUN_URL: ${{ fromJson(steps.get_failed_check_run.outputs.data).check_runs[0].html_url }}
```

# Send Emails with GitHub Actions

[![Build Status](https://api.cirrus-ci.com/github/cirrus-actions/email.svg)](https://cirrus-ci.com/github/cirrus-actions/email) [![](https://images.microbadger.com/badges/version/cirrusactions/email.svg)](https://microbadger.com/images/cirrusactions/email) [![](https://images.microbadger.com/badges/image/cirrusactions/email.svg)](https://microbadger.com/images/cirrusactions/email)

This is a simple GitHub action that allows to send emails when a GitHub Check Suite completes. This requires a few 
environment variables:
  * `APP_NAME` - Name of an application for which to send emails for.
  * `MAIL_FROM` - email address to send emails from.
  * `MAIL_HOST` - SMTP host to send emails to.
  * `MAIL_USERNAME` and `MAIL_PASSWORD` - username and password to authorize with the SMTP server.
  * `GITHUB_TOKEN` - is standard environment variable for GitHub actions.
  * optional `IGNORED_CONCLUSIONS` to secify conclusions to report. By default only `success` and `neutral` checks are ignored.

Now your action can look liker this in your `.github/main.workflow` workflow file:

```
action "Cirrus CI Email" {
  uses = "docker://cirrusactions/email:latest"
  env = {
    APP_NAME = "Cirrus CI"
  }
  secrets = ["GITHUB_TOKEN", "MAIL_FROM", "MAIL_HOST", "MAIL_USERNAME", "MAIL_PASSWORD"]
}
```
