## Enlang
A programming language, that is just english

### Installing

**Clone the github repo**
```bash
git clone https://github.com/Nameless-Productions/enlang
```

**Build it**
(you must have go installed)

```bash
go build .
```

**Move it to your bin folder** (Linux only)

```bash
sudo mv enlang /bin
```

### Usage

You must set the `CLAUDE` env variable to your Claude API key

```bash
CLAUDE=your_api_key_here enlang examples/hello-world/main.el
./out/binary
```
