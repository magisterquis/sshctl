sshctl
======

Go library to automate commands via ssh.  Not really production code at the moment.

TODO:
-----

List of username@host:port tuples, re-read on HUP...
Ability to define sets of hosts, via regular expression...
Send commands to a set or host, get output.
Custom ssh key, not just the usual $HOME/.ssh/id\_?sa{,.pub}

So, it looks vaguely like this could be used for malicious purposes.  Don't do
that.

Host file format:
=================
```user@host[:port][\tgroup]

foo@bar.com
tridge@baaz:2222
instance001@localhost	localgroup
instance002@localhost	localgroup
instance001@remote:99	remgroup
instance004@remote	remgroup```
