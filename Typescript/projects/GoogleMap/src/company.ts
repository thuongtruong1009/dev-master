import { faker } from '@faker-js/faker';
import { ICompanyMap } from './interface';

export default class Company implements ICompanyMap {
    companyName: string;
    catchPhrase: string;
    location: {
        lat: number;
        lng: number;
    };
    color = 'red';

    constructor() {
        this.companyName = faker.company.name();
        this.catchPhrase = faker.company.catchPhrase();
        this.location = {
            lat: +faker.address.latitude(),
            lng: +faker.address.longitude(),
        };
    }

    markerContent(): string {
        return `
      <div>
        <h1>Company Name: ${this.companyName}</h1>
        <h3>Catchphrase: ${this.catchPhrase}</h3>
      </div>
      `;
    }
}
