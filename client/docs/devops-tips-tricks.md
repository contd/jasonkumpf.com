---
title: Devops Tips and Tricks
lang: en-US
tags:
  - devops
---

# {{ $page.title }}

## Check the Status of All Services (upstart)

In a Linux server that uses upstart, you can quickly check the status of a number of different services by running: `

```bash
[sudo] service --status-all
```

This command will be invoked for all services under the `/etc/init.d/` directory.

The `-` symbol means it isn't running, the `+` symbol means it is up, and the `?` symbol means that it cannot determine the status.

If you want to check the status of something like `nginx` or `apache` with more details, just run:

```bash
[sudo] service nginx status
```

## Check the Status of All Services (systemd)

Severs that utilize systemd can take advantage of the `systemctl` command. On newe Ubuntu distributions like 16.04 or higher, the `service` command is just a alias to `systemctl`.  For controlling and checking a service here is what the commands would look like:

```bash
sudo systemctl status nginx.service
sudo systemctl start nginx.service
sudo systemctl stop nging.service
sudo systemctl restart nginx.service
```

The files for these services are usually located in `/etc/systemd/system` and depending on their requriemets they might be in a sub directory from there.  You can enable or disable a service like so:

```bash
sudo systemctl enable nginx.service
sudo systemctl disable nginx.service
```

This no only keeps it from automatically starting but also can't be started until it's enabled.


## Check the Syntax of nginx Files

The syntax of [`nginx`](https://www.nginx.com/) configuration files can be a bit finicky. On top of that, some `nginx` server commands can fail silently. Get more confidence by using the following command to check for syntax errors in those files:

```bash
$ [sudo] nginx -t
```

If there is an error, you might see something like this:

```bash
$ sudo nginx -t
nginx: [emerg] unexpected ";" in /etc/nginx/nginx.conf:16
nginx: configuration file /etc/nginx/nginx.conf test failed
```

If all looks good, then you'll see this:

```bash
$ sudo nginx -t
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
```

## Determine the IP Address of a Domain

The `dig` (domain information grouper) command can be used to get more information about a domain name. To discover the IP address for a given domain, invoke `dig` with the domain as an argument.

```bash
$ dig jasonkumpf.com

; <<>> DiG 9.10.3-P4-Ubuntu <<>> jasonkumpf.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 39939
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 13, ADDITIONAL: 3

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;jasonkumpf.com.			IN	A

;; ANSWER SECTION:
jasonkumpf.com.		599	IN	A	50.63.202.13

;; AUTHORITY SECTION:
.			3199	IN	NS	i.root-servers.net.
.			3199	IN	NS	b.root-servers.net.
.			3199	IN	NS	a.root-servers.net.
.			3199	IN	NS	e.root-servers.net.
.			3199	IN	NS	k.root-servers.net.
.			3199	IN	NS	j.root-servers.net.
.			3199	IN	NS	g.root-servers.net.
.			3199	IN	NS	h.root-servers.net.
.			3199	IN	NS	d.root-servers.net.
.			3199	IN	NS	c.root-servers.net.
.			3199	IN	NS	m.root-servers.net.
.			3199	IN	NS	l.root-servers.net.
.			3199	IN	NS	f.root-servers.net.

;; ADDITIONAL SECTION:
E.ROOT-SERVERS.net.	30	IN	AAAA	2001:500:a8::e
G.ROOT-SERVERS.net.	22938	IN	AAAA	2001:500:12::d0d

;; Query time: 101 msec
;; SERVER: 127.0.1.1#53(127.0.1.1)
;; WHEN: Fri Jul 07 18:10:30 EDT 2017
;; MSG SIZE  rcvd: 343

```

The *answer section* tells me that the IP address for `jasonkumpf.com` is `50.63.202.13`.

See `man dig` for more details.

## Running Out of inode Space

Unix systems have two types of storage limitations. The first, and more common, is a limitation on physical storage used for storing the contents of files. The second is a limitation on `inode` space which represents file location and other data.

Though it is uncommon, it is possible to run out of `inode` space before running out of disk space (run `df` and `df -i` to see the levels of each). When this happens, the system will complain that there is `No space left on device`. Both `inode` space and disk space are needed to create a new file.

How can this happen? If lots of directories with lots of empty, small, or duplicate files are being created, then the `inode` space can be used up disproportionately to the amount of respective disk space. You'll need to clean up some of those files before you can continue.

Sources: [this](http://blog.scoutapp.com/articles/2014/10/08/understanding-disk-inodes)
and [this](http://www.linux.org/threads/intro-to-inodes.4130/)
