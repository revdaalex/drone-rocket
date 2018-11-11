package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "1.0.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "rocket plugin"
	app.Usage = "rocket plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "webhook",
			Usage:  "rocket webhook url",
			EnvVar: "ROCKET_WEBHOOK,PLUGIN_WEBHOOK",
		},
		cli.StringFlag{
			Name:   "channel",
			Usage:  "rocket channel",
			EnvVar: "PLUGIN_CHANNEL",
		},
		cli.StringFlag{
			Name:   "username",
			Usage:  "rocket username",
			EnvVar: "PLUGIN_USERNAME",
		},
		cli.StringFlag{
			Name:   "text",
			Usage:  "rocket text",
			EnvVar: "PLUGIN_TEXT",
		},
		cli.StringFlag{
			Name:   "color",
			Usage:  "rocket color",
			EnvVar: "PLUGIN_COLOR",
		},
		cli.StringFlag{
			Name:   "emoji",
			Usage:  "rocket emoji",
			EnvVar: "PLUGIN_EMOJI",
		},
		cli.StringFlag{
			Name:   "att.title",
			Usage:  "rocket att title",
			EnvVar: "PLUGIN_ATT_TITLE",
		},
		cli.StringFlag{
			Name:   "att.title.link",
			Usage:  "rocket att title link",
			EnvVar: "PLUGIN_ATT_TITLE_LINK",
		},
		cli.StringFlag{
			Name:   "att.text",
			Usage:  "rocket att text",
			EnvVar: "PLUGIN_ATT_TEXT",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
			Value:  "00000000",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.pull",
			Usage:  "git pull request",
			EnvVar: "DRONE_PULL_REQUEST",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "build.deployTo",
			Usage:  "environment deployed to",
			EnvVar: "DRONE_DEPLOY_TO",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Tag:      c.String("build.tag"),
			Number:   c.Int("build.number"),
			Event:    c.String("build.event"),
			Status:   c.String("build.status"),
			Commit:   c.String("commit.sha"),
			Ref:      c.String("commit.ref"),
			Branch:   c.String("commit.branch"),
			Author:   c.String("commit.author"),
			Pull:     c.String("commit.pull"),
			Message:  c.String("commit.message"),
			DeployTo: c.String("build.deployTo"),
			Link:     c.String("build.link"),
			Started:  c.Int64("build.started"),
			Created:  c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			Webhook:      c.String("webhook"),
			Channel:      c.String("channel"),
			Username:     c.String("username"),
			EmojiIcon:    c.String("icon.emoji"),
			Text:         c.String("text"),
			Color:        c.String("color"),
			AttTitle:     c.String("att.title"),
			AttTitleLink: c.String("att.title.link"),
			AttText:      c.String("att.text"),
		},
	}

	return plugin.Exec()
}
