
const mergeAlternately = function (word1, word2) {
    const maxValue = Math.max(word1.length, word2.length);
    const response = [];
   for(let i=0; i<maxValue; i++){
       word1[i] && response.push(word1[i]);
       word2[i] && response.push(word2[i])
   }
   return response.join("");
};

const  word1 = "abc"
const word2 ="pqr"

console.log(mergeAlternately(word1,word2));