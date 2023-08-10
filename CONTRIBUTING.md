# Contributing

All code you commit and submit by pull-request should follow these simple guidelines.

## Branching

Create a branch using the following naming prefixes:

f = feature

b = bug fix

d = documentation

t = tests

td = technical debt

v = dependencies ("vendoring" previously)

Some indicative example branch names would be f-update_assets_to_EDCv0.2.0 or td-staticcheck-st1008

## Commit messages

Follow [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) standard.

## Pull requests

1. Open the pull request giving a title following [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/). Keep the title short (~50 characters). Always try to give a scope to the title.

  Following this convention will help keep the repository in order and create sleek changelogs.

2. Try to fill the pull request template as closely as possible. Provide a short description that answers these questions:

  - "how this pull request will change the code?"
  - "what is the purpose of these changes?"
  
  Take the opporutnity to reflect on your learnings and share them with everyone else.

3. For any task, try to provide a way to test its results; it would greatly help the QA process.

4. Assign the pull request to yourself.

5. Upon approval, you are in charge of merging the code back. It's crucial to own the code we create and ensure it **is released**. If you wish your code is merged upon approval, please, specify it in the pull request description.

6. Upon merge (ndr. squash and merge), remember to delete the merged working branch; let's keep a clean repository.

## New features & bug fixes

Within this project, we follow a _Test-Driven Development_ (TDD) approach.
