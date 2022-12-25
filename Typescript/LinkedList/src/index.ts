import CharacterCollection from './CharacterCollection';
import NumberCollection from './NumberCollection';
import LinkedList from './LinkedList';

const characters = new CharacterCollection('Xaba');
const numbers = new NumberCollection([23, -23, 10, 1]);

characters.sort();
numbers.sort();

console.log(characters.collection);
console.log(numbers.collection);
