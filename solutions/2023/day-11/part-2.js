const fs = require("fs");

const input = fs
  .readFileSync("input", "utf-8")
  .split("\r\n")
  .reduce((acc, e) => {
    acc.push(e.split(""));
    return acc;
  }, []);

console.log(input);

const colMask = Array(input.length).fill(0);
const rowMask = Array(input[0].length).fill(0);

const gals = [];
let index = 0;
for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[i].length; j++) {
    if (input[i][j] == "#") {
      index++;
      rowMask[i] = 1;
      colMask[j] = 1;
      gals.push({ row: i, col: j, index: index });
    }
  }
}

console.log(colMask);
console.log(rowMask);
console.log("Number of galaxies", gals.length);

// find each pair of galaxies and determine how many intersects with colMask & rowMask

let count = 0;
let range;
const FACTOR = 1000000;
let pair = 0;
let inters = 0;
for (let i = 0; i < gals.length; i++) {
  for (let j = i + 1; j < gals.length; j++) {
    pair++;
    count += getDist(gals[i], gals[j]);
    range = getRange(gals[i], gals[j]);
    temp = addIntercepts(range);
    inters += temp;

    count += temp * (FACTOR - 1);
  }
}

console.log("Pairs of galaxies", pair);
console.log("intersections with expansions", inters);
console.log("Total distance:", count);

function getRange(c0, c1) {
  const r = {
    row: [c0.row, c1.row],
    col: [c0.col, c1.col],
    index: { a: c0.index, b: c1.index },
  };
  return r;
}

function getDist(c0, c1) {
  const dist = Math.abs(c0.row - c1.row) + Math.abs(c0.col - c1.col);
  return dist;
}

function addIntercepts(range) {
  let inters = 0;

  range.row[0];
  rDist = Math.abs(range.row[0] - range.row[1]);
  cDist = Math.abs(range.col[0] - range.col[1]);
  rStart = Math.min(range.row[0], range.row[1]);
  cStart = Math.min(range.col[0], range.col[1]);
  //rows
  for (let i = rStart; i < rStart + rDist; i++) {
    if (rowMask[i] == 0) {
      inters++;
    }
  }

  //cols
  for (let i = cStart; i < cStart + cDist; i++) {
    if (colMask[i] == 0) {
      inters++;
    }
  }

  /*console.log("Star a:",range.row[0],"-",range.col[0],"Star b:",range.row[1],"-",range.col[1],"intersects:",inters);*/

  return inters;
}
