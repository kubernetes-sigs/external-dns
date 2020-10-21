# Merging from GitLab to Github for promoting F5 DNS external-dns changes to github

The Branching structure for the external-dns project is as follows

[github.com/kubernetes-sigs/external-dns](https://github.com/kubernetes-sigs/external-dns) is the main repo
[github.com/F5Networks/external-dns](https://github.com/F5Networks/external-dns) is forked from the above repo
 * [master] () This master branch will be periodically synced from the k8s external-dns project.
 * [release] () All F5 related changes would be pushed from F5 internal repos to this branch using mirroring

 ## Initial Mirror from github to gitlab

 Gitlab Mirroring is setup to Pull the github changes from F5Networks/external-dns to gitswarm.f5net.com/f5aas/external-dns

 This is first time only to setup the local gitswarm repo with the fork from github.

 After the first pull, the mirroring is disabled. This can be enabled to periodically sync changes from github back to gitlab.

 ## Branch structure in gitlab

 Following branches are created in gitlab 
 * release - To push changes from gitlab to github
 * github (child of release) - Intermediary branch between release and gitlab branch
 * gitlab (child of github) - All F5 DNS external-dns development should be merged in this branch. This forms the parent for any feature/bugfixes.

 ## Subsequent F5 DNS LB external-dns push from gitlab to github

Gitlab Mirroring is setup to Push the gitlab changes from gitswarm.f5net.com/f5aas/external-dns 'release' branch to F5Networks/external-dns 'release' branch.

## F5 DNS LB enhancements and bugfixes

All enhancements and bugfixes should be done in branch pulled off gitlab branch. Merge requests once approved should be merged into gitlab branch.

## Release to github

We do not want to push F5 specific files to github. So we have to follow a specific merge process to release changes to github.
1. Merge gitlab branch into github branch - This prevents the specific files to be skipped from the merge
    ```console
    Checkout github branch and do a git pull
    git merge --no-commit --no-ff gitlab
    git reset HEAD -- Makefile.f5cs .gitlab-ci.yml Dockerfile PROJECT Makefile VERSION GITLAB_TO_GITHUB_MERGE.md
    git commit -m "Merge $VERSION into github"
    git push -u origin github
    ```

2. Merge github branch to release branch. The mirroring is already setup on release branch only to push the changes to github.com/F5Netowkrs/external-dns repo.

3. Check  [github.com/F5Networks/external-dns](https://github.com/F5Networks/external-dns) is updated with the latest push.

4. In [github.com/F5Networks/external-dns](https://github.com/F5Networks/external-dns) now merge the changes to master branch.

5. Raise a Pull Request to merge the changes from [github.com/F5Networks/external-dns](https://github.com/F5Networks/external-dns) to [github.com/kubernetes-sigs/external-dns](https://github.com/kubernetes-sigs/external-dns)
