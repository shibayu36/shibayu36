# shibayu36
## How to Release

### Release by TravisCI

1. `git tag vx.y.z`
2. git push --tags
3. Wait to build at https://travis-ci.org/shibayu36/shibayu36
4. See https://github.com/shibayu36/shibayu36/releases

### Release by manually

1. Install dependencies for release by `make setup`
2. `git tag vx.y.z`
3. GITHUB_TOKEN=... script/release
4. See https://github.com/shibayu36/shibayu36/releases
