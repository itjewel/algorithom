
const  canPlaceFlowers = function(flowerbed, n) {
    let length = flowerbed.length;
    let count = 0;
   for(let i = 0; i < length; i++){
       if(flowerbed[i] === 0 && (i === 0 || flowerbed[i -1] === 0) && (flowerbed[i+1] === 0 || i === length-1)){
           flowerbed[i] = 1;
           n--
       }
       if(n <= 0) return true;
   }
 return n <= 0;
};

const flowerbed = [1,0,0,0,1];
const n = 1;
console.log(canPlaceFlowers(flowerbed,n))