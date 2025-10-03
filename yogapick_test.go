package yogapick_test

import (
	"bytes"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/bitfield/yogapick"
)

func TestLoadPoses_ReadsAllLinesInFile(t *testing.T) {
	t.Parallel()
	got, err := yogapick.LoadPoses("testdata/poses.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"pose 1", "pose 2", "pose 3"}
	if !slices.Equal(want, got) {
		t.Errorf("wrong poses: %q", got)
	}
}

func TestLoadPoses_CreatesFileWithDefaultContentsIfMissingAndReturnsDefaultData(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/poses.txt"
	got, err := yogapick.LoadPoses(path)
	if err != nil {
		t.Fatal(err)
	}
	want := strings.Split(string(yogapick.DefaultPoses), "\n")
	if !slices.Equal(want, got) {
		t.Errorf("wrong poses: %q", got)
	}
	gotBytes, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	wantBytes, err := os.ReadFile("poses_default.txt")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wantBytes, gotBytes) {
		t.Errorf("wrong file contents: %q", gotBytes)
	}
}

func TestSuggestYieldsNoMoreThanCountLines(t *testing.T) {
	t.Parallel()
	input := []string{"pose 1", "pose 2", "pose 3"}
	got := slices.Collect(yogapick.Suggest(input, 2))
	if len(got) != 2 {
		t.Errorf("wrong result: %q", got)
	}
}

func TestSuggestYieldsAllLinesWhenFewerThanCount(t *testing.T) {
	t.Parallel()
	input := []string{"pose 1", "pose 2", "pose 3"}
	got := slices.Collect(yogapick.Suggest(input, 4))
	if len(got) != 3 {
		t.Errorf("wrong result: %q", got)
	}
}

func TestSuggestYieldsNoLinesWhenCountNegative(t *testing.T) {
	t.Parallel()
	input := []string{"pose 1", "pose 2", "pose 3"}
	got := slices.Collect(yogapick.Suggest(input, -1))
	if len(got) != 0 {
		t.Errorf("wrong result: %q", got)
	}
}
