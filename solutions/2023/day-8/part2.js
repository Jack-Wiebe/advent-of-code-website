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

const paths = [];
const endPoints = [];
for (const [key] of Object.entries(map)) {
  //console.log(key);
  if (key.slice(-1) == "A") paths.push(key);
  if (key.slice(-1) == "Z") endPoints.push(key);
}

let count = 0;
let end = false;

while (!end) {
  end = moveToNext();
}
console.log(count);

function moveToNext() {
  let check = true;
  for (let i = 0; i < paths.length; i++) {
    // console.log(direction[count % (direction.length - 1)]);
    // console.log(paths[i]);

    if (direction[count % (direction.length - 1)] === "R") {
      paths[i] = map[paths[i]].right;
    } else if (direction[count % (direction.length - 1)] === "L") {
      paths[i] = map[paths[i]].left;
    }

    //console.log(paths[i]);

    if (paths[i].slice(-1) != "Z") check = false;
  }
  count++;
  return check;
}
