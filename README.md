# gh-scheduler

A GitHub (`gh`) CLI extension to schedule gh CLI commands to run sometime in the future.

- ⏰ Avoid setting reminders for yourself to merge a PR at the end of the day if no one else comments on it
- ⏲️ Automatically close an issue after a few hours of inactivity
- 😪 Create a pull request automatically during normal business hours that you worked on in the middle of the night (😅)

<img width="1207" alt="gh-scheduler" src="https://user-images.githubusercontent.com/17363579/180591592-818fe832-8bd5-4b75-a396-b9b0f2e4e1be.png">

## Installation

1. Install the `gh` CLI - see the [installation guide](https://github.com/cli/cli#installation)

_Installation requires a minimum version (2.0.0) of the the GitHub CLI that supports extensions._

2. Install this extension:
  
```sh
gh extension install philip-gai/gh-scheduler
```

## Usage

Run:

```sh
gh scheduler
```

Schedule something:

```sh
<gh_cli_cmd> in <duration>
```

- `<gh_cli_cmd>`: This can be any gh cli command, such as `gh pr merge <url> --auto --squash` or `gh issue close <url>`, or _anything else you might want to schedule_.
- `<duration>`: A duration string is a sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". If you leave of the duration, the job will run immediately.

Note that exiting out of the scheduler or killing the terminal session will kill the scheduler and all of its jobs. You must keep the session open for jobs to get processed.
