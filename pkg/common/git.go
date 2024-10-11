package common
import "github.com/go-git/go-git/v5"

func GetWorkTree() (*git.Worktree, error) {
	r, err := git.PlainOpen(".")
    if err != nil {
        return nil, err
    }

    // Get the worktree
    w, err := r.Worktree()
    if err != nil {
        return nil, err
    }

	return w, nil
}