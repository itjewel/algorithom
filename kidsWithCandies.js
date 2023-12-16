
const  kidsWithCandies = function(candies, extraCandies) {
    const maxValue = Math.max(...candies)
    const res = [];
    for(let i = 0; i<=maxValue; i++){
        candies[i] && ((candies[i] + extraCandies) >= maxValue ? res.push(true) : res.push(false));
    }
    return res;
};

const candies = [4,2,1,1,2];
const  extraCandies = 1
console.log(kidsWithCandies(candies,extraCandies))