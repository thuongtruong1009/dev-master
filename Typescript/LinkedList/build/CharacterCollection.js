"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const Sorter_1 = __importDefault(require("./Sorter"));
class CharacterCollection extends Sorter_1.default {
    constructor(collection) {
        super();
        this.collection = collection;
    }
    get length() {
        return this.collection.length;
    }
    compare(leftindex, rightindex) {
        return (this.collection[leftindex].toLocaleLowerCase() >=
            this.collection[rightindex].toLocaleLowerCase());
    }
    swap(leftindex, rightindex) {
        const characters = this.collection.split('');
        const temp = characters[leftindex];
        characters[leftindex] = characters[rightindex];
        characters[rightindex] = temp;
        this.collection = characters.join('');
    }
}
exports.default = CharacterCollection;
