# glob-demo

Mock files (tag `bal`, content: baked alaska instructions with web sources) that
probe how CODEOWNERS and a ruleset's required-reviewer `file_patterns` read
ambiguous glob forms. Each `probe/*` PR edits exactly one file; compare whether
CODEOWNERS auto-requests a review and whether the ruleset blocks the merge.
Results are filled in below once the PRs are open.

## Review Rules Demo

This repo checks how a suffix glob like `**/*_bal.go` behaves in two
places that decide who must review a change:

- **CODEOWNERS**, which requests a review from the owner.

- **A branch ruleset's required-reviewer `file_patterns`**, which
  blocks the merge until a reviewer approves.

Suffix globs let a file's name, not its location, decide who owns
it. That only works if both systems read the pattern the same
way. Both must match a file at the repo root and a file nested at any
depth.

### How it is tested

Each probe PR edits exactly one file. The sole CODEOWNERS owner is
`@tychoish`, so a review request means the file matched a rule. A
blocked merge means the ruleset pattern matched.

| Probe PR                                                    | Pattern shape       | File touched                         | Result                                                                                 |
|-------------------------------------------------------------|---------------------|--------------------------------------|----------------------------------------------------------------------------------------|
| [#3](https://github.com/tycho-glob-demo/glob-demo/pull/3)   | `/**/*_bal.go`      | `meringue_bal.go` (repo root)        | Review requested and merge blocked. `/**/` matches zero directories.                   |
| [#4](https://github.com/tycho-glob-demo/glob-demo/pull/4)   | same                | `kitchen/prep/ice_cream_bal.go`      | Review requested and merge blocked.                                                    |
| [#5](https://github.com/tycho-glob-demo/glob-demo/pull/5)   | `**/*_bal_test.go`  | `kitchen/prep/ice_cream_bal_test.go` | Review requested and merge blocked.                                                    |
| [#39](https://github.com/tycho-glob-demo/glob-demo/pull/39) | `/**/*_bal_test.go` | `kitchen/prep/ice_cream_bal_test.go` | Review requested. The base branch is outside the ruleset, so nothing blocks the merge. |
| [#16](https://github.com/tycho-glob-demo/glob-demo/pull/16) | none (control)      | `unowned/control.md`                 | No review requested and merge state clean.                                             |

Probe #39 closes a gap left by #5. The leading-slash form
`/**/*_bal_test.go` is the one used in CODEOWNERS, and #39 shows it
matches a nested test file. Its base branch carries that CODEOWNERS
line, because CODEOWNERS is read from the PR's base branch. That
branch isn't in any ruleset's ref conditions, so #39 tests CODEOWNERS
only. The ruleset form has no leading slash and was already measured
in #5.

The full results table is in [PR #19](https://github.com/tycho-glob-demo/glob-demo/pull/19).
Other glob forms are in [PRs #21 to #38](https://github.com/tycho-glob-demo/glob-demo/pull/38),
which are based on the branch in [PR #20](https://github.com/tycho-glob-demo/glob-demo/pull/20).

### Writing the same rule in both places

- CODEOWNERS can anchor a pattern with a leading `/`, as in
  `/**/*_bal.go`. The ruleset pattern is the same text without the
  slash, as in `**/*_bal.go`.

- A suffix glob needs one line per suffix. `*_bal.go` does not match
  `x_bal_test.go`, so tests need their own line.

- The ruleset is edited in GitHub, not in the repo. When CODEOWNERS
  gains or loses a pattern, the ruleset must be updated by hand. Until
  then the two lists disagree.

### Not covered

- The probes don't test a ruleset with several reviewer teams or a cap
  on patterns per team.

- No probe combines the leading-slash form and the ruleset in one
  PR. The two halves were measured separately, in #39 and #5.
