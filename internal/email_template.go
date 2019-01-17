package internal

const DefaultSubjectTemplate = `
{{.CheckSuite.App.Name}} check for {{.Repo.FullName}}#{{.CheckSuite.HeadBranch}} {{.CheckSuite.Status}}
`

const DefauleEmailMarkdownTemplate = `
Check run [{{ .CheckRun.Name }}]({{ .CheckRun.HTMLURL }}) {{ .CheckRun.Status }} in {{ .Duration }}
`
