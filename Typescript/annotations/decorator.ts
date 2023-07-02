// class
function LogClass(target: Function) {
    console.log(`Class ${target.name} was defined!`);
}

@LogClass
class Person {
    constructor() {}
}

// method
function LogMethod(target: any, key: string, descriptor: PropertyDescriptor) {
    console.log(`Method ${key} was called!`);
}

class MyClass {
    @LogMethod
    myMethod() {
        console.log('Inside myMethod!');
    }
}

const instance = new MyClass();
instance.myMethod();

// property
function DefaultValue(value: any) {
    return (target: any, key: string) => {
        target[key] = value;
    };
}

class MyClass2 {
    @DefaultValue(10)
    myProperty: number;
}

const instance2 = new MyClass2();
console.log(instance.myProperty);

// parameter
function LogParameter(target: any, key: string, parameterIndex: number) {
    console.log(
        `Parameter at index ${parameterIndex} of method ${key} was called!`,
    );
}

class MyClass3 {
    myMethod(@LogParameter value: number) {
        console.log(`Inside myMethod with value ${value}`);
    }
}

const instance3 = new MyClass3();
instance3.myMethod(5);
