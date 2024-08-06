# Rebasing Procedure

This fork needs to be periodically synced with upstream changes. To do this we
need to run the following steps:

1. Ensure the *upstream*, *downstream* and your own fork of the repository are  
   configured correctly as remotes.
   ```bash
   git remote -v
   # if upstream or downstream are not present then add them
   git remote add upstream https://github.com/kubernetes-sigs/external-dns
   git remote add downstream https://github.com/openshift/external-dns
   ```

2. Update the remotes
   ```bash
   git remote update
   ```

3. Ensure that the `master` branch is up-to-date with the *downstream* `master`
   branch.
   ```bash
   git switch master
   git pull downstream master
   ```

4. Checkout the commit from upstream to which you want to rebase the `master`
   branch.
   ```bash
   TAG_OR_COMMIT_ID="v0.14.2"
   git checkout ${TAG_OR_COMMIT_ID}
   git switch -c rebase-${TAG_OR_COMMIT_ID}
   ```

5. Merge the `master` branch into the current branch with
   the [ours](https://git-scm.com/docs/merge-strategies#_merge_strategies)
   merge strategy.
   ```bash
   git merge -s ours master
   ```
   _Note:_ This strategy should not be confused with the `ours` option for 
   the `ort` merge strategy.

6. Determine the changes present in the `master` branch which need to be
   cherry-picked.
   ```bash
   git log --oneline --no-merges ${TAG_OR_COMMIT_ID}..master
   ```

7. Cherry-pick the needed changes, try to build and test the code.
   ```bash
   make build test
   ```
    _Notes:_
    - Downstream fork uses `vendor` directory:
        - Drop carry patches with vendored dependencies from previous rebases.
        - Add a carry patch with lastest dependencies vendored.
    - Squash/fixup carry patches which touch the same files (e.g. `DOWNSTREAM_OWNERS`, `.ci-operator.yaml`).
    - Use [message format for cherry-picked commits](#message-format-for-cherry-picked-commits).

8. If any failure was found fix them and add as new carry patches. Then push the rebase branch.
   ```bash
   git push origin
   ```

## Message format for cherry-picked commits

The commits which are cherry-picked into the rebase branch should have the 
following format:

* `UPSTREAM: <carry>: $MESSAGE`
  A persistent carry that should probably be picked for the subsequent rebase
  branch. In general, these commits are used to modify behavior from upstream.

* `UPSTREAM: <drop>: $MESSAGE`
  A carry that should probably not be picked for the subsequent rebase branch.
  In general, these commits are used to maintain the codebase in ways that are
  branch-specific, like the update of generated files or dependencies.

* `UPSTREAM: 77870: $MESSAGE`
  The number identifies a PR in the upstream repository(
  i.e. https://github.com/kubernetes-sigs/external-dns/pull/<pr_id>)
  A commit with this message should only be picked into the subsequent rebase
  branch if the commits of the referenced PR are not included in the upstream
  branch. To check if a given commit is included in the upstream branch, open
  the referenced upstream PR and check any of its commits for the release tag
  targeted by the new rebase branch.
