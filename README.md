# Like

Utilities for finding similar things.

## like

Returns the line of input that most closely matches the given text.

```bash
cat names | like wombat
```

## find-like

Returns the filename in a set of directories that most closely matches the given text.

```bash
find-like app/models/user.rb tests spec
```

## lsort

Sorts lines from stdin by Levenshtein distance.

```bash
cat names | lsort
```

Options:

- `-w` calculate Levenshtein distance using words instead of letters.

## lgraph

Creates a graph of lines from stdin by Levenshtein distance.

Default output is 3 integers: a line number, the line number of the line that most closely matches the line at the first line number, and the Levenshtein distance between the two lines.

```bash
cat names | lgraph
```

Options:

- `-w` calculate Levenshtein distance using words instead of letters
- `-l` print whole lines instead of line numbers (takes four lines per edge rather than one)
