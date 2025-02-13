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
const map = array.reduce((acc, word) => {
  acc[word] = (acc[word] || 0) + 1;
  // console.log(acc[word]);
  return acc;
}, {});

//const filter = array.filter((value) => map[value] > 2);
console.log(map);
