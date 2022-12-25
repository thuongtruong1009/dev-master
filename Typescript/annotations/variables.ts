// variable type
let color: string[] = ['red'];

// class type
class Car {}
let car: Car = new Car();

// object type
let point: { x: number; y: number } = { x: 20, y: 10 };

const json = '{"x": 10, "y": 20}';
const data: { x: number; y: number } = JSON.parse(json);
console.log(data.x);

//function type
const logNum: (i: number) => void = (i) => {
    console.log(i);
};

const words = ['red', 'green', 'yellow'];
let foundWord: boolean;
for (let i = 0; i < words.length; i++) {
    if (words[i] === 'green') {
        foundWord = true;
        break;
    }
}

// Value whose type cannot be infered correctly
const numbers = [-10, 45, 98];
let numberAboveZero: number | boolean = false;

for (let i = 0; i < numbers.length; i++) {
    if (numbers[i] > 0) {
        numberAboveZero = numbers[i];
    } else {
        numberAboveZero = false;
    }
}
