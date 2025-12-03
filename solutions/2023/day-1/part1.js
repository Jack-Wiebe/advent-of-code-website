const fs = require("fs");
const path = require("path");

module.exports = {
  run: function () {
    const filePath = path.join(__dirname, "input");
    const input = fs.readFileSync(filePath, "utf-8").split("\n");

    console.log(input);

    let sum = 0;

    const regex = /\D/g;

    for (var i = 0; i < input.length; i++) {
      console.log(typeof input[i]);
      const digits = input[i].replace(regex, "");
      if (digits) console.log(digits[0] + digits[digits.length - 1]);
      sum += parseInt(digits[0] + digits[digits.length - 1]);
    }

    console.log(sum);
    return sum;

    // let result = {
    //   Part1: sum,
    // };

    // console.log(JSON.stringify(result));
  },
};
