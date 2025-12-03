const fs = require("fs");

//12 red cubes, 13 green cubes, and 14 blue

const input = fs
  .readFileSync("input", "utf-8")
  .split("\r\n")
  .reduce((acc, e) => {
    acc.push(e.replace("Game ", "").replace(":", ";").split(";"));
    return acc;
  }, []);

let count = 0;

let max = {
  red: 0,
  green: 0,
  blue: 0,
};

for (let i = 0; i < input.length; i++) {
  //game
  for (let j = 1; j < input[i].length; j++) {
    let rounds = input[i][j].split(",");
    console.log(rounds);
    //round
    for (let k = 0; k < rounds.length; k++) {
      let hand = rounds[k].split(" ");
      let num = parseInt(hand[1]);
      let color = hand[2];
      if (max[color] < num) {
        //console.log(num, "is greater than", max[color]);
        max[color] = num;
      }
    }
  }
  //console.log(max.red, max.blue, max.green);
  count += max.red * max.blue * max.green;
  max = {
    red: 0,
    green: 0,
    blue: 0,
  };
}

console.log(count);
