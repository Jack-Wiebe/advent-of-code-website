const fs = require("fs");

const w = 200000000000000;
const h = 400000000000000;
//px py pz @ vx vy vz

input = fs
  .readFileSync("input", "utf-8")
  .replaceAll("@", ",")
  .split("\r\n")
  .reduce((acc, e) => {
    acc.push(e.replaceAll(" ", "").split(","));
    return acc;
  }, [])
  .reduce((acc, e) => {
    acc.push({
      px: parseInt(e[0]),
      py: parseInt(e[1]),
      pz: parseInt(e[2]),
      vx: parseInt(e[3]),
      vy: parseInt(e[4]),
      vz: parseInt(e[5]),
    });
    return acc;
  }, []);

console.log(input, typeof input);

let count = 0;

//loop and compare every hailstorm
for (let i = 0; i <= input.length - 1; i++) {
  console.log("check hailstone pairs", i);
  for (let j = i + 1; j <= input.length - 1; j++) {
    calculateIntersect(input[i], input[j]);
  }
}

console.log("Final Count:", count);

function calculateIntersect(h0, h1) {
  const l0 = getLine(h0);
  const l1 = getLine(h1);

  if (l0.a * l1.b == l0.b * l1.a) {
    console.log("Parallel");
    return;
  }
  const x = (l0.c * l1.b - l1.c * l0.b) / (l0.a * l1.b - l1.a * l0.b);
  const y = (l1.c * l0.a - l0.c * l1.a) / (l0.a * l1.b - l1.a * l0.b);

  if (
    !(
      (x - h0.px) * h0.vx >= 0 &&
      (y - h0.py) * h0.vy >= 0 &&
      (x - h1.px) * h1.vx >= 0 &&
      (y - h1.py) * h1.vy >= 0
    )
  ) {
    console.log("Past");
    return;
  }

  if (w <= x && x <= h && w <= y && y <= h) {
    console.log("Inersection", x, y);
    count++;
  }
}

function getLine(h) {
  return { a: h.vy, b: -h.vx, c: h.vy * h.px - h.vx * h.py };
}
