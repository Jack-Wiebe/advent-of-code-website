const fs = require("fs");

const input = fs
  .readFileSync("input", "utf-8")
  .split("\r\n")
  .reduce((acc, e) => {
    acc.push(e.split(" "));
    return acc;
  }, []);

var grid = [];
let rows, cols;

rows = 1000;
cols = 1000;

let area = 0;
let perimeter = 0;
let y = 0;

for (let i = 0; i < rows; i++) {
  grid[i] = [];
  for (let j = 0; j < cols; j++) {
    grid[i][j] = ".";
  }
}
console.log(grid);

const directions = {
  U: [-1, 0],
  R: [0, 1],
  D: [1, 0],
  L: [0, -1],
};

const coord = { x: 500, y: 500 };

//populate grid
input.forEach((e) => {
  const dir = directions[e[0]];
  console.log(dir);
  for (let i = 0; i < parseInt(e[1]); i++) {
    coord.x += dir[0];
    coord.y += dir[1];
    console.log(coord);
    grid[coord.x][coord.y] = "#";
    perimeter++;
  }
});

//draw grid
let str = "";
for (let r = 0; r < grid.length; r++) {
  for (let c = 0; c < grid[r].length; c++) {
    str += grid[r][c];
  }
  console.log(str);
  str = "";
}

addInnerArea();
console.log(area + perimeter / 2 + 1);

function addInnerArea() {
  input.forEach((e) => {
    if (e[0] == "R") {
      area += y * parseInt(e[1]);
    } else if (e[0] == "L") {
      area -= y * parseInt(e[1]);
    } else if (e[0] == "U") {
      y += parseInt(e[1]);
    } else if (e[0] == "D") {
      y -= parseInt(e[1]);
    }
  });
}
