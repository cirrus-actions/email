package internal

const DefaultSubjectTemplate = `
{{ CheckSuite.App.Name }} check for {{ Repository.FullName }}#{{ CheckSuite.HeadBranch }} {{ CheckSuite.Status }}
`

const DefauleEmailMarkdownTemplate = `
`
