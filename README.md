# Plato

A CLI application for recording daily task and progress.


## Installation

Clone the repo and build.

```bash
  git clone https://github.com/CaveScraps/Plato.git
  cd Plato
  go mod tidy
  go build -o plato
```
The default dir for your notes is $HOME/notes  
If you would like to change that you can do so by setting the following env variable:
```bash
  export PLATO_NOTES_DIR="your/dir/here/"
```

## Usage:

```bash
  ./plato today "Message goes here"

  # Omitting the message prompts nvim to open to record the message:
  ./plato today

  # Keywords today, tomorrow, yesterday or dates in the 'yyyy-mm-dd' format are accepted:
  ./plato 2021-12-12 "My boy was robbed"

  # Todo items can be added with the -t flag:
  ./plato -t yesterday "Go to bed at a reasonable time"
```

## Bonus Features:
- Todo items from the previous day are copied over when a journal for a new day is added.
- You have to use nvim, i haven't set up an alternative, deal with it.

## License

[The Unlicense](./LICENSE)
