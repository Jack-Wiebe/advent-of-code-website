const fs = require("fs");

const input = fs.readFileSync("input", "utf-8").split("\r\n");

console.log(input);

let cols = new Array(input[0].length).fill("");
let count = 0;

for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[i].length; j++) {
    cols[j] += input[i][j];
  }
}

console.log(cols);

for (let i = 0; i < cols.length; i++) {
  for (let j = 0; j < cols[i].length; j++) {
    if (cols[i][j] == "O") {
      let index = j;
      while (cols[i][index - 1] == ".") {
        cols[i] = swapLeft(cols[i], index);
        index--;
      }
      count += cols[i].length - index;
    }
  }
}

console.log(cols);
console.log(count);

function swapLeft(str, index) {
  let charArray = str.split("");

  [charArray[index], charArray[index - 1]] = [
    charArray[index - 1],
    charArray[index],
  ];

  return charArray.join("");
}
