"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const Sorter_1 = __importDefault(require("./Sorter"));
class NumberCollection extends Sorter_1.default {
    constructor(collection) {
        super();
        this.collection = collection;
    }
    get length() {
        return this.collection.length;
    }
    compare(leftindex, rightindex) {
        return this.collection[leftindex] >= this.collection[rightindex];
    }
    swap(leftindex, rightindex) {
        const temp = this.collection[leftindex];
        this.collection[leftindex] = this.collection[rightindex];
        this.collection[rightindex] = temp;
    }
}
exports.default = NumberCollection;
