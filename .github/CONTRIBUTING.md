# Contributing to the project

## When to contribute

Some times, you may want to contribute to the project, this is encouraged... within a few guidelines. Some times when you might contribute are:

- When you have a new feature to add
  - If adding a feature, a review of the code is required before it is added to the project
- When you have a bug to fix
  - If fixing a bug, a review of the code is required before it is fixed
- When there is a security issue
  - If fixing a security issue, a review of the code is required before it is fixed

### How to contribute

This repository is open source, so contributing is as simple as submitting a pull request. This repository uses Trunk-Based Development (TBD) which gives the following git flow:

For a new feature:

1.  `git checkout -b feature/<feature-name>`
    - Complete Work Item
2.  `git commit -m "Add feature <feature-name>"`
3.  `git push origin feature/<feature-name>`
4.  `git checkout trunk`

For a bug fix:

1.  `git checkout -b bugfix/<bugfix-name>`
    - Complete Work Item
2.  `git commit -m "Fix bug <bugfix-name>"`
3.  `git push origin bugfix/<bugfix-name>`
4.  `git checkout trunk`

For a chore (i.e. non-code change):

1.  `git checkout -b chore/<chore-name>`
    - Complete Work Item
2.  `git commit -m "Chore <chore-name>"`
3.  `git push origin chore/<chore-name>`
4.  `git checkout trunk`
