package yogapick

import (
	_ "embed"
	"errors"
	"iter"
	"math/rand"
	"os"
	"strings"
)

//go:embed poses_default.txt
var DefaultPoses []byte

// LoadPoses reads all lines from the given file.
func LoadPoses(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
		err := os.WriteFile(path, DefaultPoses, 0o644)
		if err != nil {
			return nil, err
		}
		data = DefaultPoses
	}
	s := strings.TrimSuffix(string(data), "\n")
	return strings.Split(s, "\n"), nil
}

// Suggest returns an iterator of up to count distinct poses selected at random
// without replacement from poses. If len(poses) < count, it yields all poses in
// a random order.
func Suggest(poses []string, count int) iter.Seq[string] {
	return func(yield func(string) bool) {
		if count <= 0 || len(poses) == 0 {
			return
		}
		idx := rand.Perm(len(poses))
		n := min(len(poses), count)
		for i := range n {
			p := poses[idx[i]]
			if !yield(p) {
				return
			}
		}
	}
}
