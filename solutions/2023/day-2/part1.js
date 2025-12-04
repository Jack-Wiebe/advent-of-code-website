const fs = require("fs");
const path = require("path");

module.exports = {
  run: function () {
    //12 red cubes, 13 green cubes, and 14 blue
    const MAX = {
      red: 12,
      green: 13,
      blue: 14,
    };

    const filePath = path.join(__dirname, "input");
    const input = fs
      .readFileSync(filePath, "utf-8")
      .split("\n")
      .reduce((acc, e) => {
        acc.push(e.replace("Game ", "").replace(":", ";").split(";"));
        return acc;
      }, []);

    let check = true;
    let count = 0;
    for (let i = 0; i < input.length; i++) {
      check = true;
      let val = input[i][0];
      //game
      for (let j = 1; j < input[i].length; j++) {
        let rounds = input[i][j].split(",");
        console.log(rounds);
        //round
        for (let k = 0; k < rounds.length; k++) {
          let hand = rounds[k].split(" ");
          let num = hand[1];
          let color = hand[2];
          if (MAX[color] < num) {
            check = false;
          }
        }
      }
      if (check) {
        count += parseInt(val);
      }
    }

    console.log(count);
    return count;
  },
};
