## Learning Go and some other things .. maybe?

This is a repository I'm using as I'm learning Go from the "learn-go-by-tests" created by quii

[Learn Go with Tests](https://github.com/quii/learn-go-with-tests)

[Download and install](https://go.dev/doc/install)

[Installation from source](https://go.dev/doc/install/source)

### Install Go

#### Linux (_I'm using Debian_)

Extract the archive you downloaded into
`/usr/local`
creating a Go tree in /usr/local/go.

:: Important ::
This step will remove a previous installation at
`/usr/local/go`
if there are any, prior to extracting.
Please back up any data before proceeding.

For example, run the following as root or through sudo:
`rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz`
Add `/usr/local/go/bin` to the `PATH` environment variable.
You can do this by adding the line below to your
`$HOME/.profile`
	or
`/etc/profile` (for a system-wide installation):

Add this line 
`export PATH=$PATH:/usr/local/go/bin`

_I added mine in `~/.xsessionrc` instead, works but better and faster for me using Bunsenlabs (Debian)_


#### Note
Changes made to a profile file may not apply until the next time you log into your computer.
To apply the changes immediately run the shell commands directly or
execute them from the profile using a command such as
`source $HOME/.profile`


#### Side-note
I needed both of these lines for the `godoc -http` command to work
`export PATH="$PATH:/usr/local/go/bin"`
`export PATH="$PATH:$HOME/go/bin"`

Verify that you've installed Go by opening a command prompt and typing the following command:
$ go version
Confirm that the command prints the installed version of Go.

### Initiate Go module
go mod init $URL/PATH

### Update(?) tidy mod
go mod tidy

### Tools Installation
go get -v golang.org/x/tools/gopls
go get -v golang.org/x/tools/gopkgs
go get -v golang.org/x/tools/go-outline
go get -v golang.org/x/tools/dlv
go get -v golang.org/x/tools/dlv-dap
go get -v golang.org/x/tools/staticcheck
