
# Locomoco

Locomoco is an light-weight tool designed to provide GitHub users with a swift and intuitive way to visualize their contribution history and Github activity. Simply run `locomoco` to see git contibutions in your CLI or `locomoco showme` to see a list of repos you've contributed to.



## Installation

Install locomoco with Homebrew

```bash
brew install faceyacc/tools/locomoco
```

Install locomoco locally with Go
```bash
go get github.com/faceyacc/loco-moco
```
## Usage/Examples


Add projects to .locomoco
```bash
locomoco --add User/fakeName/project
```
Initalize email and account name associated with your GitHub account
```bash
locomoco --email fake@fake.com --user fakeacc
```
To get a quick snapshot of repos
```bash
locomoco showme
```
To get a quick snapshot of repos for a different GitHub user/organization
```bash
locomoco showme --newUser newFakeAcc
```
## Roadmap

- Additional Unit test.

- Allow users to run `locomoco` command from outside an ititalize project space.

- Bettter formatting for `showme` command.

## Demo

![](https://github.com/faceyacc/locomoco/blob/main/locomoco.gif)

## Support

For support, email me at justfacey@gmail.com.

