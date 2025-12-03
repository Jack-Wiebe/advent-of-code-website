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

for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[i].length; j++) {
    if (input[i][j] == "#") {
      rowMask[i] = 1;
      colMask[j] = 1;
    }
  }
}

console.log(colMask);
console.log(rowMask);

for (let i = colMask.length - 1; i >= 0; i--) {
  if (colMask[i] == 0) {
    addCol(i);
  }
}

for (let i = rowMask.length - 1; i >= 0; i--) {
  if (rowMask[i] == 0) {
    input.splice(i, 0, Array(input.length).fill("."));
  }
}

console.log(input);

const gals = [];
let count = 0;
for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[i].length; j++) {
    if (input[i][j] == "#") {
      gals.push({ row: i, col: j });
    }
  }
}

for (let i = 0; i < gals.length; i++) {
  for (let j = i + 1; j < gals.length; j++) {
    count += getDist(gals[i], gals[j]);
  }
}

console.log("Total distance:", count);

function addCol(index) {
  for (let i = 0; i < input.length; i++) {
    input[i].splice(index, 0, ".");
  }
}

function getDist(c0, c1) {
  const dist = Math.abs(c0.row - c1.row) + Math.abs(c0.col - c1.col);
  return dist;
}
