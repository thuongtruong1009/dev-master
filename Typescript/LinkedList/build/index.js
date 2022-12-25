"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const CharacterCollection_1 = __importDefault(require("./CharacterCollection"));
const NumberCollection_1 = __importDefault(require("./NumberCollection"));
const characters = new CharacterCollection_1.default('Xaba');
const numbers = new NumberCollection_1.default([23, -23, 10, 1]);
characters.sort();
numbers.sort();
console.log(characters.collection);
console.log(numbers.collection);
