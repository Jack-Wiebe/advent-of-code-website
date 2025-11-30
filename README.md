Advent of Code Solutions Web Interface
Written in Golang and Javascript

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
├── internal/
│   ├── types/
│   │   └── types.go          # Shared type definitions
│   └── handlers/
│       └── handlers.go       # HTTP request handlers
├── solutions/
│   ├── day1/
│   │   └── solution.go       # Day 1 solution
│   ├── day2/
│   │   └── solution.go       # Day 2 solution
│   └── solver.go             # Solution registry and interface
├── web/
│   └── static/
│       ├── index.html        # Main frontend
│       ├── styles.css        # Styling
│       └── app.js           # Frontend JavaScript
└── inputs/
    ├── day1.txt             # Day 1 input file
    └── day2.txt             # Day 2 input file
```

Quick Start
Clone the repository:

```
git clone <repository-url>
cd advent-of-code-web
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
Advent of Code by Eric Wastl
([`Advent of Code`](https://https://adventofcode.com/))

Go standard library and community packages
