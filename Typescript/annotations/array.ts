const cardMaker = ['toyota', 'vinfast', 'ford'];

const cardByMake = [['prius'], ['vfe34'], ['sliverado']];

const carOne = cardByMake[0][0];

// prevent adding incompatible type to
cardMaker.push(100);

const importantDate: (Date | string)[] = [new Date()];
importantDate.push('2020-01-01');

console.log(importantDate);
