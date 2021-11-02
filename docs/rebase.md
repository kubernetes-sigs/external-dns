# Rebase

### This document explains the steps to rebase / merge upstream tag into this repo.

1. Update your git refs from both (i.e. upstream and downstream).
    ```shell
      git remote update
    ```
   
2. Checkout the Tag to be merged from upstream (e.g. tag v.0.10.1). 
    ```shell
      git checkout  v0.10.1
    ```
   
3. Create a tmp branch to do the merge work from.
    ```shell
      git checkout -b merge-tmp
    ```
   
4. Checkout downstream master, we need to be at the top of it when we merge.
    ```shell
      git checkout openshift/master    
    ```
   
5. Merge the tmp branch into our downstream master (note the output id from this cmd).
    ```shell
      echo 'merge kubernetes-sigs/external-dns v0.10.1' | git commit-tree merge-tmp^{tree} -p HEAD -p merge-tmp -F -    
    ```
6.  Create a new branch for the cherry-pick work.
    ```shell
      git branch merge-0.10.1 fa4fdf0d659cc02843a479c242bef5a6f4cbdde1 #this hash is the id from previous command, it will be different    
    ```
7.  Checkout the merge branch for cherry-picking.
    ```shell
      git checkout merge-0.10.1
    ```    
8. Execute the following command and note the most recent upstream commit.
    ```shell
      git log openshift/master
    ```
9. Identify the commit to be cherry-picked, from the most recent upstream commit, till the head of downstream commit. Write them into a local file for easier usage.
    ```shell
      git log --oneline --no-merges 35f2594745d6955625c47d7c46029cf1cb639154..openshift/master >> merges.txt
    ```
   
10. Cherry-pick the commits from the file ( bottom -> up ) one by one, and resolve conflicts.