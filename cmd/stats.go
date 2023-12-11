package cmd

import (
	"locomoco/internals"
)

func stats(email string) {

	commits := internals.ProcessRepos(email)

	internals.PrintCommitStats(commits)

}
