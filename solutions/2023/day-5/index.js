const fs = require("fs");

const input = fs
  .readFileSync("d5", "utf-8")
  .split("\r\n\r\n")
  .reduce((acc, e) => {
    acc.push(e.split("\r\n"));

    return acc;
  }, []);

console.log(input);

//const table = [];
const ranges = [];

for (let i = 1; i < input.length; i++) {
  //const map = new Map();
  const set = [];

  for (let j = 1; j < input[i].length; j++) {
    const range = input[i][j].split(" ");

    //this section takes too long and too much memory allocation - rewrite using ranges
    //if source seed is within range then process the acutal offset of the destination
    //instead of trying to map out all the data in advance
    // for (let k = 0; k < parseInt(range[2]); k++) {
    //   map.set(parseInt(range[1]) + k, parseInt(range[0]) + k);
    // }

    set.push({
      start: parseInt(range[1]),
      end: parseInt(range[1]) + parseInt(range[2]),
      offset: parseInt(range[0]) - parseInt(range[1]),
    });
  }

  ranges.push(set);
  //table.push(map);
}
console.log(ranges);
const seeds = input[0][0].split(" ");
console.log("seeds", seeds);

let next;
let location = Number.MAX_VALUE;

for (let i = 1; i < seeds.length; i++) {
  next = parseInt(seeds[i]);
  console.log("seed", next);
  for (let j = 0; j < ranges.length; j++) {
    for (let k = 0; k < ranges[j].length; k++) {
      if (next >= ranges[j][k].start && next <= ranges[j][k].end) {
        next = next + ranges[j][k].offset;
        break;
      }
    }
  }
  console.log("destination value is:", next);
  if (next < location) location = next;
}
console.log("Closest Location is:", location);
