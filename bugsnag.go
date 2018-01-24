package main

import "time"

type Something struct {
	APIKey   string          `json:"apiKey"`
	Notifier BugsnagNotifier `json:"notifier"`
	Events   []BugsnagEvent  `json:"events"`
}

type BugsnagNotifier struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	URL     string `json:"url"`
}

type BugsnagEvent struct {
	PayloadVersion string                 `json:"payloadVersion"`
	Exceptions     []BugsnagException     `json:"exceptions"`
	Breadcrumbs    []BugsnagBreadcrumb    `json:"breadcrumbs"`
	Request        BugsnagRequest         `json:"request"`
	Threads        []BugsnagThread        `json:"threads"`
	Context        string                 `json:"context"`
	GroupingHash   string                 `json:"groupingHash"`
	Unhandled      bool                   `json:"unhandled"`
	Severity       string                 `json:"severity"`
	SeverityReason BugsnagSeverityReason  `json:"severityReason"`
	User           BugsnagUser            `json:"user"`
	App            BugsnagApp             `json:"app"`
	Device         BugsnagDevice          `json:"device"`
	Session        BugsnagSession         `json:"session"`
	Metadata       map[string]interface{} `json:"metadata"`
}

type BugsnagException struct {
	ErrorClass string              `json:"errorClass"`
	Message    string              `json:"message"`
	StackTrace []BugsnagStacktrace `json:"stacktrace"`
	Type       string              `json:"type"`
}

type BugsnagStacktrace struct {
	File         string            `json:"file"`
	LineNumber   int               `json:"lineNumber"`
	ColumnNumber int               `json:"columnNumber"`
	Method       string            `json:"method"`
	InProject    bool              `json:"inProject"`
	Code         map[string]string `json:"code"`
}

type BugsnagBreadcrumb struct {
	Timestamp time.Time         `json:"timestamp"`
	Name      string            `json:"name"`
	Type      string            `json:"type"`
	Metadata  map[string]string `json:"metadata"`
}

type BugsnagRequest struct {
	ClientIP   string            `json:"clientIp"`
	Headers    map[string]string `json:"headers"`
	HTTPMethod string            `json:"httpMethod"`
	URL        string            `json:"url"`
	Referer    string            `json:"referer"`
}

type BugsnagThread struct {
}

type BugsnagSeverityReason struct {
}

type BugsnagUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BugsnagApp struct {
}

type BugsnagDevice struct {
}

type BugsnagSession struct {
	ID        string    `json:"id"`
	StartedAt time.Time `json:"startedAt"`
}
