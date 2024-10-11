/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"

	"github.com/richbai90/git-training-wheels/pkg/common"
	cmd_errs "github.com/richbai90/git-training-wheels/pkg/errors"
	"github.com/spf13/cobra"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Switch branches or restore working tree files",
	Long: `Updates files in the working tree to match the version in the index or the specified tree. If no pathspec was given, git checkout will also update HEAD to set the
       specified branch as the current branch.

       git checkout [<branch>]
           To prepare for working on <branch>, switch to it by updating the index and the files in the working tree, and by pointing HEAD at the branch. Local
           modifications to the files in the working tree are kept, so that they can be committed to the <branch>.

           If <branch> is not found but there does exist a tracking branch in exactly one remote (call it <remote>) with a matching name and --no-guess is not specified,
           treat as equivalent to

               $ git checkout -b <branch> --track <remote>/<branch>

           You could omit <branch>, in which case the command degenerates to "check out the current branch", which is a glorified no-op with rather expensive side-effects
           to show only the tracking information, if it exists, for the current branch.

       git checkout -b|-B <new-branch> [<start-point>]
           Specifying -b causes a new branch to be created as if git-branch(1) were called and then checked out. In this case you can use the --track or --no-track
           options, which will be passed to git branch. As a convenience, --track without -b implies branch creation; see the description of --track below.

           If -B is given, <new-branch> is created if it doesn’t exist; otherwise, it is reset. This is the transactional equivalent of

               $ git branch -f <branch> [<start-point>]
               $ git checkout <branch>

           that is to say, the branch is not reset/created unless "git checkout" is successful (e.g., when the branch is in use in another worktree, not just the current
           branch stays the same, but the branch is not reset to the start-point, either).

       git checkout --detach [<branch>], git checkout [--detach] <commit>
           Prepare to work on top of <commit>, by detaching HEAD at it (see "DETACHED HEAD" section), and updating the index and the files in the working tree. Local
           modifications to the files in the working tree are kept, so that the resulting working tree will be the state recorded in the commit plus the local
           modifications.

           When the <commit> argument is a branch name, the --detach option can be used to detach HEAD at the tip of the branch (git checkout <branch> would check out
           that branch without detaching HEAD).

           Omitting <branch> detaches HEAD at the tip of the current branch.

       git checkout [-f|--ours|--theirs|-m|--conflict=<style>] [<tree-ish>] [--] <pathspec>..., git checkout [-f|--ours|--theirs|-m|--conflict=<style>] [<tree-ish>]
       --pathspec-from-file=<file> [--pathspec-file-nul]
           Overwrite the contents of the files that match the pathspec. When the <tree-ish> (most often a commit) is not given, overwrite working tree with the contents
           in the index. When the <tree-ish> is given, overwrite both the index and the working tree with the contents at the <tree-ish>.

           The index may contain unmerged entries because of a previous failed merge. By default, if you try to check out such an entry from the index, the checkout
           operation will fail and nothing will be checked out. Using -f will ignore these unmerged entries. The contents from a specific side of the merge can be checked
           out of the index by using --ours or --theirs. With -m, changes made to the working tree file can be discarded to re-create the original conflicted merge
           result.

       git checkout (-p|--patch) [<tree-ish>] [--] [<pathspec>...]
           This is similar to the previous mode, but lets you use the interactive interface to show the "diff" output and choose which hunks to use in the result. See
           below for the description of --patch option.
`,
	Args:               cobra.ArbitraryArgs,
	DisableFlagParsing: true,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		safe, err := checkSafety(args)
		if err != nil {
			return err
		}

		if !safe {
			return fmt.Errorf(("unsafe operation detected"))
		}

		return nil
	},
}

func checkSafety(args []string) (bool, error) {

	wt, err := common.GetWorkTree()
	if err != nil {
		return false, err
	}

	// check if the branch is dirty
	status, err := wt.Status()
	if err != nil {
		return false, err
	}

	if !status.IsClean() {
		return false, cmd_errs.NewSafetyError("The branch is not clean", cmd_errs.GenerateErrorCode(
			cmd_errs.CheckoutCommand,
			cmd_errs.CheckoutErrorUnclean,
		))
	}

	// check if '--' is used in args
	if idx := slices.Index(args, "--"); idx != -1 {
		isDir, err := common.IsDirectory(args[idx+1])
		if err != nil {
			return false, err
		}
		if isDir {
			return false, cmd_errs.NewSafetyError("Checkout would reset a whole directory", cmd_errs.GenerateErrorCode(
				cmd_errs.CheckoutCommand,
				cmd_errs.CheckoutErrorReset,
			))
		}

	}

	return true, nil
}

func init() {
	rootCmd.AddCommand(checkoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
