const fs = require("fs");
const path = require("path");

module.exports = {
  run: function () {
    const filePath = path.join(__dirname, "input");
    const input = fs.readFileSync(filePath, "utf-8").split("\n");

    let sum = 0;

    let nums = {
      one: "1",
      two: "2",
      three: "3",
      four: "4",
      five: "5",
      six: "6",
      seven: "7",
      eight: "8",
      nine: "9",
    };

    const regex = /(\d|one|two|three|four|five|six|seven|eight|nine)/g;

    const test =
      /(?=(\d|(o(?=ne))|(t(?=wo))|(t(?=hree))|(f(?=our))|(f(?=ive))|(s(?=ix))|(s(?=even))|(e(?=ight))|(n(?=ine))))/g;

    for (var i = 0; i < input.length; i++) {
      let first = input[i].match(regex);

      let last = [...input[i].matchAll(test)]; //get index of match
      let t = input[i].slice(last[last.length - 1].index); //slice from last index into new string to match against
      last = t.match(regex);

      let num1 = first[0];
      let num2 = last[last.length - 1];

      if (/\D/.test(num1)) {
        num1 = nums[num1];
      }

      if (/\D/.test(num2)) {
        num2 = nums[num2];
      }
      console.log(num1 + num2, i + 1);
      sum += parseInt(num1 + num2);
    }

    console.log(sum);
    return sum;

    // let result = {
    //   Part2: sum,
    // };

    // console.log(JSON.stringify(result));
  },
};
