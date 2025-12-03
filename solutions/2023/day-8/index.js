const fs = require("fs");

const input = fs
  .readFileSync("./day-8", "utf-8")
  .split("\n")
  .filter((e) => e.length > 1);

const direction = input.shift();
const map = input.reduce((acc, e) => {
  const regex = /[(]|[)]|\s/g;
  const key = e.split(" = ")[0];
  const vals = e.split(" = ")[1].replace(regex, "").split(",");

  acc[key] = { left: vals[0], right: vals[1] };
  return acc;
}, {});

let step = "AAA";
let count = 0;

while (step != "ZZZ") {
  loopDirections();
}
console.log(count);

function loopDirections() {
  for (let i = 0; i < direction.length - 1; i++, count++) {
    console.log(step, count, direction[i]);
    if (step == "ZZZ") {
      return;
    } else if (direction[i] === "R") {
      step = map[step].right;
    } else if (direction[i] === "L") {
      step = map[step].left;
    }
  }
}
