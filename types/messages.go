package types

import "github.com/charmbracelet/bubbles/table"

type DeleteMsg struct {
	Err error
}

type ListReleasesMsg struct {
	Content []table.Row
	Err     error
}
type HistoryMsg struct {
	Content []table.Row
	Err     error
}

type RollbackMsg struct {
	Err error
}

type UpgradeMsg struct {
	Err error
}

type NotesMsg struct {
	Err     error
	Content string
}

type MetadataMsg struct {
	Err     error
	Content string
}

type HooksMsg struct {
	Err     error
	Content string
}

type ValuesMsg struct {
	Err     error
	Content string
}

type ManifestMsg struct {
	Err     error
	Content string
}

type RemoveMsg struct {
	Err error
}

type ListRepoMsg struct {
	Content []table.Row
	Err     error
}

type PackagesMsg struct {
	Content []table.Row
	Err     error
}

type PackageVersionsMsg struct {
	Content []table.Row
	Err     error
}

type InstallMsg struct {
	Err error
}