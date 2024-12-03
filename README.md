# T2

**T2** is a terminal-based time tracker built with Go.
The 2 T's are for Time Tracker, I'm not good naming things.
I made this thing mainly because I wanted to learn how to 
realese an app for distribution.

## Prerequisites

### on macOS/Linux
You'll need to install the package manager [`Homebrew`](https://brew.sh/). That's it.

### on Windows
You'll need to install the package manager [`Scoop`](https://scoop.sh/).
Then, you go to the Environment Variables and under the System Variables, 
go to Path and add your Scoop installation directory, usually C:\Users\YourUser\scoop\shims


## Installation

### Install via Homebrew (macOS/Linux)
```bash
brew tap franpfeiffer/t2 # create the tap
brew install t2          # install
```
Or `brew install franpfeiffer/t2/t2`.

Or, in a [`brew bundle`](https://github.com/Homebrew/homebrew-bundle) `Brewfile`:
```bash
tap "franpfeiffer/t2"
brew "t2"
```
### Install via Scoop (Windows)
```bash
scoop bucket add t2 git@github.com:franpfeiffer/scoop-t2 # create the bucket
scoop install t2                                         # install
```

## Usage

### Start T2
Open the terminal on the directory of the project you want to track the time and run:
```bash
t2 
```
It will create a .txt with the amount of minutes you spent on the session.
Every time you run it, other line will be added, like this:
```txt
Time tracked: 69 minute(s)
Time tracked: 122 minute(s)
Time tracked: 420 minute(s)
```
Every line represents a session.
