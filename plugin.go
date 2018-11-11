package main

import (
	rocket "github.com/revdaalex/go-rocket-webhook"
	"strings"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Build struct {
		Tag      string
		Event    string
		Number   int
		Commit   string
		Ref      string
		Branch   string
		Author   string
		Pull     string
		Message  string
		DeployTo string
		Status   string
		Link     string
		Started  int64
		Created  int64
	}

	Config struct {
		Webhook      string
		Channel      string
		Username     string
		EmojiIcon    string
		Text         string
		Color        string
		AttTitle     string
		AttTitleLink string
		AttText      string
	}

	Job struct {
		Started int64
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

func (p Plugin) Exec() error {
	if p.Config.AttTitle == "" {
		p.Config.AttTitle = p.Repo.Name
	}
	if p.Config.AttText == "" {
		p.Config.AttText = "tag: " + p.Build.Tag
	}

	attachment := rocket.Attachment{
		Title:     p.Config.AttTitle,
		TitleLink: p.Config.AttTitleLink,
		Text:      p.Config.AttText,
		Color:     p.Config.Color,
	}

	payload := rocket.WebHookPostPayload{}
	payload.Username = p.Config.Username
	payload.Attachments = []*rocket.Attachment{&attachment}
	payload.EmojiIcon = p.Config.EmojiIcon
	payload.Text = p.Config.Text

	if p.Config.Channel != "" {
		payload.Channel = prepend("#", p.Config.Channel)
	}

	client := rocket.NewWebHook(p.Config.Webhook)

	return client.PostMessage(&payload)
}

func prepend(prefix, s string) string {
	if !strings.HasPrefix(s, prefix) {
		return prefix + s
	}

	return s
}
