namespace MyNamespace {
    export class MyClass {
        constructor(public value: number) {}

        displayValue() {
            console.log(`The value is: ${this.value}`);
        }
    }
}

const instance1 = new MyNamespace.MyClass(5);
instance1.displayValue();

// import
import MyClass = MyNamespace.MyClass;

const instance2 = new MyClass(10);
instance2.displayValue();

//nested

namespace OuterNamespace {
    export namespace InnerNamespace {
        export class MyClass {
            constructor(public value: number) {}

            displayValue() {
                console.log(`The value is: ${this.value}`);
            }
        }
    }
}

const instance = new OuterNamespace.InnerNamespace.MyClass(15);
instance.displayValue();
