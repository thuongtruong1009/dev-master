import { faker } from '@faker-js/faker';
import { IUserMap } from './interface';

export default class User implements IUserMap {
    username: string;
    location: {
        lat: number;
        lng: number;
    };
    color: string = 'red';

    constructor() {
        this.username = faker.name.firstName();
        this.location = {
            lat: +faker.address.latitude(),
            lng: +faker.address.longitude(),
        };
    }

    markerContent(): string {
        return `Username: ${this.username}`;
    }
}
