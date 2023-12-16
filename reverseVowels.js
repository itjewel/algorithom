
const reverseVowels = function(s) {
    checkArray = ['a','A','e','E','I','i','o','O','u',"U"];
    let isVowel = (e) => checkArray.includes(e);
    let stringToArray = s.split("");
    let j = 0;
    let vowel = '';
    // get vowel
    stringToArray.map((value)=>{
        if(isVowel(value)){
            j++;
            vowel += value;
        }
    })

     // Revers word
    stringToArray.map((value,index)=>{
        if(isVowel(value)){
           stringToArray[index] = vowel[--j];
        }
    })

return stringToArray.join("")
    
};

const s = "leetcode";
console.log(reverseVowels(s))