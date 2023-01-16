let fruits = ['apple', 'banana', 'orange', 'apple', 'orange'];

let uniqueFruits1 = Array.from([...new Set(fruits)]);
console.log(uniqueFruits1); // ['apple', 'banana', 'orange']

let uniqueFruits2 = [...new Set(fruits)];
console.log(uniqueFruits2); // ['apple', 'banana', 'orange']
