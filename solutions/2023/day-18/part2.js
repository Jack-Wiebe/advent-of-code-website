const fs = require("fs");

const input = fs
  .readFileSync("input", "utf-8")
  .split("\r\n")
  .reduce((acc, e) => {
    acc.push(e.split(" "));
    return acc;
  }, []);

let area = 0;
let perimeter = 0;
let y = 0;

input.forEach((e) => {
  const hex = e[2];
  const dir = parseInt(hex[7]);
  const distance = Number("0x" + hex.slice(2, 7));
  console.log(hex, dir, distance);

  //add perimeter
  perimeter += distance;

  //calc area
  if (dir == 0) {
    area += y * distance;
  } else if (dir == 2) {
    area -= y * distance;
  } else if (dir == 3) {
    y += distance;
  } else if (dir == 1) {
    y -= distance;
  }
});

//formula for getting area is A + P/2 + 1
console.log(area + perimeter / 2 + 1);
