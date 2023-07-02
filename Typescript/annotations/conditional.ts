type NonNullable<T> = T extends null | undefined ? never : T;

const testDefinedType: NonNullable<string> = 'a';

const testUndefinedType: NonNullable<undefined> = undefined;
const testNullType: NonNullable<null> = null;
