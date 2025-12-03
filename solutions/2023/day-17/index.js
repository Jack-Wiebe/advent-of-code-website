const fs = require("fs");

const input = fs.readFileSync("./input", "utf-8").split("\n");

console.log(input);

const grid = input.map((e) => {
  return e.replace(/\D/g, "").split("");
});

const WIDTH = grid.length - 1;
const HEIGHT = grid[0].length - 1;

const directions = {
  NORTH: 0,
  EAST: 1,
  SOUTH: 2,
  WEST: 3,
};

const vectors = {
  NORTH: [0, 1],
  EAST: [1, 0],
  SOUTH: [0, -1],
  WEST: [-1, 0],
};

console.log(directions);

//const pos = pickPath(0, 0, directions.EAST);

let pos = { row: 0, col: 0, dir: directions.SOUTH };

console.log(pos);
console.log(WIDTH);
console.log(HEIGHT);
console.log(grid[WIDTH][HEIGHT]);
console.log(pos.col != grid[WIDTH][HEIGHT]);

//while (pos.col != grid[WIDTH][HEIGHT]) {}

for (let i = 0; i < 4; i++) {
  //console.log(pos.dir);
  pos = pickPath(pos.row, pos.col, pos.dir);
}
// //console.log(grid);
function pickPath(row, col, dir) {
  //grid[row + ,col];

  //check left
  //console.log(getVectorFromDirection(dir - 1));
  let count = 10;
  let r, c;

  //check left
  console.log(dir);
  console.log((dir - 1 + (directions.length - 1)) % directions.length);
  // console.log(
  //   getVectorFromDirection((dir - 1 + directions.length) % directions.length)
  // );
  console.log(getVectorFromDirection(dir));
  console.log(getVectorFromDirection(dir + 1));

  let v = getVectorFromDirection(dir - 1);
  //console.log(v);
  if (
    grid[row + v[0]] &&
    grid[row + v[0]][col + v[1]] &&
    grid[row + v[0]][col + v[1]] < count
  ) {
    count = grid[row + v[0]][col + v[1]];
    d = dir - 1;
    r = row + v[0];
    c = col + v[1];
  }

  //check forward
  v = getVectorFromDirection(dir);
  if (
    grid[row + v[0]] &&
    grid[row + v[0]][col + v[1]] &&
    grid[row + v[0]][col + v[1]] < count
  ) {
    count = grid[row + v[0]][col + v[1]];
    d = dir;
    r = row + v[0];
    c = col + v[1];
  }

  //check right
  v = getVectorFromDirection(dir + 1);
  if (
    grid[row + v[0]] &&
    grid[row + v[0]][col + v[1]] &&
    grid[row + v[0]][col + v[1]] < count
  ) {
    count = grid[row + v[0]][col + v[1]];
    d = dir + 1;
    r = row + v[0];
    c = col + v[1];
  }

  const test = Object.keys(directions).find((key) => directions[key] === d);

  //console.log(d);
  return { row: r, col: c, dir: d };
}

function getVectorFromDirection(value) {
  return vectors[
    Object.keys(directions).find((key) => directions[key] === value)
  ];
}
