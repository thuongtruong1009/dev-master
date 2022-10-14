const source = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

function removeFirstTwo(list) {

    const [a, b, ...arr] = list;

    return arr;
}
const arr = removeFirstTwo(source);