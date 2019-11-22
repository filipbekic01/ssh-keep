![ssh-keep](https://i.imgur.com/sEr1RRJ.png)

# ssh-keep
Easy way to access and keep track of your SSH connections. It runs on single configuration file which you can backup to your cloud or hard drive and share accross devices your way. Give it a shot! :blush:

## Install

Download latest release and extract file to `/usr/share/bin` with execute permission.

## Configure

Once you run `ssh-keep` command in your console, it should automaticall create `~/.ssh-keep.conf` file for you. Otherwise manually add it and write `SSH` lines with optional parameters. Make sure that one `SSH` command is in one line. Here is an example:

```
filip@filip-bekic.dev
root@filip-bekic.dev -i ~/path/to/public-key.pub
...
```

Add alias to your `~/.bashrc` for faster tool access (optional).
```
alias sk="ssh-keep"
```

## Run
Once you run `ssh-keep` after configuration, you should see next: 

![run_1](https://i.imgur.com/tXbjDjE.png)




