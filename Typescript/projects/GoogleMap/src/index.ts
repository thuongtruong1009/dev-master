/// <reference types="@types/google.maps" />

import Map from './map';
import User from './user';
import Company from './company';

const user = new User();
const company = new Company();

// console.log(company);

const map = new Map('map');
map.addMarker(user);
map.addMarker(company);

// console.log(company);
