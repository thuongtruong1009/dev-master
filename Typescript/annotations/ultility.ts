interface Person {
    name: string;
    age: number;
    email: string;
}

const partialPerson: Partial<Person> = {
    name: 'John',
    age: 20,
};

const readonlyPerson: Readonly<Person> = {
    name: 'John',
    age: 20,
    email: '',
};

const pickNamAge: Pick<Person, 'name' | 'age'> = {
    name: 'John',
    age: 20,
};

const emptyEmail: Omit<Person, 'email'> = {
    name: 'John',
    age: 20,
};

const getUser1 = (id: string): Promise<Person> => {
  return new Promise((resolve, reject)=> {
    resolve({
      name: 'John',
      age: 20,
      email: ''
    })
  })
}

type getUser2 = ReturnType<typeof getUser1>;