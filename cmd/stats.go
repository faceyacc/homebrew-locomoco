package cmd

import (
	"loco-moco/internals"
)

func stats(email string) {

	commits := internals.ProcessRepos(email)

	internals.PrintCommitStats(commits)

}
