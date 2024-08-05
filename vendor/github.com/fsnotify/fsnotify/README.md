<<<<<<< HEAD
# File system notifications for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/fsnotify/fsnotify.svg)](https://pkg.go.dev/github.com/fsnotify/fsnotify) [![Go Report Card](https://goreportcard.com/badge/github.com/fsnotify/fsnotify)](https://goreportcard.com/report/github.com/fsnotify/fsnotify) [![Maintainers Wanted](https://img.shields.io/badge/maintainers-wanted-red.svg)](https://github.com/fsnotify/fsnotify/issues/413)

fsnotify utilizes [`golang.org/x/sys`](https://pkg.go.dev/golang.org/x/sys) rather than [`syscall`](https://pkg.go.dev/syscall) from the standard library.

Cross platform: Windows, Linux, BSD and macOS.

| Adapter               | OS                               | Status                                                                                                                          |
| --------------------- | -------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| inotify               | Linux 2.6.27 or later, Android\* | Supported |
| kqueue                | BSD, macOS, iOS\*                | Supported |
| ReadDirectoryChangesW | Windows                          | Supported |
| FSEvents              | macOS                            | [Planned](https://github.com/fsnotify/fsnotify/issues/11)                                                                       |
| FEN                   | Solaris 11                       | [In Progress](https://github.com/fsnotify/fsnotify/pull/371)                                                                   |
| fanotify              | Linux 2.6.37+                    | [Maybe](https://github.com/fsnotify/fsnotify/issues/114)                                                                      |
| USN Journals          | Windows                          | [Maybe](https://github.com/fsnotify/fsnotify/issues/53)                                                                         |
| Polling               | *All*                            | [Maybe](https://github.com/fsnotify/fsnotify/issues/9)                                                                          |

\* Android and iOS are untested.

Please see [the documentation](https://pkg.go.dev/github.com/fsnotify/fsnotify) and consult the [FAQ](#faq) for usage information.

## API stability

fsnotify is a fork of [howeyc/fsnotify](https://github.com/howeyc/fsnotify) with a new API as of v1.0. The API is based on [this design document](http://goo.gl/MrYxyA).

All [releases](https://github.com/fsnotify/fsnotify/releases) are tagged based on [Semantic Versioning](http://semver.org/).

## Usage

```go
package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/tmp/foo")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
```

## Contributing

Please refer to [CONTRIBUTING][] before opening an issue or pull request.

## FAQ

**When a file is moved to another directory is it still being watched?**

No (it shouldn't be, unless you are watching where it was moved to).

**When I watch a directory, are all subdirectories watched as well?**

No, you must add watches for any directory you want to watch (a recursive watcher is on the roadmap [#18][]).

**Do I have to watch the Error and Event channels in a separate goroutine?**

As of now, yes. Looking into making this single-thread friendly (see [howeyc #7][#7])

**Why am I receiving multiple events for the same file on OS X?**

Spotlight indexing on OS X can result in multiple events (see [howeyc #62][#62]). A temporary workaround is to add your folder(s) to the *Spotlight Privacy settings* until we have a native FSEvents implementation (see [#11][]).

**How many files can be watched at once?**

There are OS-specific limits as to how many watches can be created:
* Linux: /proc/sys/fs/inotify/max_user_watches contains the limit, reaching this limit results in a "no space left on device" error.
* BSD / OSX: sysctl variables "kern.maxfiles" and "kern.maxfilesperproc", reaching these limits results in a "too many open files" error.

**Why don't notifications work with NFS filesystems or filesystem in userspace (FUSE)?**

fsnotify requires support from underlying OS to work. The current NFS protocol does not provide network level support for file notifications.

[#62]: https://github.com/howeyc/fsnotify/issues/62
[#18]: https://github.com/fsnotify/fsnotify/issues/18
[#11]: https://github.com/fsnotify/fsnotify/issues/11
[#7]: https://github.com/howeyc/fsnotify/issues/7

[contributing]: https://github.com/fsnotify/fsnotify/blob/master/CONTRIBUTING.md

## Related Projects

* [notify](https://github.com/rjeczalik/notify)
* [fsevents](https://github.com/fsnotify/fsevents)

||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
fsnotify is a Go library to provide cross-platform filesystem notifications on
Windows, Linux, macOS, BSD, and illumos.

Go 1.17 or newer is required; the full documentation is at
https://pkg.go.dev/github.com/fsnotify/fsnotify

---

Platform support:

| Backend               | OS         | Status                                                                    |
| :-------------------- | :--------- | :------------------------------------------------------------------------ |
| inotify               | Linux      | Supported                                                                 |
| kqueue                | BSD, macOS | Supported                                                                 |
| ReadDirectoryChangesW | Windows    | Supported                                                                 |
| FEN                   | illumos    | Supported                                                                 |
| fanotify              | Linux 5.9+ | [Not yet](https://github.com/fsnotify/fsnotify/issues/114)                |
| AHAFS                 | AIX        | [aix branch]; experimental due to lack of maintainer and test environment |
| FSEvents              | macOS      | [Needs support in x/sys/unix][fsevents]                                   |
| USN Journals          | Windows    | [Needs support in x/sys/windows][usn]                                     |
| Polling               | *All*      | [Not yet](https://github.com/fsnotify/fsnotify/issues/9)                  |

Linux and illumos should include Android and Solaris, but these are currently
untested.

[fsevents]:   https://github.com/fsnotify/fsnotify/issues/11#issuecomment-1279133120
[usn]:        https://github.com/fsnotify/fsnotify/issues/53#issuecomment-1279829847
[aix branch]: https://github.com/fsnotify/fsnotify/issues/353#issuecomment-1284590129

Usage
-----
A basic example:

```go
package main

import (
    "log"

    "github.com/fsnotify/fsnotify"
)

func main() {
    // Create new watcher.
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    // Start listening for events.
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                log.Println("event:", event)
                if event.Has(fsnotify.Write) {
                    log.Println("modified file:", event.Name)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    // Add a path.
    err = watcher.Add("/tmp")
    if err != nil {
        log.Fatal(err)
    }

    // Block main goroutine forever.
    <-make(chan struct{})
}
```

Some more examples can be found in [cmd/fsnotify](cmd/fsnotify), which can be
run with:

    % go run ./cmd/fsnotify

Further detailed documentation can be found in godoc:
https://pkg.go.dev/github.com/fsnotify/fsnotify

FAQ
---
### Will a file still be watched when it's moved to another directory?
No, not unless you are watching the location it was moved to.

### Are subdirectories watched?
No, you must add watches for any directory you want to watch (a recursive
watcher is on the roadmap: [#18]).

[#18]: https://github.com/fsnotify/fsnotify/issues/18

### Do I have to watch the Error and Event channels in a goroutine?
Yes. You can read both channels in the same goroutine using `select` (you don't
need a separate goroutine for both channels; see the example).

### Why don't notifications work with NFS, SMB, FUSE, /proc, or /sys?
fsnotify requires support from underlying OS to work. The current NFS and SMB
protocols does not provide network level support for file notifications, and
neither do the /proc and /sys virtual filesystems.

This could be fixed with a polling watcher ([#9]), but it's not yet implemented.

[#9]: https://github.com/fsnotify/fsnotify/issues/9

### Why do I get many Chmod events?
Some programs may generate a lot of attribute changes; for example Spotlight on
macOS, anti-virus programs, backup applications, and some others are known to do
this. As a rule, it's typically best to ignore Chmod events. They're often not
useful, and tend to cause problems.

Spotlight indexing on macOS can result in multiple events (see [#15]). A
temporary workaround is to add your folder(s) to the *Spotlight Privacy
settings* until we have a native FSEvents implementation (see [#11]).

[#11]: https://github.com/fsnotify/fsnotify/issues/11
[#15]: https://github.com/fsnotify/fsnotify/issues/15

### Watching a file doesn't work well
Watching individual files (rather than directories) is generally not recommended
as many programs (especially editors) update files atomically: it will write to
a temporary file which is then moved to to destination, overwriting the original
(or some variant thereof). The watcher on the original file is now lost, as that
no longer exists.

The upshot of this is that a power failure or crash won't leave a half-written
file.

Watch the parent directory and use `Event.Name` to filter out files you're not
interested in. There is an example of this in `cmd/fsnotify/file.go`.

Platform-specific notes
-----------------------
### Linux
When a file is removed a REMOVE event won't be emitted until all file
descriptors are closed; it will emit a CHMOD instead:

    fp := os.Open("file")
    os.Remove("file")        // CHMOD
    fp.Close()               // REMOVE

This is the event that inotify sends, so not much can be changed about this.

The `fs.inotify.max_user_watches` sysctl variable specifies the upper limit for
the number of watches per user, and `fs.inotify.max_user_instances` specifies
the maximum number of inotify instances per user. Every Watcher you create is an
"instance", and every path you add is a "watch".

These are also exposed in `/proc` as `/proc/sys/fs/inotify/max_user_watches` and
`/proc/sys/fs/inotify/max_user_instances`

To increase them you can use `sysctl` or write the value to proc file:

    # The default values on Linux 5.18
    sysctl fs.inotify.max_user_watches=124983
    sysctl fs.inotify.max_user_instances=128

To make the changes persist on reboot edit `/etc/sysctl.conf` or
`/usr/lib/sysctl.d/50-default.conf` (details differ per Linux distro; check your
distro's documentation):

    fs.inotify.max_user_watches=124983
    fs.inotify.max_user_instances=128

Reaching the limit will result in a "no space left on device" or "too many open
files" error.

### kqueue (macOS, all BSD systems)
kqueue requires opening a file descriptor for every file that's being watched;
so if you're watching a directory with five files then that's six file
descriptors. You will run in to your system's "max open files" limit faster on
these platforms.

The sysctl variables `kern.maxfiles` and `kern.maxfilesperproc` can be used to
control the maximum number of open files.
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
