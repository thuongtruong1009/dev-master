class TimestampMixin<TBase extends new (...args: any[]) => any>(Base: TBase) {

  constructor(...args: any[]) {
    super(...args);
  }

  getTimestamp() {
    return new Date();
  }
}

class MyBaseClass {
  constructor(public value: number) {}

  displayValue() {
    console.log(`The value is: ${this.value}`)
  }
}

class MyMixedClass extends TimestampMixin(MyBaseClass) {
  constructor(value: number) {
    super(value);
  }
}

const instance = new MyMixedClass(1);
instance.displayValue();

const timestamp = instance.getTimestamp();
console.log(`The timestamp is: ${timestamp}`)
