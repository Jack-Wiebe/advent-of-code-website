const fs = require("fs");

const input = fs.readFileSync("input", "utf-8").split("\r\n");

const games = input.reduce((acc, e) => {
  const str = e.split("|");
  const gameName = str[0].split(":");

  acc.push({
    gameName: gameName[0],
    wins: gameName[1].split(" ").filter((e) => e != ""),
    nums: str[1].split(" ").filter((e) => e != ""),
  });
  return acc;
}, []);

console.log(games);
let count = 0;

//recursive call to copy cards

for (let i = 0; i < games.length; i++) {
  for (let j = 0; j < games[i].wins.length; j++) {
    for (let k = 0; k < games[i].nums.length; k++) {
      if (games[i].wins[j] == games[i].nums[k]) {
        console.log(
          "game",
          games[i].gameName,
          "compare",
          games[i].wins[j],
          "to",
          games[i].nums[k]
        );
        matches++;
      }
    }
  }
}

console.log(count);
