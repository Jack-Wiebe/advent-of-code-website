const fs = require("fs");

const input = fs.readFileSync("input", "utf-8").split("\r\n");

const games = input.reduce((acc, e) => {
  const str = e.split("|");
  const card = str[0].split(":");

  acc.push({
    card: card[0],
    wins: card[1].split(" ").filter((e) => e != ""),
    nums: str[1].split(" ").filter((e) => e != ""),
  });
  return acc;
}, []);

console.log(games);
let count = 0;
let matches = 0;

for (let i = 0; i < games.length; i++) {
  for (let j = 0; j < games[i].wins.length; j++) {
    for (let k = 0; k < games[i].nums.length; k++) {
      if (games[i].wins[j] == games[i].nums[k]) {
        matches++;
      }
    }
  }

  if (matches > 0) {
    count += Math.pow(2, matches - 1);
    matches = 0;
  }

  console.log(games[i].card, "count:", count);
}

console.log(count);
