const array = [
  "apple",
  "potato",
  "apple",
  "orange",
  "lemon",
  "apple",
  "nut",
  "lemon",
];

const count = array.length;
const map = {};
for (let i = 0; i < count; i++) {
  map[array[i]] = (map[array[i]] || 0) + 1;
}

console.log(map);
