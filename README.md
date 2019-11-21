![ssh-keep](https://i.imgur.com/sEr1RRJ.png)

# ssh-keep
Easy way access and keep track of SSH connections. It runs on single configuration file you can backup to your cloud or hard drive. Share accross devices in no time. Give it a shot!

## Install

Download latest release and extract to `/usr/share/bin` or any other folder listed in your `$PATH` variable. Therefore it's executable from your console.

## Configure

Once you start `ssh-keep` it should automaticall create `~/.ssh-keep.conf` file for you. Otherwise manually create it and add `SSH` lines. Here is an example:

```
user@hostname
user@hostname -i /home/example/your-public-key.pub
user2@ipaddress
...
```

Add alias to your `~/.bashrc` for faster tool access (optional).
```
alias sk="ssh-keep"
```

## Run
Once you hit `ssh-keep` you should see something like this:

![test](https://i.imgur.com/iWPGpr7.png)



