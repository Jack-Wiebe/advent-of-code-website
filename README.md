Advent of Code Solutions Web Interface

A web-based interface for running and displaying Advent of Code solutions written in Go. This project provides a clean web UI to execute your AOC solutions and view the results in real-time.

Features

- Web interface for Advent of Code solutions

- Real-time solution execution

- Clean separation of solution values and console output

- Support for multiple days/years

- Structured JSON API

- Responsive design

Project Structure

```
advent-of-code-web/
├── go.mod
├── main.go
├── types/
│   └── types.go              #  Shared type definitions
├── handlers/
│   └── handlers.go           # HTTP request handlers
├── solutions/
│   ├── year/
│   │   └── day
│   │       ├── main.go       # Problem entrypoint
│   │       ├── part1.go      # Part 1 solution
│   │       └── part2.go      # Part 2 solution
│   └── utils
│       └── utils.go          # Shared utility functions
└── static/
    ├── index.html            # Main frontend
    ├── styles.css            # Styling
    └── app.js                # Frontend JavaScript

```

Quick Start
Clone the repository:

```
git clone https://github.com/Jack-Wiebe/advent-of-code-website
cd advent-of-code-website
```

Run the application:

```
go run main.go
```

Open your browser and navigate to:

```
http://localhost:8080
```

License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments
[`Advent of Code`](https://adventofcode.com/) by Eric Wastl
