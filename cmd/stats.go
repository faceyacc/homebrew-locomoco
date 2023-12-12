package cmd

import (
	"locomoco/internals"
)

func stats(email string) {

	commits, ok := internals.ProcessRepos(email)
	if !ok {
		return
	}
	internals.PrintCommitStats(commits)

}
