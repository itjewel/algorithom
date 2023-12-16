
const reverseWords = function(s) {
    let trimString = s.trim();
    let stringToArray = trimString.split(/[\s]+/); 
        let length = stringToArray.length;
        let newArray = [];
        stringToArray.map((value)=>{
            newArray.push(stringToArray[--length]);

        })
        return newArray.join(" ");
};

const s = "the sky is blue";
console.log(reverseWords(s))