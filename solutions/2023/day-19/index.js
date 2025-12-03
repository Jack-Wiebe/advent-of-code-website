const fs = require("fs");

const input = fs.readFileSync("input", "utf8").split("\r\n");

console.log(input);

let parts = [];
let workflows = [];

for (let i = 0; i < input.length; i++) {
  if (input[i] == "") {
    parts = input.splice(i + 1, input.length);
    workflows = input;
    workflows.pop();
  }
}

console.log(parts, workflows);
