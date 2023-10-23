type Name = 'Name A' | 'Name B' | 'Name C';

type Address = {
    id: number;
    name: string;
};

type Info = Name | Address;

type Person = {
    name: Info;
    address: Info;
};

const person = {
    name: 'Name A',
    address: {
        id: 1,
        name: 'Address B',
    },
} satisfies Person;

const updatePerson = person.name.toUpperCase();
