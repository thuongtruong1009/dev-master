export interface IMappable {
    location: {
        lat: number;
        lng: number;
    };
    color: string;
    markerContent(): string;
}

export interface IUserMap {
    username: string;
}

export interface ICompanyMap {
    companyName: string;
}
