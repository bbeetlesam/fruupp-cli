# FruuppCLI

Your head's probably asking "Why the hell we ever need this nonsense?"

Well, your question is valid. So do I.

## What even is this?

A command-line tool for reading song lyrics. In your terminal.\
Think of lyrics-in-terminal, syrics, genius-cli, or lyrics-cli, but instead of using
an API, it uses local files (yes, you have to add the lyrics files yourself) to display.

Pretty easy (and inefficient), right?

## How to install it?

Thanks, (if) you're interested to install this thing. But before that:

Prerequisites:

- Go v1.20+ (or whatever version you want).
- A terminal, obviously.
- Questionable life choices, I guess.

Building and installing from source:

```bash
git clone https://github.com/bbeetlesam/fruupp-cli.git
cd fruupp-cli
go build -o fruupp .
```

Done! Now try to run `./fruupp` to see if it works. If yes, good. If not, not good.

## (How) does it work?

Allegedly.

The command is pretty obvious: `fruupp`.

You can use `fruupp help` to display more information on how to interact with this little
binary.

These are all available commands that you can use:

- `fruupp help` to get some immediate help.
- `fruupp list <dir>` to list all available fruupp files in a directory.
- `fruupp view <filename>` to display lyrics from a fruupp file (make sure that the file is a fruupp file and it exists).
- `fruupp version` to, well, display version info (unimportant to you probably).

### Need examples?

Because reading docs is for people with attention span (it's obsolete nowadays, isn't it?),
here's some examples for you.

#### List all fruupp files in a directory

```bash
fruupp list ./lyrics
```

If *./lyrics* has some fruupp files, the expected output is:

```
lyrics/echoes
lyrics/one-of-these-days
```

If none, then the output will be:

```
No fruupp files found.
```

(oh no!)

#### View lyrics from a fruupp file

Now that you know the file's path, you can use `view` to view it.

```bash
fruupp view lyrics/one-of-these-days
```

Expected output:

```
Title: One of These Days
Artist: Pink Floyd
Album: Meddle
Released Year: 1971

One of these days, I'm gonna cut you into little pieces!
```

#### Creating your first fruupp file

Let's say you want to add "One of These Days" by Pink Floyd:

```bash
cat > one-of-these-days << EOF
#!fruupp

title: One of These Days
artist: Pink Floyd
album: Meddle
year: 1971

---
[The only part by Nick Mason]
One of these days, I'm gonna cut you down into little pieces!
EOF
```

Then, run

```bash
fruupp view one-of-these-days
```

and it should display the song!

Now, go fill your lyrics folder with prog epics.

### About fruupp files

A fruupp file is a file that starts with a shebang-like `#!fruupp` on the first line (mandatory).\
It's like how you need to add a shebang in a shell script, but this one only works for fruupp program.\
The file's extension doesn't matter, as long as it has the fruupp-shebang on the first line,
it will be detected as a fruupp file. So, you can name the files whatever you want
(fruupp, README.md, nvme0n1, .bashrc, scary_virus.sh, whatever. I don't suggest it though).

What should be inside the fruupp file?

A fruupp file consists of two main sections: metadata and lyrics, which are separated by a `---` line.
It resembles YML-like syntax, but this one can't be used for storing data.

Metadata section is the section where you can fill in the metadata, such as song's title, album,
and artist. All values are treated as strings here (no need for "" though), so you can't put
a boolean nor mathematical equation here. Like, mate why did you have to do that in the first place.

Lyrics section is to print the lyrics. It will be displayed to the terminal as-is, so if you have
some typos or unspaced words, well that's on you, not the program.

The general (I mean the idiomatic) structure for a fruupp file is like this:

```
#!fruupp

title: [anything, like Nine Feet Underground?]
artist: [hmm, I guess it's Caravan]
album: [is it In the Land of Grey and Pink?]
year: [around 1971 I think]

---

[lyrics comes here, fill this with whatever letters you want]
```

Note that the metadata order doesn't have to be exactly like that, meaning that you can order it like:

```
year:
album:
title:
artist:
```

Whatever that would satisfy you.

Named after the Irish prog band Fruupp (1971-1976), because they deserved better recognition than they got.
