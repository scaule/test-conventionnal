name: Golang workflow
on:
  pull_request:
  push:
    branches:
      - develop
      - master
  create:
    tags:
      - v*

jobs:
  commitlint:
    runs-on: [self-hosted, docker]
    name: Check commit message format
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - uses: wagoid/commitlint-github-action@v5
