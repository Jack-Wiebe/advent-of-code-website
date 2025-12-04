const part1 = require("./part1.js");
const part2 = require("./part2.js");

const result1 = part1.run();
console.log("Result:", result1);

const result2 = part2.run();
console.log("Result:", result2);

let result = {
  Part1: result1,
  Part2: result2,
};
console.log(JSON.stringify(result));
