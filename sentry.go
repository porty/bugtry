package main

import "time"

// https://docs.sentry.io/clientdev/attributes/
type SentryRequest struct {
	EventID   string    `json:"event_id"`
	Timestamp time.Time `json:"timestamp"`
	Logger    string    `json:"logger"`
	Platform  string    `json:"platform"`
	SDK       struct {
		Name         string   `json:"name"`
		Version      string   `json:"version"`
		Integrations []string `json:"integrations"`
	} `json:"sdk"`

	// Optional below here

	Level       string                 `json:"level,omitempty"`
	Culprit     string                 `json:"culprit,omitempty"`
	ServerName  string                 `json:"server_name,omitempty"`
	Release     string                 `json:"release,omitempty"`
	Tags        map[string]string      `json:"tags,omitempty"`
	Environment string                 `json:"environment,omitempty"`
	Modules     map[string]string      `json:"modules,omitempty"`
	Extra       map[string]interface{} `json:"extra,omitempty"`
	Fingerprint []string               `json:"fingerprint,omitempty"`

	// Interfaces - https://docs.sentry.io/clientdev/interfaces/

	Exception SentryExceptions `json:"exception"`
}

type SentryExceptions struct {
	Values []SentryException `json:"values"`
}

type SentryException struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Module string `json:"module,omitempty"`
	//ThreadID  ??? `json:"thread_id,omitempty"`
	StackTrace *SentryStacktrace `json:"stacktrace,omitempty"`
}

type SentryStacktrace struct {
	Frames        []SentryFrame `json:"frames"`
	FramesOmitted []int         `json:"frames_omitted,omitempty"`
}

type SentryFrame struct {
	// must have at least one of filename/function/module
	Filename string `json:"filename,omitempty"`
	Function string `json:"function,omitempty"`
	Module   string `json:"module,omitempty"`

	// optionals
	Lineno      int                    `json:"lineno,omitempty"`
	Colno       int                    `json:"colno,omitempty"`
	AbsPath     string                 `json:"abs_path,omitempty"`
	ContextLine string                 `json:"context_line,omitempty"`
	PreContext  []string               `json:"pre_context,omitempty"`
	PostContext []string               `json:"post_context,omitempty"`
	InApp       bool                   `json:"in_app,omitempty"`
	Vars        map[string]interface{} `json:"vars,omitempty"`

	// super optionals
	// Package           string                 `json:"package,omitempty"`
	// Platform          string                 `json:"platform,omitempty"`
	// ImageAddr         string                 `json:"image_addr,omitempty"`
	// InstructionAddr   string                 `json:"instruction_addr,omitempty"`
	// SymbolAddr        string                 `json:"symbol_addr,omitempty"`
	// InstructionOffset int                    `json:"instruction_offset,omitempty"`
}
